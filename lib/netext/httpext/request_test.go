package httpext

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/loadimpact/k6/lib"
	"github.com/loadimpact/k6/stats"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type reader func([]byte) (int, error)

func (r reader) Read(a []byte) (int, error) ***REMOVED***
	return ((func([]byte) (int, error))(r))(a)
***REMOVED***

const badReadMsg = "bad read error for test"
const badCloseMsg = "bad close error for test"

func badReadBody() io.Reader ***REMOVED***
	return reader(func(_ []byte) (int, error) ***REMOVED***
		return 0, errors.New(badReadMsg)
	***REMOVED***)
***REMOVED***

type closer func() error

func (c closer) Close() error ***REMOVED***
	return ((func() error)(c))()
***REMOVED***

func badCloseBody() io.ReadCloser ***REMOVED***
	return struct ***REMOVED***
		io.Reader
		io.Closer
	***REMOVED******REMOVED***
		Reader: reader(func(_ []byte) (int, error) ***REMOVED***
			return 0, io.EOF
		***REMOVED***),
		Closer: closer(func() error ***REMOVED***
			return errors.New(badCloseMsg)
		***REMOVED***),
	***REMOVED***
***REMOVED***

func TestCompressionBodyError(t *testing.T) ***REMOVED***
	var algos = []CompressionType***REMOVED***CompressionTypeGzip***REMOVED***
	t.Run("bad read body", func(t *testing.T) ***REMOVED***
		_, _, err := compressBody(algos, ioutil.NopCloser(badReadBody()))
		require.Error(t, err)
		require.Equal(t, err.Error(), badReadMsg)
	***REMOVED***)

	t.Run("bad close body", func(t *testing.T) ***REMOVED***
		_, _, err := compressBody(algos, badCloseBody())
		require.Error(t, err)
		require.Equal(t, err.Error(), badCloseMsg)
	***REMOVED***)
***REMOVED***

func TestMakeRequestError(t *testing.T) ***REMOVED***
	var ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	t.Run("bad compression algorithm body", func(t *testing.T) ***REMOVED***
		var req, err = http.NewRequest("GET", "https://wont.be.used", nil)

		require.NoError(t, err)
		var badCompressionType = CompressionType(13)
		require.False(t, badCompressionType.IsACompressionType())
		var preq = &ParsedHTTPRequest***REMOVED***
			Req:          req,
			Body:         new(bytes.Buffer),
			Compressions: []CompressionType***REMOVED***badCompressionType***REMOVED***,
		***REMOVED***
		_, err = MakeRequest(ctx, preq)
		require.Error(t, err)
		require.Equal(t, err.Error(), "unknown compressionType CompressionType(13)")
	***REMOVED***)

	t.Run("invalid upgrade response", func(t *testing.T) ***REMOVED***
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) ***REMOVED***
			w.Header().Add("Connection", "Upgrade")
			w.Header().Add("Upgrade", "h2c")
			w.WriteHeader(http.StatusSwitchingProtocols)
		***REMOVED***))
		defer srv.Close()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		state := &lib.State***REMOVED***
			Options:   lib.Options***REMOVED***RunTags: &stats.SampleTags***REMOVED******REMOVED******REMOVED***,
			Transport: srv.Client().Transport,
		***REMOVED***
		ctx = lib.WithState(ctx, state)
		req, _ := http.NewRequest("GET", srv.URL, nil)
		var preq = &ParsedHTTPRequest***REMOVED***Req: req, URL: &URL***REMOVED***u: req.URL***REMOVED***, Body: new(bytes.Buffer)***REMOVED***

		res, err := MakeRequest(ctx, preq)

		assert.Nil(t, res)
		assert.EqualError(t, err, "unsupported response status: 101 Switching Protocols")
	***REMOVED***)
***REMOVED***

func TestURL(t *testing.T) ***REMOVED***
	t.Run("Clean", func(t *testing.T) ***REMOVED***
		testCases := []struct ***REMOVED***
			url      string
			expected string
		***REMOVED******REMOVED***
			***REMOVED***"https://example.com/", "https://example.com/"***REMOVED***,
			***REMOVED***"https://example.com/$***REMOVED******REMOVED***", "https://example.com/$***REMOVED******REMOVED***"***REMOVED***,
			***REMOVED***"https://user@example.com/", "https://****@example.com/"***REMOVED***,
			***REMOVED***"https://user:pass@example.com/", "https://****:****@example.com/"***REMOVED***,
			***REMOVED***"https://user:pass@example.com/path?a=1&b=2", "https://****:****@example.com/path?a=1&b=2"***REMOVED***,
			***REMOVED***"https://user:pass@example.com/$***REMOVED******REMOVED***/$***REMOVED******REMOVED***", "https://****:****@example.com/$***REMOVED******REMOVED***/$***REMOVED******REMOVED***"***REMOVED***,
			***REMOVED***"@malformed/url", "@malformed/url"***REMOVED***,
			***REMOVED***"not a url", "not a url"***REMOVED***,
		***REMOVED***

		for _, tc := range testCases ***REMOVED***
			tc := tc
			t.Run(tc.url, func(t *testing.T) ***REMOVED***
				u, err := url.Parse(tc.url)
				require.NoError(t, err)
				ut := URL***REMOVED***u: u, URL: tc.url***REMOVED***
				require.Equal(t, tc.expected, ut.Clean())
			***REMOVED***)
		***REMOVED***
	***REMOVED***)
***REMOVED***

func BenchmarkWrapDecompressionError(b *testing.B) ***REMOVED***
	err := errors.New("error")
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ ***REMOVED***
		_ = wrapDecompressionError(err)
	***REMOVED***
***REMOVED***
