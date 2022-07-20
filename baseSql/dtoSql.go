// 通过start 获得Basesql对象 交与operate层操作
package basesql

import "github.com/lkfll/lksql/analyze"

// 进行数据库操作的sql语句
type DtoSql struct {
	EntitySqls map[string]*EntitySql // 多个小型的entitySql
	SelectSql  SelectSql             // dto查询语句
}

// DtoSql 工厂
func DtoStart(ana analyze.Analyze) *DtoSql {
	var ret DtoSql
	ret.EntitySqls = make(map[string]*EntitySql)
	for _, table := range ana.GetAllTableName() { // 分表
		var entitySql EntitySql

		// 增加
		insertField := make([]*analyze.Field, 0)
		for _, f := range ana.GetInsertField() {
			if f.TableName == table { // 符合当前表名的
				insertField = append(insertField, f)
			}
		}
		// 修改
		updateField := make([]*analyze.Field, 0)
		for _, f := range ana.GetUpdateField() {
			if f.TableName == table { // 符合当前表名的
				updateField = append(updateField, f)
			}
		}
		entitySql.Insert = CreateInsertSQL(table, insertField...)
		entitySql.Update = CreateUpdateSQL(table, updateField...)
		entitySql.Delete = CreateDeleteSQL(table)
		// ret.Select == nil

		ret.EntitySqls[table] = &entitySql
	}

	// 生成dto查询sql
	ret.SelectSql = CreateSelectSQL(ana.GetTableName(), ana.GetJsonSqls(), ana.GetSelectField())

	return &ret
}
