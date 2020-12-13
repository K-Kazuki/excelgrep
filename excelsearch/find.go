package excelsearch

import (
	"fmt"
	"strings"

	"github.com/mattn/go-zglob"
)

func Find(path string) ([]string, error) {
	var (
		res []string
		err error
	)

	// パスが渡されたら *.xlsx を結合して Glob
	// デフォルトはカレントディレクトリ配下すべて
	if len(path) > 0 {
		// パスの最後のスラッシュを考慮
		if strings.HasSuffix(path, "/") {
			path += "*.xlsx"
		} else {
			path += "/*.xlsx"
		}
		res, err = zglob.GlobFollowSymlinks(path)
	} else {
		res, err = zglob.GlobFollowSymlinks("./**/*.xlsx")
		fmt.Println("path len 0")
	}
	if err != nil {
		return nil, err
	}

	var files []string
	for _, path := range res {
		splitePath := strings.Split(path, "/")
		file := splitePath[len(splitePath)-1]
		if !strings.HasPrefix(file, "~$") {
			files = append(files, path)
		}
	}
	return files, nil
}
