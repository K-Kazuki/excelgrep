package excelsearch

import (
	"reflect"
	"sort"
	"testing"
)

func TestFind(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name:    "case 1",
			args:    args{""},
			want:    []string{},
			wantErr: false,
		},
		{
			name: "case 2",
			args: args{"/path/to/sample"},
			want: []string{
				"sample_files/sample.xlsx",
				"sample_files/sample2.xlsx",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Find(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %T : %v, wantErr %T : %v", err, err, tt.wantErr, tt.wantErr)
				return
			}
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\nFind()\t%T : %v\nwant\t%T : %v\n", got, got, tt.want, tt.want)
			}
		})
	}
}
