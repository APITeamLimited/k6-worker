package libWorker

import (
	"crypto/tls"
	"encoding/json"
	"net"
	"reflect"
	"testing"
	"time"

	"github.com/mstoykov/envconfig"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/guregu/null.v3"

	"github.com/APITeamLimited/globe-test/worker/libWorker/types"
	"github.com/APITeamLimited/globe-test/worker/workerMetrics"
)

func TestOptions(t *testing.T) ***REMOVED***
	t.Parallel()
	t.Run("Paused", func(t *testing.T) ***REMOVED***
		t.Parallel()
		opts := Options***REMOVED******REMOVED***.Apply(Options***REMOVED***Paused: null.BoolFrom(true)***REMOVED***)
		assert.True(t, opts.Paused.Valid)
		assert.True(t, opts.Paused.Bool)
	***REMOVED***)
	t.Run("VUs", func(t *testing.T) ***REMOVED***
		t.Parallel()
		opts := Options***REMOVED******REMOVED***.Apply(Options***REMOVED***VUs: null.IntFrom(12345)***REMOVED***)
		assert.True(t, opts.VUs.Valid)
		assert.Equal(t, int64(12345), opts.VUs.Int64)
	***REMOVED***)
	t.Run("Duration", func(t *testing.T) ***REMOVED***
		t.Parallel()
		opts := Options***REMOVED******REMOVED***.Apply(Options***REMOVED***Duration: types.NullDurationFrom(2 * time.Minute)***REMOVED***)
		assert.True(t, opts.Duration.Valid)
		assert.Equal(t, "2m0s", opts.Duration.String())
	***REMOVED***)
	t.Run("Iterations", func(t *testing.T) ***REMOVED***
		t.Parallel()
		opts := Options***REMOVED******REMOVED***.Apply(Options***REMOVED***Iterations: null.IntFrom(1234)***REMOVED***)
		assert.True(t, opts.Iterations.Valid)
		assert.Equal(t, int64(1234), opts.Iterations.Int64)
	***REMOVED***)
	t.Run("Stages", func(t *testing.T) ***REMOVED***
		t.Parallel()
		opts := Options***REMOVED******REMOVED***.Apply(Options***REMOVED***Stages: []Stage***REMOVED***
			***REMOVED***Duration: types.NullDurationFrom(1 * time.Second), Target: null.IntFrom(10)***REMOVED***,
			***REMOVED***Duration: types.NullDurationFrom(2 * time.Second), Target: null.IntFrom(20)***REMOVED***,
		***REMOVED******REMOVED***)
		assert.NotNil(t, opts.Stages)
		assert.Len(t, opts.Stages, 2)
		assert.Equal(t, 1*time.Second, opts.Stages[0].Duration.TimeDuration())
		assert.Equal(t, int64(10), opts.Stages[0].Target.Int64)
		assert.Equal(t, 2*time.Second, opts.Stages[1].Duration.TimeDuration())
		assert.Equal(t, int64(20), opts.Stages[1].Target.Int64)

		emptyStages := []Stage***REMOVED******REMOVED***
		assert.Equal(t, emptyStages, Options***REMOVED******REMOVED***.Apply(Options***REMOVED***Stages: []Stage***REMOVED******REMOVED******REMOVED******REMOVED******REMOVED***).Stages)
		assert.Equal(t, emptyStages, Options***REMOVED******REMOVED***.Apply(Options***REMOVED***Stages: []Stage***REMOVED******REMOVED******REMOVED***).Stages)
		assert.Equal(t, emptyStages, opts.Apply(Options***REMOVED***Stages: []Stage***REMOVED******REMOVED******REMOVED***).Stages)
		assert.Equal(t, emptyStages, opts.Apply(Options***REMOVED***Stages: []Stage***REMOVED******REMOVED******REMOVED******REMOVED******REMOVED***).Stages)

		assert.Equal(t, opts.Stages, opts.Apply(opts).Stages)

		oneStage := []Stage***REMOVED******REMOVED***Duration: types.NullDurationFrom(5 * time.Second), Target: null.IntFrom(50)***REMOVED******REMOVED***
		assert.Equal(t, oneStage, opts.Apply(Options***REMOVED***Stages: oneStage***REMOVED***).Stages)
		assert.Equal(t, oneStage, Options***REMOVED******REMOVED***.Apply(opts).Apply(Options***REMOVED***Stages: oneStage***REMOVED***).Apply(Options***REMOVED***Stages: oneStage***REMOVED***).Stages)
	***REMOVED***)
	// Execution overwriting is tested by the config consolidation test in cmd
	t.Run("RPS", func(t *testing.T) ***REMOVED***
		t.Parallel()
		opts := Options***REMOVED******REMOVED***.Apply(Options***REMOVED***RPS: null.IntFrom(12345)***REMOVED***)
		assert.True(t, opts.RPS.Valid)
		assert.Equal(t, int64(12345), opts.RPS.Int64)
	***REMOVED***)
	t.Run("MaxRedirects", func(t *testing.T) ***REMOVED***
		t.Parallel()
		opts := Options***REMOVED******REMOVED***.Apply(Options***REMOVED***MaxRedirects: null.IntFrom(12345)***REMOVED***)
		assert.True(t, opts.MaxRedirects.Valid)
		assert.Equal(t, int64(12345), opts.MaxRedirects.Int64)
	***REMOVED***)
	t.Run("UserAgent", func(t *testing.T) ***REMOVED***
		t.Parallel()
		opts := Options***REMOVED******REMOVED***.Apply(Options***REMOVED***UserAgent: null.StringFrom("foo")***REMOVED***)
		assert.True(t, opts.UserAgent.Valid)
		assert.Equal(t, "foo", opts.UserAgent.String)
	***REMOVED***)
	t.Run("Batch", func(t *testing.T) ***REMOVED***
		t.Parallel()
		opts := Options***REMOVED******REMOVED***.Apply(Options***REMOVED***Batch: null.IntFrom(12345)***REMOVED***)
		assert.True(t, opts.Batch.Valid)
		assert.Equal(t, int64(12345), opts.Batch.Int64)
	***REMOVED***)
	t.Run("BatchPerHost", func(t *testing.T) ***REMOVED***
		t.Parallel()
		opts := Options***REMOVED******REMOVED***.Apply(Options***REMOVED***BatchPerHost: null.IntFrom(12345)***REMOVED***)
		assert.True(t, opts.BatchPerHost.Valid)
		assert.Equal(t, int64(12345), opts.BatchPerHost.Int64)
	***REMOVED***)
	t.Run("HTTPDebug", func(t *testing.T) ***REMOVED***
		t.Parallel()
		opts := Options***REMOVED******REMOVED***.Apply(Options***REMOVED***HTTPDebug: null.StringFrom("foo")***REMOVED***)
		assert.True(t, opts.HTTPDebug.Valid)
		assert.Equal(t, "foo", opts.HTTPDebug.String)
	***REMOVED***)
	t.Run("InsecureSkipTLSVerify", func(t *testing.T) ***REMOVED***
		t.Parallel()
		opts := Options***REMOVED******REMOVED***.Apply(Options***REMOVED***InsecureSkipTLSVerify: null.BoolFrom(true)***REMOVED***)
		assert.True(t, opts.InsecureSkipTLSVerify.Valid)
		assert.True(t, opts.InsecureSkipTLSVerify.Bool)
	***REMOVED***)
	t.Run("TLSCipherSuites", func(t *testing.T) ***REMOVED***
		t.Parallel()
		for suiteName, suiteID := range SupportedTLSCipherSuites ***REMOVED***
			t.Run(suiteName, func(t *testing.T) ***REMOVED***
				t.Parallel()
				opts := Options***REMOVED******REMOVED***.Apply(Options***REMOVED***TLSCipherSuites: &TLSCipherSuites***REMOVED***suiteID***REMOVED******REMOVED***)

				assert.NotNil(t, opts.TLSCipherSuites)
				assert.Len(t, *(opts.TLSCipherSuites), 1)
				assert.Equal(t, suiteID, (*opts.TLSCipherSuites)[0])
			***REMOVED***)
		***REMOVED***

		t.Run("JSON", func(t *testing.T) ***REMOVED***
			t.Parallel()
			t.Run("String", func(t *testing.T) ***REMOVED***
				t.Parallel()
				var opts Options
				jsonStr := `***REMOVED***"tlsCipherSuites":["TLS_ECDHE_RSA_WITH_RC4_128_SHA"]***REMOVED***`
				assert.NoError(t, json.Unmarshal([]byte(jsonStr), &opts))
				assert.Equal(t, &TLSCipherSuites***REMOVED***tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA***REMOVED***, opts.TLSCipherSuites)

				t.Run("Roundtrip", func(t *testing.T) ***REMOVED***
					t.Parallel()
					data, err := json.Marshal(opts.TLSCipherSuites)
					assert.NoError(t, err)
					assert.Equal(t, `["TLS_ECDHE_RSA_WITH_RC4_128_SHA"]`, string(data))
					var vers2 TLSCipherSuites
					assert.NoError(t, json.Unmarshal(data, &vers2))
					assert.Equal(t, &vers2, opts.TLSCipherSuites)
				***REMOVED***)
			***REMOVED***)
			t.Run("Not a string", func(t *testing.T) ***REMOVED***
				t.Parallel()
				var opts Options
				jsonStr := `***REMOVED***"tlsCipherSuites":[1.2]***REMOVED***`
				assert.Error(t, json.Unmarshal([]byte(jsonStr), &opts))
			***REMOVED***)
			t.Run("Unknown cipher", func(t *testing.T) ***REMOVED***
				t.Parallel()
				var opts Options
				jsonStr := `***REMOVED***"tlsCipherSuites":["foo"]***REMOVED***`
				assert.Error(t, json.Unmarshal([]byte(jsonStr), &opts))
			***REMOVED***)
		***REMOVED***)
	***REMOVED***)
	t.Run("TLSVersion", func(t *testing.T) ***REMOVED***
		t.Parallel()
		versions := TLSVersions***REMOVED***Min: tls.VersionSSL30, Max: tls.VersionTLS12***REMOVED***
		opts := Options***REMOVED******REMOVED***.Apply(Options***REMOVED***TLSVersion: &versions***REMOVED***)

		assert.NotNil(t, opts.TLSVersion)
		assert.Equal(t, opts.TLSVersion.Min, TLSVersion(tls.VersionSSL30))
		assert.Equal(t, opts.TLSVersion.Max, TLSVersion(tls.VersionTLS12))

		t.Run("JSON", func(t *testing.T) ***REMOVED***
			t.Parallel()
			t.Run("Object", func(t *testing.T) ***REMOVED***
				t.Parallel()
				var opts Options
				jsonStr := `***REMOVED***"tlsVersion":***REMOVED***"min":"tls1.0","max":"tls1.2"***REMOVED******REMOVED***`
				assert.NoError(t, json.Unmarshal([]byte(jsonStr), &opts))
				assert.Equal(t, &TLSVersions***REMOVED***
					Min: TLSVersion(tls.VersionTLS10),
					Max: TLSVersion(tls.VersionTLS12),
				***REMOVED***, opts.TLSVersion)

				t.Run("Roundtrip", func(t *testing.T) ***REMOVED***
					t.Parallel()
					data, err := json.Marshal(opts.TLSVersion)
					assert.NoError(t, err)
					assert.Equal(t, `***REMOVED***"min":"tls1.0","max":"tls1.2"***REMOVED***`, string(data))
					var vers2 TLSVersions
					assert.NoError(t, json.Unmarshal(data, &vers2))
					assert.Equal(t, &vers2, opts.TLSVersion)
				***REMOVED***)
			***REMOVED***)
			t.Run("String", func(t *testing.T) ***REMOVED***
				t.Parallel()
				var opts Options
				jsonStr := `***REMOVED***"tlsVersion":"tls1.2"***REMOVED***`
				assert.NoError(t, json.Unmarshal([]byte(jsonStr), &opts))
				assert.Equal(t, &TLSVersions***REMOVED***
					Min: TLSVersion(tls.VersionTLS12),
					Max: TLSVersion(tls.VersionTLS12),
				***REMOVED***, opts.TLSVersion)
			***REMOVED***)
			t.Run("Blank", func(t *testing.T) ***REMOVED***
				t.Parallel()
				var opts Options
				jsonStr := `***REMOVED***"tlsVersion":""***REMOVED***`
				assert.NoError(t, json.Unmarshal([]byte(jsonStr), &opts))
				assert.Equal(t, &TLSVersions***REMOVED******REMOVED***, opts.TLSVersion)
			***REMOVED***)
			t.Run("Not a string", func(t *testing.T) ***REMOVED***
				t.Parallel()
				var opts Options
				jsonStr := `***REMOVED***"tlsVersion":1.2***REMOVED***`
				assert.Error(t, json.Unmarshal([]byte(jsonStr), &opts))
			***REMOVED***)
			t.Run("Unsupported version", func(t *testing.T) ***REMOVED***
				t.Parallel()
				var opts Options
				jsonStr := `***REMOVED***"tlsVersion":"-1"***REMOVED***`
				assert.Error(t, json.Unmarshal([]byte(jsonStr), &opts))

				jsonStr = `***REMOVED***"tlsVersion":"ssl3.0"***REMOVED***`
				assert.Error(t, json.Unmarshal([]byte(jsonStr), &opts))
			***REMOVED***)
		***REMOVED***)
	***REMOVED***)
	t.Run("TLSAuth", func(t *testing.T) ***REMOVED***
		t.Parallel()
		tlsAuth := []*TLSAuth***REMOVED***
			***REMOVED***TLSAuthFields***REMOVED***
				Domains: []string***REMOVED***"example.com", "*.example.com"***REMOVED***,
				Cert: "-----BEGIN CERTIFICATE-----\n" +
					"MIIBoTCCAUegAwIBAgIUQl0J1Gkd6U2NIMwMDnpfH8c1myEwCgYIKoZIzj0EAwIw\n" +
					"EDEOMAwGA1UEAxMFTXkgQ0EwHhcNMTcwODE1MTYxODAwWhcNMTgwODE1MTYxODAw\n" +
					"WjAQMQ4wDAYDVQQDEwV1c2VyMTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABLaf\n" +
					"xEOmBHkzbqd9/0VZX/39qO2yQq2Gz5faRdvy38kuLMCV+9HYrfMx6GYCZzTUIq6h\n" +
					"8QXOrlgYTixuUVfhJNWjfzB9MA4GA1UdDwEB/wQEAwIFoDAdBgNVHSUEFjAUBggr\n" +
					"BgEFBQcDAQYIKwYBBQUHAwIwDAYDVR0TAQH/BAIwADAdBgNVHQ4EFgQUxmQiq5K3\n" +
					"KUnVME945Byt3Ysvkh8wHwYDVR0jBBgwFoAU3qEhcpRgpsqo9V+LFns9a+oZIYww\n" +
					"CgYIKoZIzj0EAwIDSAAwRQIgSGxnJ+/cLUNTzt7fhr/mjJn7ShsTW33dAdfLM7H2\n" +
					"z/gCIQDyVf8DePtxlkMBScTxZmIlMQdNc6+6VGZQ4QscruVLmg==\n" +
					"-----END CERTIFICATE-----",
				Key: "-----BEGIN EC PRIVATE KEY-----\n" +
					"MHcCAQEEIAfJeoc+XgcqmYV0b4owmofx0LXwPRqOPXMO+PUKxZSgoAoGCCqGSM49\n" +
					"AwEHoUQDQgAEtp/EQ6YEeTNup33/RVlf/f2o7bJCrYbPl9pF2/LfyS4swJX70dit\n" +
					"8zHoZgJnNNQirqHxBc6uWBhOLG5RV+Ek1Q==\n" +
					"-----END EC PRIVATE KEY-----",
			***REMOVED***, nil***REMOVED***,
			***REMOVED***TLSAuthFields***REMOVED***
				Domains: []string***REMOVED***"sub.example.com"***REMOVED***,
				Cert: "-----BEGIN CERTIFICATE-----\n" +
					"MIIBojCCAUegAwIBAgIUWMpVQhmGoLUDd2x6XQYoOOV6C9AwCgYIKoZIzj0EAwIw\n" +
					"EDEOMAwGA1UEAxMFTXkgQ0EwHhcNMTcwODE1MTYxODAwWhcNMTgwODE1MTYxODAw\n" +
					"WjAQMQ4wDAYDVQQDEwV1c2VyMTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABBfF\n" +
					"85gu8fDbNGNlsrtnO+4HvuiP4IXA041jjGczD5kUQ8aihS7hg81tSrLNd1jgxkkv\n" +
					"Po+3TQjzniysiunG3iKjfzB9MA4GA1UdDwEB/wQEAwIFoDAdBgNVHSUEFjAUBggr\n" +
					"BgEFBQcDAQYIKwYBBQUHAwIwDAYDVR0TAQH/BAIwADAdBgNVHQ4EFgQUU0JfPCQb\n" +
					"2YpQZV4j1yiRXBa7J64wHwYDVR0jBBgwFoAU3qEhcpRgpsqo9V+LFns9a+oZIYww\n" +
					"CgYIKoZIzj0EAwIDSQAwRgIhANYDaM18sXAdkjybHccH8xTbBWUNpOYvoHhrGW32\n" +
					"Ov9JAiEA7QKGpm07tQl8p+t7UsOgZu132dHNZUtfgp1bjWfcapU=\n" +
					"-----END CERTIFICATE-----",
				Key: "-----BEGIN EC PRIVATE KEY-----\n" +
					"MHcCAQEEINVilD5qOBkSy+AYfd41X0QPB5N3Z6OzgoBj8FZmSJOFoAoGCCqGSM49\n" +
					"AwEHoUQDQgAEF8XzmC7x8Ns0Y2Wyu2c77ge+6I/ghcDTjWOMZzMPmRRDxqKFLuGD\n" +
					"zW1Kss13WODGSS8+j7dNCPOeLKyK6cbeIg==\n" +
					"-----END EC PRIVATE KEY-----",
			***REMOVED***, nil***REMOVED***,
		***REMOVED***
		opts := Options***REMOVED******REMOVED***.Apply(Options***REMOVED***TLSAuth: tlsAuth***REMOVED***)
		assert.Equal(t, tlsAuth, opts.TLSAuth)

		t.Run("Roundtrip", func(t *testing.T) ***REMOVED***
			t.Parallel()
			optsData, err := json.Marshal(opts)
			assert.NoError(t, err)

			var opts2 Options
			assert.NoError(t, json.Unmarshal(optsData, &opts2))
			if assert.Len(t, opts2.TLSAuth, len(opts.TLSAuth)) ***REMOVED***
				for i := 0; i < len(opts2.TLSAuth); i++ ***REMOVED***
					assert.Equal(t, opts.TLSAuth[i].TLSAuthFields, opts2.TLSAuth[i].TLSAuthFields)
					cert, err := opts2.TLSAuth[i].Certificate()
					assert.NoError(t, err)
					assert.NotNil(t, cert)
				***REMOVED***
			***REMOVED***
		***REMOVED***)

		t.Run("Invalid JSON", func(t *testing.T) ***REMOVED***
			t.Parallel()
			var opts Options
			jsonStr := `***REMOVED***"tlsAuth":["invalid"]***REMOVED***`
			assert.Error(t, json.Unmarshal([]byte(jsonStr), &opts))
		***REMOVED***)

		t.Run("Certificate error", func(t *testing.T) ***REMOVED***
			t.Parallel()
			var opts Options
			jsonStr := `***REMOVED***"tlsAuth":[***REMOVED***"Cert":""***REMOVED***]***REMOVED***`
			assert.Error(t, json.Unmarshal([]byte(jsonStr), &opts))
		***REMOVED***)
	***REMOVED***)
	t.Run("TLSAuth with", func(t *testing.T) ***REMOVED***
		t.Parallel()
		domains := []string***REMOVED***"example.com", "*.example.com"***REMOVED***
		cert := "-----BEGIN CERTIFICATE-----\n" +
			"MIIBoTCCAUegAwIBAgIUQl0J1Gkd6U2NIMwMDnpfH8c1myEwCgYIKoZIzj0EAwIw\n" +
			"EDEOMAwGA1UEAxMFTXkgQ0EwHhcNMTcwODE1MTYxODAwWhcNMTgwODE1MTYxODAw\n" +
			"WjAQMQ4wDAYDVQQDEwV1c2VyMTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABLaf\n" +
			"xEOmBHkzbqd9/0VZX/39qO2yQq2Gz5faRdvy38kuLMCV+9HYrfMx6GYCZzTUIq6h\n" +
			"8QXOrlgYTixuUVfhJNWjfzB9MA4GA1UdDwEB/wQEAwIFoDAdBgNVHSUEFjAUBggr\n" +
			"BgEFBQcDAQYIKwYBBQUHAwIwDAYDVR0TAQH/BAIwADAdBgNVHQ4EFgQUxmQiq5K3\n" +
			"KUnVME945Byt3Ysvkh8wHwYDVR0jBBgwFoAU3qEhcpRgpsqo9V+LFns9a+oZIYww\n" +
			"CgYIKoZIzj0EAwIDSAAwRQIgSGxnJ+/cLUNTzt7fhr/mjJn7ShsTW33dAdfLM7H2\n" +
			"z/gCIQDyVf8DePtxlkMBScTxZmIlMQdNc6+6VGZQ4QscruVLmg==\n" +
			"-----END CERTIFICATE-----"
		tests := []struct ***REMOVED***
			name         string
			privateKey   string
			password     string
			hasError     bool
			errorMessage string
		***REMOVED******REMOVED***
			***REMOVED***
				name: "encrypted key and invalid password",
				privateKey: "-----BEGIN EC PRIVATE KEY-----\n" +
					"Proc-Type: 4,ENCRYPTED\n" +
					"DEK-Info: AES-256-CBC,DF2445CBFE2E5B112FB2B721063757E5\n" +
					"o/VKNZjQcRM2hatqUkQ0dTolL7i2i5hJX9XYsl+TMsq8ZkC83uY/JdR986QS+W2c\n" +
					"EoQGtVGVeL0KGvGpzjTX3YAKXM7Lg5btAeS8GvJ9S7YFd8s0q1pqDdffl2RyjJav\n" +
					"t1jx6XvLu2nBrOUARvHqjkkJQCTdRf2a34GJdbZqE+4=\n" +
					"-----END EC PRIVATE KEY-----",
				password:     "iZfYGcrgFHOg4nweEo7ufT",
				hasError:     true,
				errorMessage: "x509: decryption password incorrect",
			***REMOVED***,
			***REMOVED***
				name: "encrypted key and valid password",
				privateKey: "-----BEGIN EC PRIVATE KEY-----\n" +
					"Proc-Type: 4,ENCRYPTED\n" +
					"DEK-Info: AES-256-CBC,DF2445CBFE2E5B112FB2B721063757E5\n" +
					"o/VKNZjQcRM2hatqUkQ0dTolL7i2i5hJX9XYsl+TMsq8ZkC83uY/JdR986QS+W2c\n" +
					"EoQGtVGVeL0KGvGpzjTX3YAKXM7Lg5btAeS8GvJ9S7YFd8s0q1pqDdffl2RyjJav\n" +
					"t1jx6XvLu2nBrOUARvHqjkkJQCTdRf2a34GJdbZqE+4=\n" +
					"-----END EC PRIVATE KEY-----",
				password:     "12345",
				hasError:     false,
				errorMessage: "",
			***REMOVED***,
			***REMOVED***
				name: "encrypted pks8 format key and valid password",
				privateKey: "-----BEGIN ENCRYPTED PRIVATE KEY-----\n" +
					"MIHsMFcGCSqGSIb3DQEFDTBKMCkGCSqGSIb3DQEFDDAcBAjcfarGfrRgUgICCAAw\n" +
					"DAYIKoZIhvcNAgkFADAdBglghkgBZQMEASoEEFmtmKEFmThbkbpxmC6iBvoEgZCE\n" +
					"pDCpH/yCLmSpjdi/PC74I794nzHyCWf/oS0JhM0Q7J+abZP+p5pnreKft1f15Dbw\n" +
					"QG9alfoM6EffJcVo3gf1tgQrpGGFMwczc4VhQgSGDy0XjZSbd2K0QCFGSmD2ZIR1\n" +
					"qPG3WepWjKmIsYffGeKZx+FjXHSFeGk7RnssNAyKcPruDQIdWWyXxX1+ugBKuBw=\n" +
					"-----END ENCRYPTED PRIVATE KEY-----\n",
				password:     "12345",
				hasError:     true,
				errorMessage: "encrypted pkcs8 formatted key is not supported",
			***REMOVED***,
			***REMOVED***
				name: "non encrypted key and password",
				privateKey: "-----BEGIN EC PRIVATE KEY-----\n" +
					"MHcCAQEEINVilD5qOBkSy+AYfd41X0QPB5N3Z6OzgoBj8FZmSJOFoAoGCCqGSM49\n" +
					"AwEHoUQDQgAEF8XzmC7x8Ns0Y2Wyu2c77ge+6I/ghcDTjWOMZzMPmRRDxqKFLuGD\n" +
					"zW1Kss13WODGSS8+j7dNCPOeLKyK6cbeIg==\n" +
					"-----END EC PRIVATE KEY-----",
				password:     "12345",
				hasError:     true,
				errorMessage: "x509: no DEK-Info header in block",
			***REMOVED***,
		***REMOVED***
		for _, tc := range tests ***REMOVED***
			t.Run(tc.name, func(t *testing.T) ***REMOVED***
				tlsAuth := []*TLSAuth***REMOVED***
					***REMOVED***TLSAuthFields***REMOVED***
						Domains:  domains,
						Cert:     cert,
						Key:      tc.privateKey,
						Password: null.StringFrom(tc.password),
					***REMOVED***, nil***REMOVED***,
				***REMOVED***
				opts := Options***REMOVED******REMOVED***.Apply(Options***REMOVED***TLSAuth: tlsAuth***REMOVED***)
				assert.Equal(t, tlsAuth, opts.TLSAuth)

				t.Run("Roundtrip", func(t *testing.T) ***REMOVED***
					optsData, err := json.Marshal(opts)
					assert.NoError(t, err)

					var opts2 Options
					err = json.Unmarshal(optsData, &opts2)
					if tc.hasError ***REMOVED***
						assert.Error(t, err)
						assert.Contains(t, err.Error(), tc.errorMessage)
					***REMOVED*** else ***REMOVED***
						assert.NoError(t, err)
					***REMOVED***
				***REMOVED***)
			***REMOVED***)
		***REMOVED***
	***REMOVED***)
	t.Run("NoConnectionReuse", func(t *testing.T) ***REMOVED***
		t.Parallel()
		opts := Options***REMOVED******REMOVED***.Apply(Options***REMOVED***NoConnectionReuse: null.BoolFrom(true)***REMOVED***)
		assert.True(t, opts.NoConnectionReuse.Valid)
		assert.True(t, opts.NoConnectionReuse.Bool)
	***REMOVED***)
	t.Run("NoVUConnectionReuse", func(t *testing.T) ***REMOVED***
		t.Parallel()
		opts := Options***REMOVED******REMOVED***.Apply(Options***REMOVED***NoVUConnectionReuse: null.BoolFrom(true)***REMOVED***)
		assert.True(t, opts.NoVUConnectionReuse.Valid)
		assert.True(t, opts.NoVUConnectionReuse.Bool)
	***REMOVED***)
	t.Run("NoCookiesReset", func(t *testing.T) ***REMOVED***
		t.Parallel()
		opts := Options***REMOVED******REMOVED***.Apply(Options***REMOVED***NoCookiesReset: null.BoolFrom(true)***REMOVED***)
		assert.True(t, opts.NoCookiesReset.Valid)
		assert.True(t, opts.NoCookiesReset.Bool)
	***REMOVED***)
	t.Run("BlacklistIPs", func(t *testing.T) ***REMOVED***
		t.Parallel()
		opts := Options***REMOVED******REMOVED***.Apply(Options***REMOVED***
			BlacklistIPs: []*IPNet***REMOVED******REMOVED***
				IPNet: net.IPNet***REMOVED***
					IP:   net.IPv4bcast,
					Mask: net.CIDRMask(31, 32),
				***REMOVED***,
			***REMOVED******REMOVED***,
		***REMOVED***)
		assert.NotNil(t, opts.BlacklistIPs)
		assert.NotEmpty(t, opts.BlacklistIPs)
		assert.Equal(t, net.IPv4bcast, opts.BlacklistIPs[0].IP)
		assert.Equal(t, net.CIDRMask(31, 32), opts.BlacklistIPs[0].Mask)

		t.Run("JSON", func(t *testing.T) ***REMOVED***
			t.Parallel()

			b, err := json.Marshal(opts)
			require.NoError(t, err)

			var uopts Options
			err = json.Unmarshal(b, &uopts)
			require.NoError(t, err)
			require.Len(t, uopts.BlacklistIPs, 1)
			require.Equal(t, "255.255.255.254/31", uopts.BlacklistIPs[0].String())
		***REMOVED***)
	***REMOVED***)
	t.Run("BlockedHostnames", func(t *testing.T) ***REMOVED***
		t.Parallel()
		blockedHostnames, err := types.NewNullHostnameTrie([]string***REMOVED***"test.k6.io", "*valid.pattern"***REMOVED***)
		require.NoError(t, err)
		opts := Options***REMOVED******REMOVED***.Apply(Options***REMOVED***BlockedHostnames: blockedHostnames***REMOVED***)
		assert.NotNil(t, opts.BlockedHostnames)
		assert.Equal(t, blockedHostnames, opts.BlockedHostnames)
	***REMOVED***)

	t.Run("Hosts", func(t *testing.T) ***REMOVED***
		t.Parallel()
		host, err := NewHostAddress(net.ParseIP("192.0.2.1"), "80")
		assert.NoError(t, err)

		opts := Options***REMOVED******REMOVED***.Apply(Options***REMOVED***Hosts: map[string]*HostAddress***REMOVED***
			"test.loadimpact.com": host,
		***REMOVED******REMOVED***)
		assert.NotNil(t, opts.Hosts)
		assert.NotEmpty(t, opts.Hosts)
		assert.Equal(t, "192.0.2.1:80", opts.Hosts["test.loadimpact.com"].String())
	***REMOVED***)

	t.Run("Throws", func(t *testing.T) ***REMOVED***
		t.Parallel()
		opts := Options***REMOVED******REMOVED***.Apply(Options***REMOVED***Throw: null.BoolFrom(true)***REMOVED***)
		assert.True(t, opts.Throw.Valid)
		assert.Equal(t, true, opts.Throw.Bool)
	***REMOVED***)

	t.Run("Thresholds", func(t *testing.T) ***REMOVED***
		t.Parallel()
		opts := Options***REMOVED******REMOVED***.Apply(Options***REMOVED***Thresholds: map[string]workerMetrics.Thresholds***REMOVED***
			"metric": ***REMOVED***
				Thresholds: []*workerMetrics.Threshold***REMOVED******REMOVED******REMOVED******REMOVED***,
			***REMOVED***,
		***REMOVED******REMOVED***)
		assert.NotNil(t, opts.Thresholds)
		assert.NotEmpty(t, opts.Thresholds)
	***REMOVED***)
	t.Run("External", func(t *testing.T) ***REMOVED***
		t.Parallel()
		ext := map[string]json.RawMessage***REMOVED***"a": json.RawMessage("1")***REMOVED***
		opts := Options***REMOVED******REMOVED***.Apply(Options***REMOVED***External: ext***REMOVED***)
		assert.Equal(t, ext, opts.External)
	***REMOVED***)

	t.Run("JSON", func(t *testing.T) ***REMOVED***
		t.Parallel()
		data, err := json.Marshal(Options***REMOVED******REMOVED***)
		assert.NoError(t, err)
		var opts Options
		assert.NoError(t, json.Unmarshal(data, &opts))
		assert.Equal(t, Options***REMOVED******REMOVED***, opts)
	***REMOVED***)
	t.Run("SystemTags", func(t *testing.T) ***REMOVED***
		t.Parallel()
		opts := Options***REMOVED******REMOVED***.Apply(Options***REMOVED***SystemTags: workerMetrics.NewSystemTagSet(workerMetrics.TagProto)***REMOVED***)
		assert.NotNil(t, opts.SystemTags)
		assert.NotEmpty(t, opts.SystemTags)
		assert.True(t, opts.SystemTags.Has(workerMetrics.TagProto))

		t.Run("JSON", func(t *testing.T) ***REMOVED***
			t.Parallel()
			t.Run("Array", func(t *testing.T) ***REMOVED***
				t.Parallel()
				var opts Options
				jsonStr := `***REMOVED***"systemTags":["url"]***REMOVED***`
				assert.NoError(t, json.Unmarshal([]byte(jsonStr), &opts))
				assert.Equal(t, *workerMetrics.NewSystemTagSet(workerMetrics.TagURL), *opts.SystemTags)

				t.Run("Roundtrip", func(t *testing.T) ***REMOVED***
					t.Parallel()
					data, err := json.Marshal(opts.SystemTags)
					assert.NoError(t, err)
					assert.Equal(t, `["url"]`, string(data))
					var vers2 workerMetrics.SystemTagSet
					assert.NoError(t, json.Unmarshal(data, &vers2))
					assert.Equal(t, vers2, *opts.SystemTags)
				***REMOVED***)
			***REMOVED***)
			t.Run("Blank", func(t *testing.T) ***REMOVED***
				t.Parallel()
				var opts Options
				jsonStr := `***REMOVED***"systemTags":[]***REMOVED***`
				assert.NoError(t, json.Unmarshal([]byte(jsonStr), &opts))
				assert.Equal(t, workerMetrics.SystemTagSet(0), *opts.SystemTags)
			***REMOVED***)
		***REMOVED***)
	***REMOVED***)
	t.Run("SummaryTrendStats", func(t *testing.T) ***REMOVED***
		t.Parallel()
		stats := []string***REMOVED***"myStat1", "myStat2"***REMOVED***
		opts := Options***REMOVED******REMOVED***.Apply(Options***REMOVED***SummaryTrendStats: stats***REMOVED***)
		assert.Equal(t, stats, opts.SummaryTrendStats)
	***REMOVED***)
	t.Run("RunTags", func(t *testing.T) ***REMOVED***
		t.Parallel()
		tags := map[string]string***REMOVED***"myTag": "hello"***REMOVED***
		opts := Options***REMOVED******REMOVED***.Apply(Options***REMOVED***RunTags: tags***REMOVED***)
		assert.Equal(t, tags, opts.RunTags)
	***REMOVED***)
	t.Run("DiscardResponseBodies", func(t *testing.T) ***REMOVED***
		t.Parallel()
		opts := Options***REMOVED******REMOVED***.Apply(Options***REMOVED***DiscardResponseBodies: null.BoolFrom(true)***REMOVED***)
		assert.True(t, opts.DiscardResponseBodies.Valid)
		assert.True(t, opts.DiscardResponseBodies.Bool)
	***REMOVED***)
	t.Run("ClientIPRanges", func(t *testing.T) ***REMOVED***
		t.Parallel()
		clientIPRanges := types.NullIPPool***REMOVED******REMOVED***
		err := clientIPRanges.UnmarshalText([]byte("129.112.232.12,123.12.0.0/32"))
		require.NoError(t, err)
		opts := Options***REMOVED******REMOVED***.Apply(Options***REMOVED***LocalIPs: clientIPRanges***REMOVED***)
		assert.NotNil(t, opts.LocalIPs)
	***REMOVED***)
