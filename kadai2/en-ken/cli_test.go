package main

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gopherdojo/dojo6/kadai2/en-ken/mock_imgcnv"
)

func TestExecuteSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		argString string
		inputDir  string
		inputExt  string
		outputDir string
		outputExt string
	}{
		{
			argString: "./kadai1 ./testdata",
			inputDir:  "./testdata", inputExt: ".jpg",
			outputDir: "./testdata", outputExt: ".png",
		},
		{
			argString: "./kadai1 ./testdata -input-ext .png",
			inputDir:  "./testdata", inputExt: ".png",
			outputDir: "./testdata", outputExt: ".png",
		},
		{
			argString: "./kadai1 ./testdata -input-ext .png -output-ext .jpg",
			inputDir:  "./testdata", inputExt: ".png",
			outputDir: "./testdata", outputExt: ".jpg",
		},
		{
			argString: "./kadai1 ./testdata -output-dir ./out",
			inputDir:  "./testdata", inputExt: ".jpg",
			outputDir: "./out", outputExt: ".png",
		},
	}

	for _, test := range tests {
		testName := "input: " + test.argString
		t.Run(testName, func(t *testing.T) {

			mockDirPath := mock_imgcnv.NewMockDirPath(ctrl)

			inputAbsDir, _ := filepath.Abs(test.inputDir)
			inputPath1 := inputAbsDir + "/test1" + test.inputExt
			inputPath2 := inputAbsDir + "/test2" + test.inputExt

			mockDirPath.
				EXPECT().
				AllFilePaths(gomock.Eq(inputAbsDir), gomock.Eq(test.inputExt)).
				Return([]string{inputPath1, inputPath2}, nil)

			mockImg := mock_imgcnv.NewMockImageFile(ctrl)
			gomock.InOrder(
				mockImg.
					EXPECT().
					AbsPath().
					Return(inputPath1),
				mockImg.
					EXPECT().
					AbsPath().
					Return(inputPath2),
			)

			outputPath1 := test.outputDir + "/test1" + test.outputExt
			outputPath2 := test.outputDir + "/test2" + test.outputExt
			gomock.InOrder(
				mockImg.
					EXPECT().
					SaveAs(gomock.Eq(outputPath1)),
				mockImg.
					EXPECT().
					SaveAs(gomock.Eq(outputPath2)),
			)

			mockImgFactory := mock_imgcnv.NewMockImageFileFactory(ctrl)
			mockImgFactory.
				EXPECT().
				Create(gomock.Any()).
				Return(mockImg, nil).
				AnyTimes()

			cli := &CLI{
				dirPath:      mockDirPath,
				imageFactory: mockImgFactory,
			}
			cli.Execute(strings.Split(test.argString, " "))

		})
	}
}
