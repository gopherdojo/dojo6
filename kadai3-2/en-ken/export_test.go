package divdl

var DivideIntoRanges = func(contentLength int64, numOfDivision int) (int, [][]*TestRange) {
	n, result := divideIntoRanges(contentLength, numOfDivision)
	rngs := make([][]*TestRange, numOfDivision)
	for i, rp := range result {
		for _, r := range rp {
			rngs[i] = append(rngs[i], &TestRange{
				ID:   r.id,
				From: r.from,
				To:   r.to,
			})
		}
	}
	return n, rngs
}

type TestRange struct {
	ID   int
	From int64
	To   int64
}

var MaxRangeSize = maxRangeSize

func SetMaxRangeSize(size int64) {
	maxRangeSize = size
}
