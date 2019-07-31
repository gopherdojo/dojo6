package main

import (
	"fmt"
	"os"

	"github.com/gopherdojo/dojo6/kadai2/en-ken/cli"
	"github.com/gopherdojo/dojo6/kadai2/en-ken/imgcnv"
)

func main() {
	cli := &cli.CLI{
		AllFilePaths: imgcnv.IAllFilePaths(imgcnv.AllFilePaths),
		NewImageFIle: imgcnv.INewImageFile(imgcnv.NewImageFile),
	}

	err := cli.Execute(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERR:%v", err)
		os.Exit(-1)
	}
}
