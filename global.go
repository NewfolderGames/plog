package plog

import "os"

var instance = New(os.Stdout, 0)

// SetGlobalInstance sets the global logger instance to the provided Plog instance.
func SetGlobalInstance(i *Plog) {

	instance = i

}

// SetFlag sets the logging flag for the global logger instance.
func SetFlag(flag int64) {

	instance.SetFlag(flag)

}

// SetCallerDepth sets the runtime caller depth for the global logger instance.
func SetCallerDepth(depth int) {

	instance.SetCallerDepth(depth)

}

// Info prints an info log.
func Info(format string, args ...any) int {

	return instance.base(infoString, format, args...)

}

// Debug prints a debug log.
func Debug(format string, args ...any) int {

	return instance.base(debugString, format, args...)

}

// Warn prints a warning log.
func Warn(format string, args ...any) int {

	return instance.base(warnString, format, args...)

}

// Error prints an error log.
func Error(format string, args ...any) int {

	return instance.base(errorString, format, args...)

}
