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
package ws

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/dop251/goja"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.k6.io/k6/js/common"
	"go.k6.io/k6/lib"
	"go.k6.io/k6/lib/metrics"
	"go.k6.io/k6/lib/testutils/httpmultibin"
	"go.k6.io/k6/stats"
)

func assertSessionMetricsEmitted(t *testing.T, sampleContainers []stats.SampleContainer, subprotocol, url string, status int, group string) ***REMOVED***
	seenSessions := false
	seenSessionDuration := false
	seenConnecting := false

	for _, sampleContainer := range sampleContainers ***REMOVED***
		for _, sample := range sampleContainer.GetSamples() ***REMOVED***
			tags := sample.Tags.CloneTags()
			if tags["url"] == url ***REMOVED***
				switch sample.Metric.Name ***REMOVED***
				case metrics.WSConnectingName:
					seenConnecting = true
				case metrics.WSSessionDurationName:
					seenSessionDuration = true
				case metrics.WSSessionsName:
					seenSessions = true
				***REMOVED***

				assert.Equal(t, strconv.Itoa(status), tags["status"])
				assert.Equal(t, subprotocol, tags["subproto"])
				assert.Equal(t, group, tags["group"])
			***REMOVED***
		***REMOVED***
	***REMOVED***
	assert.True(t, seenConnecting, "url %s didn't emit Connecting", url)
	assert.True(t, seenSessions, "url %s didn't emit Sessions", url)
	assert.True(t, seenSessionDuration, "url %s didn't emit SessionDuration", url)
***REMOVED***

func assertMetricEmitted(t *testing.T, metricName string, sampleContainers []stats.SampleContainer, url string) ***REMOVED***
	seenMetric := false

	for _, sampleContainer := range sampleContainers ***REMOVED***
		for _, sample := range sampleContainer.GetSamples() ***REMOVED***
			surl, ok := sample.Tags.Get("url")
			assert.True(t, ok)
			if surl == url ***REMOVED***
				if sample.Metric.Name == metricName ***REMOVED***
					seenMetric = true
				***REMOVED***
			***REMOVED***
		***REMOVED***
	***REMOVED***
	assert.True(t, seenMetric, "url %s didn't emit %s", url, metricName)
***REMOVED***

