package main

import (
	"os"

	"github.com/gopherdojo/dojo6/kadai2/en-ken/cli"
	"github.com/gopherdojo/dojo6/kadai2/en-ken/imgcnv"
)

func main() {
	dirPath := imgcnv.NewDirPath()
	factory := imgcnv.NewImageFileFactory()
	cli := &cli.CLI{
		DirPath:      dirPath,
		ImageFactory: factory,
	}

	cli.Execute(os.Args)
}
