package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/naoty/contrib/command"
)

func main() {
	app := cli.NewApp()
	app.Name = "contrib"
	app.Usage = "Show contributors for specified files"
	app.Version = "0.1.0"
	app.Author = "Naoto Kaneko"
	app.Email = "naoty.k@gmail.com"
	app.Action = command.Contrib
	app.Run(os.Args)
}
