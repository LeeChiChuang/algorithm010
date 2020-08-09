package dp

import (
	"reflect"
	"testing"
)

func TestMinimumTotal(t *testing.T) {
	type args struct {
		a [][]int
	}

	arg1 := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{a:arg1}, 11 },
	}

	for _, tt := range tests  {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinimumTotal(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinimumTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}
