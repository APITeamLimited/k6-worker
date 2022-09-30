package orchestrator

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/APITeamLimited/redis/v9"
)

func (e *ExecutionList) addJob(job map[string]string) {
	e.mutex.Lock()
	e.currentJobs[job["id"]] = job
	e.mutex.Unlock()
}

func (e *ExecutionList) removeJob(jobId string) {
	e.mutex.Lock()
	delete(e.currentJobs, jobId)
	e.mutex.Unlock()
}

func fetchScope(ctx context.Context, scopesClient *redis.Client, scopeId string) (map[string]string, error) {
	scope, err := scopesClient.Get(ctx, fmt.Sprintf("scope__id:%s", scopeId)).Result()
	if err != nil {
		return nil, err
	}

	// Check scope not empty
	if len(scope) == 0 {
		return nil, fmt.Errorf("scope %s is empty", scopeId)
	}

	// Parse scope as map[string]string
	parsedScope := make(map[string]string)
	err = json.Unmarshal([]byte(scope), &parsedScope)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling scope %s", scopeId)
	}

	return parsedScope, nil
}
