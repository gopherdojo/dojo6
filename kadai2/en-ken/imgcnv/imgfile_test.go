package imgcnv

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const inputDir = "../testdata/"
const outputDir = "../out/"

func TestMain(m *testing.M) {

	os.RemoveAll(outputDir)
	m.Run()
	os.RemoveAll(outputDir)
}

func TestConstructImageFileSuccess(t *testing.T) {
	paths := []string{"layer1/layer2/Mandrill.jpg", "layer1/layer2/Mandrill.png"}

	for _, path := range paths {
		inputPath := filepath.Join(inputDir, path)
		_, err := NewImageFile(inputPath)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestConstructImageFileFailureWithInvalidPath(t *testing.T) {
	inputPath := filepath.Join(inputDir, "layer1/layer2/layer3/foo.jpg")
	_, err := NewImageFile(inputPath)
	if err == nil {
		t.Error(err)
	}
}

func TestConstructImageFileFailureWithOtherFormatFile(t *testing.T) {
	inputPath := filepath.Join(inputDir, "lenna_color.gif")
	_, err := NewImageFile(inputPath)
	if err == nil {
		t.Error(err)
	}
}

func TestAbsPathSuccess(t *testing.T) {
	filePath := "layer1/layer2/Mandrill.jpg"
	inputPath := filepath.Join(inputDir, filePath)
	image, _ := NewImageFile(inputPath)

	absPath := image.AbsPath()

	if !strings.Contains(absPath, filePath) {
		t.Error("Unexpected path", absPath)
	}
	if !exists(absPath) {
		t.Error("Path does not exist", absPath)
	}
}

func TestSaveAsSuccess(t *testing.T) {
	tests := []struct {
		fileName    string
		ext         string
		expectedExt string
	}{
		{fileName: "lenna_color", ext: ".png", expectedExt: ".jpg"},
		{fileName: "lenna_gray", ext: ".png", expectedExt: ".jpg"},
		{fileName: "lenna_color", ext: ".jpg", expectedExt: ".png"},
		{fileName: "lenna_gray", ext: ".jpg", expectedExt: ".png"},
		{fileName: "layer1/girl_color", ext: ".png", expectedExt: ".jpg"},
		{fileName: "layer1/girl_gray", ext: ".png", expectedExt: ".jpg"},
		{fileName: "layer1/girl_color", ext: ".jpg", expectedExt: ".png"},
		{fileName: "layer1/girl_gray", ext: ".jpg", expectedExt: ".png"},
		{fileName: "layer1/layer2/Mandrill", ext: ".png", expectedExt: ".jpg"},
		{fileName: "layer1/layer2/Mandrill", ext: ".jpg", expectedExt: ".png"},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("Test data: %v %v %v", test.fileName, test.ext, test.expectedExt)
		t.Run(testName, func(t *testing.T) {

			inputPath := filepath.Join(inputDir, test.fileName+test.ext)
			outputPath := filepath.Join(outputDir, test.fileName+test.expectedExt)

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
		})
	}

}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
