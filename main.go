package main

import (
	"os"
	"fmt"
	"github.com/urfave/cli"
	"clynapses/cmd"
)

func main() {
	app := cli.NewApp()
	app.Name = "clynapses"
	app.Usage = "CLI client for https://synapses.polytechnique.fr/"
	app.EnableBashCompletion = true
	app.Version = "0.1.0"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "day, d",
			Usage: "get calendar of the day",
		},
		cli.BoolFlag{
			Name:  "tomorrow, t",
			Usage: "get calendar for the following day",
		},
		cli.BoolFlag{
			Name:  "week, w",
			Usage: "get calendar for the week",
		},
		cli.BoolFlag{
			Name:  "nextweek, nw",
			Usage: "get calendar for the next week",
		},
		cli.BoolFlag{
			Name:  "month, m",
			Usage: "get calendar for the month",
		},
		cli.BoolFlag{
			Name:  "nextmonth, nm",
			Usage: "get calendar for the next month",
		},
		cli.BoolFlag{
			Name:  "all, a",
			Usage: "get all the calendar",
		},
	}
	app.Action = cmd.Action
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("clynapses: error: ", err)
	}
}
