package dp

import (
	"reflect"
	"testing"
)

func TestCoinChange(t *testing.T) {
	type args struct {
		a []int
	}

	arg1 := []int{1, 2, 5}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{a:arg1}, 3 },
	}

	for _, tt := range tests  {
		t.Run(tt.name, func(t *testing.T) {
			if got := CoinChange(tt.args.a, 11); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CoinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoinChange2(t *testing.T) {
	type args struct {
		a []int
	}

	arg1 := []int{1, 2, 5}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{a:arg1}, 3 },
	}

	for _, tt := range tests  {
		t.Run(tt.name, func(t *testing.T) {
			if got := CoinChange2(tt.args.a, 11); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CoinChange2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoinChange3(t *testing.T) {
	type args struct {
		a []int
	}

	arg1 := []int{1, 2, 5}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{a:arg1}, 3 },
	}

	for _, tt := range tests  {
		t.Run(tt.name, func(t *testing.T) {
			if got := CoinChange3(tt.args.a, 11); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CoinChange3() = %v, want %v", got, tt.want)
			}
		})
	}
}
