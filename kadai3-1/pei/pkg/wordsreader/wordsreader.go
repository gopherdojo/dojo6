package wordsreader

import (
	"bufio"
	"os"
)

// WordsReader has FileName
type WordsReader struct {
	FileName string
}

// Read file
func (wr WordsReader) Read() ([]string, error) {
	var words []string

	fp, err := os.Open(wr.FileName)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}
