package log

import "fmt"

const defaultScope = ""
const defaultSkip = 3

func Debug(messages ...interface{}) {
	log.print(DebugLevel, defaultScope, defaultSkip, messages)
}

func Debugf(format string, args ...interface{}) {
	log.printf(DebugLevel, defaultScope, defaultSkip, format, args)
}

func Debugv(message string, keysValues ...interface{}) {
	log.printv(DebugLevel, defaultScope, defaultSkip, message, keysValues)
}

func Info(messages ...interface{}) {
	log.print(InfoLevel, defaultScope, defaultSkip, messages)
}

func Infof(format string, args ...interface{}) {
	log.printf(InfoLevel, defaultScope, defaultSkip, format, args)
}

func Infov(message string, keysValues ...interface{}) {
	log.printv(InfoLevel, defaultScope, defaultSkip, message, keysValues)
}

func Warn(messages ...interface{}) {
	log.print(WarnLevel, defaultScope, defaultSkip, messages)
}

func Warnf(format string, args ...interface{}) {
	log.printf(WarnLevel, defaultScope, defaultSkip, format, args)
}

func Warnv(message string, keysValues ...interface{}) {
	log.printv(WarnLevel, defaultScope, defaultSkip, message, keysValues)
}

func Error(messages ...interface{}) {
	log.print(ErrorLevel, defaultScope, defaultSkip, messages)
}

func Errorf(format string, args ...interface{}) {
	log.printf(ErrorLevel, defaultScope, defaultSkip, format, args)
}

func Errorv(message string, keysValues ...interface{}) {
	log.printv(ErrorLevel, defaultScope, defaultSkip, message, keysValues)
}

func Fatal(messages ...interface{}) {
	log.print(FatalLevel, defaultScope, defaultSkip, messages)
}

func Fatalf(format string, args ...interface{}) {
	log.printf(FatalLevel, defaultScope, defaultSkip, format, args)
}

func Fatalv(message string, keysValues ...interface{}) {
	log.printv(FatalLevel, defaultScope, defaultSkip, message, keysValues)
}

func Panic(messages ...interface{}) {
	log.print(PanicLevel, defaultScope, defaultSkip, messages)
	panic(messages)
}

func Panicf(format string, args ...interface{}) {
	log.printf(PanicLevel, defaultScope, defaultSkip, format, args)
	panic(fmt.Sprintf(format, args...))
}

func Panicv(message string, keysValues ...interface{}) {
	log.printv(PanicLevel, defaultScope, defaultSkip, message, keysValues)
	panic(message)
}
