package exec

import (
	"database/sql"
	"fmt"
)

// SQLCommon 是进行sql操作的公共的接口 ，
// 提取了 *sql.DB 和 *sql.Tx 公共操作
// sql处理函数获取连接应该是这个接口,可以实现事务的处理
var _ SQLCommon = (*sql.DB)(nil) // 接口检测
var _ SQLCommon = (*sql.Tx)(nil) // 接口检测
type SQLCommon interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// // 执行接口 保存
// type Exec interface {
// 	Go(db SQLCommon, param ...interface{}) Go
// }

type QueryGo func(db SQLCommon, param ...interface{}) ([]interface{}, error)
type Go func(db SQLCommon, param ...interface{}) (sql.Result, error)

var Hook func(string) = func(sql string) {
	fmt.Printf("sql: %v\n", sql)
}
