package main

import (
	"flag"
	"fmt"
	"os"
	"path"
)

type Args struct {
	Port     uint
	Addr     string
	Filepath string
}

var args Args = Args{
	Port: 3000,
}

func init() {
	var usage string

	usage = "Specify the web server's port."
	flag.UintVar(&args.Port, "port", args.Port, usage)
	flag.UintVar(&args.Port, "p", args.Port, usage)

	flag.Parse()
	flag.Usage = help

	args.Addr = fmt.Sprintf(":%d", args.Port)

	if args.Filepath = flag.Arg(0); args.Filepath == "" {
		flag.Usage()
		os.Exit(0)
	}
}

func help() {
	program := path.Base(os.Args[0])
	output := "Usage of %s:\n" +
		"  web [flags] /path/to/json\n" +
		"\n"

	fmt.Fprintf(flag.CommandLine.Output(), output, program)
	flag.PrintDefaults()
}
