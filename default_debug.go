// +build !release

package log

func Trace(messages ...interface{}) {
	log.print(TraceLevel, defaultScope, defaultSkip, messages)
}

func Tracef(format string, args ...interface{}) {
	log.printf(TraceLevel, defaultScope, defaultSkip, format, args)
}

func Tracev(message string, keysValues ...interface{}) {
	log.printv(TraceLevel, defaultScope, defaultSkip, message, keysValues)
}
