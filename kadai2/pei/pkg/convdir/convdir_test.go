package convdir

import (
	"os"
	"reflect"
	"testing"

	"github.com/gopherdojo/dojo6/kadai2/pei/pkg/imgconv"
)

func TestConverterWithDir_Convert(t *testing.T) {
	cd := ConverterWithDir{
		Dir:             "../../testdata",
		InputExtension:  imgconv.JPEG,
		OutputExtension: imgconv.PNG,
		LeaveInput:      true,
	}
	got := cd.Convert()
	want := []ConvertedResult{
		{OutputPath: "../../testdata/sample.png"},
		{OutputPath: "../../testdata/testdir1/testdir2/sample.png"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}

	os.Remove(want[0].OutputPath)
	os.Remove(want[1].OutputPath)
}
