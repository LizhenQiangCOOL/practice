package go_safety_code

import (
	"fmt"
	"log"
	"net/http"
)

// Xss
// 主要常见于贴心的报错反馈上
// 1. XXX文件不存在  2. XXX错误等
// 输入： <script>alert("xss")</script>
// 应该 前后端数据分开渲染，使用Vue、React等前端开发框架自带的Xss防御机制
func Handle(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatal(nil)
	}
	user_pro := req.FormValue("user")
	fmt.Fprintf(w, "%s\n", user_pro)
}
