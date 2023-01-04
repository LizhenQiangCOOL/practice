package internal

import (
	"encoding/base64"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Update struct {
	Url           string
	ClashFilePath string
}

func (u *Update) getData() (string, error) {
	resp, err := http.Get(u.Url)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		fmt.Sprintf("get data error, StatusCode: %v", resp.StatusCode)
	}
	//base64 decode body
	b, _ := base64.StdEncoding.DecodeString(string(body))
	return string(b), nil
}

func (u *Update) Main() error {
	d, err := u.getData()
	if err != nil {
		return err
	}
	ss := make([]ClashSS, 0)
	vmess := make([]ClashVmess, 0)
	s := strings.Fields(d)
	for _, k := range s {
		//可以使用简单工厂进行优化
		if strings.HasPrefix(k, "ss://") {
			sstmp := ClashSS{}
			nk := strings.TrimPrefix(k, "ss://")
			res, err := sstmp.Parse(nk)
			if err != nil {
				fmt.Sprintf("%v ss解析失败", k)
			}
			ss = append(ss, *res)
		} else if strings.HasPrefix(k, "vmess://") {
			vmesstmp := ClashVmess{}
			nk := strings.TrimPrefix(k, "vmess://")
			res, err := vmesstmp.Parse(nk)
			if err != nil {
				fmt.Sprintf("%v vmes解析失败", k)
			}
			vmess = append(vmess, *res)
		}
	}
	res := make([]interface{}, len(ss)+len(vmess))
	for i, k := range ss {
		res[i] = k
	}
	for i, k := range vmess {
		res[i+len(ss)] = k
	}
	c := ClashConfig{
		Port:               7890,
		SocksPort:          7891,
		RedirPort:          7892,
		AllowLan:           true,
		Mode:               "Rule",
		LogLevel:           "info",
		ExternalController: "127.0.0.1:9090",
		CfwConnBreakStrategy: CfwConnBreakStrategy{
			Proxy:   "all",
			Profile: true,
			mode:    true,
		},
		CfwLatencyTimeout: "5000",
		Proxies:           res,
	}
	r, err := yaml.Marshal(c)
	if err != nil {
		fmt.Sprintf("yaml marsh error: %v", err)
	}
	os.WriteFile(u.ClashFilePath, r, 0644)
	return nil
}
