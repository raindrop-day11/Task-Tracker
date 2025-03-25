package logger

import (
	"os"

	"github.com/fatih/color"
)

func ErrorExit(err error) {
	color.New(color.BgHiRed).Fprintln(os.Stdout, err.Error())
	os.Exit(1)
}

func WarningExitIF(msg string, err error) {
	if err != nil {
		color.New(color.BgHiRed).Fprintln(os.Stdout, msg)
		os.Exit(1)
	}
}

func WarningExit(msg string) {
	color.New(color.BgHiRed).Fprintln(os.Stdout, msg)
	os.Exit(1)
}
