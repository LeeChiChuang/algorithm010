package dp

import (
	"reflect"
	"testing"
)

func TestMinPathSum(t *testing.T) {
	type args struct {
		a [][]int
	}
	arg1 := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	arg2 := [][]int{
		{1, 3, 5},
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{a: arg1}, 7},
		{"test2", args{a: arg2}, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinPathSum(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinPathSum2(t *testing.T) {
	type args struct {
		a [][]int
	}
	arg1 := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	arg2 := [][]int{
		{1, 3, 5},
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{a: arg1}, 7},
		{"test2", args{a: arg2}, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinPathSum2(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minPathSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}