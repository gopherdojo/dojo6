/*
JPEG形式の画像をPNG形式に変換するパッケージです。
*/
package convert

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
)

// 変換する前と後のファイル名を持った構造体です。
type ConvertFile struct {
	Old string // 変換前のファイル名
	New string // 変換後のファイル名
}

// ConvertToPng は ファイルのパスを受け取り、JPGならPNG形式に画像を変換します
func (cnv *ConvertFile) ConvertToPng() error {
	file, err := os.Open(cnv.Old)
	if err != nil {
		fmt.Printf("Can not open : %v ", err)
		return err
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil {
		fmt.Printf("Can not Decode : %v ", err)
		return err
	}
	if format != "jpeg" {
		return nil
	}

	out, err := os.Create(cnv.New)
	defer out.Close()

	err = png.Encode(out, img)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
