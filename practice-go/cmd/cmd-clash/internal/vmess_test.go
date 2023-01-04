package internal

import (
	"reflect"
	"testing"
)

func TestClashVmess_Parse(t *testing.T) {
	type fields struct {
		Name    string
		Type    string
		Server  string
		Port    string
		Uuid    string
		AlertId string
		Cipher  string
		Network string
	}
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ClashVmess
		wantErr bool
	}{
		{
			name:   "vmess 解析测试",
			fields: fields{},
			args:   args{s: "eyJwcyI6IkpNUyIsInBvcnQiOiIxMTExMSIsImlkIjoiMS0xLTEtMS0xIiwiYWlkIjowLCJuZXQiOiJ0Y3AiLCJ0eXBlIjoibm9uZSIsInRscyI6Im5vbmUiLCJhZGQiOiIxLjEuMS4xIn0="},
			want: &ClashVmess{
				Name:    "JMS",
				Type:    "vmess",
				Server:  "1.1.1.1",
				Port:    "11111",
				Uuid:    "1-1-1-1-1",
				AlterId: "0",
				Cipher:  "auto",
				Network: "tcp",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vm := &ClashVmess{
				Name:    tt.fields.Name,
				Type:    tt.fields.Type,
				Server:  tt.fields.Server,
				Port:    tt.fields.Port,
				Uuid:    tt.fields.Uuid,
				AlterId: tt.fields.AlertId,
				Cipher:  tt.fields.Cipher,
				Network: tt.fields.Network,
			}
			got, err := vm.Parse(tt.args.s)
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
