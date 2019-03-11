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
	"net"
	"net/http"
	"net/http/httptrace"
	"strconv"

	"github.com/loadimpact/k6/lib"
	"github.com/loadimpact/k6/lib/netext"
	"github.com/loadimpact/k6/stats"
	"github.com/pkg/errors"
)

// transport is an implemenation of http.RoundTripper that will measure different metrics for each
// roundtrip
type transport struct ***REMOVED***
	roundTripper http.RoundTripper
	// TODO: maybe just take the SystemTags field as it is the only thing used
	options   *lib.Options
	tags      map[string]string
	trail     *Trail
	errorMsg  string
	errorCode errCode
	tlsInfo   netext.TLSInfo
	samplesCh chan<- stats.SampleContainer
***REMOVED***

var _ http.RoundTripper = &transport***REMOVED******REMOVED***

// NewTransport returns a new Transport wrapping around the provide Roundtripper and will send
// samples on the provided channel adding the tags in accordance to the Options provided
func newTransport(
	roundTripper http.RoundTripper,
	samplesCh chan<- stats.SampleContainer,
	options *lib.Options,
	tags map[string]string,
) *transport ***REMOVED***
	return &transport***REMOVED***
		roundTripper: roundTripper,
		tags:         tags,
		options:      options,
		samplesCh:    samplesCh,
	***REMOVED***
***REMOVED***

// SetOptions sets the options that should be used
func (t *transport) SetOptions(options *lib.Options) ***REMOVED***
	t.options = options
***REMOVED***

// GetTrail returns the Trail for the last request through the Transport
func (t *transport) GetTrail() *Trail ***REMOVED***
	return t.trail
***REMOVED***

// TLSInfo returns the TLSInfo of the last tls request through the transport
func (t *transport) TLSInfo() netext.TLSInfo ***REMOVED***
	return t.tlsInfo
***REMOVED***

// RoundTrip is the implementation of http.RoundTripper
func (t *transport) RoundTrip(req *http.Request) (res *http.Response, err error) ***REMOVED***
	if t.roundTripper == nil ***REMOVED***
		return nil, errors.New("no roundtrip defined")
	***REMOVED***

	tags := map[string]string***REMOVED******REMOVED***
	for k, v := range t.tags ***REMOVED***
		tags[k] = v
	***REMOVED***

	ctx := req.Context()
	tracer := Tracer***REMOVED******REMOVED***
	reqWithTracer := req.WithContext(httptrace.WithClientTrace(ctx, tracer.Trace()))

	resp, err := t.roundTripper.RoundTrip(reqWithTracer)
	trail := tracer.Done()
	if err != nil ***REMOVED***
		t.errorCode, t.errorMsg = errorCodeForError(err)
		if t.options.SystemTags["error"] ***REMOVED***
			tags["error"] = t.errorMsg
		***REMOVED***

		if t.options.SystemTags["error_code"] ***REMOVED***
			tags["error_code"] = strconv.Itoa(int(t.errorCode))
		***REMOVED***

		if t.options.SystemTags["status"] ***REMOVED***
			tags["status"] = "0"
		***REMOVED***
	***REMOVED*** else ***REMOVED***
		if t.options.SystemTags["url"] ***REMOVED***
			tags["url"] = req.URL.String()
		***REMOVED***
		if t.options.SystemTags["status"] ***REMOVED***
			tags["status"] = strconv.Itoa(resp.StatusCode)
		***REMOVED***
		if resp.StatusCode >= 400 ***REMOVED***
			if t.options.SystemTags["error_code"] ***REMOVED***
				t.errorCode = errCode(1000 + resp.StatusCode)
				tags["error_code"] = strconv.Itoa(int(t.errorCode))
			***REMOVED***
		***REMOVED***
		if t.options.SystemTags["proto"] ***REMOVED***
			tags["proto"] = resp.Proto
		***REMOVED***

		if resp.TLS != nil ***REMOVED***
			tlsInfo, oscp := netext.ParseTLSConnState(resp.TLS)
			if t.options.SystemTags["tls_version"] ***REMOVED***
				tags["tls_version"] = tlsInfo.Version
			***REMOVED***
			if t.options.SystemTags["ocsp_status"] ***REMOVED***
				tags["ocsp_status"] = oscp.Status
			***REMOVED***

			t.tlsInfo = tlsInfo
		***REMOVED***
	***REMOVED***
	if t.options.SystemTags["ip"] && trail.ConnRemoteAddr != nil ***REMOVED***
		var ip string
		if ip, _, err = net.SplitHostPort(trail.ConnRemoteAddr.String()); err == nil ***REMOVED***
			tags["ip"] = ip
		***REMOVED***
	***REMOVED***

	t.trail = trail
	trail.SaveSamples(stats.IntoSampleTags(&tags))
	stats.PushIfNotCancelled(ctx, t.samplesCh, trail)

	return resp, err
***REMOVED***
