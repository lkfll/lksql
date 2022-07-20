package basesql

import "github.com/lkfll/lksql/analyze"

// 实体sql
type EntitySql struct {
	Insert    InsertSql // 增加语句
	Delete    DeleteSql // 删除语句
	Update    UpdateSql // 修改语句
	SelectSql SelectSql // 查询语句
	// TODO 扩展
}

// EntitySql 工厂
func EntitySqlStart(analyze analyze.Analyze) *EntitySql {
	var ret EntitySql

	// 生成增删改查sql
	ret.Insert = CreateInsertSQL(analyze.GetTableName(), analyze.GetInsertField()...)
	ret.Update = CreateUpdateSQL(analyze.GetTableName(), analyze.GetUpdateField()...)
	ret.Delete = CreateDeleteSQL(analyze.GetTableName())
	ret.SelectSql = CreateSelectSQL(analyze.GetTableName(), analyze.GetJsonSqls(), analyze.GetSelectField())

	return &ret
}
