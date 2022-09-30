package execution

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"testing"
	"time"

	"github.com/APITeamLimited/globe-test/worker/errext"
	"github.com/APITeamLimited/globe-test/worker/js/common"
	"github.com/APITeamLimited/globe-test/worker/js/modulestest"
	"github.com/APITeamLimited/globe-test/worker/libWorker"
	"github.com/APITeamLimited/globe-test/worker/libWorker/executor"
	"github.com/APITeamLimited/globe-test/worker/libWorker/testutils"
	"github.com/APITeamLimited/globe-test/worker/libWorker/types"
	"github.com/dop251/goja"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/guregu/null.v3"
)

func setupTagsExecEnv(t *testing.T) *modulestest.Runtime {
	testRuntime := modulestest.NewRuntime(t)
	m, ok := New().NewModuleInstance(testRuntime.VU).(*ModuleInstance)
	require.True(t, ok)
	require.NoError(t, testRuntime.VU.Runtime().Set("exec", m.Exports().Default))

	return testRuntime
}

func TestVUTags(t *testing.T) {
	t.Parallel()

	t.Run("Get", func(t *testing.T) {
		t.Parallel()

		tenv := setupTagsExecEnv(t)
		tenv.MoveToVUContext(&libWorker.State{Tags: libWorker.NewTagMap(map[string]string{"vu": "42"})})
		tag, err := tenv.VU.Runtime().RunString(`exec.vu.tags["vu"]`)
		require.NoError(t, err)
		assert.Equal(t, "42", tag.String())

		// not found
		tag, err = tenv.VU.Runtime().RunString(`exec.vu.tags["not-existing-tag"]`)
		require.NoError(t, err)
		assert.Equal(t, "undefined", tag.String())
	})

	t.Run("JSONEncoding", func(t *testing.T) {
		t.Parallel()

		tenv := setupTagsExecEnv(t)
		tenv.MoveToVUContext(&libWorker.State{
			Options: libWorker.Options{
				SystemTags: workerMetrics.NewSystemTagSet(workerMetrics.TagVU),
			},
			Tags: libWorker.NewTagMap(map[string]string{
				"vu": "42",
			}),
		})
		state := tenv.VU.State()
		state.Tags.Set("custom-tag", "mytag1")

		encoded, err := tenv.VU.Runtime().RunString(`JSON.stringify(exec.vu.tags)`)
		require.NoError(t, err)
		assert.JSONEq(t, `{"vu":"42","custom-tag":"mytag1"}`, encoded.String())
	})

	t.Run("Set", func(t *testing.T) {
		t.Parallel()

		t.Run("SuccessAccetedTypes", func(t *testing.T) {
			t.Parallel()

			// bool and numbers are implicitly converted into string

			tests := map[string]struct {
				v   interface{}
				exp string
			}{
				"string": {v: `"tag1"`, exp: "tag1"},
				"bool":   {v: true, exp: "true"},
				"int":    {v: 101, exp: "101"},
				"float":  {v: 3.14, exp: "3.14"},
			}

			tenv := setupTagsExecEnv(t)
			tenv.MoveToVUContext(&libWorker.State{Tags: libWorker.NewTagMap(map[string]string{"vu": "42"})})

			for _, tc := range tests {
				_, err := tenv.VU.Runtime().RunString(fmt.Sprintf(`exec.vu.tags["mytag"] = %v`, tc.v))
				require.NoError(t, err)

				val, err := tenv.VU.Runtime().RunString(`exec.vu.tags["mytag"]`)
				require.NoError(t, err)

				assert.Equal(t, tc.exp, val.String())
			}
		})

		t.Run("SuccessOverwriteSystemTag", func(t *testing.T) {
			t.Parallel()

			tenv := setupTagsExecEnv(t)
			tenv.MoveToVUContext(&libWorker.State{Tags: libWorker.NewTagMap(map[string]string{"vu": "42"})})

			_, err := tenv.VU.Runtime().RunString(`exec.vu.tags["vu"] = "vu101"`)
			require.NoError(t, err)
			val, err := tenv.VU.Runtime().RunString(`exec.vu.tags["vu"]`)
			require.NoError(t, err)
			assert.Equal(t, "vu101", val.String())
		})

		t.Run("DiscardWrongTypeAndRaisingError", func(t *testing.T) {
			t.Parallel()

			tenv := setupTagsExecEnv(t)
			tenv.MoveToVUContext(&libWorker.State{Tags: libWorker.NewTagMap(map[string]string{"vu": "42"})})

			state := tenv.VU.State()
			state.Options.Throw = null.BoolFrom(true)
			require.NotNil(t, state)

			cases := []string{
				`[1, 3, 5]`,             // array
				`{f1: "value1", f2: 4}`, // object
			}

			for _, val := range cases {
				_, err := tenv.VU.Runtime().RunString(`exec.vu.tags["custom-tag"] = ` + val)
				require.Error(t, err)
				assert.Contains(t, err.Error(), "TypeError:")
				assert.Contains(t, err.Error(), "only String, Boolean and Number")
			}
		})

		t.Run("DiscardWrongTypeOnlyWarning", func(t *testing.T) {
			t.Parallel()
			logHook := &testutils.SimpleLogrusHook{HookedLevels: []logrus.Level{logrus.WarnLevel}}
			testLog := logrus.New()
			testLog.AddHook(logHook)
			testLog.SetOutput(ioutil.Discard)

			tenv := setupTagsExecEnv(t)
			tenv.MoveToVUContext(&libWorker.State{
				Options: libWorker.Options{
					SystemTags: workerMetrics.NewSystemTagSet(workerMetrics.TagVU),
				},
				Tags: libWorker.NewTagMap(map[string]string{
					"vu": "42",
				}),
				Logger: testLog,
			})
			_, err := tenv.VU.Runtime().RunString(`exec.vu.tags["custom-tag"] = [1, 3, 5]`)
			require.NoError(t, err)

			entries := logHook.Drain()
			require.Len(t, entries, 1)
			assert.Contains(t, entries[0].Message, "discarded")
		})

		t.Run("DiscardNullOrUndefined", func(t *testing.T) {
			t.Parallel()

			logHook := &testutils.SimpleLogrusHook{HookedLevels: []logrus.Level{logrus.WarnLevel}}
			testLog := logrus.New()
			testLog.AddHook(logHook)
			testLog.SetOutput(ioutil.Discard)

			cases := []string{"null", "undefined"}
			tenv := setupTagsExecEnv(t)
			tenv.MoveToVUContext(&libWorker.State{
				Options: libWorker.Options{
					SystemTags: workerMetrics.NewSystemTagSet(workerMetrics.TagVU),
				},
				Tags:   libWorker.NewTagMap(map[string]string{"vu": "42"}),
				Logger: testLog,
			})
			for _, val := range cases {
				_, err := tenv.VU.Runtime().RunString(`exec.vu.tags["custom-tag"] = ` + val)
				require.NoError(t, err)

				entries := logHook.Drain()
				require.Len(t, entries, 1)
				assert.Contains(t, entries[0].Message, "discarded")
			}
		})
	})
}

