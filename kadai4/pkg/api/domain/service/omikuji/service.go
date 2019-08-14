package omikuji

import (
	"context"

	entity "omikuji-app/pkg/api/domain/entity/omikuji"
	"omikuji-app/pkg/api/ocontext"
)

const dateFormat = "1/2"

var daikichiOnlyDates = map[string]struct{}{"1/1": {}, "1/2": {}, "1/3": {}}

type OmikujiService interface {
	Draw(ctx context.Context) entity.OmikujiResult
}

type omikujiService struct {
}

func New() OmikujiService {
	return &omikujiService{}
}

// Draw おみくじを引く.
func (s *omikujiService) Draw(ctx context.Context) entity.OmikujiResult {
	t := ocontext.GetAccessTime(ctx)
	rs := omikujiResults
	date := t.Format(dateFormat)
	if _, exists := daikichiOnlyDates[date]; exists {
		rs = rs.ExtractByRuck("大吉")
	}
	return rs.FindRandom()
}
