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

package local

import (
	"context"
	"fmt"
	"runtime"
	"sync/atomic"
	"time"

	"github.com/loadimpact/k6/ui/pb"

	"github.com/loadimpact/k6/lib"
	"github.com/loadimpact/k6/stats"
	"github.com/sirupsen/logrus"
)

// ExecutionScheduler is the local implementation of lib.ExecutionScheduler
type ExecutionScheduler struct ***REMOVED***
	runner  lib.Runner
	options lib.Options
	logger  *logrus.Logger

	initProgress   *pb.ProgressBar
	executors      []lib.Executor // sorted by (startTime, ID)
	executionPlan  []lib.ExecutionStep
	maxDuration    time.Duration // cached value derived from the execution plan
	maxPossibleVUs uint64        // cached value derived from the execution plan
	state          *lib.ExecutionState
***REMOVED***

// Check to see if we implement the lib.ExecutionScheduler interface
var _ lib.ExecutionScheduler = &ExecutionScheduler***REMOVED******REMOVED***

// NewExecutionScheduler creates and returns a new local lib.ExecutionScheduler
// instance, without initializing it beyond the bare minimum. Specifically, it
// creates the needed executor instances and a lot of state placeholders, but it
// doesn't initialize the executors and it doesn't initialize or run VUs.
func NewExecutionScheduler(runner lib.Runner, logger *logrus.Logger) (*ExecutionScheduler, error) ***REMOVED***
	options := runner.GetOptions()

	executionPlan := options.Execution.GetFullExecutionRequirements(options.ExecutionSegment)
	maxPlannedVUs := lib.GetMaxPlannedVUs(executionPlan)
	maxPossibleVUs := lib.GetMaxPossibleVUs(executionPlan)

	executionState := lib.NewExecutionState(options, maxPlannedVUs, maxPossibleVUs)
	maxDuration, _ := lib.GetEndOffset(executionPlan) // we don't care if the end offset is final

	executorConfigs := options.Execution.GetSortedConfigs()
	executors := make([]lib.Executor, len(executorConfigs))
	for i, sc := range executorConfigs ***REMOVED***
		s, err := sc.NewExecutor(executionState, logger.WithField("executor", sc.GetName()))
		if err != nil ***REMOVED***
			return nil, err
		***REMOVED***
		executors[i] = s
	***REMOVED***

	if options.Paused.Bool ***REMOVED***
		if err := executionState.Pause(); err != nil ***REMOVED***
			return nil, err
		***REMOVED***
	***REMOVED***

	return &ExecutionScheduler***REMOVED***
		runner:  runner,
		logger:  logger,
		options: options,

		initProgress:   pb.New(pb.WithConstLeft("Init")),
		executors:      executors,
		executionPlan:  executionPlan,
		maxDuration:    maxDuration,
		maxPossibleVUs: maxPossibleVUs,
		state:          executionState,
	***REMOVED***, nil
***REMOVED***

// GetRunner returns the wrapped lib.Runner instance.
func (e *ExecutionScheduler) GetRunner() lib.Runner ***REMOVED***
	return e.runner
***REMOVED***

// GetState returns a pointer to the execution state struct for the local
// execution scheduler. It's guaranteed to be initialized and present, though
// see the documentation in lib/execution.go for caveats about its usage. The
// most important one is that none of the methods beyond the pause-related ones
// should be used for synchronization.
func (e *ExecutionScheduler) GetState() *lib.ExecutionState ***REMOVED***
	return e.state
***REMOVED***

// GetExecutors returns the slice of configured executor instances, sorted by
// their (startTime, name) in an ascending order.
func (e *ExecutionScheduler) GetExecutors() []lib.Executor ***REMOVED***
	return e.executors
***REMOVED***

// GetInitProgressBar returns a the progress bar assotiated with the Init
// function. After the Init is done, it is "hijacked" to display real-time
// execution statistics as a text bar.
func (e *ExecutionScheduler) GetInitProgressBar() *pb.ProgressBar ***REMOVED***
	return e.initProgress
***REMOVED***

// GetExecutionPlan is a helper method so users of the local execution scheduler
// don't have to calculate the execution plan again.
func (e *ExecutionScheduler) GetExecutionPlan() []lib.ExecutionStep ***REMOVED***
	return e.executionPlan
***REMOVED***

// initVU is just a helper method that's used to both initialize the planned VUs
// in the Init() method, and also passed to executors so they can initialize
// any unplanned VUs themselves.
//TODO: actually use the context...
func (e *ExecutionScheduler) initVU(
	_ context.Context, logger *logrus.Entry, engineOut chan<- stats.SampleContainer,
) (lib.VU, error) ***REMOVED***

	vu, err := e.runner.NewVU(engineOut)
	if err != nil ***REMOVED***
		return nil, fmt.Errorf("error while initializing a VU: '%s'", err)
	***REMOVED***

	// Get the VU ID here, so that the VUs are (mostly) ordered by their
	// number in the channel buffer
	vuID := e.state.GetUniqueVUIdentifier()
	if err := vu.Reconfigure(int64(vuID)); err != nil ***REMOVED***
		return nil, fmt.Errorf("error while reconfiguring VU #%d: '%s'", vuID, err)

	***REMOVED***
	logger.Debugf("Initialized VU #%d", vuID)
	return vu, nil
