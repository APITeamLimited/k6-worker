package mockoutput

import (
	"github.com/APITeamLimited/k6-worker/lib"
	"github.com/APITeamLimited/k6-worker/metrics"
	"github.com/APITeamLimited/k6-worker/output"
)

// New exists so that the usage from tests avoids repetition, i.e. is
// mockoutput.New() instead of &mockoutput.MockOutput***REMOVED******REMOVED***
func New() *MockOutput ***REMOVED***
	return &MockOutput***REMOVED******REMOVED***
***REMOVED***

// MockOutput can be used in tests to mock an actual output.
type MockOutput struct ***REMOVED***
	SampleContainers []metrics.SampleContainer
	Samples          []metrics.Sample
	RunStatus        lib.RunStatus

	DescFn  func() string
	StartFn func() error
	StopFn  func() error
***REMOVED***

var _ output.WithRunStatusUpdates = &MockOutput***REMOVED******REMOVED***

// AddMetricSamples just saves the results in memory.
func (mo *MockOutput) AddMetricSamples(scs []metrics.SampleContainer) ***REMOVED***
	mo.SampleContainers = append(mo.SampleContainers, scs...)
	for _, sc := range scs ***REMOVED***
		mo.Samples = append(mo.Samples, sc.GetSamples()...)
	***REMOVED***
***REMOVED***

// SetRunStatus updates the RunStatus property.
func (mo *MockOutput) SetRunStatus(latestStatus lib.RunStatus) ***REMOVED***
	mo.RunStatus = latestStatus
***REMOVED***

// Description calls the supplied DescFn callback, if available.
func (mo *MockOutput) Description() string ***REMOVED***
	if mo.DescFn != nil ***REMOVED***
		return mo.DescFn()
	***REMOVED***
	return "mock"
***REMOVED***

// Start calls the supplied StartFn callback, if available.
func (mo *MockOutput) Start() error ***REMOVED***
	if mo.StartFn != nil ***REMOVED***
		return mo.StartFn()
	***REMOVED***
	return nil
***REMOVED***

// Stop calls the supplied StopFn callback, if available.
func (mo *MockOutput) Stop() error ***REMOVED***
	if mo.StopFn != nil ***REMOVED***
		return mo.StopFn()
	***REMOVED***
	return nil
***REMOVED***
