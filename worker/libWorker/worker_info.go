package libWorker

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/APITeamLimited/globe-test/lib"
	"github.com/gorilla/websocket"
)

type (
	Collection struct {
		Variables map[string]string
		Name      string
	}

	Environment struct {
		Variables map[string]string
		Name      string
	}

	KeyValueItem struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	WorkerInfo struct {
		Conn              *websocket.Conn
		JobId             string
		ChildJobId        string
		ScopeId           string
		OrchestratorId    string
		WorkerId          string
		Ctx               context.Context
		Environment       *Environment
		Collection        *Collection
		WorkerOptions     Options
		FinalRequest      map[string]interface{}
		UnderlyingRequest map[string]interface{}
		Gs                *BaseGlobalState
		VerifiedDomains   []string
		SubFraction       float64
		CreditsManager    *lib.CreditsManager
		Standalone        bool
	}

	MarkMessage struct {
		Mark    string                 `json:"mark"`
		Message map[string]interface{} `json:"message"`
	}
)

type Message struct {
	JobId       string    `json:"jobId"`
	ChildJobId  string    `json:"childJobId"`
	Time        time.Time `json:"time"`
	WorkerId    string    `json:"workerId"`
	Message     string    `json:"message"`
	MessageType string    `json:"messageType"`
}

type MessageQueue struct {
	Mutex sync.Mutex

	// The count of currently actively being sent messages
	QueueCount    int
	NewQueueCount chan int
}

func DispatchMessage(gs BaseGlobalState, message string, messageType string) {
	// go func() {
	// 	messageQueue := gs.MessageQueue()

	// 	isTerminal := messageType == "STATUS" && (message == "FAILURE" || message == "SUCCESS")
	// 	if !isTerminal {
	// 		messageQueue.Mutex.Lock()
	// 		messageQueue.QueueCount++
	// 		messageQueue.Mutex.Unlock()

	// 		gs.ConnWriteMutex().Lock()
	// 		gs.Conn().WriteJSON(formatMessage(gs, message, messageType))
	// 		gs.ConnWriteMutex().Unlock()

	// 		messageQueue.Mutex.Lock()
	// 		messageQueue.QueueCount--

	// 		// Must unlock the mutex before sending the new count to the channel
	// 		messageQueue.Mutex.Unlock()
	// 		messageQueue.NewQueueCount <- messageQueue.QueueCount

	// 		return
	// 	}

	// 	messageQueue.Mutex.Lock()
	// 	queueCount := messageQueue.QueueCount
	// 	messageQueue.Mutex.Unlock()

	// 	// If the message is terminal, we want to make sure that all messages are sent before we return
	// 	if queueCount > 0 {
	// 		for newCount := range messageQueue.NewQueueCount {
	// 			if newCount == 0 {
	// 				break
	// 			}
	// 		}
	// 	}

	// 	gs.ConnWriteMutex().Lock()
	// 	gs.Conn().WriteJSON(formatMessage(gs, message, messageType))
	// 	gs.ConnWriteMutex().Unlock()
	// }()

	// isTerminal := messageType == "STATUS" && (message == "FAILURE" || message == "SUCCESS")

	// if isTerminal {
	// 	time.Sleep(300 * time.Millisecond)
	// }

	serializedMessage, err := json.Marshal(formatMessage(gs, message, messageType))
	if err != nil {
		fmt.Println("Error serializing message: ", err.Error())
		return
	}

	gs.ConnWriteMutex().Lock()
	gs.Conn().WriteMessage(websocket.TextMessage, serializedMessage)
	gs.ConnWriteMutex().Unlock()
}

func formatMessage(gs BaseGlobalState, message string, messageType string) Message {
	return Message{
		JobId:       gs.JobId(),
		ChildJobId:  gs.ChildJobId(),
		Time:        time.Now(),
		WorkerId:    gs.WorkerId(),
		Message:     message,
		MessageType: messageType,
	}
}

func UpdateStatus(gs BaseGlobalState, status string) {
	if gs.GetWorkerStatus() != status {
		DispatchMessage(gs, status, "STATUS")
		gs.SetWorkerStatus(status)
	}
}

func HandleStringError(gs BaseGlobalState, errString string) {
	fmt.Println("HandleStringError: ", errString)
	DispatchMessage(gs, errString, "ERROR")
	UpdateStatus(gs, "FAILURE")
}

func HandleError(gs BaseGlobalState, err error) {
	fmt.Println("HandleError: ", err.Error())
	DispatchMessage(gs, err.Error(), "ERROR")
	UpdateStatus(gs, "FAILURE")
}
