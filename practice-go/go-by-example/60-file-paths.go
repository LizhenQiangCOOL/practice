package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	// 拼接文件路径
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// 路径拼接会拼接处最优路径
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// 从路径分离出 目录 和 文件名
	fmt.Println("Dir(p)", filepath.Dir(p))
	fmt.Println("Base(p)", filepath.Base(p))

	// 判断是否为绝对地址
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file")) //window下判断为非绝对地址

	filename := ".jsonconfig.json"

	// 取文件后缀名称
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// 删除后缀
	fmt.Println(strings.TrimSuffix(filename, ext))

	// targpath 在 basepath 的rel 位置距离
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
