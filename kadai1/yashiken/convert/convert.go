/*
convertパッケージは、イメージファイルの形式を変換するためのパッケージです。
jpg, jpeg, png, gifの形式変換に対応しています。
*/
package convert

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"strings"
)

// Converterは変換前後のファイル形式を表現します。
type Converter struct {
	// 変換前のファイル形式
	extSrc string
	// 変換後のファイル形式
	extCnv string
}

// Converは、受け取ったファイルパスの画像ファイルを指定した形式に変換します
func (c Converter) Convert(path string) error {
	// 変換前のファイルのデコード
	data, err := c.decode(path)
	if err != nil {
		return err
	}

	// 指定したファイル形式へのエンコード
	err = c.encode(path, data)
	if err != nil {
		return err
	}
	// 成功メッセージの表示
	fmt.Printf("%s converted to %s image\n", path, c.extCnv)
	return nil

}

// decodeは指定したファイルパスのファイルの読み取り、デコードを行います
func (c Converter) decode(path string) (image.Image, error) {
	// ファイルの読み取り
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	// ファイル形式に応じたデコード処理
	switch c.extSrc {
	// jpg, jpegの場合
	case "jpg", "jpeg":
		img, err := jpeg.Decode(file)
		if err != nil {
			return nil, err
		}

		return img, nil
	// pngの場合
	case "png":
		img, err := png.Decode(file)
		if err != nil {
			return nil, err
		}

		return img, nil
	// gifの場合
	case "gif":
		img, err := gif.Decode(file)
		if err != nil {
			return nil, err
		}

		return img, nil
	}
	// 非対応のファイル形式を指定した場合、エラーを返却
	return nil, fmt.Errorf("extension %s is not supported.", c.extSrc)

}

// encodeは、image.Image型のデータを、特定の形式にエンコードし、指定したパスにファイルを作成します。
func (c Converter) encode(filepath string, data image.Image) error {
	// 書き込み先の指定
	writer, err := os.Create(strings.TrimSuffix(filepath, path.Ext(filepath)) + "." + c.extCnv)
	if err != nil {
		return err
	}

	// 変換先のファイル形式に応じたエンコード処理とファイル生成
	switch c.extCnv {
	// jpg, jpegの場合
	case "jpg", "jpeg":
		return jpeg.Encode(writer, data, nil)
	// pngの場合
	case "png":
		return png.Encode(writer, data)
	// gifの場合
	case "gif":
		return gif.Encode(writer, data, nil)
	}
	// 非対応のファイル形式の場合
	return fmt.Errorf("convert to extension %s is not supported.", c.extCnv)
}

// NewConverterはConverter構造体を生成するためのヘルパー関数です。
func NewConverter(extSrc, extCnv string) *Converter {
	return &Converter{
		extSrc: extSrc,
		extCnv: extCnv,
	}
}