***REMOVED***

func TestOptionsEnv(t *testing.T) ***REMOVED***
	t.Parallel()
	mustNullIPPool := func(s string) types.NullIPPool ***REMOVED***
		p := types.NullIPPool***REMOVED******REMOVED***
		err := p.UnmarshalText([]byte(s))
		require.NoError(t, err)
		return p
	***REMOVED***

	testdata := map[struct***REMOVED*** Name, Key string ***REMOVED***]map[string]interface***REMOVED******REMOVED******REMOVED***
		***REMOVED***"Paused", "K6_PAUSED"***REMOVED***: ***REMOVED***
			"":      null.Bool***REMOVED******REMOVED***,
			"true":  null.BoolFrom(true),
			"false": null.BoolFrom(false),
		***REMOVED***,
		***REMOVED***"VUs", "K6_VUS"***REMOVED***: ***REMOVED***
			"":    null.Int***REMOVED******REMOVED***,
			"123": null.IntFrom(123),
		***REMOVED***,
		***REMOVED***"Duration", "K6_DURATION"***REMOVED***: ***REMOVED***
			"":    types.NullDuration***REMOVED******REMOVED***,
			"10s": types.NullDurationFrom(10 * time.Second),
		***REMOVED***,
		***REMOVED***"Iterations", "K6_ITERATIONS"***REMOVED***: ***REMOVED***
			"":    null.Int***REMOVED******REMOVED***,
			"123": null.IntFrom(123),
		***REMOVED***,
		***REMOVED***"Stages", "K6_STAGES"***REMOVED***: ***REMOVED***
			// "": []Stage***REMOVED******REMOVED***,
			"1s": []Stage***REMOVED***
				***REMOVED***
					Duration: types.NullDurationFrom(1 * time.Second),
				***REMOVED***,
			***REMOVED***,
			"1s:100": []Stage***REMOVED***
				***REMOVED***Duration: types.NullDurationFrom(1 * time.Second), Target: null.IntFrom(100)***REMOVED***,
			***REMOVED***,
			"1s,2s:100": []Stage***REMOVED***
				***REMOVED***Duration: types.NullDurationFrom(1 * time.Second)***REMOVED***,
				***REMOVED***Duration: types.NullDurationFrom(2 * time.Second), Target: null.IntFrom(100)***REMOVED***,
			***REMOVED***,
		***REMOVED***,
		***REMOVED***"MaxRedirects", "K6_MAX_REDIRECTS"***REMOVED***: ***REMOVED***
			"":    null.Int***REMOVED******REMOVED***,
			"123": null.IntFrom(123),
		***REMOVED***,
		***REMOVED***"NoSetup", "K6_NO_SETUP"***REMOVED***: ***REMOVED***
			"":      null.Bool***REMOVED******REMOVED***,
			"true":  null.BoolFrom(true),
			"false": null.BoolFrom(false),
		***REMOVED***,
		***REMOVED***"InsecureSkipTLSVerify", "K6_INSECURE_SKIP_TLS_VERIFY"***REMOVED***: ***REMOVED***
			"":      null.Bool***REMOVED******REMOVED***,
			"true":  null.BoolFrom(true),
			"false": null.BoolFrom(false),
		***REMOVED***,
		// TLSCipherSuites
		// TLSVersion
		// TLSAuth
		***REMOVED***"NoConnectionReuse", "K6_NO_CONNECTION_REUSE"***REMOVED***: ***REMOVED***
			"":      null.Bool***REMOVED******REMOVED***,
			"true":  null.BoolFrom(true),
			"false": null.BoolFrom(false),
		***REMOVED***,
		***REMOVED***"NoVUConnectionReuse", "K6_NO_VU_CONNECTION_REUSE"***REMOVED***: ***REMOVED***
			"":      null.Bool***REMOVED******REMOVED***,
			"true":  null.BoolFrom(true),
			"false": null.BoolFrom(false),
		***REMOVED***,
		***REMOVED***"UserAgent", "K6_USER_AGENT"***REMOVED***: ***REMOVED***
			"":    null.String***REMOVED******REMOVED***,
			"Hi!": null.StringFrom("Hi!"),
		***REMOVED***,
		***REMOVED***"LocalIPs", "K6_LOCAL_IPS"***REMOVED***: ***REMOVED***
			"":                 types.NullIPPool***REMOVED******REMOVED***,
			"192.168.220.2":    mustNullIPPool("192.168.220.2"),
			"192.168.220.0/24": mustNullIPPool("192.168.220.0/24"),
		***REMOVED***,
		***REMOVED***"Throw", "K6_THROW"***REMOVED***: ***REMOVED***
			"":      null.Bool***REMOVED******REMOVED***,
			"true":  null.BoolFrom(true),
			"false": null.BoolFrom(false),
		***REMOVED***,
		***REMOVED***"NoCookiesReset", "K6_NO_COOKIES_RESET"***REMOVED***: ***REMOVED***
			"":      null.Bool***REMOVED******REMOVED***,
			"true":  null.BoolFrom(true),
			"false": null.BoolFrom(false),
		***REMOVED***,
		// Thresholds
		// External
	***REMOVED***
	for field, data := range testdata ***REMOVED***
		field, data := field, data
		t.Run(field.Name, func(t *testing.T) ***REMOVED***
			t.Parallel()
			for str, val := range data ***REMOVED***
				str, val := str, val
				t.Run(`"`+str+`"`, func(t *testing.T) ***REMOVED***
					t.Parallel()
					var opts Options
					assert.NoError(t, envconfig.Process("k6", &opts, func(k string) (string, bool) ***REMOVED***
						if k == field.Key ***REMOVED***
							return str, true
						***REMOVED***
						return "", false
					***REMOVED***))
					assert.Equal(t, val, reflect.ValueOf(opts).FieldByName(field.Name).Interface())
				***REMOVED***)
			***REMOVED***
		***REMOVED***)
	***REMOVED***
