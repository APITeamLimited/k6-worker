package orchestrator

import (
	"encoding/json"
	"fmt"

	orchOptions "github.com/APITeamLimited/globe-test/orchestrator/options"
	"github.com/APITeamLimited/redis/v9"

	"github.com/APITeamLimited/globe-test/orchestrator/libOrch"
	"github.com/APITeamLimited/globe-test/worker/libWorker"
	"github.com/google/uuid"
)

const maxWorkerJobSize = 400

type jobDistribution struct {
	Jobs         []libOrch.ChildJob `json:"jobs"`
	workerClient *redis.Client
}

func determineChildJobs(healthy bool, job libOrch.Job, options *libWorker.Options,
	workerClients libOrch.WorkerClients) (map[string]jobDistribution, error) {
	// Don't run if not healthy
	if !healthy {
		return nil, nil
	}

	childJobs := make(map[string]jobDistribution)

	// Loop through options load distribution
	for _, loadZone := range options.LoadDistribution.Value {
		// Find worker client
		var workerClient *libOrch.NamedClient

		for _, client := range workerClients.Clients {
			if client.Name == loadZone.Location {
				workerClient = client
				break
			}
		}

		if workerClient == nil {
			return nil, fmt.Errorf("failed to find worker client %s, this is an internal error", loadZone.Location)
		}

		subFractions := determineSubFractions(loadZone.Fraction, job.Options.MaxPossibleVUs.Int64)

		zoneChildJobs := make([]libOrch.ChildJob, len(subFractions))

		jobNoOptions := job
		// Remove options from job
		jobNoOptions.Options = nil

		// Create child jobs
		for i, subFraction := range subFractions {
			// Need to deep copy job, json only way that seems to work
			childOptions, _ := json.Marshal(job.Options)
			parsed := libWorker.Options{}
			json.Unmarshal(childOptions, &parsed)

			zoneChildJobs[i] = libOrch.ChildJob{
				Job:               jobNoOptions,
				ChildJobId:        uuid.NewString(),
				Options:           orchOptions.DetermineChildDerivedOptions(loadZone, workerClient, parsed, subFraction),
				UnderlyingRequest: job.UnderlyingRequest,
				FinalRequest:      job.FinalRequest,
				SubFraction:       subFraction,
			}
		}

		childJobs[loadZone.Location] = jobDistribution{
			Jobs:         zoneChildJobs,
			workerClient: workerClient.Client,
		}
	}

	return childJobs, nil
}

func determineSubFractions(zoneFraction int, totalMaxVUs int64) []float64 {
	actualFraction := float64(zoneFraction) / 100
	zoneMaxVUsFloat := float64(totalMaxVUs) * actualFraction

	// Split into multiple jobs, each with a max of 500 vus and one job with the remainder

	childJobs := make([]float64, 0)

	for {
		if zoneMaxVUsFloat <= maxWorkerJobSize {
			childJobs = append(childJobs, zoneMaxVUsFloat)
			break
		}

		childJobs = append(childJobs, maxWorkerJobSize)
		zoneMaxVUsFloat -= maxWorkerJobSize
	}

	childSubFractions := make([]float64, len(childJobs))

	for i, childJob := range childJobs {
		childSubFractions[i] = childJob / float64(totalMaxVUs)
	}

	fmt.Println(childSubFractions)

	return childSubFractions
}
