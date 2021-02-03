package main

import (
	"time"

	"github.com/daneshvar/go-log"
	"github.com/daneshvar/go-log/example/rtsp"
)

func main() {
	defer log.Close()
	// log.RedirectStdLog()

	consoleEnabler := func(l log.Level, s string) bool { return true }
	influxEnabler := func(l log.Level, s string) bool { return true }
	stackEnabler := func(l log.Level, s string) bool { return l == log.ErrorLevel }

	log.Config(log.ConsoleWriter(true, stackEnabler, consoleEnabler),
		log.InfluxWriter("http://localhost:8086", "my-token", "behnama", "example-log", true, stackEnabler, influxEnabler))

	logger := log.GetScope("example")

	// log.GetLogger()
	logger.Trace("Check Trace 1")

	logger.Debug("Debug Code")
	logger.Debugf("Debug Code %s", "Hello")

	logger.Warn("Not Found config file")

	logger.Infov("GET", "url", "http://example.com/data.json")

	logger.Errorv("Fetch",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)
	test(logger)

	rtsp.GetPacketFunc()

	logger.Fatal("Fatal")
}

func test(logger *log.Logger) {
	logger.Errorv("Fetch",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)
	test2(logger)
}

func test2(logger *log.Logger) {
	logger.Errorv("Fetch",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)
}
