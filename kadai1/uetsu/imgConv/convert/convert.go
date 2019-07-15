package convert

import (
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// ImgConv ファイルをjpegからpngに変換する。
// TODO: 他の拡張子に対応
func ImgConv(filePath string) (convertedFilePath string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	decodeImage, err := jpeg.Decode(file)
	if err != nil {
		return
	}

	convertedFilePath = strings.Join([]string{filePath[:len(filePath)-len(filepath.Ext(filePath))], "png"}, ".")

	dstfile, err := os.Create(convertedFilePath)
	if err != nil {
		return
	}
	defer dstfile.Close()

	err = png.Encode(dstfile, decodeImage)
	if err != nil {
		return
	}

	return
}
