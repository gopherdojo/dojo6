package main

import (
	"flag"
	"os"

	"github.com/gopherdojo/dojo6/kadai2/pei/pkg/convdir"
	"github.com/gopherdojo/dojo6/kadai2/pei/pkg/imgconv"
)

const (
	usageInExt  = `input extension (jpg, png, gif)`
	usageOutExt = `output extension (jpg, png, gif)`
	usageLeave  = `whether to leave input`
)

func main() {
	var (
		inputExtension  = flag.String("in_ext", "jpg", usageInExt)
		outputExtension = flag.String("out_ext", "png", usageOutExt)
		leaveInput      = flag.Bool("leave", false, usageLeave)
	)
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		os.Exit(1)
	}

	cd := convdir.ConverterWithDir{
		Dir:             args[0],
		InputExtension:  imgconv.ParseImgExtension(*inputExtension),
		OutputExtension: imgconv.ParseImgExtension(*outputExtension),
		LeaveInput:      *leaveInput,
	}
	cd.Convert()
}
