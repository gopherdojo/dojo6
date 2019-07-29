package converter

import (
	"image"
	"image/jpeg"
	"os"
)

// jpgEncoder Jpeg用エンコーダー
type jpgEncoder struct {
}

func (*jpgEncoder) encode(file *os.File, img image.Image) error {
	return jpeg.Encode(file, img, &jpeg.Options{Quality: 100})
}
