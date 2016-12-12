package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ghodss/yaml"
	"github.com/loadimpact/k6/api/v1"
	"github.com/loadimpact/k6/lib"
	"github.com/loadimpact/k6/stats"
	"gopkg.in/guregu/null.v3"
	"gopkg.in/urfave/cli.v1"
	"os"
	"strconv"
)

var commandStatus = cli.Command***REMOVED***
	Name:      "status",
	Usage:     "Looks up the status of a running test",
	ArgsUsage: " ",
	Action:    actionStatus,
	Description: `Status will print the status of a running test to stdout in YAML format.

   Use the global --address/-a flag to specify the host to connect to; the
   default is port 6565 on the local machine.

   Endpoint: /v1/status`,
***REMOVED***

var commandStats = cli.Command***REMOVED***
	Name:      "stats",
	Usage:     "Prints stats for a running test",
	ArgsUsage: "[name]",
	Action:    actionStats,
	Description: `Stats will print metrics about a running test to stdout in YAML format.

   The result is a dictionary of metrics. If a name is specified, only that one
   metric is fetched, otherwise every metric is printed in no particular order.

   Endpoint: /v1/metrics
             /v1/metrics/:id`,
***REMOVED***

var commandScale = cli.Command***REMOVED***
	Name:      "scale",
	Usage:     "Scales a running test",
	ArgsUsage: "vus",
	Flags: []cli.Flag***REMOVED***
		cli.Int64Flag***REMOVED***
			Name:  "max, m",
			Usage: "update the max number of VUs allowed",
		***REMOVED***,
	***REMOVED***,
	Action: actionScale,
	Description: `Scale will change the number of active VUs of a running test.

   It is an error to scale a test beyond vus-max; this is because instantiating
   new VUs is a very expensive operation, which may skew test results if done
   during a running test. Use --max if you want to do this.

   Endpoint: /v1/status`,
***REMOVED***

var commandPause = cli.Command***REMOVED***
	Name:      "pause",
	Usage:     "Pauses a running test",
	ArgsUsage: " ",
	Action:    actionPause,
	Description: `Pause pauses a running test.

   Running VUs will finish their current iterations, then suspend themselves
   until woken by the test's resumption. A sleeping VU will consume no CPU
   cycles, but will still occupy memory.

   Endpoint: /v1/status`,
***REMOVED***

var commandStart = cli.Command***REMOVED***
	Name:      "start",
	Usage:     "Starts a paused test",
	ArgsUsage: " ",
	Action:    actionStart,
	Description: `Start starts a paused test.

   This is the opposite of the pause command, and will do nothing to an already
   running test.

   Endpoint: /v1/status`,
***REMOVED***

func dumpYAML(v interface***REMOVED******REMOVED***) error ***REMOVED***
	bytes, err := yaml.Marshal(v)
	if err != nil ***REMOVED***
		log.WithError(err).Error("Serialization Error")
		return err
	***REMOVED***
	_, _ = os.Stdout.Write(bytes)
	return nil
***REMOVED***

func actionStatus(cc *cli.Context) error ***REMOVED***
	client, err := v1.NewClient(cc.GlobalString("address"))
	if err != nil ***REMOVED***
		log.WithError(err).Error("Couldn't create a client")
		return err
	***REMOVED***

	status, err := client.Status()
	if err != nil ***REMOVED***
		log.WithError(err).Error("Error")
		return err
	***REMOVED***
	return dumpYAML(status)
***REMOVED***

func actionStats(cc *cli.Context) error ***REMOVED***
	client, err := v1.NewClient(cc.GlobalString("address"))
	if err != nil ***REMOVED***
		log.WithError(err).Error("Couldn't create a client")
		return err
	***REMOVED***

	if len(cc.Args()) > 0 ***REMOVED***
		metric, err := client.Metric(cc.Args()[0])
		if err != nil ***REMOVED***
			log.WithError(err).Error("Error")
			return err
		***REMOVED***
		return dumpYAML(metric)
	***REMOVED***

	metricList, err := client.Metrics()
	if err != nil ***REMOVED***
		log.WithError(err).Error("Error")
		return err
	***REMOVED***

	metrics := make(map[string]stats.Metric)
	for _, metric := range metricList ***REMOVED***
		metrics[metric.Name] = metric
	***REMOVED***
	return dumpYAML(metrics)
***REMOVED***

func actionScale(cc *cli.Context) error ***REMOVED***
	args := cc.Args()
	if len(args) != 1 ***REMOVED***
		return cli.NewExitError("Wrong number of arguments!", 1)
	***REMOVED***
	vus, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil ***REMOVED***
		log.WithError(err).Error("Error")
		return err
	***REMOVED***

	client, err := v1.NewClient(cc.GlobalString("address"))
	if err != nil ***REMOVED***
		log.WithError(err).Error("Couldn't create a client")
		return err
	***REMOVED***

	update := lib.Status***REMOVED***VUs: null.IntFrom(vus)***REMOVED***
	if cc.IsSet("max") ***REMOVED***
		update.VUsMax = null.IntFrom(cc.Int64("max"))
	***REMOVED***

	status, err := client.UpdateStatus(update)
	if err != nil ***REMOVED***
		log.WithError(err).Error("Error")
		return err
	***REMOVED***
	return dumpYAML(status)
***REMOVED***

func actionPause(cc *cli.Context) error ***REMOVED***
	client, err := v1.NewClient(cc.GlobalString("address"))
	if err != nil ***REMOVED***
		log.WithError(err).Error("Couldn't create a client")
		return err
	***REMOVED***

	status, err := client.UpdateStatus(lib.Status***REMOVED***Running: null.BoolFrom(false)***REMOVED***)
	if err != nil ***REMOVED***
		log.WithError(err).Error("Error")
		return err
	***REMOVED***
	return dumpYAML(status)
***REMOVED***

func actionStart(cc *cli.Context) error ***REMOVED***
	client, err := v1.NewClient(cc.GlobalString("address"))
	if err != nil ***REMOVED***
		log.WithError(err).Error("Couldn't create a client")
		return err
	***REMOVED***

	status, err := client.UpdateStatus(lib.Status***REMOVED***Running: null.BoolFrom(true)***REMOVED***)
	if err != nil ***REMOVED***
		log.WithError(err).Error("Error")
		return err
	***REMOVED***
	return dumpYAML(status)
***REMOVED***
