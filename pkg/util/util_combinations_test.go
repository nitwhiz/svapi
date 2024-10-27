package util

import (
	"reflect"
	"testing"
)

func TestCombinations(t *testing.T) {
	type args struct {
		values []string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1 element", args{[]string{"A"}}, []string{"A"}},
		{"2 elements", args{[]string{"A", "B"}}, []string{"A_", "_B", "A_B"}},
		{"3 elements", args{[]string{"A", "B", "C"}}, []string{"A__", "_B_", "A_B_", "__C", "A__C", "_B_C", "A_B_C"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Combinations(tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Combinations() = %v, want %v", got, tt.want)
			}
		})
	}
}
