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
	"fmt"
	"math"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/sirupsen/logrus"
	null "gopkg.in/guregu/null.v3"

	"github.com/loadimpact/k6/lib"
	"github.com/loadimpact/k6/lib/types"
	"github.com/loadimpact/k6/stats"
	"github.com/loadimpact/k6/ui/pb"
)

const constantArrivalRateType = "constant-arrival-rate"

func init() ***REMOVED***
	lib.RegisterExecutorConfigType(
		constantArrivalRateType,
		func(name string, rawJSON []byte) (lib.ExecutorConfig, error) ***REMOVED***
			config := NewConstantArrivalRateConfig(name)
			err := lib.StrictJSONUnmarshal(rawJSON, &config)
			return config, err
		***REMOVED***,
	)
***REMOVED***

// ConstantArrivalRateConfig stores config for the constant arrival-rate executor
type ConstantArrivalRateConfig struct ***REMOVED***
	BaseConfig
	Rate     null.Int           `json:"rate"`
	TimeUnit types.NullDuration `json:"timeUnit"`
	Duration types.NullDuration `json:"duration"`

	// Initialize `PreAllocatedVUs` number of VUs, and if more than that are needed,
	// they will be dynamically allocated, until `MaxVUs` is reached, which is an
	// absolutely hard limit on the number of VUs the executor will use
	PreAllocatedVUs null.Int `json:"preAllocatedVUs"`
	MaxVUs          null.Int `json:"maxVUs"`
***REMOVED***

// NewConstantArrivalRateConfig returns a ConstantArrivalRateConfig with default values
func NewConstantArrivalRateConfig(name string) ConstantArrivalRateConfig ***REMOVED***
	return ConstantArrivalRateConfig***REMOVED***
		BaseConfig: NewBaseConfig(name, constantArrivalRateType),
		TimeUnit:   types.NewNullDuration(1*time.Second, false),
	***REMOVED***
***REMOVED***

// Make sure we implement the lib.ExecutorConfig interface
var _ lib.ExecutorConfig = &ConstantArrivalRateConfig***REMOVED******REMOVED***

// GetPreAllocatedVUs is just a helper method that returns the scaled pre-allocated VUs.
func (carc ConstantArrivalRateConfig) GetPreAllocatedVUs(et *lib.ExecutionTuple) int64 ***REMOVED***
	return et.ScaleInt64(carc.PreAllocatedVUs.Int64)
***REMOVED***

// GetMaxVUs is just a helper method that returns the scaled max VUs.
func (carc ConstantArrivalRateConfig) GetMaxVUs(et *lib.ExecutionTuple) int64 ***REMOVED***
	return et.ScaleInt64(carc.MaxVUs.Int64)
***REMOVED***

// GetDescription returns a human-readable description of the executor options
func (carc ConstantArrivalRateConfig) GetDescription(et *lib.ExecutionTuple) string ***REMOVED***
	preAllocatedVUs, maxVUs := carc.GetPreAllocatedVUs(et), carc.GetMaxVUs(et)
	maxVUsRange := fmt.Sprintf("maxVUs: %d", preAllocatedVUs)
	if maxVUs > preAllocatedVUs ***REMOVED***
		maxVUsRange += fmt.Sprintf("-%d", maxVUs)
	***REMOVED***

	timeUnit := time.Duration(carc.TimeUnit.Duration)
	var arrRatePerSec float64
	if maxVUs != 0 ***REMOVED*** // TODO: do something better?
		ratio := big.NewRat(maxVUs, carc.MaxVUs.Int64)
		arrRate := big.NewRat(carc.Rate.Int64, int64(timeUnit))
		arrRate.Mul(arrRate, ratio)
		arrRatePerSec, _ = getArrivalRatePerSec(arrRate).Float64()
	***REMOVED***

	return fmt.Sprintf("%.2f iterations/s for %s%s", arrRatePerSec, carc.Duration.Duration,
		carc.getBaseInfo(maxVUsRange))
***REMOVED***

