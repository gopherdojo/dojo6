package main

import (
	"flag"
	"log"
	"path/filepath"
	"strings"

	"github.com/gopherdojo/dojo6/kadai1/en-ken/imgcnv"
)

func main() {
	var (
		inputExt  = flag.String("input-ext", ".jpg", "input extension (.jpg/.png)")
		outputExt = flag.String("output-ext", ".png", "output extension (.jpg/.png)")
		output    = flag.String("output", "", "output directory")
		inputDir  string
		outputDir string
	)

	flag.Parse()
	inputDir, err := filepath.Abs(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	outputDir = *output
	if outputDir == "" {
		outputDir = inputDir
	}

	paths, err := imgcnv.AllFileInfo(inputDir, *inputExt)
	if err != nil {
		log.Fatal(err)
	}

	for _, path := range paths {
		img, err := imgcnv.NewImageFile(path)
		if err != nil {
			log.Fatal(err)
		}

		// Copy the hierarchy of the input dir to that of the output dir.
		outputPath := strings.Replace(img.AbsPath(), inputDir, outputDir, -1)
		outputPath = strings.Replace(outputPath, *inputExt, *outputExt, 1)

		err = img.SaveAs(outputPath)
		if err != nil {
			log.Fatal(err)
		}
	}
}