***REMOVED***

// getRunStats is a helper function that can be used as the execution
// scheduler's progressbar substitute (i.e. hijack).
func (e *ExecutionScheduler) getRunStats() string ***REMOVED***
	status := "running"
	if e.state.IsPaused() ***REMOVED***
		status = "paused"
	***REMOVED***
	if e.state.HasStarted() ***REMOVED***
		dur := e.state.GetCurrentTestRunDuration()
		status = fmt.Sprintf("%s (%s)", status, pb.GetFixedLengthDuration(dur, e.maxDuration))
	***REMOVED***

	vusFmt := pb.GetFixedLengthIntFormat(int64(e.maxPossibleVUs))
	return fmt.Sprintf(
		"%s, "+vusFmt+"/"+vusFmt+" VUs, %d complete and %d interrupted iterations",
		status, e.state.GetCurrentlyActiveVUsCount(), e.state.GetInitializedVUsCount(),
		e.state.GetFullIterationCount(), e.state.GetPartialIterationCount(),
	)
***REMOVED***

// Init concurrently initializes all of the planned VUs and then sequentially
// initializes all of the configured executors.
func (e *ExecutionScheduler) Init(ctx context.Context, engineOut chan<- stats.SampleContainer) error ***REMOVED***
	logger := e.logger.WithField("phase", "local-execution-scheduler-init")

	vusToInitialize := lib.GetMaxPlannedVUs(e.executionPlan)
	logger.WithFields(logrus.Fields***REMOVED***
		"neededVUs":      vusToInitialize,
		"executorsCount": len(e.executors),
	***REMOVED***).Debugf("Start of initialization")

	doneInits := make(chan error, vusToInitialize) // poor man's early-return waitgroup
	//TODO: make this an option?
	initConcurrency := runtime.NumCPU()
	limiter := make(chan struct***REMOVED******REMOVED***, initConcurrency)
	subctx, cancel := context.WithCancel(ctx)
	defer cancel()

	initPlannedVU := func() ***REMOVED***
		newVU, err := e.initVU(ctx, logger, engineOut)
		if err == nil ***REMOVED***
			e.state.AddInitializedVU(newVU)
			<-limiter
		***REMOVED***
		doneInits <- err
	***REMOVED***

	go func() ***REMOVED***
		for vuNum := uint64(0); vuNum < vusToInitialize; vuNum++ ***REMOVED***
			select ***REMOVED***
			case limiter <- struct***REMOVED******REMOVED******REMOVED******REMOVED***:
				go initPlannedVU()
			case <-subctx.Done():
				return
			***REMOVED***
		***REMOVED***
	***REMOVED***()

	initializedVUs := new(uint64)
	vusFmt := pb.GetFixedLengthIntFormat(int64(vusToInitialize))
	e.initProgress.Modify(
		pb.WithProgress(func() (float64, string) ***REMOVED***
			doneVUs := atomic.LoadUint64(initializedVUs)
			return float64(doneVUs) / float64(vusToInitialize),
				fmt.Sprintf(vusFmt+"/%d VUs initialized", doneVUs, vusToInitialize)
		***REMOVED***),
	)

	for vuNum := uint64(0); vuNum < vusToInitialize; vuNum++ ***REMOVED***
		select ***REMOVED***
		case err := <-doneInits:
			if err != nil ***REMOVED***
				return err
			***REMOVED***
			atomic.AddUint64(initializedVUs, 1)
		case <-ctx.Done():
			return ctx.Err()
		***REMOVED***
	***REMOVED***

	e.state.SetInitVUFunc(func(ctx context.Context, logger *logrus.Entry) (lib.VU, error) ***REMOVED***
		return e.initVU(ctx, logger, engineOut)
	***REMOVED***)

	logger.Debugf("Finished initializing needed VUs, start initializing executors...")
	for _, executor := range e.executors ***REMOVED***
		executorConfig := executor.GetConfig()

		if err := executor.Init(ctx); err != nil ***REMOVED***
			return fmt.Errorf("error while initializing executor %s: %s", executorConfig.GetName(), err)
		***REMOVED***
		logger.Debugf("Initialized executor %s", executorConfig.GetName())
	***REMOVED***

	logger.Debugf("Initization completed")
	return nil
***REMOVED***

