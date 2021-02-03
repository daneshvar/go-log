# Log

```
package main

import "github.com/daneshvar/go-log"

func main() {
	defer log.Sync()
	log.RedirectStdLog()
	log.Config(log.DebugLevel, true)

	// log.GetLogger()
	// log.SetCaller(true)

	log.Warn("Not Found config file")

	log.Infov("GET",
		"url", "http://example.com/data.json",
	)

	log.Error("Fetch",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)
}
```
