package cmd

import "flag"

var UpdateCmd *flag.FlagSet

func init() {
	UpdateCmd = flag.NewFlagSet("Update", flag.ExitOnError)
}
