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
	"strings"
	"time"

	"github.com/loadimpact/k6/lib"
	"github.com/loadimpact/k6/lib/consts"
	"github.com/loadimpact/k6/lib/types"
	"github.com/loadimpact/k6/stats"
	"github.com/loadimpact/k6/ui"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	null "gopkg.in/guregu/null.v3"
)

var (
	ErrTagEmptyName   = errors.New("Invalid tag, empty name")
	ErrTagEmptyValue  = errors.New("Invalid tag, empty value")
	ErrTagEmptyString = errors.New("Invalid tag, empty string")
)

func optionFlagSet() *pflag.FlagSet ***REMOVED***
	flags := pflag.NewFlagSet("", 0)
	flags.SortFlags = false
	flags.Int64P("vus", "u", 1, "number of virtual users")
	flags.Int64P("max", "m", 0, "max available virtual users")
	flags.DurationP("duration", "d", 0, "test duration limit")
	flags.Int64P("iterations", "i", 0, "script total iteration limit (among all VUs)")
	flags.StringSliceP("stage", "s", nil, "add a `stage`, as `[duration]:[target]`")
	flags.BoolP("paused", "p", false, "start the test in a paused state")
	flags.Int64("max-redirects", 10, "follow at most n redirects")
	flags.Int64("batch", 20, "max parallel batch reqs")
	flags.Int64("batch-per-host", 20, "max parallel batch reqs per host")
	flags.Int64("rps", 0, "limit requests per second")
	flags.String("user-agent", fmt.Sprintf("k6/%s (https://k6.io/)", consts.Version), "user agent for http requests")
	flags.String("http-debug", "", "log all HTTP requests and responses. Excludes body by default. To include body use '--http-debug=full'")
	flags.Lookup("http-debug").NoOptDefVal = "headers"
	flags.Bool("insecure-skip-tls-verify", false, "skip verification of TLS certificates")
	flags.Bool("no-connection-reuse", false, "disable keep-alive connections")
	flags.Bool("no-vu-connection-reuse", false, "don't reuse connections between iterations")
	flags.Duration("min-iteration-duration", 0, "minimum amount of time k6 will take executing a single iteration")
	flags.BoolP("throw", "w", false, "throw warnings (like failed http requests) as errors")
	flags.StringSlice("blacklist-ip", nil, "blacklist an `ip range` from being called")

	// The comment about system-tags also applies for summary-trend-stats. The default values
	// are set in applyDefault().
	sumTrendStatsHelp := fmt.Sprintf(
		"define `stats` for trend metrics (response times), one or more as 'avg,p(95),...' (default '%s')",
		strings.Join(lib.DefaultSummaryTrendStats, ","),
	)
	flags.StringSlice("summary-trend-stats", nil, sumTrendStatsHelp)
	flags.String("summary-time-unit", "", "define the time unit used to display the trend stats. Possible units are: 's', 'ms' and 'us'")
	// system-tags must have a default value, but we can't specify it here, otherwiese, it will always override others.
	// set it to nil here, and add the default in applyDefault() instead.
	systemTagsCliHelpText := fmt.Sprintf(
		"only include these system tags in metrics (default %q)",
		stats.DefaultSystemTagSet.SetString(),
	)
	flags.StringSlice("system-tags", nil, systemTagsCliHelpText)
	flags.StringSlice("tag", nil, "add a `tag` to be applied to all samples, as `[name]=[value]`")
	flags.String("console-output", "", "redirects the console logging to the provided output file")
	flags.Bool("discard-response-bodies", false, "Read but don't process or save HTTP response bodies")
	return flags
***REMOVED***

