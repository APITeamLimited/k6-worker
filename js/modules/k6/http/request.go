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
	"bytes"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/dop251/goja"
	"gopkg.in/guregu/null.v3"

	"go.k6.io/k6/js/common"
	"go.k6.io/k6/lib/netext/httpext"
	"go.k6.io/k6/lib/types"
)

// ErrHTTPForbiddenInInitContext is used when a http requests was made in the init context
var ErrHTTPForbiddenInInitContext = common.NewInitContextError("Making http requests in the init context is not supported")

// ErrBatchForbiddenInInitContext is used when batch was made in the init context
var ErrBatchForbiddenInInitContext = common.NewInitContextError("Using batch in the init context is not supported")

func (c *Client) getMethodClosure(method string) func(url goja.Value, args ...goja.Value) (*Response, error) ***REMOVED***
	return func(url goja.Value, args ...goja.Value) (*Response, error) ***REMOVED***
		return c.Request(method, url, args...)
	***REMOVED***
***REMOVED***

// Request makes an http request of the provided `method` and returns a corresponding response by
// taking goja.Values as arguments
func (c *Client) Request(method string, url goja.Value, args ...goja.Value) (*Response, error) ***REMOVED***
	state := c.moduleInstance.vu.State()
	if state == nil ***REMOVED***
		return nil, ErrHTTPForbiddenInInitContext
	***REMOVED***

	var body interface***REMOVED******REMOVED***
	var params goja.Value

	if len(args) > 0 ***REMOVED***
		body = args[0].Export()
	***REMOVED***
	if len(args) > 1 ***REMOVED***
		params = args[1]
	***REMOVED***

	req, err := c.parseRequest(method, url, body, params)
	if err != nil ***REMOVED***
		if state.Options.Throw.Bool ***REMOVED***
			return nil, err
		***REMOVED***
		state.Logger.WithField("error", err).Warn("Request Failed")
		r := httpext.NewResponse()
		r.Error = err.Error()
		var k6e httpext.K6Error
		if errors.As(err, &k6e) ***REMOVED***
			r.ErrorCode = int(k6e.Code)
		***REMOVED***
		return &Response***REMOVED***Response: r, client: c***REMOVED***, nil
	***REMOVED***

	resp, err := httpext.MakeRequest(c.moduleInstance.vu.Context(), state, req)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	c.processResponse(resp, req.ResponseType)
	return c.responseFromHTTPext(resp), nil
***REMOVED***

// processResponse stores the body as an ArrayBuffer if indicated by
// respType. This is done here instead of in httpext.readResponseBody to avoid
// a reverse dependency on js/common or goja.
func (c *Client) processResponse(resp *httpext.Response, respType httpext.ResponseType) ***REMOVED***
	if respType == httpext.ResponseTypeBinary && resp.Body != nil ***REMOVED***
		resp.Body = c.moduleInstance.vu.Runtime().NewArrayBuffer(resp.Body.([]byte))
	***REMOVED***
***REMOVED***

func (c *Client) responseFromHTTPext(resp *httpext.Response) *Response ***REMOVED***
	return &Response***REMOVED***Response: resp, client: c***REMOVED***
***REMOVED***

