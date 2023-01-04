package internal

import (
	"encoding/base64"
	"errors"
	"strings"
)

type ClashSS struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Server   string `yaml:"server"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Cipher   string `yaml:"cipher"`
	Network  string `yaml:"network"`
}

func (ss *ClashSS) Parse(s string) (*ClashSS, error) {
	res := strings.SplitN(s, "#", 2)
	if len(res) != 2 {
		return nil, errors.New("parse error")
	}
	d, _ := base64.RawStdEncoding.DecodeString(res[0])
	r := strings.SplitN(string(d), "@", 2)
	if len(r) != 2 {
		return nil, errors.New("parse error")
	}
	cipher := strings.SplitN(r[0], ":", 2)
	if len(cipher) != 2 {
		return nil, errors.New("parse error")
	}
	host := strings.SplitN(r[1], ":", 2)
	if len(host) != 2 {
		return nil, errors.New("parse error")
	}

	return &ClashSS{
		Name:     res[1],
		Type:     "ss",
		Server:   host[0],
		Port:     host[1],
		Cipher:   cipher[0],
		Password: cipher[1],
		Network:  "tcp",
	}, nil
}
