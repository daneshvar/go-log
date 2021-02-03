package log

import "fmt"

type Logger struct {
	scope  string
	caller bool
	skip   int
}

var defLogger Logger = Logger{
	scope:  defaultScope,
	caller: true,
}

func GetLogger() *Logger {
	return &defLogger
}

func GetScope(name string) *Logger {
	return &Logger{
		scope: name,
		skip:  defaultSkip,
	}
}

func GetCustom(name string, skip int) *Logger {
	return &Logger{
		scope: name,
		skip:  skip + defaultSkip,
	}
}

func (l *Logger) Print(messages ...interface{}) {
	log.print(InfoLevel, l.scope, l.skip, messages)
}

func (l *Logger) Println(messages ...interface{}) {
	log.print(InfoLevel, l.scope, l.skip, messages)
}

func (l *Logger) Printf(format string, args ...interface{}) {
	log.printf(InfoLevel, l.scope, l.skip, format, args)
}

func (l *Logger) Debug(messages ...interface{}) {
	log.print(DebugLevel, l.scope, l.skip, messages)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	log.printf(DebugLevel, l.scope, l.skip, format, args)
}

func (l *Logger) Debugv(message string, keysValues ...interface{}) {
	log.printv(DebugLevel, l.scope, l.skip, message, keysValues)
}

func (l *Logger) Info(messages ...interface{}) {
	log.print(InfoLevel, l.scope, l.skip, messages)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	log.printf(InfoLevel, l.scope, l.skip, format, args)
}

func (l *Logger) Infov(message string, keysValues ...interface{}) {
	log.printv(InfoLevel, l.scope, l.skip, message, keysValues)
}

func (l *Logger) Warn(messages ...interface{}) {
	log.print(WarnLevel, l.scope, l.skip, messages)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	log.printf(WarnLevel, l.scope, l.skip, format, args)
}

func (l *Logger) Warnv(message string, keysValues ...interface{}) {
	log.printv(WarnLevel, l.scope, l.skip, message, keysValues)
}

func (l *Logger) Error(messages ...interface{}) {
	log.print(ErrorLevel, l.scope, l.skip, messages)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	log.printf(ErrorLevel, l.scope, l.skip, format, args)
}

func (l *Logger) Errorv(message string, keysValues ...interface{}) {
	log.printv(ErrorLevel, l.scope, l.skip, message, keysValues)
}

func (l *Logger) Fatal(messages ...interface{}) {
	log.print(FatalLevel, l.scope, l.skip, messages)
}

func (l *Logger) Fatalln(messages ...interface{}) {
	log.print(FatalLevel, l.scope, l.skip, messages)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	log.printf(FatalLevel, l.scope, l.skip, format, args)
}

func (l *Logger) Fatalv(message string, keysValues ...interface{}) {
	log.printv(FatalLevel, l.scope, l.skip, message, keysValues)
}

func (l *Logger) Panic(messages ...interface{}) {
	log.print(PanicLevel, l.scope, l.skip, messages)
	panic(messages)
}

func (l *Logger) Panicln(messages ...interface{}) {
	log.print(PanicLevel, l.scope, l.skip, messages)
	panic(messages)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	log.printf(PanicLevel, l.scope, l.skip, format, args)
	panic(fmt.Sprintf(format, args...))
}

func (l *Logger) Panicv(message string, keysValues ...interface{}) {
	log.printv(PanicLevel, l.scope, l.skip, message, keysValues)
	panic(message)
}
