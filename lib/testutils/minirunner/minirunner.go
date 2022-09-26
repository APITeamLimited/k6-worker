package minirunner

import (
	"context"
	"io"

	"github.com/APITeamLimited/k6-worker/lib"
	"github.com/APITeamLimited/k6-worker/metrics"
)

// Ensure mock implementations conform to the interfaces.
var (
	_ lib.Runner        = &MiniRunner***REMOVED******REMOVED***
	_ lib.InitializedVU = &VU***REMOVED******REMOVED***
	_ lib.ActiveVU      = &ActiveVU***REMOVED******REMOVED***
)

// MiniRunner partially implements the lib.Runner interface, but instead of
// using a real JS runtime, it allows us to directly specify the options and
// functions with Go code.
type MiniRunner struct ***REMOVED***
	Fn              func(ctx context.Context, state *lib.State, out chan<- metrics.SampleContainer) error
	SetupFn         func(ctx context.Context, out chan<- metrics.SampleContainer) ([]byte, error)
	TeardownFn      func(ctx context.Context, out chan<- metrics.SampleContainer) error
	HandleSummaryFn func(context.Context, *lib.Summary) (map[string]io.Reader, error)

	SetupData []byte

	Group   *lib.Group
	Options lib.Options
***REMOVED***

// MakeArchive isn't implemented, it always returns nil and is just here to
// satisfy the lib.Runner interface.
func (r MiniRunner) MakeArchive() *lib.Archive ***REMOVED***
	return nil
***REMOVED***

// NewVU returns a new VU with an incremental ID.
func (r *MiniRunner) NewVU(idLocal, idGlobal uint64, out chan<- metrics.SampleContainer, workerInfo *lib.WorkerInfo) (lib.InitializedVU, error) ***REMOVED***
	state := &lib.State***REMOVED***VUID: idLocal, VUIDGlobal: idGlobal, Iteration: int64(-1)***REMOVED***
	return &VU***REMOVED***
		R:            r,
		Out:          out,
		ID:           idLocal,
		IDGlobal:     idGlobal,
		state:        state,
		scenarioIter: make(map[string]uint64),
	***REMOVED***, nil
***REMOVED***

// Setup calls the supplied mock setup() function, if present.
func (r *MiniRunner) Setup(ctx context.Context, out chan<- metrics.SampleContainer) (err error) ***REMOVED***
	if fn := r.SetupFn; fn != nil ***REMOVED***
		r.SetupData, err = fn(ctx, out)
	***REMOVED***
	return
***REMOVED***

// GetSetupData returns json representation of the setup data if setup() is
// specified and was ran, nil otherwise.
func (r MiniRunner) GetSetupData() []byte ***REMOVED***
	return r.SetupData
***REMOVED***

// SetSetupData saves the externally supplied setup data as JSON in the runner.
func (r *MiniRunner) SetSetupData(data []byte) ***REMOVED***
	r.SetupData = data
***REMOVED***

// Teardown calls the supplied mock teardown() function, if present.
func (r MiniRunner) Teardown(ctx context.Context, out chan<- metrics.SampleContainer) error ***REMOVED***
	if fn := r.TeardownFn; fn != nil ***REMOVED***
		return fn(ctx, out)
	***REMOVED***
	return nil
***REMOVED***

// GetDefaultGroup returns the default group.
func (r MiniRunner) GetDefaultGroup() *lib.Group ***REMOVED***
	if r.Group == nil ***REMOVED***
		r.Group = &lib.Group***REMOVED******REMOVED***
	***REMOVED***
	return r.Group
***REMOVED***

// IsExecutable satisfies lib.Runner, but is mocked for MiniRunner since
// it doesn't deal with JS.
func (r MiniRunner) IsExecutable(name string) bool ***REMOVED***
	return true
***REMOVED***

// GetOptions returns the supplied options struct.
func (r MiniRunner) GetOptions() lib.Options ***REMOVED***
	return r.Options
***REMOVED***

// SetOptions allows you to override the runner options.
func (r *MiniRunner) SetOptions(opts lib.Options) error ***REMOVED***
	r.Options = opts
	return nil
