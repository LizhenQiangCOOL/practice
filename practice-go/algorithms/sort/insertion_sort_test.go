package sort

import "testing"

func TestInsertionSort(t *testing.T) {
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
				A: RandomSlice(100),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertionSort(tt.args.A)
			if !CheckIsSorted(tt.args.A) {
				t.Errorf("sort error, Slice: %v", tt.args.A)
			}
		})
	}
}
