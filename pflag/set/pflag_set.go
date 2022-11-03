package main

import (
	"fmt"
	"github.com/spf13/pflag"
)

func main() {
	var version bool
	// custom flagSet
	// flagSet := pflag.NewFlagSet("test", pflag.ContinueOnError)
	// flagSet.BoolVarP(&version, "version", "v", false, "print version")
	// flagSet.Parse(os.Args[1:])

	// global flagSet
	pflag.BoolVarP(&version, "version", "v", false, "print version")
	pflag.Parse()
	fmt.Printf("version=%v\n", version)
}
