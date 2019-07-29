package main

import (
	"dojo6/kadai2/hisamura/convert"
	"flag"
	"fmt"
)

func main() {

	var (
		dir    = flag.String("dir", "./", "変換したいディレクトリ配下")
		src    = flag.String("s", "png", "変換前の拡張子")
		dest   = flag.String("d", "jpg", "変換後の拡張子")
		remove = flag.Bool("r", false, "変換前の拡張子ファイルを削除するかのflag")
	)

	flag.Parse()

	flagOps := convert.FlagOps{Dir: *dir, Src: *src, Dest: *dest, Remove: *remove}

	result := validation(flagOps)
	if result {
		fmt.Print("invalid extension")
		return
	}

	convertFile, err := convert.Convert(flagOps)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("convert to %v", convertFile)
}

func validation(flagOps convert.FlagOps) bool {
	if flagOps.Dest != "png" && flagOps.Dest != "jpg" && flagOps.Dest != "jpeg" && flagOps.Dest != "gif" {
		return true
	}
	return false
}
