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
	"errors"
	"fmt"
	"github.com/loadimpact/k6/stats"
	"github.com/robertkrimen/otto"
	"github.com/spf13/afero"
	"path/filepath"
	"strings"
)

type InitAPI struct ***REMOVED***
	r  *Runtime
	fs afero.Fs

	fileCache map[string]string
***REMOVED***

func (i *InitAPI) NewMetric(it int, name string, isTime bool) *stats.Metric ***REMOVED***
	t := stats.MetricType(it)
	vt := stats.Default
	if isTime ***REMOVED***
		vt = stats.Time
	***REMOVED***

	if m, ok := i.r.Metrics[name]; ok ***REMOVED***
		if m.Type != t ***REMOVED***
			throw(i.r.VM, errors.New(fmt.Sprintf("attempted to redeclare %s with a different type (%s != %s)", name, m.Type, t)))
			return nil
		***REMOVED***
		if m.Contains != vt ***REMOVED***
			throw(i.r.VM, errors.New(fmt.Sprintf("attempted to redeclare %s with a different kind of value (%s != %s)", name, m.Contains, vt)))
		***REMOVED***
		return m
	***REMOVED***

	m := stats.New(name, t, vt)
	i.r.Metrics[name] = m
	return m
***REMOVED***

func (i *InitAPI) Require(name string) otto.Value ***REMOVED***
	if !strings.HasPrefix(name, ".") ***REMOVED***
		exports, err := i.r.loadLib(name + ".js")
		if err != nil ***REMOVED***
			throw(i.r.VM, err)
		***REMOVED***
		return exports
	***REMOVED***

	exports, err := i.r.loadFile(name+".js", i.fs)
	if err != nil ***REMOVED***
		throw(i.r.VM, err)
	***REMOVED***
	return exports
***REMOVED***

func (i *InitAPI) Open(name string) string ***REMOVED***
	if i.fileCache == nil ***REMOVED***
		i.fileCache = make(map[string]string)
	***REMOVED***

	path, err := filepath.Abs(name)
	if err != nil ***REMOVED***
		throw(i.r.VM, err)
	***REMOVED***

	if data, ok := i.fileCache[path]; ok ***REMOVED***
		return data
	***REMOVED***

	data, err := afero.ReadFile(i.fs, path)
	if err != nil ***REMOVED***
		throw(i.r.VM, err)
	***REMOVED***

	s := string(data)
	i.fileCache[path] = s
	return s
***REMOVED***
