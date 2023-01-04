package internal

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type ClashVmess struct {
	Name    string `yaml:"name"`
	Type    string `yaml:"type"`
	Server  string `yaml:"server"`
	Port    string `yaml:"port"`
	Uuid    string `yaml:"uuid"`
	AlterId string `yaml:"alterId"`
	Cipher  string `yaml:"cipher"`
	Network string `yaml:"network"`
}

type VmessConfig struct {
	Ps   string `json:"ps"`
	Port string `json:"port"`
	Id   string `json:"id"`
	Aid  int    `json:"aid"`
	Net  string `json:"net"`
	Type string `json:"json"`
	Tls  string `json:"tls"`
	Add  string `json:"add"`
}

func (vm *ClashVmess) Parse(s string) (*ClashVmess, error) {
	var p VmessConfig
	//要忽略base64字符串后面==
	b, _ := base64.RawStdEncoding.DecodeString(s)
	err := json.Unmarshal(b, &p)
	if err != nil {
		return nil, err
	}
	return &ClashVmess{
		Name:    p.Ps,
		Type:    "vmess",
		Server:  p.Add,
		Port:    p.Port,
		Uuid:    p.Id,
		AlterId: fmt.Sprintf("%d", p.Aid),
		Cipher:  "auto",
		Network: p.Net,
	}, err
}
