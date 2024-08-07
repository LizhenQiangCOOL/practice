package sort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
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
				A: RandomSlice(6),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, r := 0, len(tt.args.A)-1
			QuickSort(tt.args.A, p, r)
			if !CheckIsSorted(tt.args.A) {
				t.Errorf("sort error, Slice: %v", tt.args.A)
			}
		})
	}
}
