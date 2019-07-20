package convert

import (
	"os"
	"testing"
)

// TestMainは、テストで変換した画像ファイルを消去するための関数です。
func TestMain(m *testing.M) {
	code := m.Run()

	os.Remove("test-images/Moon_cvt.bmp")
	os.Remove("test-images/Moon_cvt.gif")
	os.Remove("test-images/Moon_cvt.jpeg")
	os.Remove("test-images/Moon_cvt.jpg")
	os.Remove("test-images/Moon_cvt.png")
	os.Remove("test-images/Moon.bmp")

	os.Exit(code)
}

// decodeのテストです。
func TestDecode(t *testing.T) {
	// 引数のファイルが存在しない場合のテストケース
	path := "noexist"
	c := &Converter{}
	// エラーを返すことの確認
	_, err := c.decode(path)
	if err == nil {
		t.Error("decodeはnilでないerrorを返すべきです")
	}

	// jpgファイルを変換する場合
	c1 := &Converter{extSrc: "jpg"}
	// 引数で指定したファイルの拡張子がxxx.jpgであるにも関わらず、
	// ファイル形式がjpgでない場合のテストケース
	_, err = c1.decode("test-images/dummy.jpg")
	if err == nil {
		t.Error("decodeはnilでないerrorを返すべきです")
	}
	// 正常にファイルを指定した場合、処理が成功することの確認
	_, err = c1.decode("test-images/Moon.jpg")
	if err != nil {
		t.Error("decodeはnil以外のerrorを返すべきではありません")
	}

	// jpegファイルを変換する場合
	c2 := &Converter{extSrc: "jpeg"}
	// 引数で指定したファイルの拡張子がxxx.jpegであるにも関わらず、
	// ファイル形式がjpegでない場合のテストケース
	_, err = c2.decode("test-images/dummy.jpeg")
	if err == nil {
		t.Error("decodeはnilでないerrorを返すべきです")
	}
	// 正常にファイルを指定した場合、処理が成功することの確認
	_, err = c2.decode("test-images/Moon.jpeg")
	if err != nil {
		t.Error("decodeはnil以外のerrorを返すべきではありません")
	}

	// pngファイルを変換する場合
	c3 := &Converter{extSrc: "png"}
	// 引数で指定したファイルの拡張子がxxx.pngであるにも関わらず、
	// ファイル形式がpngでない場合のテストケース
	_, err = c3.decode("test-images/dummy.png")
	if err == nil {
		t.Error("decodeはnilでないerrorを返すべきです")
	}
	// 正常にファイルを指定した場合、処理が成功することの確認
	_, err = c3.decode("test-images/Moon.png")
	if err != nil {
		t.Error("decodeはnil以外のerrorを返すべきではありません")
	}

	// gifファイルを変換する場合
	c4 := &Converter{extSrc: "gif"}
	// 引数で指定したファイルの拡張子がxxx.gifであるにも関わらず、
	// ファイル形式がgifでない場合のテストケース
	_, err = c4.decode("test-images/dummy.gif")
	if err == nil {
		t.Error("decodeはnilでないerrorを返すべきです")
	}
	// 正常にファイルを指定した場合、処理が成功することの確認
	_, err = c4.decode("test-images/Moon.gif")
	if err != nil {
		t.Error("decodeはnil以外のerrorを返すべきではありません")
	}

	// 対応していないファイル形式を指定した場合のテストケース
	c5 := &Converter{extSrc: "bmp"}

	// エラーを返すことを確認
	_, err = c5.decode("test-images/test.bmp")
	if err == nil {
		t.Error("decodeはnilでないerrorを返すべきです")
	}

}

// encodeのテストです
func TestEncode(t *testing.T) {
	// テストファイル（jpg）から、デコードしたデータを取得
	c := &Converter{extSrc: "jpg"}
	srcPath := "test-images/Moon.jpg"
	data, _ := c.decode(srcPath)

	// jpgファイルへのエンコードが成功するか確認
	c1 := &Converter{extCnv: "jpg"}
	if err := c1.encode("test-images/Moon_cvt.jpg", data); err != nil {
		t.Error("encodeが返すerrorはnilであるべきです")
	}

	// jpegファイルへのエンコードが成功するか確認
	c2 := &Converter{extCnv: "jpeg"}
	if err := c2.encode("test-images/Moon_cvt.jpg", data); err != nil {
		t.Error("encodeが返すerrorはnilであるべきです")
	}

	// pngファイルへのエンコードが成功するか確認
	c3 := &Converter{extCnv: "png"}
	if err := c3.encode("test-images/Moon_cvt.jpg", data); err != nil {
		t.Error("encodeが返すerrorはnilであるべきです")
	}

	// gifファイルへのエンコードが成功するか確認
	c4 := &Converter{extCnv: "gif"}
	if err := c4.encode("test-images/Moon_cvt.jpg", data); err != nil {
		t.Error("encodeが返すerrorはnilであるべきです")
	}

	// 対応していないファイル形式を指定した場合、処理が失敗することの確認
	c5 := &Converter{extCnv: "bmp"}
	if err := c5.encode("test-images/Moon_cvt.jpg", data); err == nil {
		t.Error("encodeはnilでないerrorを返すべきです")
	}
}

// Convertのテストです。
func TestConvert(t *testing.T) {
	// 存在しないファイルを指定した場合にエラーを返すことの確認
	path := "noexist"
	c1 := &Converter{extSrc: "jpg", extCnv: "png"}
	if err := c1.Convert(path); err == nil {
		t.Error("Convertはnilでないerrorを返すべきです")
	}

	// 対応していないファイル形式を指定した場合にエラーを返すことの確認
	path = "test-images/Moon.jpg"
	c2 := &Converter{extSrc: "jpg", extCnv: "bmp"}
	if err := c2.Convert(path); err == nil {
		t.Error("Convertはnilでないerrorを返すべきです")
	}

	// 対応するファイル形式間での変換処理が正常に終了することの確認
	c3 := &Converter{extSrc: "jpg", extCnv: "png"}
	if err := c3.Convert(path); err != nil {
		t.Error("encodeが返すerrorはnilであるべきです")
	}
}
