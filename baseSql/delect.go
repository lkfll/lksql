package basesql

import "fmt"

// 查询sql结构体
type DeleteSql struct {
	Sql string // sql模板
}

// 创建删除basesql
const DeleteSQL string = "DELETE FROM %s " // 删除语句
func CreateDeleteSQL(tableName string) DeleteSql {
	var ret DeleteSql
	ret.Sql = fmt.Sprintf(DeleteSQL, tableName)
	return ret
}
