package exec

type SaveQuery struct {
	query *Query
}

//  go 执行sql
// 执行sql并且扫描结果
// 将GetResult 和 Scan 合并
func (sq *SaveQuery) Go(db SQLCommon, param ...interface{}) ([]interface{}, error) {
	Hook_QueryGoBefore(sq.query, param...) // 钩子
	rows, err := db.Query(sq.query.Sql, param...)
	if err != nil {
		Hook_QueryGoAfter(nil, err) // 钩子
		return nil, err
	}
	defer rows.Close() // 结束关闭rows
	i, err := sq.query.Scan(rows)
	Hook_QueryGoAfter(i, err) // 钩子
	return i, err
}

//
func (sq *SaveQuery) GetSql() string {
	return sq.query.Sql
}