func TestSession(t *testing.T) ***REMOVED***
	// TODO: split and paralelize tests
	t.Parallel()
	tb := httpmultibin.NewHTTPMultiBin(t)
	sr := tb.Replacer.Replace

	root, err := lib.NewGroup("", nil)
	assert.NoError(t, err)

	rt := goja.New()
	rt.SetFieldNameMapper(common.FieldNameMapper***REMOVED******REMOVED***)
	samples := make(chan stats.SampleContainer, 1000)
	state := &lib.State***REMOVED***
		Group:  root,
		Dialer: tb.Dialer,
		Options: lib.Options***REMOVED***
			SystemTags: stats.NewSystemTagSet(
				stats.TagURL,
				stats.TagProto,
				stats.TagStatus,
				stats.TagSubproto,
			),
		***REMOVED***,
		Samples:        samples,
		TLSConfig:      tb.TLSClientConfig,
		BuiltinMetrics: metrics.RegisterBuiltinMetrics(metrics.NewRegistry()),
	***REMOVED***

	ctx := context.Background()
	ctx = lib.WithState(ctx, state)
	ctx = common.WithRuntime(ctx, rt)

	rt.Set("ws", common.Bind(rt, New(), &ctx))

	t.Run("connect_ws", func(t *testing.T) ***REMOVED***
		_, err := rt.RunString(sr(`
		var res = ws.connect("WSBIN_URL/ws-echo", function(socket)***REMOVED***
			socket.close()
		***REMOVED***);
		if (res.status != 101) ***REMOVED*** throw new Error("connection failed with status: " + res.status); ***REMOVED***
		`))
		assert.NoError(t, err)
	***REMOVED***)
	assertSessionMetricsEmitted(t, stats.GetBufferedSamples(samples), "", sr("WSBIN_URL/ws-echo"), 101, "")

	t.Run("connect_wss", func(t *testing.T) ***REMOVED***
		_, err := rt.RunString(sr(`
		var res = ws.connect("WSSBIN_URL/ws-echo", function(socket)***REMOVED***
			socket.close()
		***REMOVED***);
		if (res.status != 101) ***REMOVED*** throw new Error("TLS connection failed with status: " + res.status); ***REMOVED***
		`))
		assert.NoError(t, err)
	***REMOVED***)
	assertSessionMetricsEmitted(t, stats.GetBufferedSamples(samples), "", sr("WSSBIN_URL/ws-echo"), 101, "")

	t.Run("open", func(t *testing.T) ***REMOVED***
		_, err := rt.RunString(sr(`
		var opened = false;
		var res = ws.connect("WSBIN_URL/ws-echo", function(socket)***REMOVED***
			socket.on("open", function() ***REMOVED***
				opened = true;
				socket.close()
			***REMOVED***)
		***REMOVED***);
		if (!opened) ***REMOVED*** throw new Error ("open event not fired"); ***REMOVED***
		`))
		assert.NoError(t, err)
	***REMOVED***)
	assertSessionMetricsEmitted(t, stats.GetBufferedSamples(samples), "", sr("WSBIN_URL/ws-echo"), 101, "")

	t.Run("send_receive", func(t *testing.T) ***REMOVED***
		_, err := rt.RunString(sr(`
		var res = ws.connect("WSBIN_URL/ws-echo", function(socket)***REMOVED***
			socket.on("open", function() ***REMOVED***
				socket.send("test")
			***REMOVED***)
			socket.on("message", function (data)***REMOVED***
				if (!data=="test") ***REMOVED***
					throw new Error ("echo'd data doesn't match our message!");
				***REMOVED***
				socket.close()
			***REMOVED***);
		***REMOVED***);
		`))
		assert.NoError(t, err)
	***REMOVED***)

	samplesBuf := stats.GetBufferedSamples(samples)
	assertSessionMetricsEmitted(t, samplesBuf, "", sr("WSBIN_URL/ws-echo"), 101, "")
	assertMetricEmitted(t, metrics.WSMessagesSentName, samplesBuf, sr("WSBIN_URL/ws-echo"))
	assertMetricEmitted(t, metrics.WSMessagesReceivedName, samplesBuf, sr("WSBIN_URL/ws-echo"))

	t.Run("interval", func(t *testing.T) ***REMOVED***
		_, err := rt.RunString(sr(`
		var counter = 0;
		var res = ws.connect("WSBIN_URL/ws-echo", function(socket)***REMOVED***
			socket.setInterval(function () ***REMOVED***
				counter += 1;
				if (counter > 2) ***REMOVED*** socket.close(); ***REMOVED***
			***REMOVED***, 100);
		***REMOVED***);
		if (counter < 3) ***REMOVED***throw new Error ("setInterval should have been called at least 3 times, counter=" + counter);***REMOVED***
		`))
		assert.NoError(t, err)
	***REMOVED***)
	assertSessionMetricsEmitted(t, stats.GetBufferedSamples(samples), "", sr("WSBIN_URL/ws-echo"), 101, "")
	t.Run("bad interval", func(t *testing.T) ***REMOVED***
		_, err := rt.RunString(sr(`
		var counter = 0;
		var res = ws.connect("WSBIN_URL/ws-echo", function(socket)***REMOVED***
			socket.setInterval(function () ***REMOVED***
				counter += 1;
				if (counter > 2) ***REMOVED*** socket.close(); ***REMOVED***
			***REMOVED***, -1.23);
		***REMOVED***);
		`))
		require.Error(t, err)
		require.Contains(t, err.Error(), "setInterval requires a >0 timeout parameter, received -1.23 ")
	***REMOVED***)

	t.Run("timeout", func(t *testing.T) ***REMOVED***
		_, err := rt.RunString(sr(`
		var start = new Date().getTime();
		var ellapsed = new Date().getTime() - start;
		var res = ws.connect("WSBIN_URL/ws-echo", function(socket)***REMOVED***
			socket.setTimeout(function () ***REMOVED***
				ellapsed = new Date().getTime() - start;
				socket.close();
			***REMOVED***, 500);
		***REMOVED***);
		if (ellapsed > 3000 || ellapsed < 500) ***REMOVED***
			throw new Error ("setTimeout occurred after " + ellapsed + "ms, expected 500<T<3000");
		***REMOVED***
		`))
		assert.NoError(t, err)
	***REMOVED***)

	t.Run("bad timeout", func(t *testing.T) ***REMOVED***
		_, err := rt.RunString(sr(`
		var start = new Date().getTime();
		var ellapsed = new Date().getTime() - start;
		var res = ws.connect("WSBIN_URL/ws-echo", function(socket)***REMOVED***
			socket.setTimeout(function () ***REMOVED***
				ellapsed = new Date().getTime() - start;
				socket.close();
			***REMOVED***, 0);
		***REMOVED***);
		`))
		require.Error(t, err)
		require.Contains(t, err.Error(), "setTimeout requires a >0 timeout parameter, received 0.00 ")
	***REMOVED***)
	assertSessionMetricsEmitted(t, stats.GetBufferedSamples(samples), "", sr("WSBIN_URL/ws-echo"), 101, "")

	t.Run("ping", func(t *testing.T) ***REMOVED***
		_, err := rt.RunString(sr(`
		var pongReceived = false;
		var res = ws.connect("WSBIN_URL/ws-echo", function(socket)***REMOVED***
			socket.on("open", function(data) ***REMOVED***
				socket.ping();
			***REMOVED***);
			socket.on("pong", function() ***REMOVED***
				pongReceived = true;
				socket.close();
			***REMOVED***);
			socket.setTimeout(function ()***REMOVED***socket.close();***REMOVED***, 3000);
		***REMOVED***);
		if (!pongReceived) ***REMOVED***
			throw new Error ("sent ping but didn't get pong back");
		***REMOVED***
		`))
		assert.NoError(t, err)
	***REMOVED***)

	samplesBuf = stats.GetBufferedSamples(samples)
	assertSessionMetricsEmitted(t, samplesBuf, "", sr("WSBIN_URL/ws-echo"), 101, "")
	assertMetricEmitted(t, metrics.WSPingName, samplesBuf, sr("WSBIN_URL/ws-echo"))

	t.Run("multiple_handlers", func(t *testing.T) ***REMOVED***
		_, err := rt.RunString(sr(`
		var pongReceived = false;
		var otherPongReceived = false;

		var res = ws.connect("WSBIN_URL/ws-echo", function(socket)***REMOVED***
			socket.on("open", function(data) ***REMOVED***
				socket.ping();
			***REMOVED***);
			socket.on("pong", function() ***REMOVED***
				pongReceived = true;
				if (otherPongReceived) ***REMOVED***
					socket.close();
				***REMOVED***
			***REMOVED***);
			socket.on("pong", function() ***REMOVED***
				otherPongReceived = true;
				if (pongReceived) ***REMOVED***
					socket.close();
				***REMOVED***
			***REMOVED***);
			socket.setTimeout(function ()***REMOVED***socket.close();***REMOVED***, 3000);
		***REMOVED***);
		if (!pongReceived || !otherPongReceived) ***REMOVED***
			throw new Error ("sent ping but didn't get pong back");
		***REMOVED***
		`))
		assert.NoError(t, err)
	***REMOVED***)

	samplesBuf = stats.GetBufferedSamples(samples)
	assertSessionMetricsEmitted(t, samplesBuf, "", sr("WSBIN_URL/ws-echo"), 101, "")
	assertMetricEmitted(t, metrics.WSPingName, samplesBuf, sr("WSBIN_URL/ws-echo"))

	t.Run("client_close", func(t *testing.T) ***REMOVED***
		_, err := rt.RunString(sr(`
		var closed = false;
		var res = ws.connect("WSBIN_URL/ws-echo", function(socket)***REMOVED***
			socket.on("open", function() ***REMOVED***
							socket.close()
			***REMOVED***)
			socket.on("close", function() ***REMOVED***
							closed = true;
			***REMOVED***)
		***REMOVED***);
		if (!closed) ***REMOVED*** throw new Error ("close event not fired"); ***REMOVED***
		`))
		assert.NoError(t, err)
	***REMOVED***)
	assertSessionMetricsEmitted(t, stats.GetBufferedSamples(samples), "", sr("WSBIN_URL/ws-echo"), 101, "")

	serverCloseTests := []struct ***REMOVED***
		name     string
		endpoint string
	***REMOVED******REMOVED***
		***REMOVED***"server_close_ok", "/ws-echo"***REMOVED***,
		// Ensure we correctly handle invalid WS server
		// implementations that close the connection prematurely
		// without sending a close control frame first.
		***REMOVED***"server_close_invalid", "/ws-close-invalid"***REMOVED***,
	***REMOVED***

	for _, tc := range serverCloseTests ***REMOVED***
		tc := tc
		t.Run(tc.name, func(t *testing.T) ***REMOVED***
			_, err := rt.RunString(sr(fmt.Sprintf(`
			var closed = false;
			var res = ws.connect("WSBIN_URL%s", function(socket)***REMOVED***
				socket.on("open", function() ***REMOVED***
					socket.send("test");
				***REMOVED***)
				socket.on("close", function() ***REMOVED***
					closed = true;
				***REMOVED***)
			***REMOVED***);
			if (!closed) ***REMOVED*** throw new Error ("close event not fired"); ***REMOVED***
			`, tc.endpoint)))
			assert.NoError(t, err)
		***REMOVED***)
	***REMOVED***