func TestAbortTest(t *testing.T) { //nolint:tparallel
	t.Parallel()

	var (
		rt    = goja.New()
		state = &libWorker.State{}
		ctx   = context.Background()
	)

	m, ok := New().NewModuleInstance(
		&modulestest.VU{
			RuntimeField: rt,
			InitEnvField: &common.InitEnvironment{},
			CtxField:     ctx,
			StateField:   state,
		},
	).(*ModuleInstance)
	require.True(t, ok)
	require.NoError(t, rt.Set("exec", m.Exports().Default))

	prove := func(t *testing.T, script, reason string) {
		_, err := rt.RunString(script)
		require.NotNil(t, err)
		var x *goja.InterruptedError
		assert.ErrorAs(t, err, &x)
		v, ok := x.Value().(*errext.InterruptError)
		require.True(t, ok)
		require.Equal(t, v.Reason, reason)
	}

	t.Run("default reason", func(t *testing.T) { //nolint:paralleltest
		prove(t, "exec.test.abort()", errext.AbortTest)
	})
	t.Run("custom reason", func(t *testing.T) { //nolint:paralleltest
		prove(t, `exec.test.abort("mayday")`, fmt.Sprintf("%s: mayday", errext.AbortTest))
	})
}

func TestOptionsTestFull(t *testing.T) {
	t.Parallel()

	expected := `{"paused":true,"scenarios":{"const-vus":{"executor":"constant-vus","startTime":"10s","gracefulStop":"30s","env":{"FOO":"bar"},"exec":"default","tags":{"tagkey":"tagvalue"},"vus":50,"duration":"10m0s"}},"executionSegment":"0:1/4","executionSegmentSequence":"0,1/4,1/2,1","noSetup":true,"setupTimeout":"1m0s","noTeardown":true,"teardownTimeout":"5m0s","rps":100,"dns":{"ttl":"1m","select":"roundRobin","policy":"any"},"maxRedirects":3,"userAgent":"k6-user-agent","batch":15,"batchPerHost":5,"httpDebug":"full","insecureSkipTLSVerify":true,"tlsCipherSuites":["TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256"],"tlsVersion":{"min":"tls1.2","max":"tls1.3"},"tlsAuth":[{"domains":["example.com"],"cert":"mycert.pem","key":"mycert-key.pem","password":"mypwd"}],"throw":true,"thresholds":{"http_req_duration":[{"threshold":"rate>0.01","abortOnFail":true,"delayAbortEval":"10s"}]},"blacklistIPs":["192.0.2.0/24"],"blockHostnames":["test.k6.io","*.example.com"],"hosts":{"test.k6.io":"1.2.3.4:8443"},"noConnectionReuse":true,"noVUConnectionReuse":true,"minIterationDuration":"10s","ext":{"ext-one":{"rawkey":"rawvalue"}},"summaryTrendStats":["avg","min","max"],"summaryTimeUnit":"ms","systemTags":["iter","vu"],"tags":null,"metricSamplesBufferSize":8,"noCookiesReset":true,"discardResponseBodies":true,"consoleOutput":"loadtest.log","tags":{"runtag-key":"runtag-value"},"localIPs":"192.168.20.12-192.168.20.15,192.168.10.0/27"}`

	var (
		rt    = goja.New()
		state = &libWorker.State{
			Options: libWorker.Options{
				Paused: null.BoolFrom(true),
				Scenarios: map[string]libWorker.ExecutorConfig{
					"const-vus": executor.ConstantVUsConfig{
						BaseConfig: executor.BaseConfig{
							Name:         "const-vus",
							Type:         "constant-vus",
							StartTime:    types.NullDurationFrom(10 * time.Second),
							GracefulStop: types.NullDurationFrom(30 * time.Second),
							Env: map[string]string{
								"FOO": "bar",
							},
							Exec: null.StringFrom("default"),
							Tags: map[string]string{
								"tagkey": "tagvalue",
							},
						},
						VUs:      null.IntFrom(50),
						Duration: types.NullDurationFrom(10 * time.Minute),
					},
				},
				ExecutionSegment: func() *libWorker.ExecutionSegment {
					seg, err := libWorker.NewExecutionSegmentFromString("0:1/4")
					require.NoError(t, err)
					return seg
				}(),
				ExecutionSegmentSequence: func() *libWorker.ExecutionSegmentSequence {
					seq, err := libWorker.NewExecutionSegmentSequenceFromString("0,1/4,1/2,1")
					require.NoError(t, err)
					return &seq
				}(),
				NoSetup:               null.BoolFrom(true),
				NoTeardown:            null.BoolFrom(true),
				NoConnectionReuse:     null.BoolFrom(true),
				NoVUConnectionReuse:   null.BoolFrom(true),
				InsecureSkipTLSVerify: null.BoolFrom(true),
				Throw:                 null.BoolFrom(true),
				NoCookiesReset:        null.BoolFrom(true),
				DiscardResponseBodies: null.BoolFrom(true),
				RPS:                   null.IntFrom(100),
				MaxRedirects:          null.IntFrom(3),
				UserAgent:             null.StringFrom("k6-user-agent"),
				Batch:                 null.IntFrom(15),
				BatchPerHost:          null.IntFrom(5),
				SetupTimeout:          types.NullDurationFrom(1 * time.Minute),
				TeardownTimeout:       types.NullDurationFrom(5 * time.Minute),
				MinIterationDuration:  types.NullDurationFrom(10 * time.Second),
				HTTPDebug:             null.StringFrom("full"),
				DNS: types.DNSConfig{
					TTL:    null.StringFrom("1m"),
					Select: types.NullDNSSelect{DNSSelect: types.DNSroundRobin, Valid: true},
					Policy: types.NullDNSPolicy{DNSPolicy: types.DNSany, Valid: true},
					Valid:  true,
				},
				TLSVersion: &libWorker.TLSVersions{
					Min: tls.VersionTLS12,
					Max: tls.VersionTLS13,
				},
				TLSAuth: []*libWorker.TLSAuth{
					{
						TLSAuthFields: libWorker.TLSAuthFields{
							Cert:     "mycert.pem",
							Key:      "mycert-key.pem",
							Password: null.StringFrom("mypwd"),
							Domains:  []string{"example.com"},
						},
					},
				},
				TLSCipherSuites: &libWorker.TLSCipherSuites{
					tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				},
				BlacklistIPs: []*libWorker.IPNet{
					{
						IPNet: func() net.IPNet {
							_, ipv4net, err := net.ParseCIDR("192.0.2.1/24")
							require.NoError(t, err)
							return *ipv4net
						}(),
					},
				},
				Thresholds: map[string]workerMetrics.Thresholds{
					"http_req_duration": {
						Thresholds: []*workerMetrics.Threshold{
							{
								Source:           "rate>0.01",
								LastFailed:       true,
								AbortOnFail:      true,
								AbortGracePeriod: types.NullDurationFrom(10 * time.Second),
							},
						},
					},
				},
				BlockedHostnames: func() types.NullHostnameTrie {
					bh, err := types.NewNullHostnameTrie([]string{"test.k6.io", "*.example.com"})
					require.NoError(t, err)
					return bh
				}(),
				Hosts: map[string]*libWorker.HostAddress{
					"test.k6.io": {
						IP:   []byte{0x01, 0x02, 0x03, 0x04},
						Port: 8443,
					},
				},
				External: map[string]json.RawMessage{
					"ext-one": json.RawMessage(`{"rawkey":"rawvalue"}`),
				},
				SummaryTrendStats: []string{"avg", "min", "max"},
				SummaryTimeUnit:   null.StringFrom("ms"),
				SystemTags: func() *workerMetrics.SystemTagSet {
					sysm := workerMetrics.TagIter | workerMetrics.TagVU
					return &sysm
				}(),
				RunTags:                 map[string]string{"runtag-key": "runtag-value"},
				MetricSamplesBufferSize: null.IntFrom(8),
				ConsoleOutput:           null.StringFrom("loadtest.log"),
				LocalIPs: func() types.NullIPPool {
					npool := types.NullIPPool{}
					err := npool.UnmarshalText([]byte("192.168.20.12-192.168.20.15,192.168.10.0/27"))
					require.NoError(t, err)
					return npool
				}(),

				// The following fields are not expected to be
				// in the final test.options object
				VUs:        null.IntFrom(50),
				Iterations: null.IntFrom(100),
				Duration:   types.NullDurationFrom(10 * time.Second),
				Stages: []libWorker.Stage{
					{
						Duration: types.NullDurationFrom(2 * time.Second),
						Target:   null.IntFrom(2),
					},
				},
			},
		}
		ctx = context.Background()
	)

	m, ok := New().NewModuleInstance(
		&modulestest.VU{
			RuntimeField: rt,
			CtxField:     ctx,
			StateField:   state,
		},
	).(*ModuleInstance)
	require.True(t, ok)
	require.NoError(t, rt.Set("exec", m.Exports().Default))

	opts, err := rt.RunString(`JSON.stringify(exec.test.options)`)
	require.NoError(t, err)
	require.NotNil(t, opts)
	assert.JSONEq(t, expected, opts.String())
}

