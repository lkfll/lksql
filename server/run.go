package server

import (
	"database/sql"
	"net/http"

	"github.com/lkfll/lksql/server/controller"
)

// 启动服务器
func Run(DB *sql.DB, port string) {
	var html controller.Html
	http.HandleFunc("/index", html.Index)

	var ddl controller.DDL
	ddl.DB = DB
	http.HandleFunc("/createTable", ddl.CreateTable)
	http.HandleFunc("/deseTable", ddl.DeseTable)
	http.HandleFunc("/deseTables", ddl.DeseTables)
	http.HandleFunc("/showTables", ddl.ShowTables)
	http.HandleFunc("/dropTable", ddl.DropTable)

	http.ListenAndServe(":"+port, nil)
}
