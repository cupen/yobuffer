package subcmds

import (
	"fmt"
	"io"
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
			// 	Name:  "with-test",
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
			if len(files) <= 0 {
				return fmt.Errorf("no files")
			}
			for _, f := range files {
				log.Printf("%s found", f)

				data, err := readFile(f)
				p := parser.New()
				c, err := p.Parse(data)
				if err != nil {
					return err
				}

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

func readFile(fpath string) ([]byte, error) {
	f, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(f)

}
