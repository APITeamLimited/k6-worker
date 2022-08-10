package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"go.k6.io/k6/lib/consts"
)

func getCmdVersion(globalState *globalState) *cobra.Command ***REMOVED***
	// versionCmd represents the version command.
	return &cobra.Command***REMOVED***
		Use:   "version",
		Short: "Show application version",
		Long:  `Show the application version and exit.`,
		Run: func(_ *cobra.Command, _ []string) ***REMOVED***
			printToStdout(globalState, fmt.Sprintf("k6 v%s\n", consts.FullVersion()))
		***REMOVED***,
	***REMOVED***
***REMOVED***
