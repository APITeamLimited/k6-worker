/*
 *
 * k6 - a next-generation load testing tool
 * Copyright (C) 2019 Load Impact
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

package httpext

import (
	"context"
	"net"
	"net/http"
	"net/http/httptrace"
	"strconv"
	"sync"

	"github.com/loadimpact/k6/lib"
	"github.com/loadimpact/k6/lib/netext"
	"github.com/loadimpact/k6/stats"
)

// transport is an implemenation of http.RoundTripper that will measure and emit
// different metrics for each roundtrip
type transport struct ***REMOVED***
	state *lib.State
	tags  map[string]string

	lastRequest     *unfinishedRequest
	lastRequestLock *sync.Mutex
***REMOVED***

// unfinishedRequest stores the request and the raw result returned from the
// underlying http.RoundTripper, but before its body has been read
type unfinishedRequest struct ***REMOVED***
	ctx      context.Context
	tracer   *Tracer
	request  *http.Request
	response *http.Response
	err      error
***REMOVED***

// finishedRequest is produced once the request has been finalized; it is
// triggered either by a subsequent RoundTrip, or for the last request in the
// chain - by the MakeRequest function manually calling the transport method
// processLastSavedRequest(), after reading the HTTP response body.
type finishedRequest struct ***REMOVED***
	*unfinishedRequest
	trail     *Trail
	tlsInfo   netext.TLSInfo
	errorCode errCode
	errorMsg  string
***REMOVED***

var _ http.RoundTripper = &transport***REMOVED******REMOVED***

// NewTransport returns a new Transport wrapping around the provide Roundtripper and will send
// samples on the provided channel adding the tags in accordance to the Options provided
func newTransport(
	state *lib.State,
	tags map[string]string,
) *transport ***REMOVED***
	return &transport***REMOVED***
		state:           state,
		tags:            tags,
		lastRequestLock: new(sync.Mutex),
	***REMOVED***
***REMOVED***

// Helper method to finish the tracer trail, assemble the tag values and emits
// the metric samples for the supplied unfinished request.
func (t *transport) measureAndEmitMetrics(unfReq *unfinishedRequest) *finishedRequest ***REMOVED***
	trail := unfReq.tracer.Done()

	tags := map[string]string***REMOVED******REMOVED***
	for k, v := range t.tags ***REMOVED***
		tags[k] = v
	***REMOVED***

	result := &finishedRequest***REMOVED***
		unfinishedRequest: unfReq,
		trail:             trail,
	***REMOVED***

	enabledTags := t.state.Options.SystemTags
	if unfReq.err != nil ***REMOVED***
		result.errorCode, result.errorMsg = errorCodeForError(unfReq.err)
		if enabledTags["error"] ***REMOVED***
			tags["error"] = result.errorMsg
		***REMOVED***

		if enabledTags["error_code"] ***REMOVED***
			tags["error_code"] = strconv.Itoa(int(result.errorCode))
		***REMOVED***

		if enabledTags["status"] ***REMOVED***
			tags["status"] = "0"
		***REMOVED***
	***REMOVED*** else ***REMOVED***
		if enabledTags["url"] ***REMOVED***
			tags["url"] = unfReq.request.URL.String()
		***REMOVED***
		if enabledTags["status"] ***REMOVED***
			tags["status"] = strconv.Itoa(unfReq.response.StatusCode)
		***REMOVED***
		if unfReq.response.StatusCode >= 400 ***REMOVED***
			if enabledTags["error_code"] ***REMOVED***
				result.errorCode = errCode(1000 + unfReq.response.StatusCode)
				tags["error_code"] = strconv.Itoa(int(result.errorCode))
			***REMOVED***
		***REMOVED***
		if enabledTags["proto"] ***REMOVED***
			tags["proto"] = unfReq.response.Proto
		***REMOVED***

		if unfReq.response.TLS != nil ***REMOVED***
			tlsInfo, oscp := netext.ParseTLSConnState(unfReq.response.TLS)
			if enabledTags["tls_version"] ***REMOVED***
				tags["tls_version"] = tlsInfo.Version
			***REMOVED***
			if enabledTags["ocsp_status"] ***REMOVED***
				tags["ocsp_status"] = oscp.Status
			***REMOVED***
			result.tlsInfo = tlsInfo
		***REMOVED***
	***REMOVED***
	if enabledTags["ip"] && trail.ConnRemoteAddr != nil ***REMOVED***
		if ip, _, err := net.SplitHostPort(trail.ConnRemoteAddr.String()); err == nil ***REMOVED***
			tags["ip"] = ip
		***REMOVED***
	***REMOVED***

	trail.SaveSamples(stats.IntoSampleTags(&tags))
	stats.PushIfNotCancelled(unfReq.ctx, t.state.Samples, trail)

	return result
***REMOVED***

func (t *transport) saveCurrentRequest(currentRequest *unfinishedRequest) ***REMOVED***
	t.lastRequestLock.Lock()
	unprocessedRequest := t.lastRequest
	t.lastRequest = currentRequest
	t.lastRequestLock.Unlock()

	if unprocessedRequest != nil ***REMOVED***
		// This shouldn't happen, since we have one transport per request, but just in case...
		t.state.Logger.Warnf("TracerTransport: unexpected unprocessed request for %s", unprocessedRequest.request.URL)
		t.measureAndEmitMetrics(unprocessedRequest)
	***REMOVED***
***REMOVED***

func (t *transport) processLastSavedRequest(lastErr error) *finishedRequest ***REMOVED***
	t.lastRequestLock.Lock()
	unprocessedRequest := t.lastRequest
	t.lastRequest = nil
	t.lastRequestLock.Unlock()

	if unprocessedRequest != nil ***REMOVED***
		// We don't want to overwrite any previous errors, but if there were
		// none and we (i.e. the MakeRequest() function) have one, save it
		// before we emit the metrics.
		if unprocessedRequest.err == nil && lastErr != nil ***REMOVED***
			unprocessedRequest.err = lastErr
		***REMOVED***

		return t.measureAndEmitMetrics(unprocessedRequest)
	***REMOVED***
	return nil
***REMOVED***

// RoundTrip is the implementation of http.RoundTripper
func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) ***REMOVED***
	t.processLastSavedRequest(nil)

	ctx := req.Context()
	tracer := &Tracer***REMOVED******REMOVED***
	reqWithTracer := req.WithContext(httptrace.WithClientTrace(ctx, tracer.Trace()))
	resp, err := t.state.Transport.RoundTrip(reqWithTracer)

	t.saveCurrentRequest(&unfinishedRequest***REMOVED***
		ctx:      ctx,
		tracer:   tracer,
		request:  req,
		response: resp,
		err:      err,
	***REMOVED***)

	return resp, err
***REMOVED***
