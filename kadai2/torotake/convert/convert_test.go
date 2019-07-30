package convert_test

import (
	"io"
	"os"

	"testing"

	"github.com/gopherdojo/dojo6/kadai2/torotake/convert"
)

func setupTestConvert(t *testing.T) func(t *testing.T) {
	// テスト用のディレクトリを作成し、テストデータをコピーする
	if err := os.Mkdir("testdata/test", 0777); err != nil {
		t.Fatalf("setup failed : can't create test dir")
	}
	copyFile("testdata/test_jpg.jpg", "testdata/test/test_jpg.jpg")
	copyFile("testdata/test_png.png", "testdata/test/test_png.png")
	copyFile("testdata/test_bmp.bmp", "testdata/test/test_bmp.bmp")
	copyFile("testdata/test_gif.gif", "testdata/test/test_gif.gif")

	// teardown テスト用のディレクトリを削除する
	return func(t *testing.T) {
		if err := os.RemoveAll("testdata/test"); err != nil {
			t.Errorf("teardown failed : can't remove test dir")
		}
	}
}

// TestConvert convert.Convert()の指定のフォーマットへの画像変換をテスト
func TestConvert(t *testing.T) {
	teardownTestConvert := setupTestConvert(t)
	defer teardownTestConvert(t)

	var cases = []struct {
		name   string
		src    []string
		format convert.Format
		actual string
	}{
		{
			"ConvertJpegToPng",
			[]string{"testdata/test/test_jpg.jpg"},
			convert.PNG,
			"testdata/test/test_jpg.png",
		},
		{
			"ConvertJpegToBmp",
			[]string{"testdata/test/test_jpg.jpg"},
			convert.BMP,
			"testdata/test/test_jpg.bmp",
		},
		{
			"ConvertJpegToGif",
			[]string{"testdata/test/test_jpg.jpg"},
			convert.GIF,
			"testdata/test/test_jpg.gif",
		},
		{
			"ConvertPngToJpeg",
			[]string{"testdata/test/test_png.png"},
			convert.JPEG,
			"testdata/test/test_png.jpg",
		},
		{
			"ConvertPngToBmp",
			[]string{"testdata/test/test_png.png"},
			convert.BMP,
			"testdata/test/test_png.bmp",
		},
		{
			"ConvertPngToGif",
			[]string{"testdata/test/test_png.png"},
			convert.GIF,
			"testdata/test/test_png.gif",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Helper()
			opts := convert.Options{SrcFiles: c.src, OutputFormat: c.format}
			convert.Convert(opts)
			_, err := os.Stat(c.actual)
			if err != nil {
				t.Fatalf("%v is not found", c.actual)
			}
		})
	}
}

func copyFile(src string, dst string) error {
	srcFp, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFp.Close()

	dstFp, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFp.Close()

	_, err = io.Copy(dstFp, srcFp)
	return err
}
