// +build !release

package logger

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

var loggerInstance logger = &silentLogger{}

type logger interface {
	Debugln(v ...interface{})
	Debugf(format string, v ...interface{})
}

type logType int

const (
	Verbose logType = iota
	Silent
)

// ロガーを取得する
func SetLogger(lt logType) {
	switch lt {
	case Verbose:
		loggerInstance = &verboseLogger{}
	default:
		// &silentLogger を設定済み
	}
}

// 何も出力しない空のロガー
type silentLogger struct{}

func (sl *silentLogger) Debugln(v ...interface{})               {}
func (sl *silentLogger) Debugf(format string, v ...interface{}) {}

// log パッケージのラッパー
type verboseLogger struct{}

func (vl *verboseLogger) Debugln(v ...interface{}) {
	// 現在のスタックから情報を取得
	file, line, ok := getColler()
	if ok {
		s := getLogStyle(file, line)
		var i []interface{}
		v = append(i, s, v)
	}

	log.Println(v...)
}

func (vl *verboseLogger) Debugf(format string, v ...interface{}) {
	f := format

	// 現在のスタックから情報を取得
	file, line, ok := getColler()
	if ok {
		s := getLogStyle(file, line)
		f = fmt.Sprintf("%s %s", s, format)
	}

	log.Printf(f, v...)
}

// 現在のスタックから情報を取得する
func getColler() (string, int, bool) {
	_, file, line, ok := runtime.Caller(3)
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

func Debugln(v ...interface{}) {
	loggerInstance.Debugln(v)
}

func Debugf(format string, v ...interface{}) {
	loggerInstance.Debugf(format, v)
}
