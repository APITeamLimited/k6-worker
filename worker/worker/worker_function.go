package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/APITeamLimited/globe-test/lib"
	"github.com/APITeamLimited/globe-test/orchestrator/libOrch"
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/google/uuid"
)

func RunWorkerFunction(w http.ResponseWriter, r *http.Request, isDebug bool) {
	ctx := context.Background()
	workerId := uuid.NewString()
	client := getWorkerClient(true)

	creditsClient := lib.GetCreditsClient(true)

	// Ensure is POST request
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	// Get the childJob from the request body
	decoder := json.NewDecoder(r.Body)
	var childJob libOrch.ChildJob

	err := decoder.Decode(&childJob)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Printf("Worker %s executing child job %s\n", workerId, childJob.ChildJobId)

	successfullExecution := handleExecution(ctx, client, childJob, workerId, creditsClient, true)

	fmt.Printf("Worker %s finished executing child job %s with success: %t\n", workerId, childJob.ChildJobId, successfullExecution)

	// if !successfullExecution {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte("Error executing child job"))
	// 	return
	// }

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))

	if !isDebug {
		// Force the function to exit
		os.Exit(0)
	}
}

func RunDevWorkerServer() {
	devWorkerServerPort := lib.GetEnvVariableRaw("DEV_WORKER_FUNCTION_PORT", "8090", true)
	fmt.Printf("Starting dev worker function on port %s\n", devWorkerServerPort)
	os.Setenv("FUNCTION_TARGET", "WorkerCloud")

	runFunction := func(w http.ResponseWriter, r *http.Request) {
		RunWorkerFunction(w, r, true)
	}

	functions.HTTP("WorkerCloud", runFunction)

	if err := funcframework.Start(devWorkerServerPort); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
