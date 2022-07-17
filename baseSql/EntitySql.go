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
func EntitySqlStart(analyzeType *analyze.AnalyzeType) *EntitySql {
	var ret EntitySql

	// 生成增删改查sql
	ret.Insert = CreateInsertSQL(analyzeType.Key.TableName, analyzeType.InsertField...)
	ret.Update = CreateUpdateSQL(analyzeType.Key.TableName, analyzeType.UpdateField...)
	ret.Delete = CreateDeleteSQL(analyzeType.Key.TableName)
	ret.SelectSql = CreateSelectSQL(analyzeType, analyzeType.SelectField...)

	return &ret
}
