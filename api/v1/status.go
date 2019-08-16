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

package v1

import (
	"github.com/loadimpact/k6/core"
	"gopkg.in/guregu/null.v3"
)

type Status struct ***REMOVED***
	Paused null.Bool `json:"paused" yaml:"paused"`
	VUs    null.Int  `json:"vus" yaml:"vus"`
	VUsMax null.Int  `json:"vus-max" yaml:"vus-max"`

	// Readonly.
	Running bool `json:"running" yaml:"running"`
	Tainted bool `json:"tainted" yaml:"tainted"`
***REMOVED***

func NewStatus(engine *core.Engine) Status ***REMOVED***
	executionState := engine.ExecutionScheduler.GetState()
	return Status***REMOVED***
		Running: executionState.HasStarted(),
		Paused:  null.BoolFrom(executionState.IsPaused()),
		VUs:     null.IntFrom(executionState.GetCurrentlyActiveVUsCount()),
		VUsMax:  null.IntFrom(executionState.GetInitializedVUsCount()),
		Tainted: engine.IsTainted(),
	***REMOVED***
***REMOVED***

func (s Status) GetName() string ***REMOVED***
	return "status"
***REMOVED***

func (s Status) GetID() string ***REMOVED***
	return "default"
***REMOVED***

func (s Status) SetID(id string) error ***REMOVED***
	return nil
***REMOVED***
