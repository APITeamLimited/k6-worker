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

package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/fatih/color"
	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/loadimpact/k6/lib/consts"
)

var BannerColor = color.New(color.FgCyan)

//TODO: remove these global variables
//nolint:gochecknoglobals
var (
	outMutex  = &sync.Mutex***REMOVED******REMOVED***
	stdoutTTY = isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd())
	stderrTTY = isatty.IsTerminal(os.Stderr.Fd()) || isatty.IsCygwinTerminal(os.Stderr.Fd())
	stdout    = &consoleWriter***REMOVED***colorable.NewColorableStdout(), stdoutTTY, outMutex, nil***REMOVED***
	stderr    = &consoleWriter***REMOVED***colorable.NewColorableStderr(), stderrTTY, outMutex, nil***REMOVED***
)

const defaultConfigFileName = "config.json"

//TODO: remove these global variables
//nolint:gochecknoglobals
var defaultConfigFilePath = defaultConfigFileName // Updated with the user's config folder in the init() function below
//nolint:gochecknoglobals
var configFilePath = os.Getenv("K6_CONFIG") // Overridden by `-c`/`--config` flag!

var (
	//TODO: have environment variables for configuring these? hopefully after we move away from global vars though...
	verbose bool
	quiet   bool
	noColor bool
	logFmt  string
	address string
)

// RootCmd represents the base command when called without any subcommands.
var RootCmd = &cobra.Command***REMOVED***
	Use:           "k6",
	Short:         "a next-generation load generator",
	Long:          BannerColor.Sprintf("\n%s", consts.Banner),
	SilenceUsage:  true,
	SilenceErrors: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) ***REMOVED***
		setupLoggers(logFmt)
		if noColor ***REMOVED***
			// TODO: figure out something else... currently, with the wrappers
			// below, we're stripping any colors from the output after we've
			// added them. The problem is that, besides being very inefficient,
			// this actually also strips other special characters from the
			// intended output, like the progressbar formatting ones, which
			// would otherwise be fine (in a TTY).
			//
			// It would be much better if we avoid messing with the output and
			// instead have a parametrized instance of the color library. It
			// will return colored output if colors are enabled and simply
			// return the passed input as-is (i.e. be a noop) if colors are
			// disabled...
			stdout.Writer = colorable.NewNonColorable(os.Stdout)
			stderr.Writer = colorable.NewNonColorable(os.Stderr)
		***REMOVED***
		log.SetOutput(logrus.StandardLogger().Writer())
		logrus.Debugf("k6 version: v%s", consts.FullVersion())
	***REMOVED***,
***REMOVED***

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() ***REMOVED***
	if err := RootCmd.Execute(); err != nil ***REMOVED***
		code := -1
		var logger logrus.FieldLogger = logrus.StandardLogger()
		if e, ok := err.(ExitCode); ok ***REMOVED***
			code = e.Code
			if e.Hint != "" ***REMOVED***
				logger = logger.WithField("hint", e.Hint)
			***REMOVED***
		***REMOVED***
		logger.Error(err)
		os.Exit(code)
	***REMOVED***
***REMOVED***

func rootCmdPersistentFlagSet() *pflag.FlagSet ***REMOVED***
	flags := pflag.NewFlagSet("", pflag.ContinueOnError)
	//TODO: figure out a better way to handle the CLI flags - global variables are not very testable... :/
	flags.BoolVarP(&verbose, "verbose", "v", false, "enable debug logging")
	flags.BoolVarP(&quiet, "quiet", "q", false, "disable progress updates")
	flags.BoolVar(&noColor, "no-color", false, "disable colored output")
	flags.StringVar(&logFmt, "logformat", "", "log output format")
	flags.StringVarP(&address, "address", "a", "localhost:6565", "address for the api server")

	//TODO: Fix... This default value needed, so both CLI flags and environment variables work
	flags.StringVarP(&configFilePath, "config", "c", configFilePath, "JSON config file")
	// And we also need to explicitly set the default value for the usage message here, so things
	// like `K6_CONFIG="blah" k6 run -h` don't produce a weird usage message
	flags.Lookup("config").DefValue = defaultConfigFilePath
	must(cobra.MarkFlagFilename(flags, "config"))
	return flags
***REMOVED***

func init() ***REMOVED***
	confDir, err := os.UserConfigDir()
	if err != nil ***REMOVED***
		logrus.WithError(err).Warn("could not get config directory")
		confDir = ".config"
	***REMOVED***
	defaultConfigFilePath = filepath.Join(
		confDir,
		"loadimpact",
		"k6",
		defaultConfigFileName,
	)

	RootCmd.PersistentFlags().AddFlagSet(rootCmdPersistentFlagSet())
***REMOVED***

// fprintf panics when where's an error writing to the supplied io.Writer
func fprintf(w io.Writer, format string, a ...interface***REMOVED******REMOVED***) (n int) ***REMOVED***
	n, err := fmt.Fprintf(w, format, a...)
	if err != nil ***REMOVED***
		panic(err.Error())
	***REMOVED***
	return n
***REMOVED***

// RawFormatter it does nothing with the message just prints it
type RawFormater struct***REMOVED******REMOVED***

// Format renders a single log entry
func (f RawFormater) Format(entry *logrus.Entry) ([]byte, error) ***REMOVED***
	return append([]byte(entry.Message), '\n'), nil
***REMOVED***

func setupLoggers(logFmt string) ***REMOVED***
	if verbose ***REMOVED***
		logrus.SetLevel(logrus.DebugLevel)
	***REMOVED***
	logrus.SetOutput(stderr)

	switch logFmt ***REMOVED***
	case "raw":
		logrus.SetFormatter(&RawFormater***REMOVED******REMOVED***)
		logrus.Debug("Logger format: RAW")
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter***REMOVED******REMOVED***)
		logrus.Debug("Logger format: JSON")
	default:
		logrus.SetFormatter(&logrus.TextFormatter***REMOVED***ForceColors: stderrTTY, DisableColors: noColor***REMOVED***)
		logrus.Debug("Logger format: TEXT")
	***REMOVED***

***REMOVED***
