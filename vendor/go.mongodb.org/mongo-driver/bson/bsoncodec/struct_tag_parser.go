// Copyright (C) MongoDB, Inc. 2017-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package bsoncodec

import (
	"reflect"
	"strings"
)

// StructTagParser returns the struct tags for a given struct field.
type StructTagParser interface ***REMOVED***
	ParseStructTags(reflect.StructField) (StructTags, error)
***REMOVED***

// StructTagParserFunc is an adapter that allows a generic function to be used
// as a StructTagParser.
type StructTagParserFunc func(reflect.StructField) (StructTags, error)

// ParseStructTags implements the StructTagParser interface.
func (stpf StructTagParserFunc) ParseStructTags(sf reflect.StructField) (StructTags, error) ***REMOVED***
	return stpf(sf)
***REMOVED***

// StructTags represents the struct tag fields that the StructCodec uses during
// the encoding and decoding process.
//
// In the case of a struct, the lowercased field name is used as the key for each exported
// field but this behavior may be changed using a struct tag. The tag may also contain flags to
// adjust the marshalling behavior for the field.
//
// The properties are defined below:
//
//     OmitEmpty  Only include the field if it's not set to the zero value for the type or to
//                empty slices or maps.
//
//     MinSize    Marshal an integer of a type larger than 32 bits value as an int32, if that's
//                feasible while preserving the numeric value.
//
//     Truncate   When unmarshaling a BSON double, it is permitted to lose precision to fit within
//                a float32.
//
//     Inline     Inline the field, which must be a struct or a map, causing all of its fields
//                or keys to be processed as if they were part of the outer struct. For maps,
//                keys must not conflict with the bson keys of other struct fields.
//
//     Skip       This struct field should be skipped. This is usually denoted by parsing a "-"
//                for the name.
//
// TODO(skriptble): Add tags for undefined as nil and for null as nil.
type StructTags struct ***REMOVED***
	Name      string
	OmitEmpty bool
	MinSize   bool
	Truncate  bool
	Inline    bool
	Skip      bool
***REMOVED***

// DefaultStructTagParser is the StructTagParser used by the StructCodec by default.
// It will handle the bson struct tag. See the documentation for StructTags to see
// what each of the returned fields means.
//
// If there is no name in the struct tag fields, the struct field name is lowercased.
// The tag formats accepted are:
//
//     "[<key>][,<flag1>[,<flag2>]]"
//
//     `(...) bson:"[<key>][,<flag1>[,<flag2>]]" (...)`
//
// An example:
//
//     type T struct ***REMOVED***
//         A bool
//         B int    "myb"
//         C string "myc,omitempty"
//         D string `bson:",omitempty" json:"jsonkey"`
//         E int64  ",minsize"
//         F int64  "myf,omitempty,minsize"
//     ***REMOVED***
//
// A struct tag either consisting entirely of '-' or with a bson key with a
// value consisting entirely of '-' will return a StructTags with Skip true and
// the remaining fields will be their default values.
var DefaultStructTagParser StructTagParserFunc = func(sf reflect.StructField) (StructTags, error) ***REMOVED***
	key := strings.ToLower(sf.Name)
	tag, ok := sf.Tag.Lookup("bson")
	if !ok && !strings.Contains(string(sf.Tag), ":") && len(sf.Tag) > 0 ***REMOVED***
		tag = string(sf.Tag)
	***REMOVED***
	return parseTags(key, tag)
***REMOVED***

func parseTags(key string, tag string) (StructTags, error) ***REMOVED***
	var st StructTags
	if tag == "-" ***REMOVED***
		st.Skip = true
		return st, nil
	***REMOVED***

	for idx, str := range strings.Split(tag, ",") ***REMOVED***
		if idx == 0 && str != "" ***REMOVED***
			key = str
		***REMOVED***
		switch str ***REMOVED***
		case "omitempty":
			st.OmitEmpty = true
		case "minsize":
			st.MinSize = true
		case "truncate":
			st.Truncate = true
		case "inline":
			st.Inline = true
		***REMOVED***
	***REMOVED***

	st.Name = key

	return st, nil
***REMOVED***

// JSONFallbackStructTagParser has the same behavior as DefaultStructTagParser
// but will also fallback to parsing the json tag instead on a field where the
// bson tag isn't available.
var JSONFallbackStructTagParser StructTagParserFunc = func(sf reflect.StructField) (StructTags, error) ***REMOVED***
	key := strings.ToLower(sf.Name)
	tag, ok := sf.Tag.Lookup("bson")
	if !ok ***REMOVED***
		tag, ok = sf.Tag.Lookup("json")
	***REMOVED***
	if !ok && !strings.Contains(string(sf.Tag), ":") && len(sf.Tag) > 0 ***REMOVED***
		tag = string(sf.Tag)
	***REMOVED***

	return parseTags(key, tag)
***REMOVED***