***REMOVED***

func TestSocketSendBinary(t *testing.T) ***REMOVED*** //nolint: tparallel
	t.Parallel()
	tb := httpmultibin.NewHTTPMultiBin(t)
	sr := tb.Replacer.Replace

	root, err := lib.NewGroup("", nil)
	assert.NoError(t, err)

	rt := goja.New()
	rt.SetFieldNameMapper(common.FieldNameMapper***REMOVED******REMOVED***)
	samples := make(chan stats.SampleContainer, 1000)
	state := &lib.State***REMOVED*** //nolint: exhaustivestruct
		Group:  root,
		Dialer: tb.Dialer,
		Options: lib.Options***REMOVED*** //nolint: exhaustivestruct
			SystemTags: stats.NewSystemTagSet(
				stats.TagURL,
				stats.TagProto,
				stats.TagStatus,
				stats.TagSubproto,
			),
		***REMOVED***,
		Samples:        samples,
		TLSConfig:      tb.TLSClientConfig,
		BuiltinMetrics: metrics.RegisterBuiltinMetrics(metrics.NewRegistry()),
	***REMOVED***

	ctx := context.Background()
	ctx = lib.WithState(ctx, state)
	ctx = common.WithRuntime(ctx, rt)

	err = rt.Set("ws", common.Bind(rt, New(), &ctx))
	assert.NoError(t, err)

	t.Run("ok", func(t *testing.T) ***REMOVED***
		_, err = rt.RunString(sr(`
		var gotMsg = false;
		var res = ws.connect('WSBIN_URL/ws-echo', function(socket)***REMOVED***
			var data = new Uint8Array([104, 101, 108, 108, 111]); // 'hello'

			socket.on('open', function() ***REMOVED***
				socket.sendBinary(data.buffer);
			***REMOVED***)
			socket.on('binaryMessage', function(msg) ***REMOVED***
				gotMsg = true;
				let decText = String.fromCharCode.apply(null, new Uint8Array(msg));
				decText = decodeURIComponent(escape(decText));
				if (decText !== 'hello') ***REMOVED***
					throw new Error('received unexpected binary message: ' + decText);
				***REMOVED***
				socket.close()
			***REMOVED***);
		***REMOVED***);
		if (!gotMsg) ***REMOVED***
			throw new Error("the 'binaryMessage' handler wasn't called")
		***REMOVED***
		`))
		assert.NoError(t, err)
	***REMOVED***)

	errTestCases := []struct ***REMOVED***
		in, expErrType string
	***REMOVED******REMOVED***
		***REMOVED***"", ""***REMOVED***,
		***REMOVED***"undefined", "undefined"***REMOVED***,
		***REMOVED***"null", "null"***REMOVED***,
		***REMOVED***"true", "Boolean"***REMOVED***,
		***REMOVED***"1", "Number"***REMOVED***,
		***REMOVED***"3.14", "Number"***REMOVED***,
		***REMOVED***"'str'", "String"***REMOVED***,
		***REMOVED***"[1, 2, 3]", "Array"***REMOVED***,
		***REMOVED***"new Uint8Array([1, 2, 3])", "Object"***REMOVED***,
		***REMOVED***"Symbol('a')", "Symbol"***REMOVED***,
		***REMOVED***"function() ***REMOVED******REMOVED***", "Function"***REMOVED***,
	***REMOVED***

	for _, tc := range errTestCases ***REMOVED*** //nolint: paralleltest
		tc := tc
		t.Run(fmt.Sprintf("err_%s", tc.expErrType), func(t *testing.T) ***REMOVED***
			_, err = rt.RunString(fmt.Sprintf(sr(`
			var res = ws.connect('WSBIN_URL/ws-echo', function(socket)***REMOVED***
				socket.on('open', function() ***REMOVED***
					socket.sendBinary(%s);
				***REMOVED***)
			***REMOVED***);
		`), tc.in))
			require.Error(t, err)
			if tc.in == "" ***REMOVED***
				assert.Contains(t, err.Error(), "missing argument, expected ArrayBuffer")
			***REMOVED*** else ***REMOVED***
				assert.Contains(t, err.Error(), fmt.Sprintf("expected ArrayBuffer as argument, received: %s", tc.expErrType))
			***REMOVED***
		***REMOVED***)
	***REMOVED***
