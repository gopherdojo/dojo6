package converter

import (
	"image"
	"log"
	"os"
)

// IConverter コンバーターインターフェース
type IConverter interface {
	convert(file *os.File, img image.Image) error
	validate() bool
}

type converter struct {
	src, dst string
}

var (
	convertibleExts = []string{"jpeg", "png", "gif"}
)

// Validate フィールドのバリデーション処理を行う.
func (c converter) validate() bool {
	return validate(c.src) && validate(c.dst)
}

func validate(ext string) bool {
	for _, e := range convertibleExts {
		if ext == e {
			return true
		}
	}
	return false
}

// Convert 画像フォーマットの変換を行う.
func Convert(path string, c IConverter, fileName string) error {

	src, err := os.Open(path)
	if err != nil {
		log.Fatalf("Faild to open file. err = %v\n", err)
	}
	defer src.Close()

	var img image.Image
	img, _, err = image.Decode(src)
	if err != nil {
		log.Fatalf("Faild to decode file. err = %v\n", err)
	}

	var dst *os.File
	dst, err = os.Create(fileName)
	if err != nil {
		log.Fatalf("Faild to create file. err = %v\n", err)
	}
	defer dst.Close()

	return c.convert(dst, img)
}

// GetConverter 拡張子に対応したコンバーターを取得する.
func GetConverter(src, dst string) IConverter {
	c := converter{src: src, dst: dst}
	c.validate()
	switch dst {
	case "png":
		return pngConverter{c}
	case "jpeg":
		return jpegConverter{c}
	case "gif":
		return gifConverter{c}
	default:
		log.Fatalf("Unsupported extension. dst = %v\n", dst)
	}
	return nil
}
