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

// 拼接sql
func (delete *Delete) MakeSql() {
	delete.Sql = fmt.Sprint(delete.Sql, "\n", delete.SqlClause_Where, ";")
}

func (delete *Delete) Go() Go {
	delete.MakeSql()
	return func(db SQLCommon, param ...interface{}) (sql.Result, error) {
		Hook(delete.Sql)
		return db.Exec(delete.Sql, param...)
	}
}
