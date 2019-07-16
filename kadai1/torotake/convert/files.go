package convert

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// ListFiles 指定のディレクトリ(dir)以下を再帰的に探索し、指定フォーマット(format)の画像ファイルを列挙します。
func ListFiles(dir string, format Format) ([]string, error) {
	if !isDir(dir) {
		return nil, errors.New(dir + " is not directory.")
	}

	var files []string
	filepath.Walk(dir, generateWalkFunc(&files, format))

	return files, nil
}

func isDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	return info.IsDir()
}

func generateWalkFunc(files *[]string, format Format) func(string, os.FileInfo, error) error {
	targetExts := getTargetExts(format)
	return func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		ext := strings.ToLower(filepath.Ext(path))
		for _, v := range targetExts {
			if v == ext {
				*files = append(*files, path)
			}
		}
		return nil
	}
}

func getTargetExts(format Format) []string {
	switch format {
	case JPEG:
		return []string{".jpeg", ".jpg"}
	case PNG:
		return []string{".png"}
	case GIF:
		return []string{".gif"}
	case BMP:
		return []string{".bmp"}
	}

	return []string{}
}
