package main

import (
	"context"
	"fmt"
	"github.com/nurana88/online-shopping/pkg/command"
	"github.com/urfave/cli"
	"log"
	"os"
)

// Build version. Sent as a linker flag in Makefile
var version string

func main() {

	fmt.Println("In the main")
	version = "1"
	app := cli.NewApp()
	app.Version = version
	app.Name = "Online shopping platform"
	app.Usage = "App provides online shopping and makes life easier for lovely customers ;)"

	basecommand := command.NewBaseCommand(context.TODO())
	httpCommand := command.NewHTTPCommand(basecommand)

	app.Commands = []cli.Command{
		{
			Name:   "http",
			Usage:  "Registers customer",
			Action: httpCommand.Run,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("At the end main")

}
