/*
 *
 * k6 - a next-generation load testing tool
 * Copyright (C) 2020 Load Impact
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

package modules

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/dop251/goja"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules/k6/http"
	"go.k6.io/k6/lib"
)

const extPrefix string = "k6/x/"

//nolint:gochecknoglobals
var (
	modules = make(map[string]interface***REMOVED******REMOVED***)
	mx      sync.RWMutex
)

// Register the given mod as an external JavaScript module that can be imported
// by name. The name must be unique across all registered modules and must be
// prefixed with "k6/x/", otherwise this function will panic.
func Register(name string, mod interface***REMOVED******REMOVED***) ***REMOVED***
	if !strings.HasPrefix(name, extPrefix) ***REMOVED***
		panic(fmt.Errorf("external module names must be prefixed with '%s', tried to register: %s", extPrefix, name))
	***REMOVED***

	mx.Lock()
	defer mx.Unlock()

	if _, ok := modules[name]; ok ***REMOVED***
		panic(fmt.Sprintf("module already registered: %s", name))
	***REMOVED***
	modules[name] = mod
***REMOVED***

// HasModuleInstancePerVU should be implemented by all native Golang modules that
// would require per-VU state. k6 will call their NewModuleInstancePerVU() methods
// every time a VU imports the module and use its result as the returned object.
type HasModuleInstancePerVU interface ***REMOVED***
	NewModuleInstancePerVU() interface***REMOVED******REMOVED***
***REMOVED***

// IsModuleV2 is the interface js modules should implement to get the version 2 of the system
type IsModuleV2 interface ***REMOVED***
	// NewModuleInstance will get InstanceCore that should provide the module with *everything* it needs and return an
	// Instance implementation (embedding the InstanceCore).
	// This method will be called for *each* require/import and return an object for VUs.
	NewModuleInstance(InstanceCore) Instance
***REMOVED***

// checks that modules implement HasModuleInstancePerVU
// this is done here as otherwise there will be a loop if the module imports this package
var _ HasModuleInstancePerVU = http.New()

// GetJSModules returns a map of all registered js modules
func GetJSModules() map[string]interface***REMOVED******REMOVED*** ***REMOVED***
	mx.Lock()
	defer mx.Unlock()
	result := make(map[string]interface***REMOVED******REMOVED***, len(modules))

	for name, module := range modules ***REMOVED***
		result[name] = module
	***REMOVED***

	return result
***REMOVED***

// Instance is what a module needs to return
type Instance interface ***REMOVED***
	InstanceCore
	Exports() Exports
***REMOVED***

func getInterfaceMethods() []string ***REMOVED***
	var t Instance
	T := reflect.TypeOf(&t).Elem()
	result := make([]string, T.NumMethod())

	for i := range result ***REMOVED***
		result[i] = T.Method(i).Name
	***REMOVED***

	return result
***REMOVED***

// InstanceCore is something that will be provided to modules and they need to embed it in ModuleInstance
type InstanceCore interface ***REMOVED***
	GetContext() context.Context

	// GetInitEnv returns common.InitEnvironment instance if present
	GetInitEnv() *common.InitEnvironment

	// GetState returns lib.State if any is present
	GetState() *lib.State

	// GetRuntime returns the goja.Runtime for the current VU
	GetRuntime() *goja.Runtime

	// sealing field will help probably with pointing users that they just need to embed this in their Instance
	// implementations
***REMOVED***

// Exports is representation of ESM exports of a module
type Exports struct ***REMOVED***
	// Default is what will be the `default` export of a module
	Default interface***REMOVED******REMOVED***
	// Named is the named exports of a module
	Named map[string]interface***REMOVED******REMOVED***
***REMOVED***

// GenerateExports generates an Exports from a module akin to how common.Bind does now.
// it also skips anything that is expected will not want to be exported such as methods and fields coming from
// interfaces defined in this package.
func GenerateExports(v interface***REMOVED******REMOVED***) Exports ***REMOVED***
	exports := make(map[string]interface***REMOVED******REMOVED***)
	val := reflect.ValueOf(v)
	typ := val.Type()
	badNames := getInterfaceMethods()
outer:
	for i := 0; i < typ.NumMethod(); i++ ***REMOVED***
		meth := typ.Method(i)
		for _, badname := range badNames ***REMOVED***
			if meth.Name == badname ***REMOVED***
				continue outer
			***REMOVED***
		***REMOVED***
		name := common.MethodName(typ, meth)

		fn := val.Method(i)
		exports[name] = fn.Interface()
	***REMOVED***

	// If v is a pointer, we need to indirect it to access its fields.
	if typ.Kind() == reflect.Ptr ***REMOVED***
		val = val.Elem()
		typ = val.Type()
	***REMOVED***
	var mic InstanceCore // TODO move this out
	for i := 0; i < typ.NumField(); i++ ***REMOVED***
		field := typ.Field(i)
		if field.Type == reflect.TypeOf(&mic).Elem() ***REMOVED***
			continue
		***REMOVED***
		name := common.FieldName(typ, field)
		if name != "" ***REMOVED***
			exports[name] = val.Field(i).Interface()
		***REMOVED***
	***REMOVED***
	return Exports***REMOVED***Default: exports, Named: exports***REMOVED***
***REMOVED***
