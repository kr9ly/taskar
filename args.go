package tasker

import "flag"

type Args struct {
	ConfigPath string
}

func Parse(args string) Args {
	a := Args{}
	f := flag.NewFlagSet(args, flag.ExitOnError)
	f.StringVar(&a.ConfigPath, "c", "task.yml", "ConfigFile Path")
	return a
}