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
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/guregu/null.v3"

	"go.k6.io/k6/lib/types"
	"go.k6.io/k6/loader"
)

// Use these when interacting with fs and writing to terminal, makes a command testable
//nolint:gochecknoglobals
var (
	defaultFs               = afero.NewOsFs()
	defaultWriter io.Writer = os.Stdout
)

// Panic if the given error is not nil.
func must(err error) ***REMOVED***
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
***REMOVED***

// TODO: refactor the CLI config so these functions aren't needed - they
// can mask errors by failing only at runtime, not at compile time
func getNullBool(flags *pflag.FlagSet, key string) null.Bool ***REMOVED***
	v, err := flags.GetBool(key)
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	return null.NewBool(v, flags.Changed(key))
***REMOVED***

func getNullInt64(flags *pflag.FlagSet, key string) null.Int ***REMOVED***
	v, err := flags.GetInt64(key)
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	return null.NewInt(v, flags.Changed(key))
***REMOVED***

func getNullDuration(flags *pflag.FlagSet, key string) types.NullDuration ***REMOVED***
	// TODO: use types.ParseExtendedDuration? not sure we should support
	// unitless durations (i.e. milliseconds) here...
	v, err := flags.GetDuration(key)
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	return types.NullDuration***REMOVED***Duration: types.Duration(v), Valid: flags.Changed(key)***REMOVED***
***REMOVED***

func getNullString(flags *pflag.FlagSet, key string) null.String ***REMOVED***
	v, err := flags.GetString(key)
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	return null.NewString(v, flags.Changed(key))
***REMOVED***

func exactArgsWithMsg(n int, msg string) cobra.PositionalArgs ***REMOVED***
	return func(cmd *cobra.Command, args []string) error ***REMOVED***
		if len(args) != n ***REMOVED***
			return fmt.Errorf("accepts %d arg(s), received %d: %s", n, len(args), msg)
		***REMOVED***
		return nil
	***REMOVED***
***REMOVED***

// readSource is a small wrapper around loader.ReadSource returning
// result of the load and filesystems map
func readSource(filename string, logger *logrus.Logger) (*loader.SourceData, map[string]afero.Fs, error) ***REMOVED***
	pwd, err := os.Getwd()
	if err != nil ***REMOVED***
		return nil, nil, err
	***REMOVED***

	filesystems := loader.CreateFilesystems()
	src, err := loader.ReadSource(logger, filename, pwd, filesystems, os.Stdin)
	return src, filesystems, err
***REMOVED***

// TODO: consider moving this out as a method of SourceData ?
func getRunType(src *loader.SourceData) string ***REMOVED***
	typ := runType
	if typ == "" ***REMOVED***
		typ = detectType(src.Data)
	***REMOVED***
	return typ
***REMOVED***

func detectType(data []byte) string ***REMOVED***
	if _, err := tar.NewReader(bytes.NewReader(data)).Next(); err == nil ***REMOVED***
		return typeArchive
	***REMOVED***
	return typeJS
***REMOVED***

// fprintf panics when where's an error writing to the supplied io.Writer
func fprintf(w io.Writer, format string, a ...interface***REMOVED******REMOVED***) (n int) ***REMOVED***
	n, err := fmt.Fprintf(w, format, a...)
	if err != nil ***REMOVED***
		panic(err.Error())
	***REMOVED***
	return n
***REMOVED***
