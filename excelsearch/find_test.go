package excelsearch

import (
	"fmt"
	"testing"
)

func Test_linuxFind(t *testing.T) {
	type args struct {
		path string
		sep  string
	}
	tests := []struct {
		name string
		args args
	}{
		{"case1", args{".", "*"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := linuxFind(tt.args.path, tt.args.sep)
			fmt.Printf("ls result: \n%s", string(out))
		})
	}
}
