package imgcnv

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

// ImageFile expresses I/F to ImageFileStruct
type ImageFile interface {
	AbsPath() string
	SaveAs(path string) error
}

// ImageFileStruct expresses the converting image
type ImageFileStruct struct {
	image *image.Image
	path  string
	ImageFile
}

// ImageFileFactory is I/F of ImageFileFactoryStruct
type ImageFileFactory interface {
	Create(path string) (ImageFile, error)
}

// ImageFileFactoryStruct is for DI
type ImageFileFactoryStruct struct {
}

// NewImageFileFactory is a constructor of NewImageFileFactory
func NewImageFileFactory() ImageFileFactory {
	return &ImageFileFactoryStruct{}
}

// Create generates ImageFile
func (factory *ImageFileFactoryStruct) Create(path string) (ImageFile, error) {
	return NewImageFile(path)
}

// NewImageFile is a constructor of ImageFile
func NewImageFile(path string) (ImageFile, error) {
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

	return &ImageFileStruct{
		image: &image,
		path:  absPath,
	}, nil
}

// AbsPath returns the absolute path of the input file
func (img *ImageFileStruct) AbsPath() string {
	return img.path
}

// SaveAs oututs a file to the specified path after convering to the specified exteinsion.
func (img *ImageFileStruct) SaveAs(path string) error {
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
