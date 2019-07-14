package info

import (
	"reflect"
	"testing"
)

func TestGetFilePath(t *testing.T) {
	// ファイルが存在しない場合のテストケース
	dir := "noexist"
	ext := "jpg"
	_, err := GetFilePath(dir, ext)
	if err == nil {
		t.Error("存在しないディレクトリを引数に指定した場合、GetFilePathはnilでないerrorを返すべきです")
	}

	//
	dir = "test.txt"
	_, err = GetFilePath(dir, ext)
	if err == nil {
		t.Error("ディレクトリでないファイルを引数に指定した場合、GetFilePathはnilでないerrorを返すべきです")
	}

	dir = "test-images"
	paths, err := GetFilePath(dir, ext)
	if err != nil {
		t.Error("ディレクトリを指定した場合、errorはnilであるべきです")
	}
	if !reflect.DeepEqual(paths, []string{"test-images/Moon.jpg", "test-images/test2-images/momiji.jpg"}) {
		t.Error("GetFilePathが正しいファイルパスが格納されたスライスを返していません")
	}

}
