package worker

import (
	"sync"

	"github.com/APITeamLimited/globe-test/orchestrator/libOrch"
	"github.com/APITeamLimited/globe-test/worker/libWorker"
)

type ExecutionList struct ***REMOVED***
	currentJobs      map[string]libOrch.ChildJob
	mutex            sync.Mutex
	maxJobs          int
	maxVUs           int64
	currentJobsCount int
	currentVUsCount  int64
***REMOVED***

func (executionList *ExecutionList) addJob(job libOrch.ChildJob) ***REMOVED***
	executionList.currentJobs[job.ChildJobId] = job

	executionList.currentJobsCount++
	executionList.currentVUsCount += job.Options.MaxPossibleVUs.ValueOrZero()
***REMOVED***

func (executionList *ExecutionList) removeJob(childJobId string) ***REMOVED***
	executionList.mutex.Lock()
	managedVUsFreed := executionList.currentJobs[childJobId].Options.MaxPossibleVUs.ValueOrZero()

	executionList.currentVUsCount -= managedVUsFreed
	executionList.currentJobsCount--

	delete(executionList.currentJobs, childJobId)

	executionList.mutex.Unlock()
***REMOVED***

// Checks if the exectutor has the physical capacity to execute this job, this does
// not concern whether the user has the required credits to execute the job.
func (executionList *ExecutionList) checkExecutionCapacity(options libWorker.Options) bool ***REMOVED***
	if executionList.maxJobs >= 0 && executionList.currentJobsCount >= executionList.maxJobs ***REMOVED***
		return false
	***REMOVED***

	// If more than max permissible managed VUs, return false
	if executionList.maxVUs >= 0 && executionList.currentVUsCount+options.MaxPossibleVUs.ValueOrZero() > executionList.maxVUs ***REMOVED***
		return false
	***REMOVED***

	return true
***REMOVED***
