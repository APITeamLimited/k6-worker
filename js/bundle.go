/*
 *
 * k6 - a next-generation load testing tool
 * Copyright (C) 2016 Load Impact
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package js

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"runtime"

	"github.com/dop251/goja"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"gopkg.in/guregu/null.v3"

	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/compiler"
	"go.k6.io/k6/lib"
	"go.k6.io/k6/lib/consts"
	"go.k6.io/k6/lib/metrics"
	"go.k6.io/k6/loader"
)

// A Bundle is a self-contained bundle of scripts and resources.
// You can use this to produce identical BundleInstance objects.
type Bundle struct ***REMOVED***
	Filename *url.URL
	Source   string
	Program  *goja.Program
	Options  lib.Options

	BaseInitContext *InitContext

	RuntimeOptions    lib.RuntimeOptions
	CompatibilityMode lib.CompatibilityMode // parsed value
	registry          *metrics.Registry

	exports map[string]goja.Callable
***REMOVED***

// A BundleInstance is a self-contained instance of a Bundle.
type BundleInstance struct ***REMOVED***
	Runtime *goja.Runtime
	Context *context.Context

	// TODO: maybe just have a reference to the Bundle? or save and pass rtOpts?
	env map[string]string

	exports map[string]goja.Callable
***REMOVED***

// NewBundle creates a new bundle from a source file and a filesystem.
func NewBundle(
	logger logrus.FieldLogger, src *loader.SourceData, filesystems map[string]afero.Fs, rtOpts lib.RuntimeOptions,
	registry *metrics.Registry,
) (*Bundle, error) ***REMOVED***
	compatMode, err := lib.ValidateCompatibilityMode(rtOpts.CompatibilityMode.String)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***

	// Compile sources, both ES5 and ES6 are supported.
	code := string(src.Data)
	c := compiler.New(logger)
	c.Options = compiler.Options***REMOVED***
		CompatibilityMode: compatMode,
		Strict:            true,
		SourceMapLoader:   generateSourceMapLoader(logger, filesystems),
	***REMOVED***
	pgm, _, err := c.Compile(code, src.URL.String(), true)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	// Make a bundle, instantiate it into a throwaway VM to populate caches.
	rt := goja.New()
	bundle := Bundle***REMOVED***
		Filename: src.URL,
		Source:   code,
		Program:  pgm,
		BaseInitContext: NewInitContext(logger, rt, c, compatMode, new(context.Context),
			filesystems, loader.Dir(src.URL)),
		RuntimeOptions:    rtOpts,
		CompatibilityMode: compatMode,
		exports:           make(map[string]goja.Callable),
		registry:          registry,
	***REMOVED***
	if err = bundle.instantiate(logger, rt, bundle.BaseInitContext, 0); err != nil ***REMOVED***
		return nil, err
	***REMOVED***

	err = bundle.getExports(logger, rt, true)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***

	return &bundle, nil
***REMOVED***

// NewBundleFromArchive creates a new bundle from an lib.Archive.
func NewBundleFromArchive(
	logger logrus.FieldLogger, arc *lib.Archive, rtOpts lib.RuntimeOptions, registry *metrics.Registry,
) (*Bundle, error) ***REMOVED***
	if arc.Type != "js" ***REMOVED***
		return nil, fmt.Errorf("expected bundle type 'js', got '%s'", arc.Type)
	***REMOVED***

	if !rtOpts.CompatibilityMode.Valid ***REMOVED***
		// `k6 run --compatibility-mode=whatever archive.tar` should override
		// whatever value is in the archive
		rtOpts.CompatibilityMode = null.StringFrom(arc.CompatibilityMode)
	***REMOVED***
	compatMode, err := lib.ValidateCompatibilityMode(rtOpts.CompatibilityMode.String)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***

	c := compiler.New(logger)
	c.Options = compiler.Options***REMOVED***
		Strict:            true,
		CompatibilityMode: compatMode,
		SourceMapLoader:   generateSourceMapLoader(logger, arc.Filesystems),
	***REMOVED***
	pgm, _, err := c.Compile(string(arc.Data), arc.FilenameURL.String(), true)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	rt := goja.New()
	initctx := NewInitContext(logger, rt, c, compatMode,
		new(context.Context), arc.Filesystems, arc.PwdURL)

	env := arc.Env
	if env == nil ***REMOVED***
		// Older archives (<=0.20.0) don't have an "env" property
		env = make(map[string]string)
	***REMOVED***
	for k, v := range rtOpts.Env ***REMOVED***
		env[k] = v
	***REMOVED***
	rtOpts.Env = env

	bundle := &Bundle***REMOVED***
		Filename:          arc.FilenameURL,
		Source:            string(arc.Data),
		Program:           pgm,
		Options:           arc.Options,
		BaseInitContext:   initctx,
		RuntimeOptions:    rtOpts,
		CompatibilityMode: compatMode,
		exports:           make(map[string]goja.Callable),
		registry:          registry,
	***REMOVED***

	if err = bundle.instantiate(logger, rt, bundle.BaseInitContext, 0); err != nil ***REMOVED***
		return nil, err
	***REMOVED***

	// Grab exported objects, but avoid overwriting options, which would
	// be initialized from the metadata.json at this point.
	err = bundle.getExports(logger, rt, false)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***

	return bundle, nil
***REMOVED***

func (b *Bundle) makeArchive() *lib.Archive ***REMOVED***
	arc := &lib.Archive***REMOVED***
		Type:              "js",
		Filesystems:       b.BaseInitContext.filesystems,
		Options:           b.Options,
		FilenameURL:       b.Filename,
		Data:              []byte(b.Source),
		PwdURL:            b.BaseInitContext.pwd,
		Env:               make(map[string]string, len(b.RuntimeOptions.Env)),
		CompatibilityMode: b.CompatibilityMode.String(),
		K6Version:         consts.Version,
		Goos:              runtime.GOOS,
	***REMOVED***
	// Copy env so changes in the archive are not reflected in the source Bundle
	for k, v := range b.RuntimeOptions.Env ***REMOVED***
		arc.Env[k] = v
	***REMOVED***

	return arc
***REMOVED***

// getExports validates and extracts exported objects
func (b *Bundle) getExports(logger logrus.FieldLogger, rt *goja.Runtime, options bool) error ***REMOVED***
	exportsV := rt.Get("exports")
	if goja.IsNull(exportsV) || goja.IsUndefined(exportsV) ***REMOVED***
		return errors.New("exports must be an object")
	***REMOVED***
	exports := exportsV.ToObject(rt)

	for _, k := range exports.Keys() ***REMOVED***
		v := exports.Get(k)
		if fn, ok := goja.AssertFunction(v); ok && k != consts.Options ***REMOVED***
			b.exports[k] = fn
			continue
		***REMOVED***
		switch k ***REMOVED***
		case consts.Options:
			if !options ***REMOVED***
				continue
			***REMOVED***
			data, err := json.Marshal(v.Export())
			if err != nil ***REMOVED***
				return err
			***REMOVED***
			dec := json.NewDecoder(bytes.NewReader(data))
			dec.DisallowUnknownFields()
			if err := dec.Decode(&b.Options); err != nil ***REMOVED***
				if uerr := json.Unmarshal(data, &b.Options); uerr != nil ***REMOVED***
					return uerr
				***REMOVED***
				logger.WithError(err).Warn("There were unknown fields in the options exported in the script")
			***REMOVED***
		case consts.SetupFn:
			return errors.New("exported 'setup' must be a function")
		case consts.TeardownFn:
			return errors.New("exported 'teardown' must be a function")
		***REMOVED***
	***REMOVED***

	if len(b.exports) == 0 ***REMOVED***
		return errors.New("no exported functions in script")
	***REMOVED***

	return nil
***REMOVED***

// Instantiate creates a new runtime from this bundle.
func (b *Bundle) Instantiate(
	logger logrus.FieldLogger, vuID uint64, vuImpl *moduleVUImpl,
) (bi *BundleInstance, instErr error) ***REMOVED***
	// Instantiate the bundle into a new VM using a bound init context. This uses a context with a
	// runtime, but no state, to allow module-provided types to function within the init context.
	vuImpl.runtime = goja.New()
	init := newBoundInitContext(b.BaseInitContext, vuImpl)
	if err := b.instantiate(logger, vuImpl.runtime, init, vuID); err != nil ***REMOVED***
		return nil, err
	***REMOVED***

	rt := vuImpl.runtime
	bi = &BundleInstance***REMOVED***
		Runtime: rt,
		Context: vuImpl.ctxPtr,
		exports: make(map[string]goja.Callable),
		env:     b.RuntimeOptions.Env,
	***REMOVED***

	// Grab any exported functions that could be executed. These were
	// already pre-validated in cmd.validateScenarioConfig(), just get them here.
	exports := rt.Get("exports").ToObject(rt)
	for k := range b.exports ***REMOVED***
		fn, _ := goja.AssertFunction(exports.Get(k))
		bi.exports[k] = fn
	***REMOVED***

	jsOptions := rt.Get("options")
	var jsOptionsObj *goja.Object
	if jsOptions == nil || goja.IsNull(jsOptions) || goja.IsUndefined(jsOptions) ***REMOVED***
		jsOptionsObj = rt.NewObject()
		rt.Set("options", jsOptionsObj)
	***REMOVED*** else ***REMOVED***
		jsOptionsObj = jsOptions.ToObject(rt)
	***REMOVED***
	b.Options.ForEachSpecified("json", func(key string, val interface***REMOVED******REMOVED***) ***REMOVED***
		if err := jsOptionsObj.Set(key, val); err != nil ***REMOVED***
			instErr = err
		***REMOVED***
	***REMOVED***)

	return bi, instErr
***REMOVED***

// Instantiates the bundle into an existing runtime. Not public because it also messes with a bunch
// of other things, will potentially thrash data and makes a mess in it if the operation fails.
func (b *Bundle) instantiate(logger logrus.FieldLogger, rt *goja.Runtime, init *InitContext, vuID uint64) error ***REMOVED***
	rt.SetFieldNameMapper(common.FieldNameMapper***REMOVED******REMOVED***)
	rt.SetRandSource(common.NewRandSource())

	exports := rt.NewObject()
	rt.Set("exports", exports)
	module := rt.NewObject()
	_ = module.Set("exports", exports)
	rt.Set("module", module)

	env := make(map[string]string, len(b.RuntimeOptions.Env))
	for key, value := range b.RuntimeOptions.Env ***REMOVED***
		env[key] = value
	***REMOVED***
	rt.Set("__ENV", env)
	rt.Set("__VU", vuID)
	rt.Set("console", common.Bind(rt, newConsole(logger), init.ctxPtr))

	if init.compatibilityMode == lib.CompatibilityModeExtended ***REMOVED***
		rt.Set("global", rt.GlobalObject())
	***REMOVED***

	// TODO: get rid of the unused ctxPtr, use a real external context (so we
	// can interrupt), build the common.InitEnvironment earlier and reuse it
	initenv := &common.InitEnvironment***REMOVED***
		Logger:      logger,
		FileSystems: init.filesystems,
		CWD:         init.pwd,
		Registry:    b.registry,
	***REMOVED***
	init.moduleVUImpl.initEnv = initenv
	ctx := common.WithInitEnv(context.Background(), initenv)
	*init.ctxPtr = common.WithRuntime(ctx, rt)
	unbindInit := common.BindToGlobal(rt, common.Bind(rt, init, init.ctxPtr))
	if _, err := rt.RunProgram(b.Program); err != nil ***REMOVED***
		var exception *goja.Exception
		if errors.As(err, &exception) ***REMOVED***
			err = &scriptException***REMOVED***inner: exception***REMOVED***
		***REMOVED***
		return err
	***REMOVED***
	unbindInit()
	*init.ctxPtr = nil

	// If we've already initialized the original VU init context, forbid
	// any subsequent VUs to open new files
	if vuID == 0 ***REMOVED***
		init.allowOnlyOpenedFiles()
	***REMOVED***

	rt.SetRandSource(common.NewRandSource())

	return nil
***REMOVED***

func generateSourceMapLoader(logger logrus.FieldLogger, filesystems map[string]afero.Fs,
) func(path string) ([]byte, error) ***REMOVED***
	return func(path string) ([]byte, error) ***REMOVED***
		u, err := url.Parse(path)
		if err != nil ***REMOVED***
			return nil, err
		***REMOVED***
		data, err := loader.Load(logger, filesystems, u, path)
		if err != nil ***REMOVED***
			return nil, err
		***REMOVED***
		return data.Data, nil
	***REMOVED***
***REMOVED***
