package convdir

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gopherdojo/dojo6/kadai2/pei/pkg/imgconv"
)

// ImgConverter With Directory struct has target directory and 2 extention and option.
type ConverterWithDir struct {
	Dir             string
	InputExtension  imgconv.ImgExtension
	OutputExtension imgconv.ImgExtension
	LeaveInput      bool
}

// Result struct has output image path.
type ConvertedResult struct {
	OutputPath string
}

// Convert image file in target directory.
func (cd ConverterWithDir) Convert() (results []ConvertedResult) {
	files, err := cd.getImageFilePaths()
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {

		inputPath := fmt.Sprintf("%s/%s", cd.Dir, f)
		outputPath := fmt.Sprintf("%s.%s", strings.TrimSuffix(inputPath, filepath.Ext(inputPath)), cd.OutputExtension)

		ic := imgconv.ImgConverter{
			InputPath:       inputPath,
			InputExtension:  cd.InputExtension,
			OutputPath:      outputPath,
			OutputExtension: cd.OutputExtension,
			LeaveInput:      cd.LeaveInput,
		}
		err = ic.Convert()
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, ConvertedResult{OutputPath: outputPath})
	}

	return
}

func (cd ConverterWithDir) getImageFilePaths() (files []string, err error) {
	_, err = os.Stat(cd.Dir)
	if err != nil {
		return
	}

	err = filepath.Walk(cd.Dir,
		func(path string, info os.FileInfo, err error) error {
			relPath, err := filepath.Rel(cd.Dir, path)
			if !info.IsDir() && err == nil && imgconv.ParseImgExtension(relPath) == cd.InputExtension {
				files = append(files, relPath)
			}
			return nil
		})

	return
}
