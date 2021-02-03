package rtsp

import (
	"time"

	"github.com/daneshvar/go-log"
)

func GetPacketFunc() {
	logger := log.GetScope("rtsp")

	//log.Info("Namitonam fetch konam")
	//
	//log.Infof("Namitonam fetch konam %s", "Hossein")
	logger.Warn("Not Found config file")

	logger.Infov("GET",
		"url", "http://example.com/data.json",
	)
	logger.Errorv("Fetch",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)

	logger.Infov("Namitonam",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)

	logger.Errorv("Namitonam",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)
}
