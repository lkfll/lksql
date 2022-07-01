package basesql

import (
	"fmt"

	"github.com/lkfll/lksql/analyze"
)

// 查询sql结构体
type UpdateSql struct {
	Sql          string           // sql模板
	UpdateFields []*analyze.Field // 修改字段
}

// 创建修改basesql
const UpdateSQL string = "UPDATE %s SET %s " // 修改语句
func CreateUpdateSQL(tableName string, field ...*analyze.Field) UpdateSql {
	var updateField string = "" // 修改字段
	for _, v := range field {   // 增加对应的字段
		updateField = fmt.Sprintf("%s %s=%s ,", updateField, v.FieldName, "?")
	}
	var ret UpdateSql
	ret.UpdateFields = make([]*analyze.Field, len(field))
	copy(ret.UpdateFields, field)
	ret.Sql = fmt.Sprintf(UpdateSQL, tableName, updateField[:len(updateField)-1])
	return ret
}
