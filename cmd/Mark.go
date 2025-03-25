package cmd

import "flag"

var MarkCmd *flag.FlagSet

func init() {
	MarkCmd = flag.NewFlagSet("mark", flag.ExitOnError)
}
