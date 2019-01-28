package http

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"net/url"
	"runtime"
	"syscall"
	"testing"

	"github.com/loadimpact/k6/lib/netext"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/http2"
)

func TestDefaultError(t *testing.T) ***REMOVED***
	testErrorCode(t, defaultErrorCode, fmt.Errorf("random error"))
***REMOVED***

func TestHTTP2Errors(t *testing.T) ***REMOVED***
	var unknownErrorCode = 220
	var connectionError = http2.ConnectionError(unknownErrorCode)
	var testTable = map[errCode]error***REMOVED***
		unknownHTTP2ConnectionErrorCode + 1: new(http2.ConnectionError),
		unknownHTTP2StreamErrorCode + 1:     new(http2.StreamError),
		unknownHTTP2GoAwayErrorCode + 1:     new(http2.GoAwayError),

		unknownHTTP2ConnectionErrorCode: &connectionError,
		unknownHTTP2StreamErrorCode:     &http2.StreamError***REMOVED***Code: 220***REMOVED***,
		unknownHTTP2GoAwayErrorCode:     &http2.GoAwayError***REMOVED***ErrCode: 220***REMOVED***,
	***REMOVED***
	testMapOfErrorCodes(t, testTable)
***REMOVED***

func TestTLSErrors(t *testing.T) ***REMOVED***
	var testTable = map[errCode]error***REMOVED***
		x509UnknownAuthorityErrorCode: new(x509.UnknownAuthorityError),
		x509HostnameErrorCode:         new(x509.HostnameError),
		defaultTLSErrorCode:           new(tls.RecordHeaderError),
	***REMOVED***
	testMapOfErrorCodes(t, testTable)
***REMOVED***

func TestDNSErrors(t *testing.T) ***REMOVED***
	var (
		defaultDNSError = new(net.DNSError)
		noSuchHostError = new(net.DNSError)
	)

	noSuchHostError.Err = "no such host" // defined as private in go stdlib
	var testTable = map[errCode]error***REMOVED***
		defaultDNSErrorCode:    defaultDNSError,
		dnsNoSuchHostErrorCode: noSuchHostError,
	***REMOVED***
	testMapOfErrorCodes(t, testTable)
***REMOVED***

func TestBlackListedIPError(t *testing.T) ***REMOVED***
	var err = netext.BlackListedIPError***REMOVED******REMOVED***
	testErrorCode(t, blackListedIPErrorCode, err)
	var errorCode, errorMsg = errorCodeForError(err)
	require.NotEqual(t, err.Error(), errorMsg)
	require.Equal(t, blackListedIPErrorCode, errorCode)
***REMOVED***

type timeoutError bool

func (t timeoutError) Timeout() bool ***REMOVED***
	return (bool)(t)
***REMOVED***

func (t timeoutError) Error() string ***REMOVED***
	return fmt.Sprintf("%t", t)
***REMOVED***

func TestUnknownNetErrno(t *testing.T) ***REMOVED***
	var err = new(net.OpError)
	err.Op = "write"
	err.Net = "tcp"
	err.Err = syscall.EBFONT // Highly unlikely to actually need to do anything with this error
	var expectedError = fmt.Sprintf(
		"write: unknown errno `%d` on %s with message `%s`",
		syscall.EBFONT, runtime.GOOS, err.Err)
	var errorCode, errorMsg = errorCodeForError(err)
	require.Equal(t, expectedError, errorMsg)
	require.Equal(t, netUnknownErrnoErrorCode, errorCode)
***REMOVED***

func TestTCPErrors(t *testing.T) ***REMOVED***
	var (
		nonTCPError       = &net.OpError***REMOVED***Net: "something", Err: errors.New("non tcp error")***REMOVED***
		econnreset        = &net.OpError***REMOVED***Net: "tcp", Op: "write", Err: syscall.ECONNRESET***REMOVED***
		epipeerror        = &net.OpError***REMOVED***Net: "tcp", Op: "write", Err: syscall.EPIPE***REMOVED***
		econnrefused      = &net.OpError***REMOVED***Net: "tcp", Op: "dial", Err: syscall.ECONNREFUSED***REMOVED***
		tcperror          = &net.OpError***REMOVED***Net: "tcp", Err: errors.New("tcp error")***REMOVED***
		timeoutedError    = &net.OpError***REMOVED***Net: "tcp", Op: "dial", Err: timeoutError(true)***REMOVED***
		notTimeoutedError = &net.OpError***REMOVED***Net: "tcp", Op: "dial", Err: timeoutError(false)***REMOVED***
	)

	var testTable = map[errCode]error***REMOVED***
		defaultNetNonTCPErrorCode: nonTCPError,
		tcpResetByPeerErrorCode:   econnreset,
		tcpBrokenPipeErrorCode:    epipeerror,
		tcpDialRefusedErrorCode:   econnrefused,
		defaultTCPErrorCode:       tcperror,
		tcpDialErrorCode:          notTimeoutedError,
		tcpDialTimeoutErrorCode:   timeoutedError,
	***REMOVED***

	testMapOfErrorCodes(t, testTable)
***REMOVED***

func testErrorCode(t *testing.T, code errCode, err error) ***REMOVED***
	t.Helper()
	result, _ := errorCodeForError(err)
	require.Equalf(t, code, result, "Wrong error code for error `%s`", err)

	result, _ = errorCodeForError(errors.WithStack(err))
	require.Equalf(t, code, result, "Wrong error code for error `%s`", err)

	result, _ = errorCodeForError(&url.Error***REMOVED***Err: err***REMOVED***)
	require.Equalf(t, code, result, "Wrong error code for error `%s`", err)
***REMOVED***

func testMapOfErrorCodes(t *testing.T, testTable map[errCode]error) ***REMOVED***
	t.Helper()
	for code, err := range testTable ***REMOVED***
		testErrorCode(t, code, err)
	***REMOVED***
***REMOVED***
