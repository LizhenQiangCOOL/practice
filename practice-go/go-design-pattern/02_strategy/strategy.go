package strategy

import (
	"fmt"
	"os"
)

//策略模式

// 存储策略
type Storage interface {
	Save(name string, data []byte) error
}

var strategys = map[string]Storage{
	"file":         &fileStorage{},
	"encrypt_file": &encryptFileStorage{},
}

func NewStorageStrategy(t string) (Storage, error) {
	s, ok := strategys[t]
	if !ok {
		return nil, fmt.Errorf("not found Storage: %s", t)
	}
	return s, nil
}

// 保存到文件的算法
type fileStorage struct {
}

func (s fileStorage) Save(name string, data []byte) error {
	// go 1.6+
	return os.WriteFile(name, data, os.ModePerm)
}

// 加密保存到文件的算法
type encryptFileStorage struct {
}

func (s encryptFileStorage) Save(name string, data []byte) error {
	//加密
	data, err := encrypt(data)
	if err != nil {
		return err
	}
	return os.WriteFile(name, data, os.ModeAppend)
}

func encrypt(data []byte) ([]byte, error) {
	// 加密具体实现
	return data, nil
}
