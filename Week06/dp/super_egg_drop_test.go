package dp

import (
	"reflect"
	"testing"
)

func TestSuperEggDrop(t *testing.T) {
	type args struct {
		a []int
	}
	arg1 := []int{1, 2, 3, 1}
	arg3 := []int{2, 1, 1, 2}
	arg4 := []int{4,1,2,7,5,3,1}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{a: arg1}, 4},
		{"test2", args{a: arg2}, 12},
		{"test3", args{a: arg3}, 4},
		{"test4", args{a: arg4}, 14},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rob(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
