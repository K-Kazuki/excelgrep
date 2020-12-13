package excelsearch

import (
	"reflect"
	"testing"
)

func Test_search(t *testing.T) {
	type args struct {
		s   string
		sep string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"ケース１",
			args{"検索対象の文字列", "検索"},
			"検索対象の文字列",
		},
		{
			"ケース2",
			args{"検索対象の文字列2", "検索"},
			"検索対象の文字列2",
		},
		{
			"ケース3",
			args{"検索対象の文字列3", "検索"},
			"検索対象の文字列3",
		},
	}
	for _, tt := range tests {
		got := []byte(search(tt.args.s, tt.args.sep))
		expected := []byte(tt.want)
		t.Run(tt.name, func(t *testing.T) {
			if reflect.DeepEqual(got, expected) {
				t.Errorf("search() = %v, want %v", got, expected)
			}
		})
	}
}
