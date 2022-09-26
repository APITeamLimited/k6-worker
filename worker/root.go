package worker

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/APITeamLimited/k6-worker/lib"
	"github.com/APITeamLimited/redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

const (
	defaultConfigFileName   = "config.json"
	waitRemoteLoggerTimeout = time.Second * 5
)

// globalFlags contains global config values that apply for all k6 sub-commands.
type globalFlags struct ***REMOVED***
	configFilePath string
	noColor        bool
	address        string
	logOutput      string
	logFormat      string
***REMOVED***

// globalState contains the globalFlags and accessors for most of the global
// process-external state like CLI arguments, env vars, standard input, output
// and error, etc. In practice, most of it is normally accessed through the `os`
// package from the Go stdlib.
//
// We group them here so we can prevent direct access to them from the rest of
// the k6 codebase. This gives us the ability to mock them and have robust and
// easy-to-write integration-like tests to check the k6 end-to-end behavior in
// any simulated conditions.
//
// `newGlobalState()` returns a globalState object with the real `os`
// parameters, while `newGlobalTestState()` can be used in tests to create
// simulated environments.
type globalState struct ***REMOVED***
	ctx context.Context

	fs      afero.Fs
	getwd   func() (string, error)
	args    []string
	envVars map[string]string

	defaultFlags, flags globalFlags

	stdOut, stdErr *consoleWriter
	stdIn          io.Reader

	osExit       func(int)
	signalNotify func(chan<- os.Signal, ...os.Signal)
	signalStop   func(chan<- os.Signal)

	logger         *logrus.Logger
	fallbackLogger *logrus.Logger
***REMOVED***

// Ideally, this should be the only function in the whole codebase where we use
// global variables and functions from the os package. Anywhere else, things
// like os.Stdout, os.Stderr, os.Stdin, os.Getenv(), etc. should be removed and
// the respective properties of globalState used instead.

// Care is needed to prevent leaking system info to malicious actors.

func newGlobalState(ctx context.Context, client *redis.Client, jobId string, workerId string) *globalState ***REMOVED***
	redisStdOut := &consoleWriter***REMOVED***ctx, client, jobId, workerId***REMOVED***
	redisStdErr := &consoleWriter***REMOVED***ctx, client, jobId, workerId***REMOVED***

	envVars := make(map[string]string)

	logger := &logrus.Logger***REMOVED***
		Out:       redisStdOut,
		Formatter: new(logrus.JSONFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.InfoLevel,
	***REMOVED***

	confDir, err := os.UserConfigDir()
	if err != nil ***REMOVED***
		logger.WithError(err).Warn("could not get config directory")
		confDir = ".config"
	***REMOVED***

	defaultFlags := getDefaultFlags(confDir)

	logrus.SetOutput(ioutil.Discard)

	return &globalState***REMOVED***
		ctx:            ctx,
		fs:             afero.NewMemMapFs(),
		getwd:          os.Getwd,
		args:           []string***REMOVED******REMOVED***,
		envVars:        envVars,
		defaultFlags:   defaultFlags,
		flags:          getFlags(defaultFlags, envVars),
		stdOut:         redisStdOut,
		stdErr:         redisStdErr,
		stdIn:          os.Stdin,
		osExit:         os.Exit,
		signalNotify:   signal.Notify,
		signalStop:     signal.Stop,
		logger:         logger,
		fallbackLogger: logger,
	***REMOVED***
***REMOVED***

func (w *consoleWriter) Write(p []byte) (n int, err error) ***REMOVED***
	origLen := len(p)

	// Intercept the write message so can assess log errors parse json
	parsed := make(map[string]interface***REMOVED******REMOVED***)
	if err := json.Unmarshal(p, &parsed); err != nil ***REMOVED***
		return origLen, err
	***REMOVED***

	// Check message level, if error then log error
	if parsed["level"] == "error" ***REMOVED***
		if parsed["error"] != nil ***REMOVED***
			go lib.HandleStringError(w.ctx, w.client, w.jobId, w.workerId, parsed["error"].(string))
		***REMOVED*** else ***REMOVED***
			go lib.HandleStringError(w.ctx, w.client, w.jobId, w.workerId, parsed["msg"].(string))
		***REMOVED***
		return
	***REMOVED***

	go lib.DispatchMessage(w.ctx, w.client, w.jobId, w.workerId, string(p), "CONSOLE")

	return origLen, err
***REMOVED***

func getDefaultFlags(homeFolder string) globalFlags ***REMOVED***
	return globalFlags***REMOVED***
		address:        "localhost:6565",
		configFilePath: filepath.Join(homeFolder, "loadimpact", "k6", defaultConfigFileName),
		logOutput:      "stderr",
	***REMOVED***
***REMOVED***

func getFlags(defaultFlags globalFlags, env map[string]string) globalFlags ***REMOVED***
	result := defaultFlags

	// TODO: add env vars for the rest of the values (after adjusting
	// rootCmdPersistentFlagSet(), of course)

	if val, ok := env["K6_CONFIG"]; ok ***REMOVED***
		result.configFilePath = val
	***REMOVED***
	if val, ok := env["K6_LOG_OUTPUT"]; ok ***REMOVED***
		result.logOutput = val
	***REMOVED***
	if val, ok := env["K6_LOG_FORMAT"]; ok ***REMOVED***
		result.logFormat = val
	***REMOVED***
	if env["K6_NO_COLOR"] != "" ***REMOVED***
		result.noColor = true
	***REMOVED***
	// Support https://no-color.org/, even an empty value should disable the
	// color output from k6.
	if _, ok := env["NO_COLOR"]; ok ***REMOVED***
		result.noColor = true
	***REMOVED***
	return result
***REMOVED***
