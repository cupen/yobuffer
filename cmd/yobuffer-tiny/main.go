package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func readFile(fpath string) (string, error) {
	f, err := os.OpenFile(fpath, os.O_RDONLY, 0)
	if err != nil {
		return "", fmt.Errorf("file open failed:%w.", err)
	}
	data, err := io.ReadAll(f)
	if err != nil {
		return "", fmt.Errorf("file read failed:%w", err)
	}
	text := string(data)
	return text, nil

}

func main() {
	var src = flag.String("src", "", "input file")
	flag.Parse()
	if *src == "" {
		log.Printf("missing --src")
		return
	}
	text, err := readFile(*src)
	if err != nil {
		log.Printf("invalid file:%v file:%s", err, *src)
		return
	}
	log.Println(text)
}
