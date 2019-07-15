package imgcnv

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

// ImageFile expresses the i
type ImageFile struct {
	image *image.Image
	path  string
}

// NewImageFile is a constructor of ImageFile
func NewImageFile(path string) (*ImageFile, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var image image.Image
	switch filepath.Ext(path) {
	case ".jpg", ".jpeg":
		image, err = jpeg.Decode(file)
	case ".png":
		image, err = png.Decode(file)
	default:
		// try to decode as jpeg
		image, err = jpeg.Decode(file)
	}
	if err != nil {
		return nil, err
	}

	return &ImageFile{
		image: &image,
		path:  absPath,
	}, nil
}

// AbsPath returns the absolute path of the input file
func (img *ImageFile) AbsPath() string {
	return img.path
}

// SaveAs oututs a file to the specified path after convering to the specified exteinsion.
func (img *ImageFile) SaveAs(path string) error {
	err := os.MkdirAll(filepath.Dir(path), 0777)
	if err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	ext := filepath.Ext(path)
	switch ext {
	case ".jpg", ".jpeg":
		return jpeg.Encode(file, *img.image, nil)
	case ".png":
		return png.Encode(file, *img.image)
	default:
		return fmt.Errorf("Unexpected extension")
	}
}
