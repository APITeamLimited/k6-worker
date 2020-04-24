package consts

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
)

// Version contains the current semantic version of k6.
var Version = "0.26.2" //nolint:gochecknoglobals

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
