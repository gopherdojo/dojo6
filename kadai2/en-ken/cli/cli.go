package cli

import (
	"flag"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gopherdojo/dojo6/kadai2/en-ken/imgcnv"
)

// CLI is for DI
type CLI struct {
	AllFilePaths imgcnv.IAllFilePaths
	NewImageFIle imgcnv.INewImageFile
}

// Execute executes this app according to options
func (cli *CLI) Execute(args []string) error {

	var inputExt, outputExt string
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	flags.StringVar(&inputExt, "in", "jpg", "input extension (jpg/png)")
	flags.StringVar(&outputExt, "out", "png", "output extension (jpg/png)")

	// Parse command args
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	if len(flags.Args()) != 2 {
		return fmt.Errorf("Either input directory or output directory not specified")
	}
	// Get input dir
	inputDir, err := filepath.Abs(flags.Arg(0))
	if err != nil {
		return err
	}
	// Get output dir
	outputDir, err := filepath.Abs(flags.Arg(1))
	if err != nil {
		return err
	}

	// Get all file paths
	paths, err := cli.AllFilePaths(inputDir, inputExt)
	if err != nil {
		return err
	}

	// Convert and save
	for _, path := range paths {
		img, err := cli.NewImageFIle(path)
		if err != nil {
			return err
		}

		// Copy the hierarchy of the input dir to that of the output dir.
		outputPath := strings.Replace(img.AbsPath(), inputDir, outputDir, -1)
		outputPath = strings.Replace(outputPath, inputExt, outputExt, 1)

		err = img.SaveAs(outputPath)
		if err != nil {
			return err
		}
	}

	return nil
}