***REMOVED***

func TestErrors(t *testing.T) ***REMOVED***
	t.Parallel()
	tb := httpmultibin.NewHTTPMultiBin(t)
	sr := tb.Replacer.Replace

	root, err := lib.NewGroup("", nil)
	assert.NoError(t, err)

	rt := goja.New()
	rt.SetFieldNameMapper(common.FieldNameMapper***REMOVED******REMOVED***)
	samples := make(chan stats.SampleContainer, 1000)
	state := &lib.State***REMOVED***
		Group:  root,
		Dialer: tb.Dialer,
		Options: lib.Options***REMOVED***
			SystemTags: &stats.DefaultSystemTagSet,
		***REMOVED***,
		Samples:        samples,
		BuiltinMetrics: metrics.RegisterBuiltinMetrics(metrics.NewRegistry()),
	***REMOVED***

	ctx := context.Background()
	ctx = lib.WithState(ctx, state)
	ctx = common.WithRuntime(ctx, rt)

	rt.Set("ws", common.Bind(rt, New(), &ctx))

	t.Run("invalid_url", func(t *testing.T) ***REMOVED***
		_, err := rt.RunString(`
		var res = ws.connect("INVALID", function(socket)***REMOVED***
			socket.on("open", function() ***REMOVED***
				socket.close();
			***REMOVED***);
		***REMOVED***);
		`)
		assert.Error(t, err)
	***REMOVED***)

	t.Run("invalid_url_message_panic", func(t *testing.T) ***REMOVED***
		// Attempting to send a message to a non-existent socket shouldn't panic
		_, err := rt.RunString(`
		var res = ws.connect("INVALID", function(socket)***REMOVED***
			socket.send("new message");
		***REMOVED***);
		`)
		assert.Error(t, err)
	***REMOVED***)

	t.Run("error_in_setup", func(t *testing.T) ***REMOVED***
		_, err := rt.RunString(sr(`
		var res = ws.connect("WSBIN_URL/ws-echo-invalid", function(socket)***REMOVED***
			throw new Error("error in setup");
		***REMOVED***);
		`))
		assert.Error(t, err)
	***REMOVED***)

	t.Run("send_after_close", func(t *testing.T) ***REMOVED***
		_, err := rt.RunString(sr(`
		var hasError = false;
		var res = ws.connect("WSBIN_URL/ws-echo-invalid", function(socket)***REMOVED***
			socket.on("open", function() ***REMOVED***
				socket.close();
				socket.send("test");
			***REMOVED***);

			socket.on("error", function(errorEvent) ***REMOVED***
				hasError = true;
			***REMOVED***);
		***REMOVED***);
		if (!hasError) ***REMOVED***
			throw new Error ("no error emitted for send after close");
		***REMOVED***
		`))
		assert.NoError(t, err)
		assertSessionMetricsEmitted(t, stats.GetBufferedSamples(samples), "", sr("WSBIN_URL/ws-echo-invalid"), 101, "")
	***REMOVED***)

	t.Run("error on close", func(t *testing.T) ***REMOVED***
		_, err := rt.RunString(sr(`
		var closed = false;
		var res = ws.connect("WSBIN_URL/ws-close", function(socket)***REMOVED***
			socket.on('open', function open() ***REMOVED***
				socket.setInterval(function timeout() ***REMOVED***
				  socket.ping();
				***REMOVED***, 1000);
			***REMOVED***);

			socket.on("ping", function() ***REMOVED***
				socket.close();
			***REMOVED***);

			socket.on("error", function(errorEvent) ***REMOVED***
				if (errorEvent == null) ***REMOVED***
					throw new Error(JSON.stringify(errorEvent));
				***REMOVED***
				if (!closed) ***REMOVED***
					closed = true;
				    socket.close();
				***REMOVED***
			***REMOVED***);
		***REMOVED***);
		`))
		assert.NoError(t, err)
		assertSessionMetricsEmitted(t, stats.GetBufferedSamples(samples), "", sr("WSBIN_URL/ws-close"), 101, "")
	***REMOVED***)
