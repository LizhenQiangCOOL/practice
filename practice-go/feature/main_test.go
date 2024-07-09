package feature

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		Arr []int
		A   int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		// TODO: Add test cases.
		{
			name:  "case 1",
			args:  args{Arr: []int{1, 2, 3, 3, 4, 5, 6, 7}, A: 6},
			want:  0,
			want1: 5,
		},
		{
			name:  "case 1",
			args:  args{Arr: []int{}, A: 6},
			want:  -1,
			want1: -1,
		},
		{
			name:  "case 1",
			args:  args{Arr: []int{1, 2}, A: 3},
			want:  0,
			want1: 1,
		}, {
			name:  "case 1",
			args:  args{Arr: []int{1, 2}, A: 3},
			want:  0,
			want1: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Solution3(tt.args.Arr, tt.args.A)
			if got != tt.want {
				t.Errorf("solution3() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("solution3() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
