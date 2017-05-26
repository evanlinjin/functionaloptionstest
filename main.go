package main

import (
	"github.com/urfave/cli"
	"os"
	"github.com/evanlinjin/functionaloptionstest/hello"
)

var (
	ArgGreeting = "Hello"
	ArgName = "Anonymous"
	ArgCount = 1
)

func main() {
	app := cli.NewApp()
	app.Name = "Hello"
	app.Usage = "Give yourself the greeting you deserve!"
	app.Version = "1.0"
	app.Author = "Evan Chih-Yu Lin <evan.linjin@gmail.com>"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "greeting",
			Usage: "greeting to use",
			Destination: &ArgGreeting,
			Value: ArgGreeting,
		},
		cli.StringFlag{
			Name: "name",
			Usage: "your name",
			Destination: &ArgName,
			Value: ArgName,
		},
		cli.IntFlag{
			Name: "count",
			Usage: "number of times to greet you",
			Destination: &ArgCount,
			Value: ArgCount,
		},
	}
	app.Action = func(c *cli.Context) error {
		greeter, e := hello.NewGreeter(
			hello.Greeting(ArgGreeting),
			hello.Name(ArgName),
			hello.Count(ArgCount),
		)
		if e != nil {
			return e
		}
		greeter.Greet()
		return nil
	}
	if e := app.Run(os.Args); e != nil {
		panic(e)
	}
}