func getOptions(flags *pflag.FlagSet) (lib.Options, error) ***REMOVED***
	opts := lib.Options***REMOVED***
		VUs:                   getNullInt64(flags, "vus"),
		VUsMax:                getNullInt64(flags, "max"),
		Duration:              getNullDuration(flags, "duration"),
		Iterations:            getNullInt64(flags, "iterations"),
		Paused:                getNullBool(flags, "paused"),
		MaxRedirects:          getNullInt64(flags, "max-redirects"),
		Batch:                 getNullInt64(flags, "batch"),
		BatchPerHost:          getNullInt64(flags, "batch-per-host"),
		RPS:                   getNullInt64(flags, "rps"),
		UserAgent:             getNullString(flags, "user-agent"),
		HTTPDebug:             getNullString(flags, "http-debug"),
		InsecureSkipTLSVerify: getNullBool(flags, "insecure-skip-tls-verify"),
		NoConnectionReuse:     getNullBool(flags, "no-connection-reuse"),
		NoVUConnectionReuse:   getNullBool(flags, "no-vu-connection-reuse"),
		MinIterationDuration:  getNullDuration(flags, "min-iteration-duration"),
		Throw:                 getNullBool(flags, "throw"),
		DiscardResponseBodies: getNullBool(flags, "discard-response-bodies"),
		// Default values for options without CLI flags:
		// TODO: find a saner and more dev-friendly and error-proof way to handle options
		SetupTimeout:    types.NullDuration***REMOVED***Duration: types.Duration(10 * time.Second), Valid: false***REMOVED***,
		TeardownTimeout: types.NullDuration***REMOVED***Duration: types.Duration(10 * time.Second), Valid: false***REMOVED***,

		MetricSamplesBufferSize: null.NewInt(1000, false),
	***REMOVED***

	// Using Changed() because GetStringSlice() doesn't differentiate between empty and no value
	if flags.Changed("stage") ***REMOVED***
		stageStrings, err := flags.GetStringSlice("stage")
		if err != nil ***REMOVED***
			return opts, err
		***REMOVED***
		opts.Stages = []lib.Stage***REMOVED******REMOVED***
		for i, s := range stageStrings ***REMOVED***
			var stage lib.Stage
			if err := stage.UnmarshalText([]byte(s)); err != nil ***REMOVED***
				return opts, errors.Wrapf(err, "stage %d", i)
			***REMOVED***
			if !stage.Duration.Valid ***REMOVED***
				return opts, fmt.Errorf("stage %d doesn't have a specified duration", i)
			***REMOVED***
			opts.Stages = append(opts.Stages, stage)
		***REMOVED***
	***REMOVED***

	if flags.Changed("system-tags") ***REMOVED***
		systemTagList, err := flags.GetStringSlice("system-tags")
		if err != nil ***REMOVED***
			return opts, err
		***REMOVED***
		opts.SystemTags = stats.ToSystemTagSet(systemTagList)
	***REMOVED***

	blacklistIPStrings, err := flags.GetStringSlice("blacklist-ip")
	if err != nil ***REMOVED***
		return opts, err
	***REMOVED***
	for _, s := range blacklistIPStrings ***REMOVED***
		net, parseErr := lib.ParseCIDR(s)
		if parseErr != nil ***REMOVED***
			return opts, errors.Wrap(parseErr, "blacklist-ip")
		***REMOVED***
		opts.BlacklistIPs = append(opts.BlacklistIPs, net)
	***REMOVED***

	if flags.Changed("summary-trend-stats") ***REMOVED***
		trendStats, errSts := flags.GetStringSlice("summary-trend-stats")
		if errSts != nil ***REMOVED***
			return opts, errSts
		***REMOVED***
		if errSts = ui.ValidateSummary(trendStats); err != nil ***REMOVED***
			return opts, errSts
		***REMOVED***
		opts.SummaryTrendStats = trendStats
	***REMOVED***

	summaryTimeUnit, err := flags.GetString("summary-time-unit")
	if err != nil ***REMOVED***
		return opts, err
	***REMOVED***
	if summaryTimeUnit != "" ***REMOVED***
		if summaryTimeUnit != "s" && summaryTimeUnit != "ms" && summaryTimeUnit != "us" ***REMOVED***
			return opts, errors.New("invalid summary time unit. Use: 's', 'ms' or 'us'")
		***REMOVED***
		opts.SummaryTimeUnit = null.StringFrom(summaryTimeUnit)
	***REMOVED***

	runTags, err := flags.GetStringSlice("tag")
	if err != nil ***REMOVED***
		return opts, err
	***REMOVED***

	if len(runTags) > 0 ***REMOVED***
		parsedRunTags := make(map[string]string, len(runTags))
		for i, s := range runTags ***REMOVED***
			name, value, err := parseTagNameValue(s)
			if err != nil ***REMOVED***
				return opts, errors.Wrapf(err, "tag %d", i)
			***REMOVED***
			parsedRunTags[name] = value
		***REMOVED***
		opts.RunTags = stats.IntoSampleTags(&parsedRunTags)
	***REMOVED***

	redirectConFile, err := flags.GetString("console-output")
	if err != nil ***REMOVED***
		return opts, err
	***REMOVED***

	if redirectConFile != "" ***REMOVED***
		opts.ConsoleOutput = null.StringFrom(redirectConFile)
	***REMOVED***

	return opts, nil
***REMOVED***

func parseTagNameValue(nv string) (string, string, error) ***REMOVED***
	if nv == "" ***REMOVED***
		return "", "", ErrTagEmptyString
	***REMOVED***

	idx := strings.IndexRune(nv, '=')

	switch idx ***REMOVED***
	case 0:
		return "", "", ErrTagEmptyName
	case -1, len(nv) - 1:
		return "", "", ErrTagEmptyValue
	default:
		return nv[:idx], nv[idx+1:], nil
	***REMOVED***
***REMOVED***
