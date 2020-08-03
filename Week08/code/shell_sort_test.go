package sort

import (
	"reflect"
	"testing"
)

func TestShellSort(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{a: []int{7, 8, 9, 2, 3, 5}}, []int{2, 3, 5, 7, 8, 9}},
		{"", args{a: []int{9, 8, 7, 6, 5, 4}}, []int{4, 5, 6, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShellSort(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShellSort() = %v, want %v", got, tt.want)
			}
		})
	}
}