package go_safety_code

import (
	"fmt"
	"os"
	"os/exec"
)

// 命令注入
// path 未经过过滤就拼接到指令中， 注入例子 path = "/ && whoami"
func errexample_exec_cmd(path string) {
	cmd_str := "ls" + path
	cmd := exec.Command("bash", "-c", cmd_str)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)
}

// 需求是列出目录, 其他命令也类似，注意命令本身会不会被利用
// 参考 https://gtfobins.github.io/
func safe_exec_cmd(path string) {
	output, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range output {
		fmt.Println(string(file.Name()))
	}
}
