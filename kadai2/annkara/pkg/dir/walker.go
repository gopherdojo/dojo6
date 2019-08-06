// Package dir provides directory walk function.
package dir

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/dojo6/kadai2/annkara/pkg/image"
)

// Walk only calls the filepath.Walk fuction and internally calls image.Convert function.
// The root parameter is the target root directory, the before is the file extension before converting image,
// the after is after coverting image.
func Walk(root, before, after string) error {

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		n := info.Name()
		if strings.HasSuffix(n, before) {
			origin, err := os.Open(path)
			if err != nil {
				return err
			}
			defer origin.Close()

			// 拡張子を含まない出力用ファイル名
			n := filepath.Base(n[:len(n)-len(filepath.Ext(n))])
			dir := filepath.Dir(path)
			out, err := os.Create(filepath.Join(dir, n+"."+after))
			if err != nil {
				return err
			}

			err = image.Convert(origin, out, after)
			if err != nil {
				// 変換処理に失敗した場合、不要なファイルが作成されてしまうため、削除する
				// ファイルを閉じた後でないと、Windowsの場合削除できないのでここでCloseする
				out.Close()
				e := os.Remove(filepath.Join(dir, n+"."+after))
				if e != nil {
					return e
				}
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
