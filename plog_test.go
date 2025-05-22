package plog

import (
	"os"
	"testing"
)

func TestOutput(t *testing.T) {

	var instance = New(os.Stdout, 0)

	instance.Info("%s", "hello")
	instance.Debug("%s", "hello")
	instance.Warn("%s", "hello")
	instance.Error("%s", "hello")

}

func TestOutputGlobal(t *testing.T) {

	Info("%s", "hello")
	Debug("%s", "hello")
	Warn("%s", "hello")
	Error("%s", "hello")

}

func TestGlobalReplace(t *testing.T) {

	SetGlobalInstance(New(os.Stdout, 0))

	instance.Info("%s", "hello")
	instance.Debug("%s", "hello")
	instance.Warn("%s", "hello")
	instance.Error("%s", "hello")

}

func TestOutputGlobalHideFile(t *testing.T) {

	SetFlag(FLAG_HIDE_FILENAME)
	Info("%s", "hello")
	Debug("%s", "hello")
	Warn("%s", "hello")
	Error("%s", "hello")

}

func TestOutputGlobalHideDate(t *testing.T) {

	SetFlag(FLAG_HIDE_DATE)
	Info("%s", "hello")
	Debug("%s", "hello")
	Warn("%s", "hello")
	Error("%s", "hello")

}

func TestCallerDepth(t *testing.T) {

	SetCallerDepth(1)
	Info("%s", "hello")
	Debug("%s", "hello")
	Warn("%s", "hello")
	Error("%s", "hello")

}
