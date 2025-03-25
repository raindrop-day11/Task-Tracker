package cmd

import (
	"flag"
)

var AddCmd *flag.FlagSet

func init() {
	//添加命令
	AddCmd = flag.NewFlagSet("Add", flag.ExitOnError)
}
