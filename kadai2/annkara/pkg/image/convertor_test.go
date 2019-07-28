package image

import (
	"os"
	"path/filepath"
	"testing"
)

var directory = "testdata"

type testcase struct {
	desc   string
	origin string
	target string
	format string
	err    bool
}

func TestConvert(t *testing.T) {

	tests := []testcase{
		{
			desc:   "jpg to png",
			origin: filepath.Join(directory, "test.jpg"),
			target: filepath.Join(directory, "test2.png"),
			format: "png",
			err:    false,
		}, {
			desc:   "png to jpg",
			origin: filepath.Join(directory, "test.png"),
			target: filepath.Join(directory, "test2.jpg"),
			format: "jpg",
			err:    false,
		}, {
			desc:   "Unknow format",
			origin: filepath.Join(directory, "test.png"),
			target: filepath.Join(directory, "test2.diff"),
			format: "gif",
			err:    true,
		},
	}

	for _, tc := range tests {
		testConvert(t, tc)
	}
}

func testConvert(t *testing.T, tc testcase) {
	in, err := os.Open(tc.origin)
	if err != nil {
		t.Fatalf("failed test %s: %#v", tc.desc, err)
	}
	defer in.Close()

	out, err := os.Create(tc.target)
	if err != nil {
		t.Fatalf("failed test %s: %#v", tc.desc, err)
	}
	// テスト用に生成されたファイルを削除
	defer os.Remove(tc.target)
	defer out.Close()

	err = Convert(in, out, tc.format)
	if !((err != nil) == tc.err) {
		t.Fatalf("failed test %s: %#v", tc.desc, err)
	}

}
