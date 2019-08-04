package main

import (
	"log"
	"os"

	"github.com/dojo6/kadai3-1/annkara/gotyping"
)

func main() {

	// 標準ロガーに日時などの情報を付加しない
	log.SetFlags(0)

	var exitCode int
	err := gotyping.Run(os.Args[1:], os.Stdout, os.Stderr)
	if err != nil {
		log.Println(err)
		exitCode = 1
	}
	os.Exit(exitCode)
}
