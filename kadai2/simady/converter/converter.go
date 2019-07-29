package converter

import (
	"fmt"
	"image"
	"os"
)

// IConverter コンバーターインターフェース
type IConverter interface {
	Convert(path string, fileName string) error
}

type converter struct {
	Encoder
}

type Encoder interface {
	encode(file *os.File, img image.Image) error
}

// NewConverter 拡張子に対応したコンバーターを生成する.
func NewConverter(dst string) (IConverter, error) {
	switch dst {
	case "png":
		return converter{&pngEncoder{}}, nil
	case "jpg":
		return converter{&jpgEncoder{}}, nil
	case "gif":
		return converter{&gifEncoder{}}, nil
	}
	return nil, ImgconvError{Message: fmt.Sprintf("Unsupported extension. dst = %v\n", dst)}
}

// Convert 画像フォーマットの変換を行う.
func (c converter) Convert(path string, fileName string) error {
	src, err := os.Open(path)
	if err != nil {
		return err
	}
	defer src.Close()

	var img image.Image
	img, _, err = image.Decode(src)
	if err != nil {
		return err
	}

	var dst *os.File
	dst, err = os.Create(fileName)
	if err != nil {
		return err
	}
	defer dst.Close()

	return c.encode(dst, img)
}
