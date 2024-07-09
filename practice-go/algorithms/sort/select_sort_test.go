package sort

import (
	"testing"
)

func TestSelectSort(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "random test",
			args: args{
				A: RandomSlice(10),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SelectSort(tt.args.A)
			if !CheckIsSorted(tt.args.A) {
				t.Errorf("sort error, Slice: %v", tt.args.A)
			}
		})
	}
}
