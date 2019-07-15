package imgcnv

import (
	"path/filepath"
	"sort"
	"testing"
)

func TestCanGetAllFileInfo(t *testing.T) {
	dirPath, _ := NewDirPath("../testdata/")
	actual, _ := dirPath.AllFilePaths("jpg")
	expected := []string{"../testdata/lenna_color.jpg", "../testdata/lenna_gray.jpg", "../testdata/layer1/girl_color.jpg", "../testdata/layer1/girl_gray.jpg", "../testdata/layer1/layer2/Mandrill.jpg"}
	for i := 0; i < len(expected); i++ {
		expected[i], _ = filepath.Abs(expected[i])
	}

	if isEqualArray(actual, expected) == false {
		t.Errorf("\nactual:%v\nexpected:%v", actual, expected)
	}
}

func compare(x string, y string) bool {
	return x < y
}
func isEqualArray(array1 []string, array2 []string) bool {
	sort.Slice(array1, func(i, j int) bool { return array1[i] > array1[j] })
	sort.Slice(array2, func(i, j int) bool { return array2[i] > array2[j] })
	if len(array1) != len(array2) {
		return false
	}
	for i := 0; i < len(array1); i++ {
		if array1[i] != array2[i] {
			return false
		}
	}
	return true
}
