package imgcnv

import (
	"os"
	"path/filepath"
	"testing"
)

const inputDir = "../testdata/"
const outputDir = "../out/"

func TestMain(m *testing.M) {

	os.RemoveAll(outputDir)
	m.Run()
	os.RemoveAll(outputDir)
}

func TestCanConstructImageFileFromJpegToPng(t *testing.T) {
	const fileName = "lenna_color"
	inputPath := filepath.Join(inputDir, fileName+".jpg")
	outputPath := filepath.Join(outputDir, fileName+".png")

	imageFile, err := NewImageFile(inputPath)
	if err != nil {
		t.Error(err)
	}
	err = imageFile.SaveAs(outputPath)
	if err != nil {
		t.Error(err)
	}
	if !exists(outputPath) {
		t.Error(err)
	}
}

func TestCanConstructImageFileFromPngToJpeg(t *testing.T) {
	const fileName = "layer1/girl_color"
	inputPath := filepath.Join(inputDir, fileName+".png")
	outputPath := filepath.Join(outputDir, fileName+".jpg")

	imageFile, err := NewImageFile(inputPath)
	if err != nil {
		t.Error(err)
	}
	err = imageFile.SaveAs(outputPath)
	if err != nil {
		t.Error(err)
	}
	if !exists(outputPath) {
		t.Error("No output files")
	}
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
