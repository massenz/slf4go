package main

import (
	"flag"
	"github.com/massenz/slf4go/logging"
)

func main() {
	logging.RootLog.Info("Program started - before any logs are available")
	var log = logging.NewLog("example")
	defer log.Trace("Exiting")

	trace := flag.Bool("trace", false, "If set it will emit trace logs")
	flag.Parse()
	if *trace {
		log.Level = logging.TRACE
	}
	log.Info("An INFO message")
	log.Debug("This will NOT be logged, unless -trace is given")

	if !*trace {
		log.Level = logging.DEBUG
	}
	log.Debug("This WILL be printed out")
	log.Info("The `trace` on exit will only be visible with the -trace option")

	logging.NullLog.Error("No one will ever see this, like, ever 😤")
}
