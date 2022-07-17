package exec

import (
	"database/sql"
	"fmt"
)

// var _ common.Exec = (*Delete)(nil)

type Delete struct {
	Sql             string
	SqlClause_Where string // where子句
}

// 构造函数
func NewDelete(sql string) func(string) *Delete {
	var ret Delete
	ret.Sql = sql
	return ret.Where
}

// 设置sql where 部分
func (delete *Delete) Where(SqlClause_Where string) *Delete {
	delete.SqlClause_Where = fmt.Sprint("WHERE ", SqlClause_Where)
	return delete
}

// 执行
func (delete *Delete) Go(db SQLCommon, param ...interface{}) (sql.Result, error) {
	delete.Sql = fmt.Sprint(delete.Sql, "\n", delete.SqlClause_Where, ";")
	return db.Exec(delete.Sql, param...)
}

// 拼接sql
func (delete *Delete) MakeSql() {
	delete.Sql = fmt.Sprint(delete.Sql, "\n", delete.SqlClause_Where)
}

// 保存 省略 隐藏细节
type SaveDelete struct {
	delete *Delete
}

func (delete *Delete) Save() *SaveDelete {
	var ret SaveDelete
	ret.delete = delete
	return &ret
}

// 执行
func (delete *SaveDelete) Go(db SQLCommon, param ...interface{}) (sql.Result, error) {
	delete.delete.Sql = fmt.Sprint(delete.delete.Sql, ";")
	return db.Exec(delete.delete.Sql, param...)
}
func (delete *SaveDelete) GetSql() string {
	return delete.delete.Sql
}
