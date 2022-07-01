package basesql

import (
	"fmt"

	"github.com/lkfll/lksql/analyze"
)

// 查询sql结构体
type SelectSql struct {
	Sql                 string // sql模板
	SqlClause_Sdistinct string // sdistinct 去重
	SqlClause_Select    string // select 结果字段
	SqlClause_Form      string // form 表名
	SqlClause_Join      string // join on 连接表
	SqlClause_Where     string // where 条件
	SqlClause_Group     string // group by 分组
	SqlClause_Having    string // having 条件
	SqlClause_Order     string // order by 排序
	SqlClause_limit     string // limit 分页
}

// select 去重 列1 as 别名1, 列2 as 别名2 ,...
// from 表1 as 表名1 left join 表2 as 表名2 on 表名1.字段=表名2.字段 ...
// where 条件1 and/or 条件2 and/or?...
// group by 列
// having 条件1 and/or 条件2 and/or ...
// order by 列1, 列2,...
// limit m,n ; 从m处开始获取n条
const SelectSQL string = // 查询语句
"SELECT %s %s\n" +       // 去重字段，结果
	"FROM  %s\n" + // 表名
	"%s\n" + // 连接表
	"%s\n" + // where
	"%s\n" + // group by分组
	"%s\n" + // having
	"%s\n" + // order bt排序
	"%s\n" //   limit 分页
// 创建查询basesql
// return: sql语句模板，select子句，表名，连接子句
func CreateSelectSQL(ana *analyze.AnalyzeType, fields ...*analyze.Field) SelectSql {
	var ret SelectSql
	ret.Sql = SelectSQL // 查询语句模板

	// 连接子句
	joins := ""
	for _, jf := range ana.Join {
		joins = fmt.Sprintf("\t%s%s\n", joins, jf.JoinSql)
	}
	ret.SqlClause_Join = joins

	// 查询 select子句
	retFields := "" // 字段
	for _, v := range fields {
		oneTableField := "" // 一个表的字段
		oneTableField = fmt.Sprintf("%s%s.%s , ", oneTableField, v.TableName, v.FieldName)
		retFields = fmt.Sprintf("%s%s", retFields, oneTableField)
	}
	ret.SqlClause_Select = retFields[:len(retFields)-2]

	// 表名 form子句
	ret.SqlClause_Form = ana.Key.TableName

	return ret
}
