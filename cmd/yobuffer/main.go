package main

import (
	"log"
	"os"

	"github.com/cupen/yobuffer/cmd/yobuffer/subcmds"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "yobuffer",
		Usage: "manage your buffers",
		Commands: []*cli.Command{
			subcmds.Compile(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
