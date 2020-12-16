// +build !release

package logger

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

func Debugln(v ...interface{}) {
	// 現在のスタックから情報を取得
	file, line, ok := getColler()
	if ok {
		s := getLogStyle(file, line)
		var i []interface{}
		v = append(i, s, v)
	}

	log.Println(v...)
}

func Debugf(format string, v ...interface{}) {
	// 現在のスタックから情報を取得
	file, line, ok := getColler()
	if ok {
		s := getLogStyle(file, line)
		var i []interface{}
		v = append(i, s, v)
	}

	log.Printf(format, v...)
}

// 現在のスタックから情報を取得する
func getColler() (string, int, bool) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return "", 0, ok
	}
	splitF := strings.Split(file, "/")
	f := splitF[len(splitF)-1]
	return f, line, ok
}

// log.LstdFlags | log.Lshortfile と同等の文字列を作成
func getLogStyle(file string, line int) string {
	// t := time.Now()
	// dt := t.Format("2006/01/02 15:04:05")
	// s := fmt.Sprintf("%s %s:%d", dt, file, line)
	s := fmt.Sprintf("%s:%d", file, line)
	return s
}
