package cli

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/gopherdojo/dojo6/kadai2/en-ken/imgcnv"
)

type ImageFileMock struct {
	t          *testing.T
	inputPath  string
	outputPath string
}

func (mock *ImageFileMock) AbsPath() string {
	return mock.inputPath
}

func (mock *ImageFileMock) SaveAs(path string) error {
	if path != mock.outputPath {
		mock.t.Errorf("SaveAs was not called correctly:\nactual: %v\nexpected: %v", path, mock.outputPath)
	}
	return nil
}

func TestExecuteSuccess(t *testing.T) {
	tests := []struct {
		argString string
		inputDir  string
		inputExt  string
		outputDir string
		outputExt string
	}{
		{
			argString: "./kadai1 ./testdata ./out",
			inputDir:  "./testdata", inputExt: "jpg",
			outputDir: "./out", outputExt: "png",
		},
		{
			argString: "./kadai1 -in png ./testdata ./out",
			inputDir:  "./testdata", inputExt: "png",
			outputDir: "./out", outputExt: "png",
		},
		{
			argString: "./kadai1 -out jpg ./testdata ./out",
			inputDir:  "./testdata", inputExt: "jpg",
			outputDir: "./out", outputExt: "jpg",
		},
		{
			argString: "./kadai1 -in png -out jpg ./testdata ./out",
			inputDir:  "./testdata", inputExt: "png",
			outputDir: "./out", outputExt: "jpg",
		},
	}

	for _, test := range tests {
		testName := "input: " + test.argString
		t.Run(testName, func(t *testing.T) {

			inputAbsDir, _ := filepath.Abs(test.inputDir)
			inputPath1 := inputAbsDir + "/test1" + test.inputExt
			inputPath2 := inputAbsDir + "/test2" + test.inputExt
			outputPath1, _ := filepath.Abs(test.outputDir + "/test1" + test.outputExt)
			outputPath2, _ := filepath.Abs(test.outputDir + "/test2" + test.outputExt)

			allFilePathsMock := func(path string, ext string) ([]string, error) {
				if path == inputAbsDir && ext == test.inputExt {
					return []string{inputPath1, inputPath2}, nil
				}
				t.Errorf("AllFilePaths was not called correctly: %v %v", path, ext)
				return nil, nil
			}

			newImageFileMock := func(path string) (imgcnv.ImageFile, error) {
				switch path {
				case inputPath1:
					return &ImageFileMock{
						t:          t,
						inputPath:  path,
						outputPath: outputPath1,
					}, nil
				case inputPath2:
					return &ImageFileMock{
						t:          t,
						inputPath:  path,
						outputPath: outputPath2,
					}, nil
				}
				return nil, nil
			}

			cli := &CLI{
				AllFilePaths: allFilePathsMock,
				NewImageFIle: newImageFileMock,
			}
			cli.Execute(strings.Split(test.argString, " "))

		})
	}
}
