package internal

import (
	"reflect"
	"testing"
)

func TestClashSS_Parse(t *testing.T) {
	type fields struct {
		Name     string
		Type     string
		Server   string
		Port     string
		Password string
		Cipher   string
		Network  string
	}
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ClashSS
		wantErr bool
	}{
		{
			name:   "ss解析测试",
			fields: fields{},
			args:   args{s: "YWVzLTI1Ni1nY206cGFzc3dvcmRAMS4xLjEuMToxMTEx#JMS"},
			want: &ClashSS{
				"JMS",
				"ss",
				"1.1.1.1",
				"1111",
				"password",
				"aes-256-gcm",
				"tcp",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := &ClashSS{
				Name:     tt.fields.Name,
				Type:     tt.fields.Type,
				Server:   tt.fields.Server,
				Port:     tt.fields.Port,
				Password: tt.fields.Password,
				Cipher:   tt.fields.Cipher,
				Network:  tt.fields.Network,
			}
			got, err := ss.Parse(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
