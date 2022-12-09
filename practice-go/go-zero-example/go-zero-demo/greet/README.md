

## 流量访问流程
get 请求 -> 路由匹配 -> GeetHandler -> 参数解析httpx.Parse[Path、Form、Headers、jsonBody] -> 函数逻辑 -> 返回结果


## 服务器启动加载流程
命令行解析参数 -> 加载配置文件 -> 初始化服务器 -> 挂在路由 -> 服务器启动


## 编码流程
1. 出入参数编写， types
2. 核心逻辑编写， logic
3. 逻辑函数挂在到路由