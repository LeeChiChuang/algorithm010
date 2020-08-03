package sort

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{a: []int{7, 8, 9, 2, 3, 5}}, []int{2, 3, 5, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectionSort(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectionSort() input = %v output = %v, want %v", tt.args.a, got, tt.want)
			}
		})
	}
}
