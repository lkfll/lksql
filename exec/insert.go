package exec

import (
	"database/sql"
	"fmt"
)

// var _ common.Insert = (*Insert)(nil)

type Insert struct {
	Sql   string
	Param []interface{}
}

// 构造函数
func NewInsert(sql string, param ...interface{}) *Insert {
	var ret Insert
	ret.Sql = sql
	ret.Param = param
	return &ret
}

// 执行
func (Insert *Insert) Go(db SQLCommon) (sql.Result, error) {
	Insert.Sql = fmt.Sprint(Insert.Sql, ";")
	return db.Exec(Insert.Sql, Insert.Param...)
}
