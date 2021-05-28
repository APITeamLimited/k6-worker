/*
 *
 * k6 - a next-generation load testing tool
 * Copyright (C) 2019 Load Impact
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

package executor

import (
	"context"
	"strconv"
	"sync/atomic"

	"github.com/sirupsen/logrus"

	"go.k6.io/k6/lib"
	"go.k6.io/k6/stats"
	"go.k6.io/k6/ui/pb"
)

// BaseExecutor is a helper struct that contains common properties and methods
// between most executors. It is intended to be used as an anonymous struct
// inside of most of the executors, for the purpose of reducing boilerplate
// code.
type BaseExecutor struct ***REMOVED***
	config         lib.ExecutorConfig
	executionState *lib.ExecutionState
	localVUID      *uint64 // counter for assigning executor-specific VU IDs
	localIter      *int64  // counter for keeping track of all VU iterations completed by this executor
	logger         *logrus.Entry
	progress       *pb.ProgressBar
***REMOVED***

// NewBaseExecutor returns an initialized BaseExecutor
func NewBaseExecutor(config lib.ExecutorConfig, es *lib.ExecutionState, logger *logrus.Entry) *BaseExecutor ***REMOVED***
	// Start at -1 so that the first iteration can be 0
	startLocalIter := int64(-1)
	return &BaseExecutor***REMOVED***
		config:         config,
		executionState: es,
		localVUID:      new(uint64),
		localIter:      &startLocalIter,
		logger:         logger,
		progress: pb.New(
			pb.WithLeft(config.GetName),
			pb.WithLogger(logger),
		),
	***REMOVED***
***REMOVED***

// Init doesn't do anything for most executors, since initialization of all
// planned VUs is handled by the executor.
func (bs *BaseExecutor) Init(_ context.Context) error ***REMOVED***
	return nil
***REMOVED***

// GetConfig returns the configuration with which this executor was launched.
func (bs BaseExecutor) GetConfig() lib.ExecutorConfig ***REMOVED***
	return bs.config
***REMOVED***

// GetNextLocalVUID increments and returns the next VU ID that's specific for
// this executor (i.e. not global like __VU).
func (bs BaseExecutor) GetNextLocalVUID() uint64 ***REMOVED***
	return atomic.AddUint64(bs.localVUID, 1)
***REMOVED***

// GetLogger returns the executor logger entry.
func (bs BaseExecutor) GetLogger() *logrus.Entry ***REMOVED***
	return bs.logger
***REMOVED***

// GetProgress just returns the progressbar pointer.
func (bs BaseExecutor) GetProgress() *pb.ProgressBar ***REMOVED***
	return bs.progress
***REMOVED***

// getMetricTags returns a tag set that can be used to emit metrics by the
// executor. The VU ID is optional.
func (bs BaseExecutor) getMetricTags(vuID *uint64) *stats.SampleTags ***REMOVED***
	tags := bs.executionState.Options.RunTags.CloneTags()
	if bs.executionState.Options.SystemTags.Has(stats.TagScenario) ***REMOVED***
		tags["scenario"] = bs.config.GetName()
	***REMOVED***
	if vuID != nil && bs.executionState.Options.SystemTags.Has(stats.TagVU) ***REMOVED***
		tags["vu"] = strconv.FormatUint(*vuID, 10)
	***REMOVED***
	return stats.IntoSampleTags(&tags)
***REMOVED***

// incrScenarioIter increments the counter of completed iterations by all VUs
// for this executor.
func (bs *BaseExecutor) incrScenarioIter() int64 ***REMOVED***
	return atomic.AddInt64(bs.localIter, 1)
***REMOVED***
