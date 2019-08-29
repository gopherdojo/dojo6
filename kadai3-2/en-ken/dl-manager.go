package divdl

import (
	"fmt"

	"golang.org/x/sync/errgroup"

	"github.com/pkg/errors"
)

// DlRange expresses range of Range request
type DlRange struct {
	id   int
	from int64
	to   int64
}

// MaxRangeSize is limit size of range
const MaxRangeSize = 1024 * 1024 //1MB

func divideIntoRanges(contentLength int64, numOfDivision int) (numOfRanges int, rngs [][]*DlRange) {
	rngs = make([][]*DlRange, numOfDivision)

	rngSize := (contentLength + int64(numOfDivision)) / int64(numOfDivision)
	if MaxRangeSize < rngSize {
		rngSize = MaxRangeSize
	}

	for j, pos := 0, int64(0); pos < contentLength; j++ {
		for i := 0; i < numOfDivision; i++ {
			id := j*numOfDivision + i
			numOfRanges = id + 1
			// Last range
			if contentLength-pos < rngSize {
				rngs[i] = append(rngs[i], &DlRange{
					id:   id,
					from: int64(id) * rngSize,
					to:   contentLength - 1,
				})
				pos = contentLength
				break
			}

			rngs[i] = append(rngs[i], &DlRange{
				id:   id,
				from: int64(id) * rngSize,
				to:   int64(id+1)*rngSize - 1,
			})
			pos += rngSize
		}
	}
	return
}

// Do manages separately downloading.
func Do(url string, fileName string, numOfDivision int) error {
	req, err := NewRequest(url)
	if err != nil {
		return errors.WithStack(err)
	}

	// Range request is not accepted
	if !req.CanAcceptRangeRequest() {
		data, err := req.Download()
		if err != nil {
			return err
		}

		return SaveFile(fileName, data)
	}

	// Range request is accepted
	n, rngs := divideIntoRanges(req.GetContentLength(), numOfDivision)

	var g errgroup.Group
	for _, rSet := range rngs {
		rSet := rSet
		g.Go(func() error {
			for _, r := range rSet {
				tmpFileName := createPartialFileName(fileName, r.id)
				data, err := req.DownloadPartially(r.from, r.to)
				if err != nil {
					return err
				}
				if err := SaveFile(tmpFileName, data); err != nil {
					return err
				}
				fmt.Printf("%v saved\n", tmpFileName)
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return errors.WithStack(err)
	}

	files := make([]string, 0)
	for i := 0; i < n; i++ {
		files = append(files, createPartialFileName(fileName, i))
	}
	return MergeFiles(files, fileName)
}

func createPartialFileName(fileName string, suffix int) string {
	return fmt.Sprintf("%v.%v", fileName, suffix)
}
