package dp

import (
	"reflect"
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	type args struct {
		a []int
	}
	arg1 := []int{10, 9, 2, 5, 3, 7, 101, 18}
	arg2 := []int{1, 3, 5}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{a: arg1}, 4},
		{"test2", args{a: arg2}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLIS(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LongestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
