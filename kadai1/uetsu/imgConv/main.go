package main

import (
	"flag"
	"fmt"

	"dojo6/kadai1/uetsu/ImgConv/convert"
	"dojo6/kadai1/uetsu/imgConv/searchfile"
)

func main() {
	flag.Parse()
	args := flag.Args()

	for _, dirPath := range args {
		filePathList := searchfile.RecursionFile(dirPath)

		for _, filePath := range filePathList {
			convertedFilePath, err := convert.ImgConv(filePath)
			if err != nil {
				fmt.Printf("%s convert fail\n", filePath)
				continue
			}
			fmt.Printf("%s convert succeed => %s\n", filePath, convertedFilePath)
		}

	}
}
