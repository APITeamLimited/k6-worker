package stats

var DefaultRegistry = Registry***REMOVED******REMOVED***
var DefaultCollector = DefaultRegistry.NewCollector()

func NewCollector() *Collector ***REMOVED***
	return DefaultRegistry.NewCollector()
***REMOVED***

func Submit() error ***REMOVED***
	return DefaultRegistry.Submit()
***REMOVED***

func Add(s Sample) ***REMOVED***
	DefaultCollector.Add(s)
***REMOVED***
