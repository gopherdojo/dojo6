package convert

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

// Information
type Info struct {
}

// ファイル名から拡張子を取ったファイル名を返します
func getFileNameWithoutExt(path string) string {
	// Fixed with a nice method given by mattn-san
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

// ConvertToPng は ファイルのパスを受け取り、JPGならPNG形式に画像を変換します
func ConvertToPng(path string) error {
	file, err := os.Open(path)
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
		fmt.Println("jpeg じゃないよ")
		return nil
	}

	filename := getFileNameWithoutExt(path) + ".png"
	out, err := os.Create(filename)
	defer out.Close()

	err = png.Encode(out, img)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
