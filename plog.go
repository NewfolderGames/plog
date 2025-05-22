package plog

import (
	"fmt"
	"io"
	"runtime"
	"time"
)

const infoString = "\x1b[32mINFO\x1b[0m"
const debugString = "\x1b[34mDBUG\x1b[0m"
const warnString = "\x1b[33mWARN\x1b[0m"
const errorString = "\x1b[31mEROR\x1b[0m"

// FLAG_HIDE_FILENAME is a flag used to hide the filename and line number in the log output when enabled.
const FLAG_HIDE_FILENAME = 1 << 0

// FLAG_HIDE_DATE is a flag used to hide the date and time in the log output when enabled.
const FLAG_HIDE_DATE = 1 << 1

// Plog is a structured logger that writes log entries to an io.Writer with optional flags for controlling log output.
type Plog struct {
	flag        int64
	callerDepth int
	writer      io.Writer
}

// New creates a new pLog instance.
func New(writer io.Writer, flag int64) *Plog {

	return &Plog{
		flag:        flag,
		callerDepth: 2,
		writer:      writer,
	}

}

// Writer returns the io.Writer of the instance.
func (plog *Plog) Writer() io.Writer {

	return plog.writer

}

// SetFlag sets the flag of the instance.
func (plog *Plog) SetFlag(flag int64) {

	plog.flag = flag

}

// SetCallerDepth sets the runtime caller depth of the instance.
func (plog *Plog) SetCallerDepth(depth int) {

	plog.callerDepth = depth

}

// base is a base method for printing a log.
func (plog *Plog) base(level string, format string, a ...any) int {

	var timeString string
	var fileString string

	if plog.flag&FLAG_HIDE_FILENAME == 0 {

		_, fileName, lineNum, _ := runtime.Caller(plog.callerDepth)
		fileString = fmt.Sprintf(" \x1b[90m%s:%d\x1b[0m", fileName, lineNum)

	}

	if plog.flag&FLAG_HIDE_DATE == 0 {

		timeString = fmt.Sprintf("\x1b[90m%s\x1b[0m ", time.Now().UTC().Format("2006-01-02T15:04:05.000Z"))

	}

	bytesWritten, _ := io.WriteString(plog.writer, fmt.Sprintf(timeString+level+fileString+"\u001B[90m:\u001B[0m "+format+"\n", a...))

	return bytesWritten

}

// Info prints an info log.
func (plog *Plog) Info(format string, args ...any) int {

	return plog.base(infoString, format, args...)

}

// Debug prints a debug log.
func (plog *Plog) Debug(format string, args ...any) int {

	return plog.base(debugString, format, args...)

}

// Warn prints a warning log.
func (plog *Plog) Warn(format string, args ...any) int {

	return plog.base(warnString, format, args...)

}

// Error prints an error log.
func (plog *Plog) Error(format string, args ...any) int {

	return plog.base(errorString, format, args...)

}
