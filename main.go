package main

import (
	"github.com/urfave/cli"
	"github.com/vyasgiridhar/gstress/src"
	"os"
	"runtime"
)

func main() {
	var (
		cpu     int
		io      int
		timeout int
		hdd     int
	)

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "cpu",
			Value:       runtime.GOMAXPROCS(0),
			Usage:       "Number of cpus to use for stress test",
			Destination: &cpu,
		},
		cli.IntFlag{
			Name:        "io",
			Value:       4,
			Usage:       "Spawn N goroutines working on Sync()",
			Destination: &io,
		},
		cli.IntFlag{
			Name:        "hdd",
			Value:       4,
			Usage:       "Spawn N goroutines on Write()",
			Destination: &hdd,
		},
		cli.IntFlag{
			Name:        "timeout",
			Value:       0,
			Usage:       "timeout after N seconds",
			Destination: &timeout,
		},
	}

	app.Action = func(c *cli.Context) error {

		cpu := c.Int("cpu")
		io := c.Int("io")
		hdd := c.Int("hdd")
		timeout := c.Int("timeout")

		src.Spawner(cpu, io, hdd, timeout)

		return nil
	}

	app.Run(os.Args)
}
