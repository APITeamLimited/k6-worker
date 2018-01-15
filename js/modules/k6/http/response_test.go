/*
 *
 * k6 - a next-generation load testing tool
 * Copyright (C) 2016 Load Impact
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

package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/dop251/goja"
	"github.com/loadimpact/k6/js/common"
	"github.com/loadimpact/k6/lib"
	"github.com/loadimpact/k6/lib/netext"
	"github.com/oxtoacart/bpool"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	null "gopkg.in/guregu/null.v3"
)

const testGetFormHTML = `
<html>
<head>
	<title>This is the title</title>
</head>
<body>
	<form method="get" id="form1">
		<input name="input_with_value" type="text" value="value"/>
		<input name="input_without_value" type="text"/>
		<select name="select_one">
			<option value="not this option">no</option>
			<option value="yes this option" selected>yes</option>
		</select>
		<select name="select_multi" multiple>
			<option>option 1</option>
			<option selected>option 2</option>
			<option selected>option 3</option>
		</select>
		<textarea name="textarea" multiple>Lorem ipsum dolor sit amet</textarea>
	</form>
</body>
`

type TestGetFormHttpHandler struct***REMOVED******REMOVED***

func (h *TestGetFormHttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) ***REMOVED***
	var body []byte
	var err error
	if r.URL.RawQuery != "" ***REMOVED***
		body, err = json.Marshal(struct ***REMOVED***
			Query url.Values `json:"query"`
		***REMOVED******REMOVED***
			Query: r.URL.Query(),
		***REMOVED***)
		if err != nil ***REMOVED***
			body = []byte(`***REMOVED***"error": "failed serializing json"***REMOVED***`)
		***REMOVED***
		w.Header().Set("Content-Type", "application/json")
	***REMOVED*** else ***REMOVED***
		w.Header().Set("Content-Type", "text/html")
		body = []byte(testGetFormHTML)
	***REMOVED***
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(body)))
	w.WriteHeader(200)
	w.Write(body)
***REMOVED***

func TestResponse(t *testing.T) ***REMOVED***
	root, err := lib.NewGroup("", nil)
	assert.NoError(t, err)

	logger := log.New()
	logger.Level = log.DebugLevel
	logger.Out = ioutil.Discard

	rt := goja.New()
	rt.SetFieldNameMapper(common.FieldNameMapper***REMOVED******REMOVED***)
	state := &common.State***REMOVED***
		Options: lib.Options***REMOVED***
			MaxRedirects: null.IntFrom(10),
			UserAgent:    null.StringFrom("TestUserAgent"),
			Throw:        null.BoolFrom(true),
		***REMOVED***,
		Logger: logger,
		Group:  root,
		HTTPTransport: &http.Transport***REMOVED***
			DialContext: (netext.NewDialer(net.Dialer***REMOVED***
				Timeout:   10 * time.Second,
				KeepAlive: 60 * time.Second,
				DualStack: true,
			***REMOVED***)).DialContext,
		***REMOVED***,
		BPool: bpool.NewBufferPool(1),
	***REMOVED***

	ctx := new(context.Context)
	*ctx = context.Background()
	*ctx = common.WithState(*ctx, state)
	*ctx = common.WithRuntime(*ctx, rt)
	rt.Set("http", common.Bind(rt, New(), ctx))

	mux := http.NewServeMux()
	mux.Handle("/forms/get", &TestGetFormHttpHandler***REMOVED******REMOVED***)
	srv := httptest.NewServer(mux)
	defer srv.Close()

	rt.Set("httpTestServerURL", srv.URL)

	t.Run("Html", func(t *testing.T) ***REMOVED***
		state.Samples = nil
		_, err := common.RunString(rt, `
		let res = http.request("GET", "https://httpbin.org/html");
		if (res.status != 200) ***REMOVED*** throw new Error("wrong status: " + res.status); ***REMOVED***
		if (res.body.indexOf("Herman Melville - Moby-Dick") == -1) ***REMOVED*** throw new Error("wrong body: " + res.body); ***REMOVED***
		`)
		assert.NoError(t, err)
		assertRequestMetricsEmitted(t, state.Samples, "GET", "https://httpbin.org/html", "", 200, "")

		t.Run("html", func(t *testing.T) ***REMOVED***
			_, err := common.RunString(rt, `
			if (res.html().find("h1").text() != "Herman Melville - Moby-Dick") ***REMOVED*** throw new Error("wrong title: " + res.body); ***REMOVED***
			`)
			assert.NoError(t, err)

			t.Run("shorthand", func(t *testing.T) ***REMOVED***
				_, err := common.RunString(rt, `
				if (res.html("h1").text() != "Herman Melville - Moby-Dick") ***REMOVED*** throw new Error("wrong title: " + res.body); ***REMOVED***
				`)
				assert.NoError(t, err)
			***REMOVED***)

			t.Run("url", func(t *testing.T) ***REMOVED***
				_, err := common.RunString(rt, `
				if (res.html().url != "https://httpbin.org/html") ***REMOVED*** throw new Error("url incorrect: " + res.html().url); ***REMOVED***
				`)
				assert.NoError(t, err)
			***REMOVED***)
		***REMOVED***)

		t.Run("group", func(t *testing.T) ***REMOVED***
			g, err := root.Group("my group")
			if assert.NoError(t, err) ***REMOVED***
				old := state.Group
				state.Group = g
				defer func() ***REMOVED*** state.Group = old ***REMOVED***()
			***REMOVED***

			state.Samples = nil
			_, err = common.RunString(rt, `
			let res = http.request("GET", "https://httpbin.org/html");
			if (res.status != 200) ***REMOVED*** throw new Error("wrong status: " + res.status); ***REMOVED***
			if (res.body.indexOf("Herman Melville - Moby-Dick") == -1) ***REMOVED*** throw new Error("wrong body: " + res.body); ***REMOVED***
			`)
			assert.NoError(t, err)
			assertRequestMetricsEmitted(t, state.Samples, "GET", "https://httpbin.org/html", "", 200, "::my group")
		***REMOVED***)
	***REMOVED***)
	t.Run("Json", func(t *testing.T) ***REMOVED***
		state.Samples = nil
		_, err := common.RunString(rt, `
		let res = http.request("GET", "https://httpbin.org/get?a=1&b=2");
		if (res.status != 200) ***REMOVED*** throw new Error("wrong status: " + res.status); ***REMOVED***
		if (res.json().args.a != "1") ***REMOVED*** throw new Error("wrong ?a: " + res.json().args.a); ***REMOVED***
		if (res.json().args.b != "2") ***REMOVED*** throw new Error("wrong ?b: " + res.json().args.b); ***REMOVED***
		`)
		assert.NoError(t, err)
		assertRequestMetricsEmitted(t, state.Samples, "GET", "https://httpbin.org/get?a=1&b=2", "", 200, "")

		t.Run("Invalid", func(t *testing.T) ***REMOVED***
			_, err := common.RunString(rt, `http.request("GET", "https://httpbin.org/html").json();`)
			assert.EqualError(t, err, "GoError: invalid character '<' looking for beginning of value")
		***REMOVED***)
	***REMOVED***)

	t.Run("SubmitForm", func(t *testing.T) ***REMOVED***
		t.Run("withoutArgs", func(t *testing.T) ***REMOVED***
			state.Samples = nil
			_, err := common.RunString(rt, `
			let res = http.request("GET", "https://httpbin.org/forms/post");
			if (res.status != 200) ***REMOVED*** throw new Error("wrong status: " + res.status); ***REMOVED***
			res = res.submitForm()
			if (res.status != 200) ***REMOVED*** throw new Error("wrong status: " + res.status); ***REMOVED***
			let data = res.json().form
			if (data.custname !== "" ||
				data.extradata !== undefined || 
				data.comments !== "" || 
				data.custemail !== "" || 
				data.custtel !== "" || 
				data.delivery !== "" 
			) ***REMOVED*** throw new Error("incorrect body: " + JSON.stringify(data, null, 4) ); ***REMOVED***
		`)
			assert.NoError(t, err)
			assertRequestMetricsEmitted(t, state.Samples, "POST", "https://httpbin.org/post", "", 200, "")
		***REMOVED***)

		t.Run("withFields", func(t *testing.T) ***REMOVED***
			state.Samples = nil
			_, err := common.RunString(rt, `
			let res = http.request("GET", "https://httpbin.org/forms/post");
			if (res.status != 200) ***REMOVED*** throw new Error("wrong status: " + res.status); ***REMOVED***
			res = res.submitForm(***REMOVED*** fields: ***REMOVED*** custname: "test", extradata: "test2" ***REMOVED*** ***REMOVED***)
			if (res.status != 200) ***REMOVED*** throw new Error("wrong status: " + res.status); ***REMOVED***
			let data = res.json().form
			if (data.custname !== "test" ||
				data.extradata !== "test2" || 
				data.comments !== "" || 
				data.custemail !== "" || 
				data.custtel !== "" || 
				data.delivery !== "" 
			) ***REMOVED*** throw new Error("incorrect body: " + JSON.stringify(data, null, 4) ); ***REMOVED***
		`)
			assert.NoError(t, err)
			assertRequestMetricsEmitted(t, state.Samples, "POST", "https://httpbin.org/post", "", 200, "")
		***REMOVED***)

		t.Run("withRequestParams", func(t *testing.T) ***REMOVED***
			state.Samples = nil
			_, err := common.RunString(rt, `
			let res = http.request("GET", "https://httpbin.org/forms/post");
			if (res.status != 200) ***REMOVED*** throw new Error("wrong status: " + res.status); ***REMOVED***
			res = res.submitForm(***REMOVED*** params: ***REMOVED*** headers: ***REMOVED*** "My-Fancy-Header": "SomeValue" ***REMOVED*** ***REMOVED******REMOVED***)
			if (res.status != 200) ***REMOVED*** throw new Error("wrong status: " + res.status); ***REMOVED***
			let headers = res.json().headers
			if (headers["My-Fancy-Header"] !== "SomeValue" ) ***REMOVED*** throw new Error("incorrect header: " + headers["My-Fancy-Header"]); ***REMOVED***
		`)
			assert.NoError(t, err)
			assertRequestMetricsEmitted(t, state.Samples, "POST", "https://httpbin.org/post", "", 200, "")
		***REMOVED***)

		t.Run("withFormSelector", func(t *testing.T) ***REMOVED***
			state.Samples = nil
			_, err := common.RunString(rt, `
			let res = http.request("GET", "https://httpbin.org/forms/post");
			if (res.status != 200) ***REMOVED*** throw new Error("wrong status: " + res.status); ***REMOVED***
			res = res.submitForm(***REMOVED*** formSelector: 'form[method="post"]' ***REMOVED***)
			if (res.status != 200) ***REMOVED*** throw new Error("wrong status: " + res.status); ***REMOVED***
			let data = res.json().form
			if (data.custname !== "" ||
				data.extradata !== undefined || 
				data.comments !== "" || 
				data.custemail !== "" || 
				data.custtel !== "" || 
				data.delivery !== "" 
			) ***REMOVED*** throw new Error("incorrect body: " + JSON.stringify(data, null, 4) ); ***REMOVED***
		`)
			assert.NoError(t, err)
			assertRequestMetricsEmitted(t, state.Samples, "POST", "https://httpbin.org/post", "", 200, "")
		***REMOVED***)

		t.Run("withNonExistentForm", func(t *testing.T) ***REMOVED***
			state.Samples = nil
			_, err := common.RunString(rt, `
			let res = http.request("GET", "https://httpbin.org/forms/post");
			if (res.status != 200) ***REMOVED*** throw new Error("wrong status: " + res.status); ***REMOVED***
			res.submitForm(***REMOVED*** formSelector: "#doesNotExist" ***REMOVED***)
		`)
			assert.EqualError(t, err, "GoError: no form found for selector '#doesNotExist' in response 'https://httpbin.org/forms/post'")
		***REMOVED***)

		t.Run("withGetMethod", func(t *testing.T) ***REMOVED***
			state.Samples = nil
			_, err := common.RunString(rt, `
			let res = http.request("GET", httpTestServerURL + "/forms/get");
			if (res.status != 200) ***REMOVED*** throw new Error("wrong status: " + res.status); ***REMOVED***
			res = res.submitForm()
			if (res.status != 200) ***REMOVED*** throw new Error("wrong status: " + res.status); ***REMOVED***
			let data = res.json().query
			if (data.input_with_value[0] !== "value" ||
				data.input_without_value[0] !== "" ||
				data.select_one[0] !== "yes this option" ||
				data.select_multi[0] !== "option 2,option 3" ||
				data.textarea[0] !== "Lorem ipsum dolor sit amet"
			) ***REMOVED*** throw new Error("incorrect body: " + JSON.stringify(data, null, 4) ); ***REMOVED***
		`)
			assert.NoError(t, err)
			assertRequestMetricsEmitted(t, state.Samples, "GET", srv.URL+"/forms/get", "", 200, "")
		***REMOVED***)
	***REMOVED***)

	t.Run("ClickLink", func(t *testing.T) ***REMOVED***
		t.Run("withoutArgs", func(t *testing.T) ***REMOVED***
			state.Samples = nil
			_, err := common.RunString(rt, `
			let res = http.request("GET", "https://httpbin.org/links/10/0");
			if (res.status != 200) ***REMOVED*** throw new Error("wrong status: " + res.status); ***REMOVED***
			res = res.clickLink()
			if (res.status != 200) ***REMOVED*** throw new Error("wrong status: " + res.status); ***REMOVED***
		`)
			assert.NoError(t, err)
			assertRequestMetricsEmitted(t, state.Samples, "GET", "https://httpbin.org/links/10/1", "", 200, "")
		***REMOVED***)

		t.Run("withSelector", func(t *testing.T) ***REMOVED***
			state.Samples = nil
			_, err := common.RunString(rt, `
			let res = http.request("GET", "https://httpbin.org/links/10/0");
			if (res.status != 200) ***REMOVED*** throw new Error("wrong status: " + res.status); ***REMOVED***
			res = res.clickLink(***REMOVED*** selector: 'a:nth-child(4)' ***REMOVED***)
			if (res.status != 200) ***REMOVED*** throw new Error("wrong status: " + res.status); ***REMOVED***
		`)
			assert.NoError(t, err)
			assertRequestMetricsEmitted(t, state.Samples, "GET", "https://httpbin.org/links/10/4", "", 200, "")
		***REMOVED***)

		t.Run("withNonExistentLink", func(t *testing.T) ***REMOVED***
			state.Samples = nil
			_, err := common.RunString(rt, `
			let res = http.request("GET", "https://httpbin.org/links/10/0");
			if (res.status != 200) ***REMOVED*** throw new Error("wrong status: " + res.status); ***REMOVED***
			res = res.clickLink(***REMOVED*** selector: 'a#doesNotExist' ***REMOVED***)
		`)
			assert.EqualError(t, err, "GoError: no element found for selector 'a#doesNotExist' in response 'https://httpbin.org/links/10/0'")
		***REMOVED***)

		t.Run("withRequestParams", func(t *testing.T) ***REMOVED***
			state.Samples = nil
			_, err := common.RunString(rt, `
			let res = http.request("GET", "https://httpbin.org");
			if (res.status != 200) ***REMOVED*** throw new Error("wrong status: " + res.status); ***REMOVED***
			res = res.clickLink(***REMOVED*** selector: 'a[href="/get"]', params: ***REMOVED*** headers: ***REMOVED*** "My-Fancy-Header": "SomeValue" ***REMOVED*** ***REMOVED*** ***REMOVED***)
			if (res.status != 200) ***REMOVED*** throw new Error("wrong status: " + res.status); ***REMOVED***
			let headers = res.json().headers
			if (headers["My-Fancy-Header"] !== "SomeValue" ) ***REMOVED*** throw new Error("incorrect header: " + headers["My-Fancy-Header"]); ***REMOVED***
		`)
			assert.NoError(t, err)
			assertRequestMetricsEmitted(t, state.Samples, "GET", "https://httpbin.org/get", "", 200, "")
		***REMOVED***)
	***REMOVED***)
***REMOVED***
