package image

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

// Convert is the function that converts an image format to another image format.
// origin is orifinal file, target is converted image, and format is target image format.
func Convert(origin io.Reader, target io.Writer, format string) error {

	var err error
	img, _, err := image.Decode(origin)
	if err != nil {
		return err
	}

	switch format {
	case "jpg", "jpeg":
		err = jpeg.Encode(target, img, nil)
	case "png":
		err = png.Encode(target, img)
	default:
		err = fmt.Errorf("対応していないフォーマットです")
	}

	if err != nil {
		return err
	}
	return nil
}
