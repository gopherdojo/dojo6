/*
Package convert は指定のディレクトリ以下に存在する指定された形式の画像ファイルを
別の形式の画像ファイルに変換するためのパッケージです。
*/
package convert

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	"golang.org/x/image/bmp"
)

// Format 画像ファイルの形式
type Format int

const (
	// UNKNOWN 形式不明
	UNKNOWN Format = iota - 1
	// JPEG JPEG形式
	JPEG
	// PNG PNG形式
	PNG
	// GIF GIF形式
	GIF
	// BMP BMP形式
	BMP
)

// Options 画像変換のオプション指定
type Options struct {
	SrcFiles     []string
	OutputFormat Format
}

// Convert optionsに指定された内容に従って画像を変換します
func Convert(options Options) {
	for _, src := range options.SrcFiles {
		dst := getDst(src, options.OutputFormat)
		fmt.Printf("convert %s -> %s ...\n", src, dst)
		err := convertFormat(src, dst, options.OutputFormat)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
		}
	}
}

func getDst(src string, f Format) string {
	return src[:len(src)-len(filepath.Ext(src))] + getDstExt(f)
}

func getDstExt(f Format) string {
	switch f {
	case JPEG:
		return ".jpg"
	case PNG:
		return ".png"
	case GIF:
		return ".gif"
	case BMP:
		return ".bmp"
	}

	return ""
}

func convertFormat(src string, dst string, format Format) error {
	srcFp, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFp.Close()

	img, _, err := image.Decode(srcFp)
	if err != nil {
		return err
	}

	dstFp, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFp.Close()

	switch format {
	case JPEG:
		return jpeg.Encode(dstFp, img, nil)
	case PNG:
		return png.Encode(dstFp, img)
	case GIF:
		return gif.Encode(dstFp, img, nil)
	case BMP:
		return bmp.Encode(dstFp, img)
	}

	return errors.New("Invalid format")
}
