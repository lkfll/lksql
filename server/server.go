package server

import (
	"fmt"
	"net/http"

	"github.com/lkfll/lksql/server/controller"
)

// 启动服务器
func Run() {
	var html controller.Html
	http.HandleFunc("/", html.Index)
	http.HandleFunc("/index", html.Index)
	http.HandleFunc("/tables", Tables)
	http.ListenAndServe(":8080", nil)
}

func Tables(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "2222222")
}
