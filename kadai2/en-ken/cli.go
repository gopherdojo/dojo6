package main

import (
	"flag"
	"path/filepath"
	"strings"

	"github.com/gopherdojo/dojo6/kadai2/en-ken/imgcnv"
)

// CLI is for DI
type CLI struct {
	dirPath      imgcnv.DirPath
	imageFactory imgcnv.ImageFileFactory
}

// Execute executes this app according to options
// arg[0] application name
// arg[1] input directory
// arg[2:] options
func (cli *CLI) Execute(args []string) error {

	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	var (
		inputExt  = flags.String("input-ext", ".jpg", "input extension (.jpg/.png)")
		outputExt = flags.String("output-ext", ".png", "output extension (.jpg/.png)")
		output    = flags.String("output-dir", args[1], "output directory")
		inputDir  string
		outputDir string
	)

	inputDir, err := filepath.Abs(args[1])
	if err != nil {
		return err
	}

	if err := flags.Parse(args[2:]); err != nil {
		return err
	}

	outputDir = *output
	if outputDir == "" {
		outputDir = inputDir
	}

	paths, err := cli.dirPath.AllFilePaths(inputDir, *inputExt)
	if err != nil {
		return err
	}

	for _, path := range paths {
		img, err := cli.imageFactory.Create(path)
		if err != nil {
			return err
		}

		// Copy the hierarchy of the input dir to that of the output dir.
		outputPath := strings.Replace(img.AbsPath(), inputDir, outputDir, -1)
		outputPath = strings.Replace(outputPath, *inputExt, *outputExt, 1)

		err = img.SaveAs(outputPath)
		if err != nil {
			return err
		}
	}

	return nil
}
