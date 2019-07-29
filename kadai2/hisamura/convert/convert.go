// Package convert is image conversion processing
//
// 画像変換処理。
// 変換前と変換後の拡張子を指定し、画像変換を行います。
// ディレクトリを選択し、再帰的にディレクトリ探索を行い変換します。
// 変換前のファイルを削除したい場合は、remove(-r)フラグを指定します。
package convert

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// FlagOps is flag option
type FlagOps struct {
	Dir    string
	Src    string
	Dest   string
	Remove bool
}

// FileDetails have Extionsion and FileName info
type FileDetails struct {
	Extension string
	FileName  string
}

// Convert is image conversion processing
func Convert(flagOps FlagOps) (string, error) {

	files := dirwalk(flagOps.Dir)
	var convertedName string

	for _, file := range files {
		fileDetails := FileDetails{
			filepath.Ext(file),
			filepath.Clean(file),
		}

		if strings.HasSuffix(fileDetails.Extension, flagOps.Src) {

			decodeImg, err := decodeImage(fileDetails)
			if err != nil {
				return "", err
			}

			if flagOps.Remove {
				defer os.Remove(fileDetails.FileName)
			}

			fileNameRemoveExt := strings.Replace(fileDetails.FileName, flagOps.Src, "", 1)
			dstFile, err := os.Create(fmt.Sprintf(fileNameRemoveExt + flagOps.Dest))
			if err != nil {
				return "", err
			}
			convertedName = dstFile.Name()
			defer dstFile.Close()

			err = encodeImage(decodeImg, flagOps, dstFile)
			if err != nil {
				return "", err
			}
		}
	}
	return convertedName, nil

}

func decodeImage(fileDetails FileDetails) (image.Image, error) {
	srcFile, err := os.Open(fileDetails.FileName)

	defer srcFile.Close()

	if err != nil {
		return nil, err
	}

	decodeImg, _, err := image.Decode(srcFile)
	return decodeImg, err
}

func encodeImage(decodeImg image.Image, flagOps FlagOps, dstFile *os.File) error {
	switch flagOps.Dest {
	case "jpeg", "jpg":
		err := jpeg.Encode(dstFile, decodeImg, nil)
		return err

	case "gif":
		err := gif.Encode(dstFile, decodeImg, nil)
		return err

	case "png":
		err := png.Encode(dstFile, decodeImg)
		return err

	default:
		return fmt.Errorf("Error: invalid extension")
	}
}

//指定したディレクトリ配下のファイルを取得
func dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}
