/*
 *
 * k6 - a next-generation load testing tool
 * Copyright (C) 2017 Load Impact
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package har

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/loadimpact/k6/js"
	"github.com/loadimpact/k6/lib"
	"github.com/loadimpact/k6/loader"
	"github.com/stretchr/testify/assert"
)

func TestBuildK6Headers(t *testing.T) ***REMOVED***
	var headers = []struct ***REMOVED***
		values   []Header
		expected []string
	***REMOVED******REMOVED***
		***REMOVED***[]Header***REMOVED******REMOVED***"name", "1"***REMOVED***, ***REMOVED***"name", "2"***REMOVED******REMOVED***, []string***REMOVED***`"name": "1"`***REMOVED******REMOVED***,
		***REMOVED***[]Header***REMOVED******REMOVED***"name", "1"***REMOVED***, ***REMOVED***"name2", "2"***REMOVED******REMOVED***, []string***REMOVED***`"name": "1"`, `"name2": "2"`***REMOVED******REMOVED***,
		***REMOVED***[]Header***REMOVED******REMOVED***":host", "localhost"***REMOVED******REMOVED***, []string***REMOVED******REMOVED******REMOVED***,
	***REMOVED***

	for _, pair := range headers ***REMOVED***
		v := buildK6Headers(pair.values)
		assert.Equal(t, len(v), len(pair.expected), fmt.Sprintf("params: %v", pair.values))
	***REMOVED***
***REMOVED***

func TestBuildK6RequestObject(t *testing.T) ***REMOVED***
	req := &Request***REMOVED***
		Method:  "get",
		URL:     "http://www.google.es",
		Headers: []Header***REMOVED******REMOVED***"accept-language", "es-ES,es;q=0.8"***REMOVED******REMOVED***,
		Cookies: []Cookie***REMOVED******REMOVED***Name: "a", Value: "b"***REMOVED******REMOVED***,
	***REMOVED***
	v, err := buildK6RequestObject(req)
	assert.NoError(t, err)
	_, err = js.New(&loader.SourceData***REMOVED***
		URL:  &url.URL***REMOVED***Path: "/script.js"***REMOVED***,
		Data: []byte(fmt.Sprintf("export default function() ***REMOVED*** res = http.batch([%v]); ***REMOVED***", v)),
	***REMOVED***, nil, lib.RuntimeOptions***REMOVED******REMOVED***)
	assert.NoError(t, err)
***REMOVED***

func TestBuildK6Body(t *testing.T) ***REMOVED***

	bodyText := "ccustemail=ppcano%40gmail.com&size=medium&topping=cheese&delivery=12%3A00&comments="

	req := &Request***REMOVED***
		Method: "post",
		URL:    "http://www.google.es",
		PostData: &PostData***REMOVED***
			MimeType: "application/x-www-form-urlencoded",
			Text:     bodyText,
		***REMOVED***,
	***REMOVED***
	postParams, plainText, err := buildK6Body(req)
	assert.NoError(t, err)
	assert.Equal(t, len(postParams), 0, "postParams should be empty")
	assert.Equal(t, bodyText, plainText)

	email := "user@mail.es"
	expectedEmailParam := fmt.Sprintf(`"email": %q`, email)

	req = &Request***REMOVED***
		Method: "post",
		URL:    "http://www.google.es",
		PostData: &PostData***REMOVED***
			MimeType: "application/x-www-form-urlencoded",
			Params: []Param***REMOVED***
				***REMOVED***Name: "email", Value: url.QueryEscape(email)***REMOVED***,
				***REMOVED***Name: "pw", Value: "hola"***REMOVED***,
			***REMOVED***,
		***REMOVED***,
	***REMOVED***
	postParams, plainText, err = buildK6Body(req)
	assert.NoError(t, err)
	assert.Equal(t, plainText, "", "expected empty plainText")
	assert.Equal(t, len(postParams), 2, "postParams should have two items")
	assert.Equal(t, postParams[0], expectedEmailParam, "expected unescaped value")
***REMOVED***
