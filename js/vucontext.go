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
	"github.com/loadimpact/k6/lib"
)

// Provides APIs and state for use in a VU context.
type VUContext struct ***REMOVED***
	// Console Object.
	Console *Console `js:"console"`

	// Environment variables.
	Env map[string]string `js:"__ENV"`
***REMOVED***

func NewVUContext(opts lib.Options) *VUContext ***REMOVED***
	return &VUContext***REMOVED***
		Console: NewConsole(),
		Env:     opts.Env,
	***REMOVED***
***REMOVED***
