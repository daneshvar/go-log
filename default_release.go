// +build release

package log

func Trace(messages ...interface{}) {}

func Tracef(format string, args ...interface{}) {}

func Tracev(message string, keysValues ...interface{}) {}
