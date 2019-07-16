package convert

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// InputExt 拡張子の変換前と変換後を格納する。
type InputExt struct {
	Src string
	Dst string
}

// ImgConv ファイルをjpgからpngに変換する。
// TODO: 他の拡張子に対応
func ImgConv(filePath string, inputExt InputExt) (convertedFilePath string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	var decodeImage image.Image
	switch inputExt.Src {
	case "jpg", "jpeg":
		decodeImage, err = jpeg.Decode(file)
	}
	if err != nil {
		return
	}

	convertedFilePath = strings.Join([]string{filePath[:len(filePath)-len(filepath.Ext(filePath))], inputExt.Dst}, ".")

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
