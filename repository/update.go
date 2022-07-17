package repository

import (
	"reflect"

	"github.com/lkfll/lksql/analyze"
	"github.com/lkfll/lksql/exec"
)

// // 按表格修改
// // UPDATE user SET  user_name='迪迦奥特曼' , password='16191411qin' , role=0  WHERE id = 1 and id = 2;
// func (Repository *Repository) UpdateByTable(tableName string, param analyze.Type) func(string) *exec.Update {
// 	sql := Repository.BaseSQL.Update[tableName] // 表对应的sql

// 	fields := Repository.BaseSQL.Update[tableName].UpdateFields // 表的字段

// 	ParamArr := make([]interface{}, 0) // 参数切片
// 	paramMap := param.ToMap()
// 	for _, f := range fields {
// 		ParamArr = append(ParamArr, paramMap[f.FieldName])
// 	}

// 	return exec.NewUpdate(sql.Sql, ParamArr...)
// }

// 按表格修改
// UPDATE user SET  user_name='迪迦奥特曼' , password='16191411qin' , role=0  WHERE id = 1 and id = 2;
func (Repository *EntityRepository) Update(param analyze.Type) func(string) *exec.Update {
	sql := Repository.EntitySql.Update // 表对应的sql

	fields := Repository.EntitySql.Update.UpdateFields // 表的字段

	ParamList := make([]interface{}, 0)  // 参数切片
	paramValue := reflect.ValueOf(param) // 获取一行参数value反射
	for _, v := range fields {
		// 通过反射获取参数列表
		ParamList = append(ParamList, paramValue.FieldByName(v.Name).Interface())
	}

	return exec.NewUpdate(sql.Sql, ParamList...)
}
