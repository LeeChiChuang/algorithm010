package dp

import (
	"reflect"
	"testing"
)

func TestMaxSumArray(t *testing.T) {
	type args struct {
		a []int
	}
	arg1 := []int{-2,1,-3,4,-1,2,1,-5,4}
	arg2 := []int{-11, -3, -5}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{a: arg1}, 6},
		{"test2", args{a: arg2}, -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSumArray(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxSumArray() = %v, want %v", got, tt.want)
			}
		})
	}
}