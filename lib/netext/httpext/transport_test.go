package httpext

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/sirupsen/logrus"
	"go.k6.io/k6/lib"
	"go.k6.io/k6/metrics"
)

func BenchmarkMeasureAndEmitMetrics(b *testing.B) ***REMOVED***
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	samples := make(chan metrics.SampleContainer, 10)
	defer close(samples)
	go func() ***REMOVED***
		for range samples ***REMOVED***
		***REMOVED***
	***REMOVED***()
	logger := logrus.New()
	logger.Level = logrus.DebugLevel

	registry := metrics.NewRegistry()
	state := &lib.State***REMOVED***
		Options: lib.Options***REMOVED***
			SystemTags: &metrics.DefaultSystemTagSet,
		***REMOVED***,
		BuiltinMetrics: metrics.RegisterBuiltinMetrics(registry),
		Samples:        samples,
		Logger:         logger,
	***REMOVED***
	t := transport***REMOVED***
		state: state,
		ctx:   ctx,
	***REMOVED***

	unfRequest := &unfinishedRequest***REMOVED***
		tracer: &Tracer***REMOVED******REMOVED***,
		response: &http.Response***REMOVED***
			StatusCode: 200,
		***REMOVED***,
		request: &http.Request***REMOVED***
			URL: &url.URL***REMOVED***
				Host:   "example.com",
				Scheme: "https",
			***REMOVED***,
		***REMOVED***,
	***REMOVED***

	b.Run("no responseCallback", func(b *testing.B) ***REMOVED***
		for i := 0; i < b.N; i++ ***REMOVED***
			t.measureAndEmitMetrics(unfRequest)
		***REMOVED***
	***REMOVED***)

	t.responseCallback = func(n int) bool ***REMOVED*** return true ***REMOVED***

	b.Run("responseCallback", func(b *testing.B) ***REMOVED***
		for i := 0; i < b.N; i++ ***REMOVED***
			t.measureAndEmitMetrics(unfRequest)
		***REMOVED***
	***REMOVED***)
***REMOVED***
