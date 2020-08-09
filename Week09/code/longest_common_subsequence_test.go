package dp

import (
	"reflect"
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{a:"abcde", b:"ace"}, 3 },
		{"test2", args{a:"abc", b:"def"}, 0 },
	}

	for _, tt := range tests  {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestCommonSubsequence(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LongestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
