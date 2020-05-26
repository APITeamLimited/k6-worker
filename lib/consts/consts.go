/*
 *
 * k6 - a next-generation load testing tool
 * Copyright (C) 2019 Load Impact
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

package consts

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
)

// Version contains the current semantic version of k6.
var Version = "0.27.0-dev" //nolint:gochecknoglobals

// VersionDetails can be set externally as part of the build process
var VersionDetails = "" // nolint:gochecknoglobals

// FullVersion returns the maximally full version and build information for
// the currently running k6 executable.
func FullVersion() string ***REMOVED***
	goVersionArch := fmt.Sprintf("%s, %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
	if VersionDetails != "" ***REMOVED***
		return fmt.Sprintf("%s (%s, %s)", Version, VersionDetails, goVersionArch)
	***REMOVED***

	if buildInfo, ok := debug.ReadBuildInfo(); ok ***REMOVED***
		return fmt.Sprintf("%s (%s, %s)", Version, buildInfo.Main.Version, goVersionArch)
	***REMOVED***

	return fmt.Sprintf("%s (dev build, %s)", Version, goVersionArch)
***REMOVED***

// Banner contains the ASCII-art banner with the k6 logo and stylized website URL
// TODO: make these into methods, only the version needs to be a variable
//nolint:gochecknoglobals
var Banner = strings.Join([]string***REMOVED***
	`          /\      |‾‾|  /‾‾/  /‾/   `,
	`     /\  /  \     |  |_/  /  / /    `,
	`    /  \/    \    |      |  /  ‾‾\  `,
	`   /          \   |  |‾\  \ | (_) | `,
	`  / __________ \  |__|  \__\ \___/ .io`,
***REMOVED***, "\n")