***REMOVED***

func TestSystemTags(t *testing.T) ***REMOVED***
	tb := httpmultibin.NewHTTPMultiBin(t)

	sr := tb.Replacer.Replace

	root, err := lib.NewGroup("", nil)
	assert.NoError(t, err)

	rt := goja.New()
	rt.SetFieldNameMapper(common.FieldNameMapper***REMOVED******REMOVED***)

	// TODO: test for actual tag values after removing the dependency on the
	// external service demos.kaazing.com (https://github.com/k6io/k6/issues/537)
	testedSystemTags := []string***REMOVED***"group", "status", "subproto", "url", "ip"***REMOVED***

	samples := make(chan stats.SampleContainer, 1000)
	state := &lib.State***REMOVED***
		Group:          root,
		Dialer:         tb.Dialer,
		Options:        lib.Options***REMOVED***SystemTags: stats.ToSystemTagSet(testedSystemTags)***REMOVED***,
		Samples:        samples,
		TLSConfig:      tb.TLSClientConfig,
		BuiltinMetrics: metrics.RegisterBuiltinMetrics(metrics.NewRegistry()),
	***REMOVED***

	ctx := context.Background()
	ctx = lib.WithState(ctx, state)
	ctx = common.WithRuntime(ctx, rt)

	rt.Set("ws", common.Bind(rt, New(), &ctx))

	for _, expectedTag := range testedSystemTags ***REMOVED***
		expectedTag := expectedTag
		t.Run("only "+expectedTag, func(t *testing.T) ***REMOVED***
			state.Options.SystemTags = stats.ToSystemTagSet([]string***REMOVED***expectedTag***REMOVED***)
			_, err := rt.RunString(sr(`
			var res = ws.connect("WSBIN_URL/ws-echo", function(socket)***REMOVED***
				socket.on("open", function() ***REMOVED***
					socket.send("test")
				***REMOVED***)
				socket.on("message", function (data)***REMOVED***
					if (!data=="test") ***REMOVED***
						throw new Error ("echo'd data doesn't match our message!");
					***REMOVED***
					socket.close()
				***REMOVED***);
			***REMOVED***);
			`))
			assert.NoError(t, err)

			for _, sampleContainer := range stats.GetBufferedSamples(samples) ***REMOVED***
				for _, sample := range sampleContainer.GetSamples() ***REMOVED***
					for emittedTag := range sample.Tags.CloneTags() ***REMOVED***
						assert.Equal(t, expectedTag, emittedTag)
					***REMOVED***
				***REMOVED***
			***REMOVED***
		***REMOVED***)
	***REMOVED***
