package main

import (
	"fmt"
	"os"

	"github.com/gopherdojo/dojo6/kadai2/en-ken/cli"
	"github.com/gopherdojo/dojo6/kadai2/en-ken/imgcnv"
)

func main() {
	cli := &cli.CLI{
		AllFilePaths: func(path string, ext string) ([]string, error) { return imgcnv.AllFilePaths(path, ext) },
		NewImageFIle: func(path string) (imgcnv.ImageFile, error) { return imgcnv.NewImageFile(path) },
	}

	err := cli.Execute(os.Args)
	if err != nil {
		fmt.Printf("%v", err)
	}
}
