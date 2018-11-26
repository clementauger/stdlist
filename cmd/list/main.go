package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func main() {

	srcPath := filepath.Join(runtime.GOROOT(), "src")

	var res []string
	err := filepath.Walk(srcPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() == false {
			return nil
		}
		p, err := filepath.Rel(srcPath, path)
		if err != nil {
			return err
		}
		if p == "." {
			return nil
		}
		if p == "cmd" {
			return filepath.SkipDir
		}
		res = append(res, p)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range res {
		fmt.Println(r)
	}

}
