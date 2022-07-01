package basesql

import (
	"fmt"

	"github.com/lkfll/lksql/analyze"
)

// 查询sql结构体
type InsertSql struct {
	Sql          string           // sql模板
	InsertFields []*analyze.Field // 插入字段
}

// 创建插入的basesql
// 创造insertSQL
const InsertSQL string = "INSERT INTO %s(%s) VALUES %s " // 增加语句
func CreateInsertSQL(tableName string, field ...*analyze.Field) InsertSql {
	var temp string = ""
	for _, v := range field { // 增加对应的字段
		temp = fmt.Sprintf("%s , %s ", temp, v.FieldName)
	}
	var ret InsertSql
	ret.Sql = fmt.Sprintf(InsertSQL, tableName, temp[2:], "%s")
	ret.InsertFields = make([]*analyze.Field, len(field))
	copy(ret.InsertFields, field)
	return ret
}
