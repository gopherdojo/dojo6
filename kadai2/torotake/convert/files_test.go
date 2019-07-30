package convert_test

import (
	"testing"

	"github.com/gopherdojo/dojo6/kadai2/torotake/convert"
)

// convert.ListFiles()が指定のフォーマットのファイルを列挙出来るかをテスト
func TestListFiles(t *testing.T) {
	files, err := convert.ListFiles("./testdata", convert.JPEG)
	if err != nil {
		t.Errorf("ListFiles returns err %#v", err)
	}
	if len(files) != 1 || files[0] != "testdata/test_jpg.jpg" {
		t.Errorf("ListFiles JPEG : failed")
	}
}