***REMOVED***

func TestTLSConfig(t *testing.T) ***REMOVED***
	root, err := lib.NewGroup("", nil)
	assert.NoError(t, err)

	tb := httpmultibin.NewHTTPMultiBin(t)

	sr := tb.Replacer.Replace

	rt := goja.New()
	rt.SetFieldNameMapper(common.FieldNameMapper***REMOVED******REMOVED***)
	samples := make(chan stats.SampleContainer, 1000)
	state := &lib.State***REMOVED***
		Group:  root,
		Dialer: tb.Dialer,
		Options: lib.Options***REMOVED***
			SystemTags: stats.NewSystemTagSet(
				stats.TagURL,
				stats.TagProto,
				stats.TagStatus,
				stats.TagSubproto,
				stats.TagIP,
			),
		***REMOVED***,
		Samples:        samples,
		BuiltinMetrics: metrics.RegisterBuiltinMetrics(metrics.NewRegistry()),
	***REMOVED***

	ctx := context.Background()
	ctx = lib.WithState(ctx, state)
	ctx = common.WithRuntime(ctx, rt)

	rt.Set("ws", common.Bind(rt, New(), &ctx))

	t.Run("insecure skip verify", func(t *testing.T) ***REMOVED***
		state.TLSConfig = &tls.Config***REMOVED***
			InsecureSkipVerify: true,
		***REMOVED***

		_, err := rt.RunString(sr(`
		var res = ws.connect("WSSBIN_URL/ws-close", function(socket)***REMOVED***
			socket.close()
		***REMOVED***);
		if (res.status != 101) ***REMOVED*** throw new Error("TLS connection failed with status: " + res.status); ***REMOVED***
		`))
		assert.NoError(t, err)
	***REMOVED***)
	assertSessionMetricsEmitted(t, stats.GetBufferedSamples(samples), "", sr("WSSBIN_URL/ws-close"), 101, "")

	t.Run("custom certificates", func(t *testing.T) ***REMOVED***
		state.TLSConfig = tb.TLSClientConfig

		_, err := rt.RunString(sr(`
			var res = ws.connect("WSSBIN_URL/ws-close", function(socket)***REMOVED***
				socket.close()
			***REMOVED***);
			if (res.status != 101) ***REMOVED***
				throw new Error("TLS connection failed with status: " + res.status);
			***REMOVED***
		`))
		assert.NoError(t, err)
	***REMOVED***)
	assertSessionMetricsEmitted(t, stats.GetBufferedSamples(samples), "", sr("WSSBIN_URL/ws-close"), 101, "")
