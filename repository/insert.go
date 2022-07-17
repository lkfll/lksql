package repository

import (
	"fmt"
	"reflect"

	"github.com/lkfll/lksql/analyze"
	"github.com/lkfll/lksql/exec"
)

// // 按表格插入
// // values 传入参数使用map 使用map代替对象
// // INSERT INTO user( user_name  , password  , role ) VALUES  ( '小明' , '16191411qin' , 0  ) , ( '小明' , '16191411qin' , 0  )
// func (Repository *Repository) InsertByTable(tableName string, param ...analyze.Type) *exec.Insert {
// 	sql := Repository.BaseSQL.Insert[tableName] // 表对应的sql

// 	value := "(" // 拼接增加的一行value
// 	for i := 0; i < len(Repository.BaseSQL.Insert[tableName].InsertFields); i++ {
// 		value = fmt.Sprintf("%s %s,", value, "?")
// 	}
// 	value = fmt.Sprintf("%s %s", value[:len(value)-1], ")") // 一行value

// 	var valuesSql string = ""
// 	ParamArr := make([]interface{}, 0) // 参数切片
// 	for i := 0; i < len(param); i++ {  // 传入参数的长度
// 		valuesSql = fmt.Sprintf("%s %s,", valuesSql, value)                   // 拼接sql
// 		paramMap := param[i].ToMap()                                          // 拼接参数
// 		for _, v := range Repository.BaseSQL.Insert[tableName].InsertFields { //
// 			ParamArr = append(ParamArr, paramMap[v.FieldName])
// 		}
// 	}
// 	tempsql := fmt.Sprintf(sql.Sql, valuesSql[:len(valuesSql)-1])

// 	return exec.NewInsert(tempsql, ParamArr...)
// }

// 按表格插入
// values 传入参数使用map 使用map代替对象
// INSERT INTO user( user_name  , password  , role ) VALUES  ( '小明' , '16191411qin' , 0  ) , ( '小明' , '16191411qin' , 0  )
func (Repository *EntityRepository) Insert(params ...analyze.Type) *exec.Insert {
	sql := Repository.EntitySql.Insert // 表对应的sql

	value := "(" // 拼接增加的一行value
	for i := 0; i < len(Repository.EntitySql.Insert.InsertFields); i++ {
		value = fmt.Sprintf("%s %s,", value, "?")
	}
	value = fmt.Sprintf("%s %s", value[:len(value)-1], ")") // 一行value

	var valuesSql string = ""
	ParamList := make([]interface{}, 0) // 参数切片
	for i := 0; i < len(params); i++ {  // 传入参数的长度
		valuesSql = fmt.Sprintf("%s %s,", valuesSql, value) // 拼接sql
		param := reflect.ValueOf(params[i])                 // 获取一行参数value反射
		for _, v := range Repository.EntitySql.Insert.InsertFields {
			// 通过反射获取参数列表
			ParamList = append(ParamList, param.FieldByName(v.Name).Interface())
		}
	}
	tempsql := fmt.Sprintf(sql.Sql, valuesSql[:len(valuesSql)-1])

	return exec.NewInsert(tempsql, ParamList...)
}
