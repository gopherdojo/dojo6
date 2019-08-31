package download

import (
	"path/filepath"
	"strings"
)

func parseDirAndFileName(path string) (dir, file string) {
	lastSlashIndex := strings.LastIndex(path, "/")
	dir = path[:lastSlashIndex+1]
	if len(dir) == len(path) {
		return dir, ""
	}

	return dir, path[(len(dir) + 1):]
}

func parseFileName(url string) string {
	return filepath.Base(url)
}
