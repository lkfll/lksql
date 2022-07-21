package exec

import (
	"database/sql"
	"fmt"
)

// var _ common.Update = (*Update)(nil)
type Update struct {
	Sql             string        // sql语句
	Param           []interface{} // 参数列表
	SqlClause_Where string
}

// 构造函数
func NewUpdate(sql string, param ...interface{}) func(string) *Update {
	var ret Update
	ret.Param = make([]interface{}, 0)
	ret.Param = append(ret.Param, param...)
	ret.Sql = sql
	return ret.Where
}

// 执行
func (Update *Update) Go(db SQLCommon, param ...interface{}) (sql.Result, error) {
	Update.Sql = fmt.Sprint(Update.Sql, "\n", Update.SqlClause_Where, ";")
	param = append(Update.Param, param...)
	Hook(Update.Sql)
	return db.Exec(Update.Sql, param...)
}

// 设置sql where 部分
func (Update *Update) Where(SqlClause_Where string) *Update {
	Update.SqlClause_Where = fmt.Sprint("WHERE ", SqlClause_Where)
	return Update
}
