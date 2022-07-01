package repository

import (
	"github.com/lkfll/lksql/analyze"
	"github.com/lkfll/lksql/exec"
)

// 按表格修改
// UPDATE user SET  user_name='迪迦奥特曼' , password='16191411qin' , role=0  WHERE id = 1 and id = 2;
func (Repository *Repository) UpdateByTable(tableName string, param analyze.Type) func(string) *exec.Update {
	sql := Repository.BaseSQL.Update[tableName] // 表对应的sql

	fields := Repository.BaseSQL.Update[tableName].UpdateFields // 表的字段

	ParamArr := make([]interface{}, 0) // 参数切片
	paramMap := param.ToMap()
	for _, f := range fields {
		ParamArr = append(ParamArr, paramMap[f.FieldName])
	}

	return exec.NewUpdate(sql.Sql, ParamArr...)
}