***REMOVED***

func TestReadPump(t *testing.T) ***REMOVED***
	var closeCode int
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) ***REMOVED***
		conn, err := (&websocket.Upgrader***REMOVED******REMOVED***).Upgrade(w, r, w.Header())
		assert.NoError(t, err)
		closeMsg := websocket.FormatCloseMessage(closeCode, "")
		_ = conn.WriteControl(websocket.CloseMessage, closeMsg, time.Now().Add(time.Second))
	***REMOVED***))
	defer srv.Close()

	closeCodes := []int***REMOVED***websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseInternalServerErr***REMOVED***

	numAsserts := 0
	srvURL := "ws://" + srv.Listener.Addr().String()

	// Ensure readPump returns the response close code sent by the server
	for _, code := range closeCodes ***REMOVED***
		code := code
		t.Run(strconv.Itoa(code), func(t *testing.T) ***REMOVED***
			closeCode = code
			conn, resp, err := websocket.DefaultDialer.Dial(srvURL, nil)
			assert.NoError(t, err)
			defer func() ***REMOVED***
				_ = resp.Body.Close()
				_ = conn.Close()
			***REMOVED***()

			msgChan := make(chan *message)
			errChan := make(chan error)
			closeChan := make(chan int)
			s := &Socket***REMOVED***conn: conn***REMOVED***
			go s.readPump(msgChan, errChan, closeChan)

		readChans:
			for ***REMOVED***
				select ***REMOVED***
				case responseCode := <-closeChan:
					assert.Equal(t, code, responseCode)
					numAsserts++
					break readChans
				case <-errChan:
					continue
				case <-time.After(time.Second):
					t.Errorf("Read timed out")
					break readChans
				***REMOVED***
			***REMOVED***
		***REMOVED***)
	***REMOVED***

	// Ensure all close code asserts passed
	assert.Equal(t, numAsserts, len(closeCodes))
***REMOVED***
