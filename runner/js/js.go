package js

import (
	"github.com/loadimpact/speedboat/runner"
	"github.com/robertkrimen/otto"
	"net/http"
	"time"
)

type JSRunner struct ***REMOVED***
	BaseVM *otto.Otto
	Script *otto.Script
***REMOVED***

func New() (r *JSRunner, err error) ***REMOVED***
	r = &JSRunner***REMOVED******REMOVED***

	// Create a base VM
	r.BaseVM = otto.New()

	// Bridge basic functions
	r.BaseVM.Set("sleep", jsSleepFactory(time.Sleep))
	r.BaseVM.Set("get", jsHTTPGetFactory(r.BaseVM, func(url string) (*http.Response, error) ***REMOVED***
		client := &http.Client***REMOVED******REMOVED***
		return client.Get(url)
	***REMOVED***))

	return r, nil
***REMOVED***

func (r *JSRunner) Load(filename, src string) (err error) ***REMOVED***
	r.Script, err = r.BaseVM.Compile(filename, src)
	return err
***REMOVED***

func (r *JSRunner) RunVU(stop <-chan interface***REMOVED******REMOVED***) <-chan interface***REMOVED******REMOVED*** ***REMOVED***
	out := make(chan interface***REMOVED******REMOVED***)

	go func() ***REMOVED***
		defer close(out)

	runLoop:
		for ***REMOVED***
			select ***REMOVED***
			case <-stop:
				break runLoop
			default:
				vm := r.BaseVM.Copy()
				for res := range r.RunIteration(vm) ***REMOVED***
					out <- res
				***REMOVED***
			***REMOVED***
		***REMOVED***
	***REMOVED***()

	return out
***REMOVED***

func (r *JSRunner) RunIteration(vm *otto.Otto) <-chan interface***REMOVED******REMOVED*** ***REMOVED***
	out := make(chan interface***REMOVED******REMOVED***)

	go func() ***REMOVED***
		defer close(out)
		defer func() ***REMOVED***
			if err := recover(); err != nil ***REMOVED***
				out <- runner.NewError(err.(error))
			***REMOVED***
		***REMOVED***()

		// Log has to be bridged here, as it needs a reference to the channel
		vm.Set("log", jsLogFactory(func(text string) ***REMOVED***
			out <- runner.NewLogEntry(text)
		***REMOVED***))

		startTime := time.Now()
		_, err := vm.Run(r.Script)
		duration := time.Since(startTime)

		if err != nil ***REMOVED***
			out <- runner.NewError(err)
		***REMOVED***

		out <- runner.NewMetric(startTime, duration)
	***REMOVED***()

	return out
***REMOVED***
