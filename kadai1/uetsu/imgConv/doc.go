// Copyright 2019 masashi.uetsu. All rights reserved.

/*
	指定したディレクトリ配下のjpegファイルを再帰的に探索してpngファイルに変換します。ディレクトリは並べて複数指定可能です。

	Usage

		imgconv dirPath dirPath

	func searchfile.RecursionFile

		RecursionFile(dirPath string) (filePathList []string)

	func convert.ImgConv

		ImgConv(filePath string) (convertedFilePath string, err error)

*/
package main