***REMOVED***

func TestCIDRUnmarshal(t *testing.T) ***REMOVED***
	t.Parallel()
	testData := []struct ***REMOVED***
		input          string
		expectedOutput *IPNet
		expectFailure  bool
	***REMOVED******REMOVED***
		***REMOVED***
			"10.0.0.0/8",
			&IPNet***REMOVED***IPNet: net.IPNet***REMOVED***
				IP:   net.IP***REMOVED***10, 0, 0, 0***REMOVED***,
				Mask: net.IPv4Mask(255, 0, 0, 0),
			***REMOVED******REMOVED***,
			false,
		***REMOVED***,
		***REMOVED***
			"fc00:1234:5678::/48",
			&IPNet***REMOVED***IPNet: net.IPNet***REMOVED***
				IP:   net.ParseIP("fc00:1234:5678::"),
				Mask: net.CIDRMask(48, 128),
			***REMOVED******REMOVED***,
			false,
		***REMOVED***,
		***REMOVED***"10.0.0.0", nil, true***REMOVED***,
		***REMOVED***"fc00:1234:5678::", nil, true***REMOVED***,
		***REMOVED***"fc00::1234::/48", nil, true***REMOVED***,
	***REMOVED***

	for _, data := range testData ***REMOVED***
		data := data
		t.Run(data.input, func(t *testing.T) ***REMOVED***
			t.Parallel()
			actualIPNet := &IPNet***REMOVED******REMOVED***
			err := actualIPNet.UnmarshalText([]byte(data.input))

			if data.expectFailure ***REMOVED***
				require.Error(t, err)
				require.Contains(t, err.Error(), "invalid CIDR address: "+data.input)
			***REMOVED*** else ***REMOVED***
				require.NoError(t, err)
				assert.Equal(t, data.expectedOutput, actualIPNet)
			***REMOVED***
		***REMOVED***)
	***REMOVED***
