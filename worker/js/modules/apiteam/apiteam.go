// Package apiteam implements the module imported as 'apiteam' from inside k6.
package apiteam

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/APITeamLimited/globe-test/worker/js/common"
	"github.com/APITeamLimited/globe-test/worker/js/modules"
	"github.com/APITeamLimited/globe-test/worker/libWorker"
)

var (
	// ErrGroupInInitContext is returned when group() are using in the init context.
	ErrGroupInInitContext = common.NewInitContextError("Using group() in the init context is not supported")

	// ErrCheckInInitContext is returned when check() are using in the init context.
	ErrCheckInInitContext = common.NewInitContextError("Using check() in the init context is not supported")
)

type (
	// RootModule is the global module instance that will create module
	// instances for each VU.
	RootModule struct***REMOVED******REMOVED***

	// APITeam represents an instance of the k6 module.
	APITeam struct ***REMOVED***
		vu modules.VU
	***REMOVED***
)

var (
	_ modules.Module   = &RootModule***REMOVED******REMOVED***
	_ modules.Instance = &APITeam***REMOVED******REMOVED***
)

// New returns a pointer to a new RootModule instance.
func New() *RootModule ***REMOVED***
	return &RootModule***REMOVED******REMOVED***
***REMOVED***

// NewModuleInstance implements the modules.Module interface to return
// a new instance for each VU.
func (*RootModule) NewModuleInstance(vu modules.VU) modules.Instance ***REMOVED***
	return &APITeam***REMOVED***vu: vu***REMOVED***
***REMOVED***

// Exports returns the exports of the apiteam module.
func (mi *APITeam) Exports() modules.Exports ***REMOVED***
	return modules.Exports***REMOVED***
		Named: map[string]interface***REMOVED******REMOVED******REMOVED***
			"context": mi.Context,
			"mark":    mi.Mark,
		***REMOVED***,
	***REMOVED***
***REMOVED***

// Info returns current info about the APITeam Execution Context.
func (mi *APITeam) Context() *libWorker.WorkerInfo ***REMOVED***
	workerInfo := mi.vu.InitEnv().WorkerInfo

	return workerInfo
***REMOVED***

type markMessage struct ***REMOVED***
	Mark    string      `json:"mark"`
	Message interface***REMOVED******REMOVED*** `json:"message"`
***REMOVED***

// Returns a marked value to the orchestrator
func (mi *APITeam) Mark(mark string, value interface***REMOVED******REMOVED***) error ***REMOVED***
	workerInfo := mi.vu.InitEnv().WorkerInfo

	// Ensure no ':' in the tag
	if strings.Contains(mark, ":") ***REMOVED***
		return fmt.Errorf("filename cannot contain ':'")
	***REMOVED***

	markMessage := markMessage***REMOVED***
		Mark:    mark,
		Message: value,
	***REMOVED***

	marshalled, err := json.Marshal(markMessage)
	if err != nil ***REMOVED***
		return err
	***REMOVED***

	libWorker.DispatchMessage(workerInfo.Ctx, workerInfo.Client, workerInfo.JobId, workerInfo.WorkerId, string(marshalled), "MARK")

	return nil
***REMOVED***
