package omikuji

import (
	"log"
	"math/rand"
	"time"
)

// おみくじ結果
type OmikujiResult struct {
	ID      int32  `json:"id"`
	Ruck    string `json:"ruck"`
	Message string `json:"message"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
	log.Println("Seed initialized...")
}

// おみくじ結果セット
type OmikujiResults []OmikujiResult

// FindRandom おみくじ結果セットからランダムに一つを取得する.
func (rs OmikujiResults) FindRandom() OmikujiResult {
	return rs[rand.Int31n(int32(len(rs)))]
}

// ExtractByRuck おみくじ結果セットから特定のruckを抽出する.
func (rs OmikujiResults) ExtractByRuck(ruck string) OmikujiResults {
	var ret OmikujiResults
	for _, r := range rs {
		if r.Ruck == ruck {
			ret = append(ret, r)
		}
	}
	return ret
}