// TODO: break this function up
//nolint: gocyclo, cyclop, funlen, gocognit
func (c *Client) parseRequest(
	method string, reqURL, body interface***REMOVED******REMOVED***, params goja.Value,
) (*httpext.ParsedHTTPRequest, error) ***REMOVED***
	rt := c.moduleInstance.vu.Runtime()
	state := c.moduleInstance.vu.State()
	if state == nil ***REMOVED***
		return nil, ErrHTTPForbiddenInInitContext
	***REMOVED***

	if urlJSValue, ok := reqURL.(goja.Value); ok ***REMOVED***
		reqURL = urlJSValue.Export()
	***REMOVED***
	u, err := httpext.ToURL(reqURL)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***

	result := &httpext.ParsedHTTPRequest***REMOVED***
		URL: &u,
		Req: &http.Request***REMOVED***
			Method: method,
			URL:    u.GetURL(),
			Header: make(http.Header),
		***REMOVED***,
		Timeout:          60 * time.Second,
		Throw:            state.Options.Throw.Bool,
		Redirects:        state.Options.MaxRedirects,
		Cookies:          make(map[string]*httpext.HTTPRequestCookie),
		Tags:             make(map[string]string),
		ResponseCallback: c.responseCallback,
	***REMOVED***

	if state.Options.DiscardResponseBodies.Bool ***REMOVED***
		result.ResponseType = httpext.ResponseTypeNone
	***REMOVED*** else ***REMOVED***
		result.ResponseType = httpext.ResponseTypeText
	***REMOVED***

	formatFormVal := func(v interface***REMOVED******REMOVED***) string ***REMOVED***
		// TODO: handle/warn about unsupported/nested values
		return fmt.Sprintf("%v", v)
	***REMOVED***

	handleObjectBody := func(data map[string]interface***REMOVED******REMOVED***) error ***REMOVED***
		if !requestContainsFile(data) ***REMOVED***
			bodyQuery := make(url.Values, len(data))
			for k, v := range data ***REMOVED***
				if arr, ok := v.([]interface***REMOVED******REMOVED***); ok ***REMOVED***
					for _, el := range arr ***REMOVED***
						bodyQuery.Add(k, formatFormVal(el))
					***REMOVED***
					continue
				***REMOVED***
				bodyQuery.Set(k, formatFormVal(v))
			***REMOVED***
			result.Body = bytes.NewBufferString(bodyQuery.Encode())
			result.Req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			return nil
		***REMOVED***

		// handling multipart request
		result.Body = &bytes.Buffer***REMOVED******REMOVED***
		mpw := multipart.NewWriter(result.Body)

		// For parameters of type common.FileData, created with open(file, "b"),
		// we write the file boundary to the body buffer.
		// Otherwise parameters are treated as standard form field.
		for k, v := range data ***REMOVED***
			switch ve := v.(type) ***REMOVED***
			case FileData:
				// writing our own part to handle receiving
				// different content-type than the default application/octet-stream
				h := make(textproto.MIMEHeader)
				escapedFilename := escapeQuotes(ve.Filename)
				h.Set("Content-Disposition",
					fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
						k, escapedFilename))
				h.Set("Content-Type", ve.ContentType)

				// this writer will be closed either by the next part or
				// the call to mpw.Close()
				fw, err := mpw.CreatePart(h)
				if err != nil ***REMOVED***
					return err
				***REMOVED***

				if _, err := fw.Write(ve.Data); err != nil ***REMOVED***
					return err
				***REMOVED***
			default:
				fw, err := mpw.CreateFormField(k)
				if err != nil ***REMOVED***
					return err
				***REMOVED***

				if _, err := fw.Write([]byte(formatFormVal(v))); err != nil ***REMOVED***
					return err
				***REMOVED***
			***REMOVED***
		***REMOVED***

		if err := mpw.Close(); err != nil ***REMOVED***
			return err
		***REMOVED***

		result.Req.Header.Set("Content-Type", mpw.FormDataContentType())
		return nil
	***REMOVED***

	if body != nil ***REMOVED***
		switch data := body.(type) ***REMOVED***
		case map[string]goja.Value:
			// TODO: fix forms submission and serialization in k6/html before fixing this..
			newData := map[string]interface***REMOVED******REMOVED******REMOVED******REMOVED***
			for k, v := range data ***REMOVED***
				newData[k] = v.Export()
			***REMOVED***
			if err := handleObjectBody(newData); err != nil ***REMOVED***
				return nil, err
			***REMOVED***
		case goja.ArrayBuffer:
			result.Body = bytes.NewBuffer(data.Bytes())
		case map[string]interface***REMOVED******REMOVED***:
			if err := handleObjectBody(data); err != nil ***REMOVED***
				return nil, err
			***REMOVED***
		case string:
			result.Body = bytes.NewBufferString(data)
		case []byte:
			result.Body = bytes.NewBuffer(data)
		default:
			return nil, fmt.Errorf("unknown request body type %T", body)
		***REMOVED***
	***REMOVED***

	result.Req.Header.Set("User-Agent", state.Options.UserAgent.String)

	if state.CookieJar != nil ***REMOVED***
		result.ActiveJar = state.CookieJar
	***REMOVED***

	// TODO: ditch goja.Value, reflections and Object and use a simple go map and type assertions?
	if params != nil && !goja.IsUndefined(params) && !goja.IsNull(params) ***REMOVED***
		params := params.ToObject(rt)
		for _, k := range params.Keys() ***REMOVED***
			switch k ***REMOVED***
			case "cookies":
				cookiesV := params.Get(k)
				if goja.IsUndefined(cookiesV) || goja.IsNull(cookiesV) ***REMOVED***
					continue
				***REMOVED***
				cookies := cookiesV.ToObject(rt)
				if cookies == nil ***REMOVED***
					continue
				***REMOVED***
				for _, key := range cookies.Keys() ***REMOVED***
					cookieV := cookies.Get(key)
					if goja.IsUndefined(cookieV) || goja.IsNull(cookieV) ***REMOVED***
						continue
					***REMOVED***
					switch cookieV.ExportType() ***REMOVED***
					case reflect.TypeOf(map[string]interface***REMOVED******REMOVED******REMOVED******REMOVED***):
						result.Cookies[key] = &httpext.HTTPRequestCookie***REMOVED***Name: key, Value: "", Replace: false***REMOVED***
						cookie := cookieV.ToObject(rt)
						for _, attr := range cookie.Keys() ***REMOVED***
							switch strings.ToLower(attr) ***REMOVED***
							case "replace":
								result.Cookies[key].Replace = cookie.Get(attr).ToBoolean()
							case "value":
								result.Cookies[key].Value = cookie.Get(attr).String()
							***REMOVED***
						***REMOVED***
					default:
						result.Cookies[key] = &httpext.HTTPRequestCookie***REMOVED***Name: key, Value: cookieV.String(), Replace: false***REMOVED***
					***REMOVED***
				***REMOVED***
			case "headers":
				headersV := params.Get(k)
				if goja.IsUndefined(headersV) || goja.IsNull(headersV) ***REMOVED***
					continue
				***REMOVED***
				headers := headersV.ToObject(rt)
				if headers == nil ***REMOVED***
					continue
				***REMOVED***
				for _, key := range headers.Keys() ***REMOVED***
					str := headers.Get(key).String()
					if strings.ToLower(key) == "host" ***REMOVED***
						result.Req.Host = str
					***REMOVED***
					result.Req.Header.Set(key, str)
				***REMOVED***
			case "jar":
				jarV := params.Get(k)
				if goja.IsUndefined(jarV) || goja.IsNull(jarV) ***REMOVED***
					continue
				***REMOVED***
				switch v := jarV.Export().(type) ***REMOVED***
				case *CookieJar:
					result.ActiveJar = v.Jar
				***REMOVED***
			case "compression":
				algosString := strings.TrimSpace(params.Get(k).ToString().String())
				if algosString == "" ***REMOVED***
					continue
				***REMOVED***
				algos := strings.Split(algosString, ",")
				var err error
				result.Compressions = make([]httpext.CompressionType, len(algos))
				for index, algo := range algos ***REMOVED***
					algo = strings.TrimSpace(algo)
					result.Compressions[index], err = httpext.CompressionTypeString(algo)
					if err != nil ***REMOVED***
						return nil, fmt.Errorf("unknown compression algorithm %s, supported algorithms are %s",
							algo, httpext.CompressionTypeValues())
					***REMOVED***
				***REMOVED***
			case "redirects":
				result.Redirects = null.IntFrom(params.Get(k).ToInteger())
			case "tags":
				tagsV := params.Get(k)
				if goja.IsUndefined(tagsV) || goja.IsNull(tagsV) ***REMOVED***
					continue
				***REMOVED***
				tagObj := tagsV.ToObject(rt)
				if tagObj == nil ***REMOVED***
					continue
				***REMOVED***
				for _, key := range tagObj.Keys() ***REMOVED***
					result.Tags[key] = tagObj.Get(key).String()
				***REMOVED***
			case "auth":
				result.Auth = params.Get(k).String()
			case "timeout":
				t, err := types.GetDurationValue(params.Get(k).Export())
				if err != nil ***REMOVED***
					return nil, fmt.Errorf("invalid timeout value: %w", err)
				***REMOVED***
				result.Timeout = t
			case "throw":
				result.Throw = params.Get(k).ToBoolean()
			case "responseType":
				responseType, err := httpext.ResponseTypeString(params.Get(k).String())
				if err != nil ***REMOVED***
					return nil, err
				***REMOVED***
				result.ResponseType = responseType
			case "responseCallback":
				v := params.Get(k).Export()
				if v == nil ***REMOVED***
					result.ResponseCallback = nil
				***REMOVED*** else if c, ok := v.(*expectedStatuses); ok ***REMOVED***
					result.ResponseCallback = c.match
				***REMOVED*** else ***REMOVED***
					return nil, fmt.Errorf("unsupported responseCallback")
				***REMOVED***
			***REMOVED***
		***REMOVED***
	***REMOVED***

	if result.ActiveJar != nil ***REMOVED***
		httpext.SetRequestCookies(result.Req, result.ActiveJar, result.Cookies)
	***REMOVED***

	return result, nil
