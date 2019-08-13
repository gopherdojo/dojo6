package omikuji

import (
	"context"

	"omikuji-app/pkg/api/ocontext"

	entity "omikuji-app/pkg/api/domain/entity/omikuji"
)

var rs = entity.OmikujiResults{
	{
		ID:      1,
		Ruck:    "大吉",
		Message: "おめでとうございます！大吉です！",
	},
	{
		ID:      2,
		Ruck:    "大吉",
		Message: "大吉でした！絶好調です！",
	},
	{
		ID:      3,
		Ruck:    "吉",
		Message: "吉です！かなりツイてます！",
	},
	{
		ID:      4,
		Ruck:    "吉",
		Message: "吉です！良い運勢ですね！",
	},
	{
		ID:      5,
		Ruck:    "中吉",
		Message: "中吉でした！いい感じです！",
	},
	{
		ID:      6,
		Ruck:    "中吉",
		Message: "なかなかいいですね！中吉です！",
	},
	{
		ID:      7,
		Ruck:    "小吉",
		Message: "小吉です！悪くないですね！",
	},
	{
		ID:      8,
		Ruck:    "小吉",
		Message: "小吉です！少しツイてます！",
	},
	{
		ID:      9,
		Ruck:    "小吉",
		Message: "小吉でした！！！",
	},
	{
		ID:      10,
		Ruck:    "凶",
		Message: "凶でした！今日も一日頑張りましょう！",
	},
}

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
	day := t.Format("1/2")
	if day == "1/1" || day == "1/2" || day == "1/3" {
		rs = rs.ExtractByRuck("大吉")
	}
	return rs.FindRandom()
}