func TestOptionsTestSetPropertyDenied(t *testing.T) {
	t.Parallel()

	rt := goja.New()
	m, ok := New().NewModuleInstance(
		&modulestest.VU{
			RuntimeField: rt,
			CtxField:     context.Background(),
			StateField: &libWorker.State{
				Options: libWorker.Options{
					Paused: null.BoolFrom(true),
				},
			},
		},
	).(*ModuleInstance)
	require.True(t, ok)
	require.NoError(t, rt.Set("exec", m.Exports().Default))

	_, err := rt.RunString(`exec.test.options.paused = false`)
	require.NoError(t, err)
	paused, err := rt.RunString(`exec.test.options.paused`)
	require.NoError(t, err)
	assert.Equal(t, true, rt.ToValue(paused).ToBoolean())
}

func TestScenarioNoAvailableInInitContext(t *testing.T) {
	t.Parallel()

	rt := goja.New()
	m, ok := New().NewModuleInstance(
		&modulestest.VU{
			RuntimeField: rt,
			CtxField:     context.Background(),
			StateField: &libWorker.State{
				Options: libWorker.Options{
					Paused: null.BoolFrom(true),
				},
			},
		},
	).(*ModuleInstance)
	require.True(t, ok)
	require.NoError(t, rt.Set("exec", m.Exports().Default))

	scenarioExportedProps := []string{"name", "executor", "startTime", "progress", "iterationInInstance", "iterationInTest"}

	for _, code := range scenarioExportedProps {
		prop := fmt.Sprintf("exec.scenario.%s", code)
		_, err := rt.RunString(prop)
		require.Error(t, err)
		require.ErrorContains(t, err, "getting scenario information outside of the VU context is not supported")
	}
}
