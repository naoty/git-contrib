package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/naoty/git-contrib/command"
)

func main() {
	app := cli.NewApp()
	app.Name = "contrib"
	app.Usage = "Show contributors for specified files"
	app.Version = "0.1.0"
	app.Author = "Naoto Kaneko"
	app.Email = "naoty.k@gmail.com"
	app.Action = command.Contrib
	app.Flags = []cli.Flag{command.ContribFlag}
	app.Run(os.Args)
}
