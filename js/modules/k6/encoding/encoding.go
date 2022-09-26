package encoding

import (
	"encoding/base64"

	"github.com/APITeamLimited/k6-worker/js/common"
	"github.com/APITeamLimited/k6-worker/js/modules"
)

type (
	// RootModule is the global module instance that will create module
	// instances for each VU.
	RootModule struct***REMOVED******REMOVED***

	// Encoding represents an instance of the encoding module.
	Encoding struct ***REMOVED***
		vu modules.VU
	***REMOVED***
)

var (
	_ modules.Module   = &RootModule***REMOVED******REMOVED***
	_ modules.Instance = &Encoding***REMOVED******REMOVED***
)

// New returns a pointer to a new RootModule instance.
func New() *RootModule ***REMOVED***
	return &RootModule***REMOVED******REMOVED***
***REMOVED***

// NewModuleInstance implements the modules.Module interface to return
// a new instance for each VU.
func (*RootModule) NewModuleInstance(vu modules.VU) modules.Instance ***REMOVED***
	return &Encoding***REMOVED***vu: vu***REMOVED***
***REMOVED***

// Exports returns the exports of the encoding module.
func (e *Encoding) Exports() modules.Exports ***REMOVED***
	return modules.Exports***REMOVED***
		Named: map[string]interface***REMOVED******REMOVED******REMOVED***
			"b64encode": e.b64Encode,
			"b64decode": e.b64Decode,
		***REMOVED***,
	***REMOVED***
***REMOVED***

// b64encode returns the base64 encoding of input as a string.
// The data type of input can be a string, []byte or ArrayBuffer.
func (e *Encoding) b64Encode(input interface***REMOVED******REMOVED***, encoding string) string ***REMOVED***
	data, err := common.ToBytes(input)
	if err != nil ***REMOVED***
		common.Throw(e.vu.Runtime(), err)
	***REMOVED***
	switch encoding ***REMOVED***
	case "rawstd":
		return base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
	case "std":
		return base64.StdEncoding.EncodeToString(data)
	case "rawurl":
		return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
	case "url":
		return base64.URLEncoding.EncodeToString(data)
	default:
		return base64.StdEncoding.EncodeToString(data)
	***REMOVED***
***REMOVED***

// b64decode returns the decoded data of the base64 encoded input string using
// the given encoding. If format is "s" it returns the data as a string,
// otherwise as an ArrayBuffer.
func (e *Encoding) b64Decode(input, encoding, format string) interface***REMOVED******REMOVED*** ***REMOVED***
	var (
		output []byte
		err    error
	)

	switch encoding ***REMOVED***
	case "rawstd":
		output, err = base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(input)
	case "std":
		output, err = base64.StdEncoding.DecodeString(input)
	case "rawurl":
		output, err = base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(input)
	case "url":
		output, err = base64.URLEncoding.DecodeString(input)
	default:
		output, err = base64.StdEncoding.DecodeString(input)
	***REMOVED***

	if err != nil ***REMOVED***
		common.Throw(e.vu.Runtime(), err)
	***REMOVED***

	var out interface***REMOVED******REMOVED***
	if format == "s" ***REMOVED***
		out = string(output)
	***REMOVED*** else ***REMOVED***
		ab := e.vu.Runtime().NewArrayBuffer(output)
		out = &ab
	***REMOVED***

	return out
***REMOVED***
