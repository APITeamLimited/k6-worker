package worker

import (
	log "github.com/Sirupsen/logrus"
	"github.com/loadimpact/speedboat/comm"
)

// A Worker executes distributed tasks, communicating over a Master.
type Worker struct ***REMOVED***
	Connector  comm.Connector
	Processors []func(*Worker) comm.Processor
***REMOVED***

// Creates a new Worker, connecting to a master listening on the given in/out addresses.
func New(inAddr string, outAddr string) (w Worker, err error) ***REMOVED***
	w.Connector, err = comm.NewClientConnector(comm.WorkerTopic, inAddr, outAddr)
	if err != nil ***REMOVED***
		return w, err
	***REMOVED***

	return w, nil
***REMOVED***

// Runs the main loop for a worker.
func (w *Worker) Run() ***REMOVED***
	in, out := w.Connector.Run()
	pInstances := w.createProcessors()
	for msg := range in ***REMOVED***
		log.WithFields(log.Fields***REMOVED***
			"type":    msg.Type,
			"payload": string(msg.Payload),
		***REMOVED***).Debug("Worker Received")

		go func() ***REMOVED***
			for m := range comm.Process(pInstances, msg) ***REMOVED***
				out <- m
			***REMOVED***
		***REMOVED***()
	***REMOVED***
***REMOVED***

func (w *Worker) createProcessors() []comm.Processor ***REMOVED***
	pInstances := []comm.Processor***REMOVED******REMOVED***
	for _, fn := range w.Processors ***REMOVED***
		pInstances = append(pInstances, fn(w))
	***REMOVED***
	return pInstances
***REMOVED***
