package actions

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/loadimpact/speedboat/actions/registry"
	"github.com/loadimpact/speedboat/common"
	"github.com/loadimpact/speedboat/master"
	"github.com/loadimpact/speedboat/message"
	"github.com/loadimpact/speedboat/worker"
	"time"
)

func init() ***REMOVED***
	registry.RegisterCommand(cli.Command***REMOVED***
		Name:   "ping",
		Usage:  "Tests master connectivity",
		Action: actionPing,
		Flags: []cli.Flag***REMOVED***
			cli.BoolFlag***REMOVED***
				Name:  "worker",
				Usage: "Pings a worker instead of the master",
			***REMOVED***,
			cli.BoolFlag***REMOVED***
				Name:  "local",
				Usage: "Allow pinging an inproc master/worker",
			***REMOVED***,
			common.MasterHostFlag,
			common.MasterPortFlag,
		***REMOVED***,
	***REMOVED***)
	registry.RegisterMasterProcessor(func(*master.Master) master.Processor ***REMOVED***
		return &PingProcessor***REMOVED******REMOVED***
	***REMOVED***)
	registry.RegisterProcessor(func(*worker.Worker) master.Processor ***REMOVED***
		return &PingProcessor***REMOVED******REMOVED***
	***REMOVED***)
***REMOVED***

// Processes pings, on both master and worker.
type PingProcessor struct***REMOVED******REMOVED***

type PingMessage struct ***REMOVED***
	Time time.Time
***REMOVED***

func (*PingProcessor) Process(msg message.Message) <-chan message.Message ***REMOVED***
	out := make(chan message.Message)

	go func() ***REMOVED***
		defer close(out)
		switch msg.Type ***REMOVED***
		case "ping.ping":
			data := PingMessage***REMOVED******REMOVED***
			if err := msg.Take(&data); err != nil ***REMOVED***
				out <- message.ToClient("error").WithError(err)
				break
			***REMOVED***
			out <- message.ToClient("ping.pong").With(data)
		***REMOVED***
	***REMOVED***()

	return out
***REMOVED***

// Pings a master or specified workers.
func actionPing(c *cli.Context) ***REMOVED***
	client, local := common.MustGetClient(c)
	if local && !c.Bool("local") ***REMOVED***
		log.Fatal("You're about to ping an in-process system, which doesn't make a lot of sense. You probably want to specify --master=..., or use --local if this is actually what you want.")
	***REMOVED***

	in, out := client.Connector.Run()

	topic := message.MasterTopic
	if c.Bool("worker") ***REMOVED***
		topic = message.WorkerTopic
	***REMOVED***
	out <- message.To(topic, "ping.ping").With(PingMessage***REMOVED***
		Time: time.Now(),
	***REMOVED***)

readLoop:
	for msg := range in ***REMOVED***
		switch msg.Type ***REMOVED***
		case "ping.pong":
			data := PingMessage***REMOVED******REMOVED***
			if err := msg.Take(&data); err != nil ***REMOVED***
				log.WithError(err).Error("Couldn't decode pong")
				break
			***REMOVED***
			log.WithField("time", data.Time).Info("Pong!")
			break readLoop
		***REMOVED***
	***REMOVED***
***REMOVED***
