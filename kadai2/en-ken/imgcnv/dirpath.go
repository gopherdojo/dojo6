package imgcnv

import (
	"os"
	"path/filepath"
)

// IAllFilePaths is I/F of AllFilePaths
type IAllFilePaths func(path string, ext string) ([]string, error)

const allocSize = 100

// AllFilePaths returns
// the paths of the files in the specified directory
// filtered by the specified extension.
func AllFilePaths(path string, ext string) ([]string, error) {

	result := make([]string, 0, allocSize)
	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			// Skip
			return nil
		}

		// Find by extension
		isMatch, _ := filepath.Match("*."+ext, info.Name())
		if isMatch {
			absPath, err := filepath.Abs(filePath)
			result = append(result, absPath)
			return err
		}
		return nil
	})

	return result, err
}
