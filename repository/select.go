package repository

import (
	"github.com/lkfll/lksql/exec"
)

// 查询全部
// SELECT  issues.id , issues.issues_name , issues.label , issues.description , issues.uid , priority.priority_name
// FROM  issues
//         LEFT JOIN priority ON issues.pid=priority.id
func (Repository *DtoRepository) Select() *exec.Query {
	return exec.NewQuery(Repository.BaseSQL.SelectSql, exec.NewScan(Repository.AnalyzeType.Type, Repository.AnalyzeType.SelectField...))
}

// 查询全部
// SELECT  issues.id , issues.issues_name , issues.label , issues.description , issues.uid , priority.priority_name
// FROM  issues
func (Repository *EntityRepository) Select() *exec.Query {
	return exec.NewQuery(Repository.EntitySql.SelectSql, exec.NewScan(Repository.AnalyzeType.Type, Repository.AnalyzeType.SelectField...))
}