***REMOVED***

func (c *Client) prepareBatchArray(requests []interface***REMOVED******REMOVED***) (
	[]httpext.BatchParsedHTTPRequest, []*Response, error,
) ***REMOVED***
	reqCount := len(requests)
	batchReqs := make([]httpext.BatchParsedHTTPRequest, reqCount)
	results := make([]*Response, reqCount)

	for i, req := range requests ***REMOVED***
		resp := httpext.NewResponse()
		parsedReq, err := c.parseBatchRequest(i, req)
		if err != nil ***REMOVED***
			resp.Error = err.Error()
			var k6e httpext.K6Error
			if errors.As(err, &k6e) ***REMOVED***
				resp.ErrorCode = int(k6e.Code)
			***REMOVED***
			results[i] = c.responseFromHTTPext(resp)
			return batchReqs, results, err
		***REMOVED***
		batchReqs[i] = httpext.BatchParsedHTTPRequest***REMOVED***
			ParsedHTTPRequest: parsedReq,
			Response:          resp,
		***REMOVED***
		results[i] = c.responseFromHTTPext(resp)
	***REMOVED***

	return batchReqs, results, nil
***REMOVED***

func (c *Client) prepareBatchObject(requests map[string]interface***REMOVED******REMOVED***) (
	[]httpext.BatchParsedHTTPRequest, map[string]*Response, error,
) ***REMOVED***
	reqCount := len(requests)
	batchReqs := make([]httpext.BatchParsedHTTPRequest, reqCount)
	results := make(map[string]*Response, reqCount)

	i := 0
	for key, req := range requests ***REMOVED***
		resp := httpext.NewResponse()
		parsedReq, err := c.parseBatchRequest(key, req)
		if err != nil ***REMOVED***
			resp.Error = err.Error()
			var k6e httpext.K6Error
			if errors.As(err, &k6e) ***REMOVED***
				resp.ErrorCode = int(k6e.Code)
			***REMOVED***
			results[key] = c.responseFromHTTPext(resp)
			return batchReqs, results, err
		***REMOVED***
		batchReqs[i] = httpext.BatchParsedHTTPRequest***REMOVED***
			ParsedHTTPRequest: parsedReq,
			Response:          resp,
		***REMOVED***
		results[key] = c.responseFromHTTPext(resp)
		i++
	***REMOVED***

	return batchReqs, results, nil
