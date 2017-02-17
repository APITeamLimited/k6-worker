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
	"encoding/json"
	"fmt"
	"github.com/GeertJohan/go.rice"
	log "github.com/Sirupsen/logrus"
	"github.com/loadimpact/k6/lib"
	"github.com/loadimpact/k6/stats"
	"github.com/pkg/errors"
	"github.com/robertkrimen/otto"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

const wrapper = "(function() ***REMOVED*** var e = ***REMOVED******REMOVED***; (function(exports) ***REMOVED***%s\n***REMOVED***)(e); return e; ***REMOVED***)();"

var (
	libBox      = rice.MustFindBox("lib")
	polyfillBox = rice.MustFindBox("node_modules/babel-polyfill")
)

type Runtime struct ***REMOVED***
	VM      *otto.Otto
	Exports map[string]otto.Value
	Metrics map[string]*stats.Metric
	Options lib.Options

	lib map[string]otto.Value
***REMOVED***

func New() (*Runtime, error) ***REMOVED***
	rt := &Runtime***REMOVED***
		VM:      otto.New(),
		Exports: make(map[string]otto.Value),
		Metrics: make(map[string]*stats.Metric),
		lib:     make(map[string]otto.Value),
	***REMOVED***

	polyfillJS, err := polyfillBox.String("dist/polyfill.js")
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	polyfill, err := rt.VM.Compile("polyfill.js", polyfillJS)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	if _, err := rt.VM.Run(polyfill); err != nil ***REMOVED***
		return nil, err
	***REMOVED***

	if _, err := rt.loadLib("_global.js"); err != nil ***REMOVED***
		return nil, err
	***REMOVED***

	__ENV := make(map[string]string)
	for _, kv := range os.Environ() ***REMOVED***
		k, v := lib.SplitKV(kv)
		__ENV[k] = v
	***REMOVED***
	if err := rt.VM.Set("__ENV", __ENV); err != nil ***REMOVED***
		return nil, err
	***REMOVED***

	return rt, nil
***REMOVED***

func (r *Runtime) Load(src *lib.SourceData) (otto.Value, error) ***REMOVED***
	if err := r.VM.Set("__initapi__", &InitAPI***REMOVED***r: r***REMOVED***); err != nil ***REMOVED***
		return otto.UndefinedValue(), err
	***REMOVED***
	exp, err := r.loadSource(src)
	if err := r.VM.Set("__initapi__", nil); err != nil ***REMOVED***
		return otto.UndefinedValue(), err
	***REMOVED***
	return exp, err
***REMOVED***

func (r *Runtime) extractOptions(exports otto.Value, opts *lib.Options) error ***REMOVED***
	expObj := exports.Object()
	if expObj == nil ***REMOVED***
		return nil
	***REMOVED***

	v, err := expObj.Get("options")
	if err != nil ***REMOVED***
		return err
	***REMOVED***
	ev, err := v.Export()
	if err != nil ***REMOVED***
		return err
	***REMOVED***

	data, err := json.Marshal(ev)
	if err != nil ***REMOVED***
		return err
	***REMOVED***
	if err := json.Unmarshal(data, opts); err != nil ***REMOVED***
		return err
	***REMOVED***

	return nil
***REMOVED***

func (r *Runtime) loadSource(src *lib.SourceData) (otto.Value, error) ***REMOVED***
	path, err := filepath.Abs(src.Filename)
	if err != nil ***REMOVED***
		return otto.UndefinedValue(), err
	***REMOVED***

	// Don't re-compile repeated includes of the same module
	if exports, ok := r.Exports[path]; ok ***REMOVED***
		return exports, nil
	***REMOVED***
	exports, err := r.load(path, src.Data)
	if err != nil ***REMOVED***
		return otto.UndefinedValue(), err
	***REMOVED***
	r.Exports[path] = exports

	log.WithField("path", path).Debug("File loaded")

	return exports, nil
***REMOVED***

func (r *Runtime) loadFile(filename string) (otto.Value, error) ***REMOVED***
	path, err := filepath.Abs(filename)
	if err != nil ***REMOVED***
		return otto.UndefinedValue(), err
	***REMOVED***

	// Don't re-compile repeated includes of the same module
	if exports, ok := r.Exports[path]; ok ***REMOVED***
		return exports, nil
	***REMOVED***

	data, err := ioutil.ReadFile(path)
	if err != nil ***REMOVED***
		return otto.UndefinedValue(), err
	***REMOVED***
	exports, err := r.load(path, data)
	if err != nil ***REMOVED***
		return otto.UndefinedValue(), err
	***REMOVED***
	r.Exports[path] = exports

	log.WithField("path", path).Debug("File loaded")

	return exports, nil
***REMOVED***

func (r *Runtime) loadLib(filename string) (otto.Value, error) ***REMOVED***
	if exports, ok := r.lib[filename]; ok ***REMOVED***
		return exports, nil
	***REMOVED***

	data, err := libBox.Bytes(filename)
	if err != nil ***REMOVED***
		return otto.UndefinedValue(), err
	***REMOVED***
	exports, err := r.load(filename, data)
	if err != nil ***REMOVED***
		return otto.UndefinedValue(), err
	***REMOVED***
	r.lib[filename] = exports

	log.WithField("filename", filename).Debug("Library loaded")

	return exports, nil
***REMOVED***

func (r *Runtime) load(filename string, data []byte) (otto.Value, error) ***REMOVED***
	nodeNames := []string***REMOVED***"node", "nodejs", "node.exe"***REMOVED***
	var nodePath string
	for _, name := range nodeNames ***REMOVED***
		path, err := exec.LookPath(name)
		if err != nil ***REMOVED***
			if e, ok := err.(*exec.Error); ok && e.Err != exec.ErrNotFound ***REMOVED***
				return otto.UndefinedValue(), err
			***REMOVED***
			continue
		***REMOVED***
		nodePath = path
		break
	***REMOVED***
	if nodePath == "" ***REMOVED***
		return otto.UndefinedValue(), errors.New(
			"Couldn't find node, make sure it's in your PATH. " +
				"This is a TEMPORARY dependency and will be removed. " +
				"See: https://github.com/loadimpact/k6/issues/14",
		)
	***REMOVED***

	// Compile the file with Babel; this subprocess invocation is TEMPORARY:
	// https://github.com/robertkrimen/otto/pull/205
	cmd := exec.Command(nodePath, babel, "--presets", "latest", "--no-babelrc")
	cmd.Dir = babelDir
	cmd.Stdin = bytes.NewReader(data)
	cmd.Stderr = os.Stderr
	src, err := cmd.Output()
	if err != nil ***REMOVED***
		return otto.UndefinedValue(), err
	***REMOVED***

	// Use a wrapper function to turn the script into an exported module
	s, err := r.VM.Compile(filename, fmt.Sprintf(wrapper, string(src)))
	if err != nil ***REMOVED***
		return otto.UndefinedValue(), err
	***REMOVED***
	exports, err := r.VM.Run(s)
	if err != nil ***REMOVED***
		return otto.UndefinedValue(), err
	***REMOVED***

	// Extract script-defined options.
	var opts lib.Options
	if err := r.extractOptions(exports, &opts); err != nil ***REMOVED***
		return exports, err
	***REMOVED***
	r.Options = r.Options.Apply(opts)

	return exports, nil
***REMOVED***
