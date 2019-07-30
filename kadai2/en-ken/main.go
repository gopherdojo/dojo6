package main

import (
	"os"

	"github.com/gopherdojo/dojo6/kadai1/en-ken/imgcnv"
)

func main() {
	dirPath := imgcnv.NewDirPath()
	factory := imgcnv.NewImageFileFactory()
	cli := &CLI{
		dirPath:      dirPath,
		imageFactory: factory,
	}

	cli.Execute(os.Args)
}
