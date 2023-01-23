package function_auth_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	urlpkg "net/url"

	"cloud.google.com/go/functions/apiv2/functionspb"
	"github.com/APITeamLimited/globe-test/lib"
	"github.com/APITeamLimited/globe-test/orchestrator/libOrch"
	"google.golang.org/api/idtoken"
	"google.golang.org/api/option"
)

func (config *FunctionAuthClient) ExecuteFunction(location string, childJob libOrch.ChildJob) (*(chan libOrch.FunctionResult), error) {
	config.liveFunctionsMutex.Lock()
	defer config.liveFunctionsMutex.Unlock()

	var liveFunction *libOrch.LiveFunction

	// Find the function
	for _, function := range config.liveFunctions {
		if function.Location == location {
			liveFunction = &function
			break
		}
	}

	specificFuncOverride := getSpecificFuncOverride(config.funcUrlOverride, location)

	if liveFunction == nil && specificFuncOverride == "NO" {
		return nil, fmt.Errorf("function at location %s not found", location)
	}

	uri := determineUri(liveFunction, specificFuncOverride)

	client, err := getClient(config, uri, specificFuncOverride != "NO")
	if err != nil {
		return nil, err
	}

	request, err := buildRequest(childJob, uri)
	if err != nil {
		return nil, err
	}

	responseCh := make(chan libOrch.FunctionResult, 1)

	go func() {
		// Make the request
		response, err := client.Do(request)

		responseCh <- libOrch.FunctionResult{
			Response: response,
			Error:    err,
		}
	}()

	return &responseCh, nil
}

func getSpecificFuncOverride(overrides []string, location string) string {
	if len(overrides) == 1 && overrides[0] == "NO" {
		return "NO"
	}

	// Loop through WORKER_0_DISPLAY_NAME, WORKER_1_DISPLAY_NAME env variables till get location
	index := 0

	for {
		locationIndex := lib.GetEnvVariable(fmt.Sprintf("WORKER_%d_DISPLAY_NAME", index), "")

		if locationIndex == "" {
			break
		}

		if locationIndex == location {
			return overrides[index]
		}
	}

	panic(fmt.Sprintf("location %s not found in env variables", location))
}

func (config *FunctionAuthClient) CheckFunctionAvailability(location string) error {
	config.liveFunctionsMutex.Lock()
	defer config.liveFunctionsMutex.Unlock()

	var liveFunction *libOrch.LiveFunction

	// Find the function
	for _, function := range config.liveFunctions {
		if function.Location == location {
			liveFunction = &function
			break
		}
	}

	// If running locally, allow all functions
	if getSpecificFuncOverride(config.funcUrlOverride, location) != "NO" {
		return nil
	}

	if liveFunction == nil {
		return fmt.Errorf("location %s not found", location)
	}

	state := (*liveFunction).State

	// Could be updating, but still live
	if state != functionspb.Function_ACTIVE && state != functionspb.Function_DEPLOYING {
		return fmt.Errorf("location %s is not active", location)
	}

	return nil
}

func buildRequest(childJob libOrch.ChildJob, uri string) (*http.Request, error) {
	url, err := urlpkg.Parse(uri)
	if err != nil {
		return nil, err
	}

	// Encode the job in the body of the request
	var buff bytes.Buffer
	err = json.NewEncoder(&buff).Encode(childJob)
	if err != nil {
		return nil, err
	}

	request := &http.Request{
		Method: "POST",
		URL:    url,
		Body:   io.NopCloser(&buff),
	}

	return request, nil
}

func getClient(config *FunctionAuthClient, uri string, urlOverride bool) (*http.Client, error) {
	// In localhost development, don't need to authenticate
	if urlOverride {
		return http.DefaultClient, nil
	}

	// Authenticate the function
	return idtoken.NewClient(config.ctx, uri, option.WithCredentialsJSON(config.serviceAccount))
}

func determineUri(liveFunction *libOrch.LiveFunction, funcUrlOverride string) string {
	if funcUrlOverride == "NO" {
		return liveFunction.Uri
	}

	fmt.Println("Using function url override", funcUrlOverride)
	return funcUrlOverride
}
