package exec

import (
	"database/sql"
	"fmt"
	"strings"

	basesql "github.com/lkfll/lksql/baseSql"
)

// 扫描方法
type ScanHandle func(...string) interface{}

// go方法的钩子函数
var Hook_QueryGoBefore func(*Query, ...interface{}) = func(query *Query, param ...interface{}) {
	fmt.Printf("Sql: %v\n", query.Sql)
	fmt.Printf("param: %v\n", param)
}

// go方法的钩子函数
var Hook_QueryGoAfter func([]interface{}, error) = func(i []interface{}, err error) {
	fmt.Printf("i: %v\n", i)
	fmt.Printf("err: %v\n", err)
}

// 查询
// (1) form
// (3) join on
// (4) where （可以使用表的别名）
// (7) select
// (8) distinct
// (9) order by
// (10) limit
type Query struct {
	IsGroup bool // 是否是分组查询
	RowNum  int  // 一行数量

	basesql.SelectSql // 查询结构体

	_scanHandle ScanHandle // 扫描方法

}

// 构造函数
func NewQuery(selectSql basesql.SelectSql, scanHandle ScanHandle) *Query {
	var ret Query
	ret.IsGroup = false // 默认不是分组查询
	ret.SelectSql = selectSql
	ret._scanHandle = scanHandle
	return &ret
}

//  go 执行sql
// 执行sql并且扫描结果
// 将GetResult 和 Scan 合并
func (query *Query) Go(db SQLCommon, param ...interface{}) ([]interface{}, error) {
	query.MakeSql()
	Hook_QueryGoBefore(query) // 钩子
	rows, err := db.Query(query.Sql, param...)
	if err != nil {
		Hook_QueryGoAfter(nil, err) // 钩子
		return nil, err
	}
	defer rows.Close() // 结束关闭rows
	i, err := query.Scan(rows)
	Hook_QueryGoAfter(i, err) // 钩子
	return i, err
}

// 获得结果集合
// 记得  r*sql.Rows.Close
// 调用了 Makesql
func (query *Query) GetResult(db SQLCommon) (*sql.Rows, error) {
	query.MakeSql()

	return db.Query(query.Sql)
}

// 组合sql
func (query *Query) MakeSql() {
	// 拼接sql
	query.Sql = fmt.Sprintf(query.Sql, query.SqlClause_Sdistinct,
		query.SqlClause_Select, query.SqlClause_Form, query.SqlClause_Join,
		query.SqlClause_Where, query.SqlClause_Group, query.SqlClause_Having,
		query.SqlClause_Order, query.SqlClause_limit)
}

// 扫描rows分析
func (query *Query) Scan(rows *sql.Rows) ([]interface{}, error) {
	query.RowNum = len(strings.Split(query.SqlClause_Select, ",")) // 设置一行的数量
	ret := make([]interface{}, 0)
	for rows.Next() { // 逐行扫描
		arr := make([]interface{}, query.RowNum)
		for i := 0; i < query.RowNum; i++ { // 设置指针
			arr[i] = ""
			arr[i] = &arr[i]
		}
		err := rows.Scan(arr...) // 扫描
		if err != nil {
			return nil, err
		}
		arrStr := make([]string, query.RowNum) // 将
		for i := 0; i < len(arr); i++ {        // TODO 想办法优化
			arrStr[i] = fmt.Sprintf("%s", arr[i])
		}
		ret = append(ret, query._scanHandle(arrStr...)) // 调用Type接口的ArrayOf方法
	}
	return ret, nil
}

// Order by 排序
func (query *Query) Order(SqlClause_Order string) *Query {
	query.SqlClause_Order = fmt.Sprint("ORDER BY ", SqlClause_Order)
	return query
}

// limit 分页
func (query *Query) Limit(SqlClause_limit string) *Query {
	query.SqlClause_limit = fmt.Sprint("LIMIT ", SqlClause_limit)
	return query
}

// sdistinct 去重
func (query *Query) Sdistinct(is bool) *Query {
	if is {
		query.SqlClause_Sdistinct = " DISTINCT "
	} else {
		query.SqlClause_Sdistinct = ""
	}
	return query
}

// 设置sql where 部分
func (Query *Query) Where(SqlClause_Where string) *Query {
	Query.SqlClause_Where = fmt.Sprint("WHERE ", SqlClause_Where)
	return Query
}

// 设置sql form 部分
func (Query *Query) Form(SqlClause_Form string) *Query {
	Query.SqlClause_Form = SqlClause_Form
	return Query
}

// 设置sql join 部分
func (Query *Query) Join(SqlClause_Join string) *Query {
	Query.SqlClause_Join = SqlClause_Join
	return Query
}

// 分组
func (query *Query) Group(groupField ...string) func(selectField ...string) func(ScanHandle) *Query {
	str := ""
	if !query.IsGroup { // 等于flase
		query.IsGroup = true // 设置为分组
		str = "GROUP BY "    // 去除开头逗号
	}
	for _, v := range groupField {
		str = fmt.Sprint(str, v, ",")
	}
	query.SqlClause_Group = fmt.Sprint(query.SqlClause_Group, str[:len(str)-1])

	return func(selectField ...string) func(ScanHandle) *Query { // 函数返回设置查询返回字段 防止用户忘记设置
		str = ""
		for _, v := range selectField {
			str = fmt.Sprint(str, v, ",")
		}
		return query.SetSelectField(str[:len(str)-1])
	}
}

// having
// 如果没有分组返回空指针
func (query *Query) Having(SqlClause_Having string) *Query {
	if query.IsGroup {
		return nil
	}
	query.SqlClause_Having = fmt.Sprint("HAVING ", SqlClause_Having)
	return query
}

// 设置扫描处理方法
// 不建议直接使用
func (query *Query) SetScanHandle(s ScanHandle) *Query {
	query._scanHandle = s
	return query
}

// 设置查询字段
// 不建议直接使用
func (query *Query) SetSelectField(SqlClause_Select string) func(ScanHandle) *Query {
	query.SqlClause_Select = SqlClause_Select
	return func(scanHandle ScanHandle) *Query {
		return query.SetScanHandle(scanHandle)
	}
}

// 保存节省
func (query *Query) Save() *SaveQuery {
	query.MakeSql() // 组合sql
	var sq SaveQuery
	sq.query = query
	return &sq
}
