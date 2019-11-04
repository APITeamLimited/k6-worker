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

package netext

import (
	"context"
	"fmt"
	"net"
	"strings"
	"sync/atomic"
	"time"

	"github.com/loadimpact/k6/lib"
	"github.com/loadimpact/k6/lib/metrics"
	"github.com/loadimpact/k6/stats"

	"github.com/viki-org/dnscache"
)

// Dialer wraps net.Dialer and provides k6 specific functionality -
// tracing, blacklists and DNS cache and aliases.
type Dialer struct ***REMOVED***
	net.Dialer

	Resolver  *dnscache.Resolver
	Blacklist []*lib.IPNet
	Hosts     map[string]net.IP

	BytesRead    int64
	BytesWritten int64
***REMOVED***

// NewDialer constructs a new Dialer and initializes its cache.
func NewDialer(dialer net.Dialer) *Dialer ***REMOVED***
	return &Dialer***REMOVED***
		Dialer:   dialer,
		Resolver: dnscache.New(0),
	***REMOVED***
***REMOVED***

// BlackListedIPError is an error that is returned when a given IP is blacklisted
type BlackListedIPError struct ***REMOVED***
	ip  net.IP
	net *lib.IPNet
***REMOVED***

func (b BlackListedIPError) Error() string ***REMOVED***
	return fmt.Sprintf("IP (%s) is in a blacklisted range (%s)", b.ip, b.net)
***REMOVED***

// DialContext wraps the net.Dialer.DialContext and handles the k6 specifics
func (d *Dialer) DialContext(ctx context.Context, proto, addr string) (net.Conn, error) ***REMOVED***
	delimiter := strings.LastIndex(addr, ":")
	host := addr[:delimiter]

	// lookup for domain defined in Hosts option before trying to resolve DNS.
	ip, ok := d.Hosts[host]
	if !ok ***REMOVED***
		var err error
		ip, err = d.Resolver.FetchOne(host)
		if err != nil ***REMOVED***
			return nil, err
		***REMOVED***
	***REMOVED***

	for _, ipnet := range d.Blacklist ***REMOVED***
		if (*net.IPNet)(ipnet).Contains(ip) ***REMOVED***
			return nil, BlackListedIPError***REMOVED***ip: ip, net: ipnet***REMOVED***
		***REMOVED***
	***REMOVED***
	ipStr := ip.String()
	if strings.ContainsRune(ipStr, ':') ***REMOVED***
		ipStr = "[" + ipStr + "]"
	***REMOVED***
	conn, err := d.Dialer.DialContext(ctx, proto, ipStr+":"+addr[delimiter+1:])
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	conn = &Conn***REMOVED***conn, &d.BytesRead, &d.BytesWritten***REMOVED***
	return conn, err
***REMOVED***

// GetTrail creates a new NetTrail instance with the Dialer
// sent and received data metrics and the supplied times and tags.
// TODO: Refactor this according to
// https://github.com/loadimpact/k6/pull/1203#discussion_r337938370
func (d *Dialer) GetTrail(
	startTime, endTime time.Time, fullIteration bool, emitIterations bool, tags *stats.SampleTags,
) *NetTrail ***REMOVED***
	bytesWritten := atomic.SwapInt64(&d.BytesWritten, 0)
	bytesRead := atomic.SwapInt64(&d.BytesRead, 0)
	samples := []stats.Sample***REMOVED***
		***REMOVED***
			Time:   endTime,
			Metric: metrics.DataSent,
			Value:  float64(bytesWritten),
			Tags:   tags,
		***REMOVED***,
		***REMOVED***
			Time:   endTime,
			Metric: metrics.DataReceived,
			Value:  float64(bytesRead),
			Tags:   tags,
		***REMOVED***,
	***REMOVED***
	if fullIteration ***REMOVED***
		samples = append(samples, stats.Sample***REMOVED***
			Time:   endTime,
			Metric: metrics.IterationDuration,
			Value:  stats.D(endTime.Sub(startTime)),
			Tags:   tags,
		***REMOVED***)
		if emitIterations ***REMOVED***
			samples = append(samples, stats.Sample***REMOVED***
				Time:   endTime,
				Metric: metrics.Iterations,
				Value:  1,
				Tags:   tags,
			***REMOVED***)
		***REMOVED***
	***REMOVED***

	return &NetTrail***REMOVED***
		BytesRead:     bytesRead,
		BytesWritten:  bytesWritten,
		FullIteration: fullIteration,
		StartTime:     startTime,
		EndTime:       endTime,
		Tags:          tags,
		Samples:       samples,
	***REMOVED***
***REMOVED***

// NetTrail contains information about the exchanged data size and length of a
// series of connections from a particular netext.Dialer
type NetTrail struct ***REMOVED***
	BytesRead     int64
	BytesWritten  int64
	FullIteration bool
	StartTime     time.Time
	EndTime       time.Time
	Tags          *stats.SampleTags
	Samples       []stats.Sample
***REMOVED***

// Ensure that interfaces are implemented correctly
var _ stats.ConnectedSampleContainer = &NetTrail***REMOVED******REMOVED***

// GetSamples implements the stats.SampleContainer interface.
func (ntr *NetTrail) GetSamples() []stats.Sample ***REMOVED***
	return ntr.Samples
***REMOVED***

// GetTags implements the stats.ConnectedSampleContainer interface.
func (ntr *NetTrail) GetTags() *stats.SampleTags ***REMOVED***
	return ntr.Tags
***REMOVED***

// GetTime implements the stats.ConnectedSampleContainer interface.
func (ntr *NetTrail) GetTime() time.Time ***REMOVED***
	return ntr.EndTime
***REMOVED***

// Conn wraps net.Conn and keeps track of sent and received data size
type Conn struct ***REMOVED***
	net.Conn

	BytesRead, BytesWritten *int64
***REMOVED***

func (c *Conn) Read(b []byte) (int, error) ***REMOVED***
	n, err := c.Conn.Read(b)
	if n > 0 ***REMOVED***
		atomic.AddInt64(c.BytesRead, int64(n))
	***REMOVED***
	return n, err
***REMOVED***

func (c *Conn) Write(b []byte) (int, error) ***REMOVED***
	n, err := c.Conn.Write(b)
	if n > 0 ***REMOVED***
		atomic.AddInt64(c.BytesWritten, int64(n))
	***REMOVED***
	return n, err
***REMOVED***