***REMOVED***

// Batch makes multiple simultaneous HTTP requests. The provideds reqsV should be an array of request
// objects. Batch returns an array of responses and/or error
func (c *Client) Batch(reqsV ...goja.Value) (interface***REMOVED******REMOVED***, error) ***REMOVED***
	state := c.moduleInstance.vu.State()
	if state == nil ***REMOVED***
		return nil, ErrBatchForbiddenInInitContext
	***REMOVED***

	if len(reqsV) == 0 ***REMOVED***
		return nil, fmt.Errorf("no argument was provided to http.batch()")
	***REMOVED*** else if len(reqsV) > 1 ***REMOVED***
		return nil, fmt.Errorf("http.batch() accepts only an array or an object of requests")
	***REMOVED***
	var (
		err       error
		batchReqs []httpext.BatchParsedHTTPRequest
		results   interface***REMOVED******REMOVED*** // either []*Response or map[string]*Response
	)

	switch v := reqsV[0].Export().(type) ***REMOVED***
	case []interface***REMOVED******REMOVED***:
		batchReqs, results, err = c.prepareBatchArray(v)
	case map[string]interface***REMOVED******REMOVED***:
		batchReqs, results, err = c.prepareBatchObject(v)
	default:
		return nil, fmt.Errorf("invalid http.batch() argument type %T", v)
	***REMOVED***

	if err != nil ***REMOVED***
		if state.Options.Throw.Bool ***REMOVED***
			return nil, err
		***REMOVED***
		state.Logger.WithField("error", err).Warn("A batch request failed")
		return results, nil
	***REMOVED***

	reqCount := len(batchReqs)
	errs := httpext.MakeBatchRequests(
		c.moduleInstance.vu.Context(), state, batchReqs, reqCount,
		int(state.Options.Batch.Int64), int(state.Options.BatchPerHost.Int64),
		c.processResponse,
	)

	for i := 0; i < reqCount; i++ ***REMOVED***
		if e := <-errs; e != nil && err == nil ***REMOVED*** // Save only the first error
			err = e
		***REMOVED***
	***REMOVED***
	return results, err
