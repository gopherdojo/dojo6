package Converter

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

var c = &Converter{
	FileInfo: &FileInfo{
		Base: DirType{
			DirName: "../jpeg",
			Extension: "jpg",
			FilePaths: []string{
				"../jpeg/abc/sample03.jpg",
				"../jpeg/sample01.jpg",
				"../jpeg/sample02.jpg",
			},
		},
		Dist: DirType{
			DirName: "output",
			Extension: "png",
		},
	},
}

func TestConverter_Decode(t *testing.T) {

	c.OpenFiles()
	c.Decode()

	fmt.Println(*c)

	// デコード後のio.Writerの個数と元ファイルの個数が一致しているかを確認
	if len(c.Imgs) != 3 {
		t.Errorf("Invalid length of c.Imgs. Length: %v, expected: 3", len(c.Imgs))
	}
}

func TestConverter_Encode(t *testing.T) {

	fmt.Println("1", c.FileInfo.Dist.DirName)

	c.Encode()

	var s []string
	err := filepath.Walk("./" + c.FileInfo.Dist.DirName, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == "." + c.FileInfo.Dist.Extension {
			s = append(s, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	// エンコード後に書き出されたファイルの個数と元のファイルの個数が同じかを確認
	if len(s) != 3 {
		t.Errorf("Invalid length of output files. Length: %v, expected: 3", len(s))
	}

	if err := os.RemoveAll(c.FileInfo.Dist.DirName); err != nil {
		fmt.Println(err)
	}

}