// Run the ExecutionScheduler, funneling all generated metric samples through the supplied
// out channel.
func (e *ExecutionScheduler) Run(ctx context.Context, engineOut chan<- stats.SampleContainer) error ***REMOVED***
	executorsCount := len(e.executors)
	logger := e.logger.WithField("phase", "local-execution-scheduler-run")
	e.initProgress.Modify(pb.WithConstLeft("Run"))

	if e.state.IsPaused() ***REMOVED***
		logger.Debug("Execution is paused, waiting for resume or interrupt...")
		e.initProgress.Modify(pb.WithConstProgress(1, "paused"))
		select ***REMOVED***
		case <-e.state.ResumeNotify():
			// continue
		case <-ctx.Done():
			return nil
		***REMOVED***
	***REMOVED***

	e.state.MarkStarted()
	defer e.state.MarkEnded()
	e.initProgress.Modify(pb.WithConstProgress(1, "running"))

	logger.WithFields(logrus.Fields***REMOVED***"executorsCount": executorsCount***REMOVED***).Debugf("Start of test run")

	runResults := make(chan error, executorsCount) // nil values are successful runs

	runCtx, cancel := context.WithCancel(ctx)
	defer cancel() // just in case, and to shut up go vet...

	// Run setup() before any executors, if it's not disabled
	if !e.options.NoSetup.Bool ***REMOVED***
		logger.Debug("Running setup()")
		e.initProgress.Modify(pb.WithConstProgress(1, "setup()"))
		if err := e.runner.Setup(runCtx, engineOut); err != nil ***REMOVED***
			logger.WithField("error", err).Debug("setup() aborted by error")
			return err
		***REMOVED***
	***REMOVED***
	e.initProgress.Modify(pb.WithHijack(e.getRunStats))

	runCtxDone := runCtx.Done()
	runExecutor := func(executor lib.Executor) ***REMOVED***
		executorConfig := executor.GetConfig()
		executorStartTime := executorConfig.GetStartTime()
		executorLogger := logger.WithFields(logrus.Fields***REMOVED***
			"executor":  executorConfig.GetName(),
			"type":      executorConfig.GetType(),
			"startTime": executorStartTime,
		***REMOVED***)
		executorProgress := executor.GetProgress()

		// Check if we have to wait before starting the actual executor execution
		if executorStartTime > 0 ***REMOVED***
			startTime := time.Now()
			executorProgress.Modify(pb.WithProgress(func() (float64, string) ***REMOVED***
				remWait := (executorStartTime - time.Since(startTime))
				return 0, fmt.Sprintf("waiting %s", pb.GetFixedLengthDuration(remWait, executorStartTime))
			***REMOVED***))

			executorLogger.Debugf("Waiting for executor start time...")
			select ***REMOVED***
			case <-runCtxDone:
				runResults <- nil // no error since executor hasn't started yet
				return
			case <-time.After(executorStartTime):
				// continue
			***REMOVED***
		***REMOVED***

		executorProgress.Modify(pb.WithConstProgress(0, "started"))
		executorLogger.Debugf("Starting executor")
		err := executor.Run(runCtx, engineOut) // executor should handle context cancel itself
		if err == nil ***REMOVED***
			executorLogger.Debugf("Executor finished successfully")
		***REMOVED*** else ***REMOVED***
			executorLogger.WithField("error", err).Errorf("Executor error")
		***REMOVED***
		runResults <- err
	***REMOVED***

	// Start all executors at their particular startTime in a separate goroutine...
	logger.Debug("Start all executors...")
	for _, executor := range e.executors ***REMOVED***
		go runExecutor(executor)
	***REMOVED***

	// Wait for all executors to finish
	var firstErr error
	for range e.executors ***REMOVED***
		err := <-runResults
		if err != nil && firstErr == nil ***REMOVED***
			firstErr = err
			cancel()
		***REMOVED***
	***REMOVED***

	// Run teardown() after all executors are done, if it's not disabled
	if !e.options.NoTeardown.Bool ***REMOVED***
		logger.Debug("Running teardown()")
		if err := e.runner.Teardown(ctx, engineOut); err != nil ***REMOVED***
			logger.WithField("error", err).Debug("teardown() aborted by error")
			return err
		***REMOVED***
	***REMOVED***

	return firstErr
***REMOVED***

// SetPaused pauses a test, if called with true. And if called with false, tries
// to start/resume it. See the lib.ExecutionScheduler interface documentation of
// the methods for the various caveats about its usage.
func (e *ExecutionScheduler) SetPaused(pause bool) error ***REMOVED***
	if !e.state.HasStarted() && e.state.IsPaused() ***REMOVED***
		if pause ***REMOVED***
			return fmt.Errorf("execution is already paused")
		***REMOVED***
		e.logger.Debug("Starting execution")
		return e.state.Resume()
	***REMOVED***

	for _, executor := range e.executors ***REMOVED***
		pausableExecutor, ok := executor.(lib.PausableExecutor)
		if !ok ***REMOVED***
			return fmt.Errorf(
				"%s executor '%s' doesn't support pause and resume operations after its start",
				executor.GetConfig().GetType(), executor.GetConfig().GetName(),
			)
		***REMOVED***
		if err := pausableExecutor.SetPaused(pause); err != nil ***REMOVED***
			return err
		***REMOVED***
	***REMOVED***
	if pause ***REMOVED***
		return e.state.Pause()
	***REMOVED***
	return e.state.Resume()
***REMOVED***
