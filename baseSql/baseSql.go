// 通过start 获得Basesql对象 交与operate层操作
package basesql

import "github.com/lkfll/lksql/analyze"

// 进行数据库操作的sql语句
// 由工厂生产出来的数据库操作对象，所拥有的操作语句
// 增加修改 默认不带id字段 id自增
// key 表名 value sql
type BaseSQL struct { // 基础语句，和又进行拼接的提交语句
	Delete    map[string]DeleteSql // 删除语句
	Insert    map[string]InsertSql // 增加语句
	Update    map[string]UpdateSql // 修改语句
	SelectSql SelectSql            // 查询语句

	// TODO 扩展
}

// baseSql 工厂
func Start(analyzeType *analyze.AnalyzeType) *BaseSQL {
	var ret BaseSQL
	ret.Update = make(map[string]UpdateSql) // 初始化map
	ret.Insert = make(map[string]InsertSql)
	ret.Delete = make(map[string]DeleteSql)
	for _, table := range analyzeType.TableNames { // 分表
		insertField := make([]*analyze.Field, 0)
		updateField := make([]*analyze.Field, 0)
		// 增加字段
		for _, f := range analyzeType.InsertField {
			if f.TableName == table { // 符合当前表名的
				insertField = append(insertField, f)
			}
		}
		// 修改字段
		for _, f := range analyzeType.UpdateField {
			if f.TableName == table { // 符合当前表名的
				updateField = append(insertField, f)
			}
		}
		ret.Insert[table] = CreateInsertSQL(table, insertField...)
		ret.Update[table] = CreateUpdateSQL(table, updateField...)
		ret.Delete[table] = CreateDeleteSQL(table)
	}

	// 生成查询sql
	ret.SelectSql = CreateSelectSQL(analyzeType, analyzeType.SelectField...)
	// TODO 扩展
	return &ret
}
