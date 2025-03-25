package logger

import "os"

func LogIF(err error) {
	if err != nil {
		os.Exit(1)
	}
}
