package or_channel

import "testing"

func Test_my(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"Test_my"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			my()
		})
	}
}