***REMOVED***

func (c *Client) parseBatchRequest(key interface***REMOVED******REMOVED***, val interface***REMOVED******REMOVED***) (*httpext.ParsedHTTPRequest, error) ***REMOVED***
	var (
		method       = http.MethodGet
		ok           bool
		body, reqURL interface***REMOVED******REMOVED***
		params       goja.Value
		rt           = c.moduleInstance.vu.Runtime()
	)

	switch data := val.(type) ***REMOVED***
	case []interface***REMOVED******REMOVED***:
		// Handling of ["GET", "http://example.com/"]
		dataLen := len(data)
		if dataLen < 2 ***REMOVED***
			return nil, fmt.Errorf("invalid batch request '%#v'", data)
		***REMOVED***
		method, ok = data[0].(string)
		if !ok ***REMOVED***
			return nil, fmt.Errorf("invalid method type '%#v'", data[0])
		***REMOVED***
		reqURL = data[1]
		if dataLen > 2 ***REMOVED***
			body = data[2]
		***REMOVED***
		if dataLen > 3 ***REMOVED***
			params = rt.ToValue(data[3])
		***REMOVED***

	case map[string]interface***REMOVED******REMOVED***:
		// Handling of ***REMOVED***method: "GET", url: "https://test.k6.io"***REMOVED***
		if _, ok := data["url"]; !ok ***REMOVED***
			return nil, fmt.Errorf("batch request %v doesn't have a url key", key)
		***REMOVED***

		reqURL = data["url"]
		body = data["body"] // It's fine if it's missing, the map lookup will return

		if newMethod, ok := data["method"]; ok ***REMOVED***
			if method, ok = newMethod.(string); !ok ***REMOVED***
				return nil, fmt.Errorf("invalid method type '%#v'", newMethod)
			***REMOVED***
			method = strings.ToUpper(method)
			if method == http.MethodGet || method == http.MethodHead ***REMOVED***
				body = nil
			***REMOVED***
		***REMOVED***

		if p, ok := data["params"]; ok ***REMOVED***
			params = rt.ToValue(p)
		***REMOVED***
	default:
		reqURL = val
	***REMOVED***

	return c.parseRequest(method, reqURL, body, params)
***REMOVED***

func requestContainsFile(data map[string]interface***REMOVED******REMOVED***) bool ***REMOVED***
	for _, v := range data ***REMOVED***
		switch v.(type) ***REMOVED***
		case FileData:
			return true
		***REMOVED***
	***REMOVED***
	return false
***REMOVED***