// Validate makes sure all options are configured and valid
func (carc ConstantArrivalRateConfig) Validate() []error ***REMOVED***
	errors := carc.BaseConfig.Validate()
	if !carc.Rate.Valid ***REMOVED***
		errors = append(errors, fmt.Errorf("the iteration rate isn't specified"))
	***REMOVED*** else if carc.Rate.Int64 <= 0 ***REMOVED***
		errors = append(errors, fmt.Errorf("the iteration rate should be more than 0"))
	***REMOVED***

	if time.Duration(carc.TimeUnit.Duration) <= 0 ***REMOVED***
		errors = append(errors, fmt.Errorf("the timeUnit should be more than 0"))
	***REMOVED***

	if !carc.Duration.Valid ***REMOVED***
		errors = append(errors, fmt.Errorf("the duration is unspecified"))
	***REMOVED*** else if time.Duration(carc.Duration.Duration) < minDuration ***REMOVED***
		errors = append(errors, fmt.Errorf(
			"the duration should be at least %s, but is %s", minDuration, carc.Duration,
		))
	***REMOVED***

	if !carc.PreAllocatedVUs.Valid ***REMOVED***
		errors = append(errors, fmt.Errorf("the number of preAllocatedVUs isn't specified"))
	***REMOVED*** else if carc.PreAllocatedVUs.Int64 < 0 ***REMOVED***
		errors = append(errors, fmt.Errorf("the number of preAllocatedVUs shouldn't be negative"))
	***REMOVED***

	if !carc.MaxVUs.Valid ***REMOVED***
		errors = append(errors, fmt.Errorf("the number of maxVUs isn't specified"))
	***REMOVED*** else if carc.MaxVUs.Int64 < carc.PreAllocatedVUs.Int64 ***REMOVED***
		errors = append(errors, fmt.Errorf("maxVUs shouldn't be less than preAllocatedVUs"))
	***REMOVED***

	return errors
***REMOVED***

// GetExecutionRequirements returns the number of required VUs to run the
// executor for its whole duration (disregarding any startTime), including the
// maximum waiting time for any iterations to gracefully stop. This is used by
// the execution scheduler in its VU reservation calculations, so it knows how
// many VUs to pre-initialize.
func (carc ConstantArrivalRateConfig) GetExecutionRequirements(et *lib.ExecutionTuple) []lib.ExecutionStep ***REMOVED***
	return []lib.ExecutionStep***REMOVED***
		***REMOVED***
			TimeOffset:      0,
			PlannedVUs:      uint64(et.ScaleInt64(carc.PreAllocatedVUs.Int64)),
			MaxUnplannedVUs: uint64(et.ScaleInt64(carc.MaxVUs.Int64) - et.ScaleInt64(carc.PreAllocatedVUs.Int64)),
		***REMOVED***, ***REMOVED***
			TimeOffset:      time.Duration(carc.Duration.Duration + carc.GracefulStop.Duration),
			PlannedVUs:      0,
			MaxUnplannedVUs: 0,
		***REMOVED***,
	***REMOVED***
***REMOVED***

// NewExecutor creates a new ConstantArrivalRate executor
func (carc ConstantArrivalRateConfig) NewExecutor(
	es *lib.ExecutionState, logger *logrus.Entry,
) (lib.Executor, error) ***REMOVED***
	return &ConstantArrivalRate***REMOVED***
		BaseExecutor: NewBaseExecutor(carc, es, logger),
		config:       carc,
	***REMOVED***, nil
***REMOVED***

// HasWork reports whether there is any work to be done for the given execution segment.
func (carc ConstantArrivalRateConfig) HasWork(et *lib.ExecutionTuple) bool ***REMOVED***
	return carc.GetMaxVUs(et) > 0
***REMOVED***

// ConstantArrivalRate tries to execute a specific number of iterations for a
// specific period.
type ConstantArrivalRate struct ***REMOVED***
	*BaseExecutor
	config ConstantArrivalRateConfig
	et     *lib.ExecutionTuple
***REMOVED***

// Make sure we implement the lib.Executor interface.
var _ lib.Executor = &ConstantArrivalRate***REMOVED******REMOVED***

// Init values needed for the execution
func (car *ConstantArrivalRate) Init(ctx context.Context) (err error) ***REMOVED***
	// err should always be nil, because Init() won't be called for executors
	// with no work, as determined by their config's HasWork() method.
	car.et, err = car.BaseExecutor.executionState.ExecutionTuple.GetNewExecutionTupleFromValue(car.config.MaxVUs.Int64)
	return
***REMOVED***

