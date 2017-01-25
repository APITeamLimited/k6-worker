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

package simple

import (
	"context"
	"errors"
	"io"
	"io/ioutil"
	"math"
	"net"
	"net/http"
	"net/http/httptrace"
	"net/url"
	"strconv"
	"time"

	"github.com/loadimpact/k6/lib"
	"github.com/loadimpact/k6/stats"
)

var (
	MetricReqs          = stats.New("http_reqs", stats.Counter)
	MetricReqDuration   = stats.New("http_req_duration", stats.Trend, stats.Time)
	MetricReqBlocked    = stats.New("http_req_blocked", stats.Trend, stats.Time)
	MetricReqLookingUp  = stats.New("http_req_looking_up", stats.Trend, stats.Time)
	MetricReqConnecting = stats.New("http_req_connecting", stats.Trend, stats.Time)
	MetricReqSending    = stats.New("http_req_sending", stats.Trend, stats.Time)
	MetricReqWaiting    = stats.New("http_req_waiting", stats.Trend, stats.Time)
	MetricReqReceiving  = stats.New("http_req_receiving", stats.Trend, stats.Time)
	ErrEmptyScheme      = errors.New("URL contained no scheme")
)

type Runner struct ***REMOVED***
	URL       *url.URL
	SrcData   *lib.SourceData
	Transport *http.Transport
	Options   lib.Options

	defaultGroup *lib.Group
***REMOVED***

func New(src *lib.SourceData, u *url.URL) (*Runner, error) ***REMOVED***
	return &Runner***REMOVED***
		URL:     u,
		SrcData: src,
		Transport: &http.Transport***REMOVED***
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer***REMOVED***
				Timeout:   10 * time.Second,
				KeepAlive: 60 * time.Second,
				DualStack: true,
			***REMOVED***).DialContext,
			MaxIdleConns:        math.MaxInt32,
			MaxIdleConnsPerHost: math.MaxInt32,
		***REMOVED***,
		defaultGroup: &lib.Group***REMOVED******REMOVED***,
	***REMOVED***, nil
***REMOVED***

func (r *Runner) NewVU() (lib.VU, error) ***REMOVED***
	tracer := &lib.Tracer***REMOVED******REMOVED***

	return &VU***REMOVED***
		Runner:    r,
		URLString: r.URL.String(),
		Request: &http.Request***REMOVED***
			Method: "GET",
			URL:    r.URL,
		***REMOVED***,
		Client: &http.Client***REMOVED***
			Transport: r.Transport,
		***REMOVED***,
		tracer: tracer,
		cTrace: tracer.Trace(),
	***REMOVED***, nil
***REMOVED***

func (r *Runner) GetSourceData() *lib.SourceData ***REMOVED***
	return r.SrcData
***REMOVED***

func (r *Runner) GetGroups() []*lib.Group ***REMOVED***
	return []*lib.Group***REMOVED******REMOVED***
***REMOVED***

func (r *Runner) GetChecks() []*lib.Check ***REMOVED***
	return []*lib.Check***REMOVED******REMOVED***
***REMOVED***

func (r Runner) GetOptions() lib.Options ***REMOVED***
	return r.Options
***REMOVED***

func (r *Runner) ApplyOptions(opts lib.Options) ***REMOVED***
	r.Options = r.Options.Apply(opts)
***REMOVED***

type VU struct ***REMOVED***
	Runner   *Runner
	ID       int64
	IDString string

	URLString string
	Request   *http.Request
	Client    *http.Client

	tracer *lib.Tracer
	cTrace *httptrace.ClientTrace
***REMOVED***

func (u *VU) RunOnce(ctx context.Context) ([]stats.Sample, error) ***REMOVED***
	resp, err := u.Client.Do(u.Request.WithContext(httptrace.WithClientTrace(ctx, u.cTrace)))
	if err != nil ***REMOVED***
		u.tracer.Done()
		return nil, err
	***REMOVED***

	_, _ = io.Copy(ioutil.Discard, resp.Body)
	_ = resp.Body.Close()
	trail := u.tracer.Done()

	tags := map[string]string***REMOVED***
		"vu":     u.IDString,
		"method": "GET",
		"url":    u.URLString,
		"status": strconv.Itoa(resp.StatusCode),
	***REMOVED***

	t := time.Now()
	return []stats.Sample***REMOVED***
		stats.Sample***REMOVED***Metric: MetricReqs, Time: t, Tags: tags, Value: 1***REMOVED***,
		stats.Sample***REMOVED***Metric: MetricReqDuration, Time: t, Tags: tags, Value: float64(trail.Duration)***REMOVED***,
		stats.Sample***REMOVED***Metric: MetricReqBlocked, Time: t, Tags: tags, Value: float64(trail.Blocked)***REMOVED***,
		stats.Sample***REMOVED***Metric: MetricReqLookingUp, Time: t, Tags: tags, Value: float64(trail.LookingUp)***REMOVED***,
		stats.Sample***REMOVED***Metric: MetricReqConnecting, Time: t, Tags: tags, Value: float64(trail.Connecting)***REMOVED***,
		stats.Sample***REMOVED***Metric: MetricReqSending, Time: t, Tags: tags, Value: float64(trail.Sending)***REMOVED***,
		stats.Sample***REMOVED***Metric: MetricReqWaiting, Time: t, Tags: tags, Value: float64(trail.Waiting)***REMOVED***,
		stats.Sample***REMOVED***Metric: MetricReqReceiving, Time: t, Tags: tags, Value: float64(trail.Receiving)***REMOVED***,
	***REMOVED***, nil
***REMOVED***

func (u *VU) Reconfigure(id int64) error ***REMOVED***
	u.ID = id
	u.IDString = strconv.FormatInt(id, 10)
	return nil
***REMOVED***
