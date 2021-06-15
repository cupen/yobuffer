package subcmds

import (
	"log"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func Compile() *cli.Command {
	return &cli.Command{
		Name:    "compile",
		Aliases: []string{"c"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "src",
				Usage:    "source files",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			src := c.String("src")
			dest := c.String("dest")
			_ = dest
			files, err := filepath.Glob(src)
			if err != nil {
				return err
			}
			for _, f := range files {
				log.Printf("%s found", f)
			}
			return nil
		},
	}
}
