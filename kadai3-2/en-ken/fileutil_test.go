package divdl_test

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	divdl "github.com/gopherdojo/dojo6/kadai3-2/en-ken"

	"github.com/google/go-cmp/cmp"
)

var tmpDir string

func TestMain(m *testing.M) {
	tmpDir, _ = ioutil.TempDir("", ".tmp")
	m.Run()
}

func TestSaveFile(t *testing.T) {
	file := filepath.Join(tmpDir, "test.txt")

	expected := []byte("foo\nbar\nbaz\n")
	divdl.SaveFile(file, expected)

	actual, _ := ioutil.ReadFile(file)
	if string(actual) != string(expected) {
		t.Errorf("failed. Diff:\n%v", cmp.Diff(actual, expected))
	}
}

func TestMergeFiles(t *testing.T) {
	strA := "0\n1\n2\n3\n4\n"
	strB := "5\n6\n7\n8\n9\n"
	fileInfo := []struct {
		fileName string
		data     string
	}{
		{
			fileName: filepath.Join(tmpDir, "merged.txt.0"),
			data:     strA,
		},
		{
			fileName: filepath.Join(tmpDir, "merged.txt.1"),
			data:     strB,
		},
		{
			fileName: filepath.Join(tmpDir, "merged.txt.2"),
			data:     strA,
		},
		{
			fileName: filepath.Join(tmpDir, "merged.txt.3"),
			data:     strB,
		},
	}

	var inputFiles []string
	var expected string
	for _, fi := range fileInfo {
		if err := divdl.SaveFile(fi.fileName, []byte(fi.data)); err != nil {
			t.Errorf("SaveFile failed: %v", err)
		}
		inputFiles = append(inputFiles, fi.fileName)
		expected += fi.data
	}

	mergedFile := filepath.Join(tmpDir, "merged.txt")
	if err := divdl.MergeFiles(inputFiles, mergedFile); err != nil {
		t.Errorf("MergeFiles failed: %v", err)
	}

	actual, _ := ioutil.ReadFile(mergedFile)
	if string(actual) != expected {
		t.Errorf("MergeFiles makes invalid data: %v", string(actual))
	}
}