// Run executes a constant number of iterations per second.
//
// TODO: Reuse the variable arrival rate method?
func (car ConstantArrivalRate) Run(ctx context.Context, out chan<- stats.SampleContainer) (err error) ***REMOVED*** //nolint:funlen
	gracefulStop := car.config.GetGracefulStop()
	duration := time.Duration(car.config.Duration.Duration)
	preAllocatedVUs := car.config.GetPreAllocatedVUs(car.executionState.ExecutionTuple)
	maxVUs := car.config.GetMaxVUs(car.executionState.ExecutionTuple)
	// TODO: refactor and simplify
	arrivalRate := getScaledArrivalRate(car.et.Segment, car.config.Rate.Int64, time.Duration(car.config.TimeUnit.Duration))
	tickerPeriod := time.Duration(getTickerPeriod(arrivalRate).Duration)
	arrivalRatePerSec, _ := getArrivalRatePerSec(arrivalRate).Float64()

	// Make sure the log and the progress bar have accurate information
	car.logger.WithFields(logrus.Fields***REMOVED***
		"maxVUs": maxVUs, "preAllocatedVUs": preAllocatedVUs, "duration": duration,
		"tickerPeriod": tickerPeriod, "type": car.config.GetType(),
	***REMOVED***).Debug("Starting executor run...")

	// Pre-allocate the VUs local shared buffer
	activeVUs := make(chan lib.ActiveVU, maxVUs)
	activeVUsCount := uint64(0)

	activeVUsWg := &sync.WaitGroup***REMOVED******REMOVED***
	defer activeVUsWg.Wait()

	startTime, maxDurationCtx, regDurationCtx, cancel := getDurationContexts(ctx, duration, gracefulStop)
	defer cancel()

	// Make sure all VUs aren't executing iterations anymore, for the cancel()
	// above to deactivate them.
	defer func() ***REMOVED***
		// activeVUsCount is modified only in the loop below, which is done here
		for i := uint64(0); i < activeVUsCount; i++ ***REMOVED***
			<-activeVUs
		***REMOVED***
	***REMOVED***()

	activateVU := func(initVU lib.InitializedVU) lib.ActiveVU ***REMOVED***
		activeVUsWg.Add(1)
		activeVU := initVU.Activate(&lib.VUActivationParams***REMOVED***
			RunContext: maxDurationCtx,
			DeactivateCallback: func() ***REMOVED***
				car.executionState.ReturnVU(initVU, true)
				activeVUsWg.Done()
			***REMOVED***,
		***REMOVED***)
		car.executionState.ModCurrentlyActiveVUsCount(+1)
		atomic.AddUint64(&activeVUsCount, 1)
		return activeVU
	***REMOVED***

	// Get the pre-allocated VUs in the local buffer
	for i := int64(0); i < preAllocatedVUs; i++ ***REMOVED***
		initVU, err := car.executionState.GetPlannedVU(car.logger, false)
		if err != nil ***REMOVED***
			return err
		***REMOVED***
		activeVUs <- activateVU(initVU)
	***REMOVED***

	vusFmt := pb.GetFixedLengthIntFormat(maxVUs)
	progIters := fmt.Sprintf(
		pb.GetFixedLengthFloatFormat(arrivalRatePerSec, 0)+" iters/s", arrivalRatePerSec)
	progresFn := func() (float64, []string) ***REMOVED***
		spent := time.Since(startTime)
		currActiveVUs := atomic.LoadUint64(&activeVUsCount)
		vusInBuffer := uint64(len(activeVUs))
		progVUs := fmt.Sprintf(vusFmt+"/"+vusFmt+" VUs",
			currActiveVUs-vusInBuffer, currActiveVUs)

		right := []string***REMOVED***progVUs, duration.String(), progIters***REMOVED***

		if spent > duration ***REMOVED***
			return 1, right
		***REMOVED***

		spentDuration := pb.GetFixedLengthDuration(spent, duration)
		progDur := fmt.Sprintf("%s/%s", spentDuration, duration)
		right[1] = progDur

		return math.Min(1, float64(spent)/float64(duration)), right
	***REMOVED***
	car.progress.Modify(pb.WithProgress(progresFn))
	go trackProgress(ctx, maxDurationCtx, regDurationCtx, &car, progresFn)

	runIterationBasic := getIterationRunner(car.executionState, car.logger)
	runIteration := func(vu lib.ActiveVU) ***REMOVED***
		runIterationBasic(maxDurationCtx, vu)
		activeVUs <- vu
	***REMOVED***

	remainingUnplannedVUs := maxVUs - preAllocatedVUs
	start, offsets, _ := car.et.GetStripedOffsets()
	startTime = time.Now()
	timer := time.NewTimer(time.Hour * 24)
	// here the we need the not scaled one
	notScaledTickerPeriod := time.Duration(
		getTickerPeriod(
			big.NewRat(
				car.config.Rate.Int64,
				int64(time.Duration(car.config.TimeUnit.Duration)),
			)).Duration)

	for li, gi := 0, start; ; li, gi = li+1, gi+offsets[li%len(offsets)] ***REMOVED***
		t := notScaledTickerPeriod*time.Duration(gi) - time.Since(startTime)
		timer.Reset(t)
		select ***REMOVED***
		case <-timer.C:
			select ***REMOVED***
			case vu := <-activeVUs:
				// ideally, we get the VU from the buffer without any issues
				go runIteration(vu)
			default:
				if remainingUnplannedVUs == 0 ***REMOVED***
					// TODO: emit an error metric?
					car.logger.Warningf("Insufficient VUs, reached %d active VUs and cannot allocate more", maxVUs)
					break
				***REMOVED***
				initVU, err := car.executionState.GetUnplannedVU(maxDurationCtx, car.logger)
				if err != nil ***REMOVED***
					return err
				***REMOVED***
				remainingUnplannedVUs--
				go runIteration(activateVU(initVU))
			***REMOVED***
		case <-regDurationCtx.Done():
			return nil
		***REMOVED***
	***REMOVED***
***REMOVED***
