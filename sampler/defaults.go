package sampler

var DefaultSampler = New()

func Get(name string) *Metric ***REMOVED***
	return DefaultSampler.Get(name)
***REMOVED***

func GetAs(name string, t int) *Metric ***REMOVED***
	return DefaultSampler.GetAs(name, t)
***REMOVED***

func Counter(name string) *Metric ***REMOVED***
	return DefaultSampler.Counter(name)
***REMOVED***

func Stats(name string) *Metric ***REMOVED***
	return DefaultSampler.Stats(name)
***REMOVED***
