package imgconv

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

// ImgConverter struct has 2 paths and 2 extensions for input and output.
type ImgConverter struct {
	// 
	InputPath string
	InputExtension ImgExtension
	OutputPath string
	OutputExtension ImgExtension
}

// Convert image file.
func (ic *ImgConverter) Convert() error {
	inFile, err := os.Open(ic.InputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer inFile.Close()

	inImg, err := decodeImg(inFile, ic.InputExtension)
	if err != nil {
		log.Fatal(err)
	}

	outFile, err := os.Create(ic.OutputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	err = encodeImg(outFile, ic.OutputExtension, &inImg)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func encodeImg(file *os.File, ext ImgExtension, img *image.Image) (err error) {
	switch ext {
	case JPEG:
		err = jpeg.Encode(file, *img, nil)
	case PNG:
		err = png.Encode(file, *img)
	case GIF:
		err = gif.Encode(file, *img, nil)
	default:
		err = fmt.Errorf("%s is not supported", ext)
	}
	return
}

func decodeImg(file *os.File, ext ImgExtension) (img image.Image, err error) {
	switch ext {
	case JPEG:
		img, err = jpeg.Decode(file)
	case PNG:
		img, err = png.Decode(file)
	case GIF:
		img, err = gif.Decode(file)
	default:
		err = fmt.Errorf("%s is not supported", ext)
	}
	return
}
