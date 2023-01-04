package internal

import "testing"

func TestUpdate_main(t *testing.T) {
	type fields struct {
		Url           string
		ClashFilePath string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"test",
			fields{
				Url:           "testurl",
				ClashFilePath: "",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Update{
				Url:           tt.fields.Url,
				ClashFilePath: tt.fields.ClashFilePath,
			}
			if err := u.Main(); (err != nil) != tt.wantErr {
				t.Errorf("main() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
