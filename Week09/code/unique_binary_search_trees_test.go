package dp

import (
	"reflect"
	"testing"
)

func TestNumTrees(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{a: 3}, 5},
		{"test2", args{a: 10}, 16796},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumTrees(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NumTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumTrees2(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{a: 3}, 5},
		{"test2", args{a: 10}, 16796},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumTrees2(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NumTrees2() = %v, want %v", got, tt.want)
			}
		})
	}
}
