package converter

// 独自エラー定義
type ImgconvError struct {
	Message string
}

// Error エラー出力
func (e ImgconvError) Error() string {
	return e.Message
}