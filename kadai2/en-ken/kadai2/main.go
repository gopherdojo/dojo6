package main

import (
	"os"

	"github.com/gopherdojo/dojo6/kadai2/en-ken/cli"
	"github.com/gopherdojo/dojo6/kadai2/en-ken/imgcnv"
)

func main() {
	cli := &cli.CLI{
		AllFilePaths: func(path string, ext string) ([]string, error) { return imgcnv.AllFilePaths(path, ext) },
		NewImageFIle: func(path string) (imgcnv.ImageFile, error) { return imgcnv.NewImageFile(path) },
	}

	cli.Execute(os.Args)
}
