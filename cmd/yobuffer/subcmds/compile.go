package subcmds

import (
	"log"
	"os"
	"path/filepath"

	"github.com/cupen/yobuffer/parser"
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
			// &cli.StringFlag{
			// 	Name:  "dest",
			// 	Usage: "destination file or directory",
			// },
			// &cli.BoolFlag{
			// 	Name:  "testcase",
			// 	Usage: "generate testcases",
			// },
		},
		Action: func(c *cli.Context) error {
			src := c.String("src")
			dest := c.String("dest")
			testcase := c.Bool("testcase")
			_ = testcase
			_ = dest
			files, err := filepath.Glob(src)
			if err != nil {
				return err
			}
			for _, f := range files {
				log.Printf("%s found", f)
				c, err := parser.Parse(f)
				if err != nil {
					return err
				}
				// log.Printf("\t%v err:%v", c, err)
				_destFile := f + ".go"
				f, err := os.OpenFile(_destFile, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0755)
				if err != nil {
					log.Printf("generated fail: err: %v", err)
					return err
				}
				err = c.WriteTo(f)
				if err != nil {
					log.Printf("generated fail: err: %v", err)
					return err
				}
				postAction := "format ok"
				// if err := exec.Command("bash", "-c", "go fmt "+_destFile).Run(); err != nil {
				// 	postAction = fmt.Sprintf("format failed:%v", err)
				// }
				log.Printf("generated ok %s: %s", postAction, _destFile)
			}
			return nil
		},
	}
}
