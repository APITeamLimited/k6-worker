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

package lib

import (
	"github.com/loadimpact/k6/lib/metrics"
	"github.com/loadimpact/k6/stats"
	"net"
	"net/http/httptrace"
	"time"
)

// A Trail represents detailed information about an HTTP request.
// You'd typically get one from a Tracer.
type Trail struct ***REMOVED***
	StartTime time.Time
	EndTime   time.Time

	// Total request duration, excluding DNS lookup and connect time.
	Duration time.Duration

	Blocked    time.Duration // Waiting to acquire a connection.
	LookingUp  time.Duration // Looking up DNS records.
	Connecting time.Duration // Connecting to remote host.
	Sending    time.Duration // Writing request.
	Waiting    time.Duration // Waiting for first byte.
	Receiving  time.Duration // Receiving response.

	// Detailed connection information.
	ConnReused     bool
	ConnRemoteAddr net.Addr
***REMOVED***

func (tr Trail) Samples(tags map[string]string) []stats.Sample ***REMOVED***
	return []stats.Sample***REMOVED***
		***REMOVED***Metric: metrics.HTTPReqs, Time: tr.EndTime, Tags: tags, Value: 1***REMOVED***,
		***REMOVED***Metric: metrics.HTTPReqDuration, Time: tr.EndTime, Tags: tags, Value: stats.D(tr.Duration)***REMOVED***,
		***REMOVED***Metric: metrics.HTTPReqBlocked, Time: tr.EndTime, Tags: tags, Value: stats.D(tr.Blocked)***REMOVED***,
		***REMOVED***Metric: metrics.HTTPReqLookingUp, Time: tr.EndTime, Tags: tags, Value: stats.D(tr.LookingUp)***REMOVED***,
		***REMOVED***Metric: metrics.HTTPReqConnecting, Time: tr.EndTime, Tags: tags, Value: stats.D(tr.Connecting)***REMOVED***,
		***REMOVED***Metric: metrics.HTTPReqSending, Time: tr.EndTime, Tags: tags, Value: stats.D(tr.Sending)***REMOVED***,
		***REMOVED***Metric: metrics.HTTPReqWaiting, Time: tr.EndTime, Tags: tags, Value: stats.D(tr.Waiting)***REMOVED***,
		***REMOVED***Metric: metrics.HTTPReqReceiving, Time: tr.EndTime, Tags: tags, Value: stats.D(tr.Receiving)***REMOVED***,
	***REMOVED***
***REMOVED***

// A Tracer wraps "net/http/httptrace" to collect granular timings for HTTP requests.
// Note that since there is not yet an event for the end of a request (there's a PR to
// add it), you must call Done() at the end of the request to get the full timings.
// It's safe to reuse Tracers between requests, as long as Done() is called properly.
// Cheers, love, the cavalry's here.
type Tracer struct ***REMOVED***
	getConn              time.Time
	gotConn              time.Time
	gotFirstResponseByte time.Time
	dnsStart             time.Time
	dnsDone              time.Time
	connectStart         time.Time
	connectDone          time.Time
	wroteRequest         time.Time

	connReused     bool
	connRemoteAddr net.Addr
***REMOVED***

// Trace() returns a premade ClientTrace that calls all of the Tracer's hooks.
func (t *Tracer) Trace() *httptrace.ClientTrace ***REMOVED***
	return &httptrace.ClientTrace***REMOVED***
		GetConn:              t.GetConn,
		GotConn:              t.GotConn,
		GotFirstResponseByte: t.GotFirstResponseByte,
		DNSStart:             t.DNSStart,
		DNSDone:              t.DNSDone,
		ConnectStart:         t.ConnectStart,
		ConnectDone:          t.ConnectDone,
		WroteRequest:         t.WroteRequest,
	***REMOVED***
***REMOVED***

// Call when the request is finished. Calculates metrics and resets the tracer.
func (t *Tracer) Done() Trail ***REMOVED***
	done := time.Now()
	trail := Trail***REMOVED***
		StartTime:  t.getConn,
		EndTime:    done,
		Duration:   done.Sub(t.getConn),
		Blocked:    t.gotConn.Sub(t.getConn),
		LookingUp:  t.dnsDone.Sub(t.dnsStart),
		Connecting: t.connectDone.Sub(t.connectStart),
		Sending:    t.wroteRequest.Sub(t.connectDone),
		Waiting:    t.gotFirstResponseByte.Sub(t.wroteRequest),
		Receiving:  done.Sub(t.gotFirstResponseByte),

		ConnReused:     t.connReused,
		ConnRemoteAddr: t.connRemoteAddr,
	***REMOVED***

	*t = Tracer***REMOVED******REMOVED***
	return trail
***REMOVED***

// GetConn event hook.
func (t *Tracer) GetConn(hostPort string) ***REMOVED***
	t.getConn = time.Now()
***REMOVED***

// GotConn event hook.
func (t *Tracer) GotConn(info httptrace.GotConnInfo) ***REMOVED***
	t.gotConn = time.Now()
	t.connReused = info.Reused
	t.connRemoteAddr = info.Conn.RemoteAddr()

	if t.connReused ***REMOVED***
		t.connectStart = t.gotConn
		t.connectDone = t.gotConn
	***REMOVED***
***REMOVED***

// GotFirstResponseByte hook.
func (t *Tracer) GotFirstResponseByte() ***REMOVED***
	t.gotFirstResponseByte = time.Now()
***REMOVED***

// DNSStart hook.
func (t *Tracer) DNSStart(info httptrace.DNSStartInfo) ***REMOVED***
	t.dnsStart = time.Now()
	t.dnsDone = t.dnsStart
***REMOVED***

// DNSDone hook.
func (t *Tracer) DNSDone(info httptrace.DNSDoneInfo) ***REMOVED***
	t.dnsDone = time.Now()
	if t.dnsStart.IsZero() ***REMOVED***
		t.dnsStart = t.dnsDone
	***REMOVED***
***REMOVED***

// ConnectStart hook.
func (t *Tracer) ConnectStart(network, addr string) ***REMOVED***
	// If using dual-stack dialing, it's possible to get this multiple times.
	if !t.connectStart.IsZero() ***REMOVED***
		return
	***REMOVED***
	t.connectStart = time.Now()
***REMOVED***

// ConnectDone hook.
func (t *Tracer) ConnectDone(network, addr string, err error) ***REMOVED***
	// If using dual-stack dialing, it's possible to get this multiple times.
	if !t.connectDone.IsZero() ***REMOVED***
		return
	***REMOVED***
	t.connectDone = time.Now()
***REMOVED***

// WroteRequest hook.
func (t *Tracer) WroteRequest(info httptrace.WroteRequestInfo) ***REMOVED***
	t.wroteRequest = time.Now()
***REMOVED***
