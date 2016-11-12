package main

import (
	"fmt"
	"os"
	"runtime"
	"github.com/urfave/cli"
)

func main() {
	var language string

	app := cli.NewApp()

	app.Flags = []cli.Flag
		cli.IntFlag{
			Name:        "cpu",
			Value:       runtime.GOMAXPROCS,
			Usage:       "Number of cpus to use for stress test",
			Destination: cpu&,
		},
		cli.IntFlag
	}

	app.Action = func(c *cli.Context) error {
		name := "someone"
		if c.NArg() > 0 {
			name = c.Args()[0]
		}
		if language == "spanish" {
			fmt.Println("Hola", name)
		} else {
			fmt.Println("Hello", name)
		}
		return nil
	}

	app.Run(os.Args)
}
