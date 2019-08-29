package divdl

import (
	"fmt"
	"io"
	"os"

	"github.com/pkg/errors"
)

// SaveFile saves data as fileName.
func SaveFile(fileName string, data []byte) error {
	file, err := os.Create(fileName)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Failed to open %v", fileName))
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// MergeFiles merges separeted files.
func MergeFiles(inputFiles []string, outputFileName string) error {

	fw, err := os.Create(outputFileName)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Failed to open output file: %v", outputFileName))
	}
	defer fw.Close()

	for _, f := range inputFiles {
		fr, err := os.Open(f)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Failed to open [%v]", f))
		}
		defer fr.Close()

		_, err = io.Copy(fw, fr)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Failed to write data from %v", f))
		}

		if err := os.Remove(f); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Failed to delete %v", f))
		}

	}

	return nil
}