***REMOVED***

// HandleSummary calls the specified summary callback, if supplied.
func (r *MiniRunner) HandleSummary(ctx context.Context, s *lib.Summary) (map[string]io.Reader, error) ***REMOVED***
	if r.HandleSummaryFn != nil ***REMOVED***
		return r.HandleSummaryFn(ctx, s)
	***REMOVED***
	return nil, nil
***REMOVED***

// VU is a mock VU, spawned by a MiniRunner.
type VU struct ***REMOVED***
	R            *MiniRunner
	Out          chan<- metrics.SampleContainer
	ID, IDGlobal uint64
	Iteration    int64
	state        *lib.State
	// count of iterations executed by this VU in each scenario
	scenarioIter map[string]uint64
***REMOVED***

// ActiveVU holds a VU and its activation parameters
type ActiveVU struct ***REMOVED***
	*VU
	*lib.VUActivationParams
	busy chan struct***REMOVED******REMOVED***

	scenarioName              string
	getNextIterations         func() (uint64, uint64)
	scIterLocal, scIterGlobal uint64
***REMOVED***

// GetID returns the unique VU ID.
func (vu *VU) GetID() uint64 ***REMOVED***
	return vu.ID
***REMOVED***

// State returns the VU's State.
func (vu *VU) State() *lib.State ***REMOVED***
	return vu.state
***REMOVED***

// Activate the VU so it will be able to run code.
func (vu *VU) Activate(params *lib.VUActivationParams) lib.ActiveVU ***REMOVED***
	ctx := params.RunContext

	vu.state.GetScenarioVUIter = func() uint64 ***REMOVED***
		return vu.scenarioIter[params.Scenario]
	***REMOVED***

	avu := &ActiveVU***REMOVED***
		VU:                 vu,
		VUActivationParams: params,
		busy:               make(chan struct***REMOVED******REMOVED***, 1),
		scenarioName:       params.Scenario,
		scIterLocal:        ^uint64(0),
		scIterGlobal:       ^uint64(0),
		getNextIterations:  params.GetNextIterationCounters,
	***REMOVED***

	vu.state.GetScenarioLocalVUIter = func() uint64 ***REMOVED***
		return avu.scIterLocal
	***REMOVED***
	vu.state.GetScenarioGlobalVUIter = func() uint64 ***REMOVED***
		return avu.scIterGlobal
	***REMOVED***

	go func() ***REMOVED***
		<-ctx.Done()

		// Wait for the VU to stop running, if it was, and prevent it from
		// running again for this activation
		avu.busy <- struct***REMOVED******REMOVED******REMOVED******REMOVED***

		if params.DeactivateCallback != nil ***REMOVED***
			params.DeactivateCallback(vu)
		***REMOVED***
	***REMOVED***()

	return avu
***REMOVED***

func (vu *ActiveVU) incrIteration() ***REMOVED***
	vu.Iteration++
	vu.state.Iteration = vu.Iteration

	if _, ok := vu.scenarioIter[vu.scenarioName]; ok ***REMOVED***
		vu.scenarioIter[vu.scenarioName]++
	***REMOVED*** else ***REMOVED***
		vu.scenarioIter[vu.scenarioName] = 0
	***REMOVED***
	vu.scIterLocal, vu.scIterGlobal = vu.getNextIterations()
***REMOVED***

// RunOnce runs the mock default function once, incrementing its iteration.
func (vu *ActiveVU) RunOnce() error ***REMOVED***
	if vu.R.Fn == nil ***REMOVED***
		return nil
	***REMOVED***

	select ***REMOVED***
	case <-vu.RunContext.Done():
		return vu.RunContext.Err() // we are done, return
	case vu.busy <- struct***REMOVED******REMOVED******REMOVED******REMOVED***:
		// nothing else can run now, and the VU cannot be deactivated
	***REMOVED***
	defer func() ***REMOVED***
		<-vu.busy // unlock deactivation again
	***REMOVED***()

	vu.incrIteration()
	return vu.R.Fn(vu.RunContext, vu.State(), vu.Out)
***REMOVED***
