package libAgent

import (
	"github.com/APITeamLimited/globe-test/lib"
	"github.com/APITeamLimited/globe-test/orchestrator/libOrch"
)

const AgentPort = 59125

const OrchestratorRedisHost = "localhost"
const OrchestratorRedisPort = "59126"

const WorkerRedisHost = "localhost"
const WorkerRedisPort = "59127"

type (
	ClientLocalTestManagerMessage struct {
		Type string `json:"type"`
	}

	ClientNewJobMessage struct {
		Type    string      `json:"type"` // "newJob"
		Message libOrch.Job `json:"message"`
	}
	ClientAbortJobMessage struct {
		Type    string `json:"type"` // "abortJob"
		Message string `json:"message"`
	}
	ClientJobUpdateMessage struct {
		Type    string                   `json:"type"` // "jobUpdate"
		Message lib.WrappedJobUserUpdate `json:"message"`
	}
)

// Server relays some messages back when successful

type (
	ServerLocalTestManagerMessage struct {
		Type string `json:"type"`
	}

	ServerAbortJobMessage struct {
		Type    string `json:"type"` // "abortJob"
		Message string `json:"message"`
	}

	ServerNewJobMessage struct {
		Type    string      `json:"type"` // "newJob"
		Message libOrch.Job `json:"message"`
	}

	ServerGlobeTestMessage struct {
		Type    string `json:"type"` // "globeTestMessage"
		Message string `json:"message"`
	}

	ServerWrappedJobUserUpdate struct {
		Type    string                   `json:"type"` // "jobUpdate"
		Message lib.WrappedJobUserUpdate `json:"message"`
	}
)