***REMOVED***

func TestHostAddressUnmarshal(t *testing.T) ***REMOVED***
	t.Parallel()
	testData := []struct ***REMOVED***
		input          string
		expectedOutput *HostAddress
		expectFailure  string
	***REMOVED******REMOVED***
		***REMOVED***
			"1.2.3.4",
			&HostAddress***REMOVED***IP: net.ParseIP("1.2.3.4")***REMOVED***,
			"",
		***REMOVED***,
		***REMOVED***
			"1.2.3.4:80",
			&HostAddress***REMOVED***IP: net.ParseIP("1.2.3.4"), Port: 80***REMOVED***,
			"",
		***REMOVED***,
		***REMOVED***
			"1.2.3.4:asdf",
			nil,
			"strconv.Atoi: parsing \"asdf\": invalid syntax",
		***REMOVED***,
		***REMOVED***
			"2001:0db8:0000:0000:0000:ff00:0042:8329",
			&HostAddress***REMOVED***IP: net.ParseIP("2001:0db8:0000:0000:0000:ff00:0042:8329")***REMOVED***,
			"",
		***REMOVED***,
		***REMOVED***
			"2001:db8::68",
			&HostAddress***REMOVED***IP: net.ParseIP("2001:db8::68")***REMOVED***,
			"",
		***REMOVED***,
		***REMOVED***
			"[2001:db8::68]:80",
			&HostAddress***REMOVED***IP: net.ParseIP("2001:db8::68"), Port: 80***REMOVED***,
			"",
		***REMOVED***,
		***REMOVED***
			"[2001:db8::68]:asdf",
			nil,
			"strconv.Atoi: parsing \"asdf\": invalid syntax",
		***REMOVED***,
	***REMOVED***

	for _, data := range testData ***REMOVED***
		data := data
		t.Run(data.input, func(t *testing.T) ***REMOVED***
			t.Parallel()
			actualHost := &HostAddress***REMOVED******REMOVED***
			err := actualHost.UnmarshalText([]byte(data.input))

			if data.expectFailure != "" ***REMOVED***
				require.EqualError(t, err, data.expectFailure)
			***REMOVED*** else ***REMOVED***
				require.NoError(t, err)
				assert.Equal(t, data.expectedOutput, actualHost)
			***REMOVED***
		***REMOVED***)
	***REMOVED***
***REMOVED***
