package gotyping

import (
	"flag"
	"io"
	"log"
)

// Run the gotyping
func Run(outStream, errStream io.Writer) error {

	log.SetOutput(errStream)

	var limits int
	flag.IntVar(&limits, "limits", 60, "制限時間")
	flag.Parse()

	return nil
}
