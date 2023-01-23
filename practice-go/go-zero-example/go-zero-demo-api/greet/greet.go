package main

import (
	"flag"
	"fmt"

	"go-zero-demo/greet/internal/config"
	"go-zero-demo/greet/internal/handler"
	"go-zero-demo/greet/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/greet-api.yaml", "the config file")

func main() {
	//解析命令
	flag.Parse()

	// 读取并映射配置文件到 config 结果体
	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 初始化服务
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 挂在 处理程序
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 启动服务
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
