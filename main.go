package main

import (
	"flag"
	"fmt"
	daemon "github.com/leprosus/golang-daemon"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var log = logrus.New()

func wrapLogger(module string) logrus.FieldLogger {
	return log.WithField("module", module)
}

func parseLogLevel(level string) logrus.Level {
	switch level {
	case "error":
		return logrus.ErrorLevel
	case "warn", "warning":
		return logrus.WarnLevel
	case "info", "notice":
		return logrus.InfoLevel
	case "debug":
		return logrus.DebugLevel
	case "trace":
		return logrus.TraceLevel
	default:
		return logrus.InfoLevel
	}
}
func main() {
	// Get cmd line parameters
	coin := flag.String("coin", "coin", "Coin")
	addr := flag.String("addr", "localhost", "Address")
	port := flag.String("port", "15500", "Port")
	loglevel := flag.String("loglevel", "info", "Logging level (debug, info, warn, error)")
	flag.Parse()

	log.SetLevel(parseLogLevel(*loglevel))

	err := daemon.Init(os.Args[0], map[string]interface{}{}, "./daemonized.pid")
	if err != nil {
		return
	}

	switch os.Args[1] {
	case "start":
		err = daemon.Start()
	case "stop":
		err = daemon.Stop()
	case "restart":
		err = daemon.Stop()
		err = daemon.Start()
	case "status":
		status := "stopped"
		if daemon.IsRun() {
			status = "started"
		}

		fmt.Printf("Application is %s\n", status)

		return
	case "":
	default:
		mainLoop(*coin)
		fmt.Println("JORM node is on: ", *addr+":"+*port)
	}
}

func mainLoop(coin string) {

	n := NewJORMnode(coin)

	GetBitNodes(n.JDB)
	ticker := time.NewTicker(23 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:

				fmt.Println("JORM node wooikos")
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

}
