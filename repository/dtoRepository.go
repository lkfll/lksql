package repository

import (
	"database/sql"

	basesql "github.com/lkfll/lksql/baseSql"
	"github.com/lkfll/lksql/exec"
)

//
type DtoRepository struct {
	*Repository
	DtoSql            *basesql.DtoSql // 生成的基础sql语句
	EntityRepositorys map[string]*EntityRepository
}

func DtoStart(repository *Repository, baseSql *basesql.DtoSql) *DtoRepository {
	var ret DtoRepository
	ret.Repository = repository
	ret.DtoSql = baseSql
	ret.EntityRepositorys = make(map[string]*EntityRepository)
	for k, v := range ret.DtoSql.EntitySqls {
		ret.EntityRepositorys[k] = EntityStart(nil, v)
	}
	return &ret
}

// 查询全部
// SELECT  issues.id , issues.issues_name , issues.label , issues.description , issues.uid , priority.priority_name
// FROM  issues
//         LEFT JOIN priority ON issues.pid=priority.id
func (repository *DtoRepository) Select() *exec.Query {
	return exec.NewQuery(repository.DtoSql.SelectSql, exec.NewScan(repository.Type, repository.SelectField...))
}

// 增加
func (repository *DtoRepository) Insert(DB *sql.DB, params ...interface{}) ([]sql.Result, error) {
	ret := make([]sql.Result, 0)
	var err error = nil
	for _, v := range repository.EntityRepositorys {
		r, err := v.Insert(DB, params...)
		if err != nil {
			return ret, err
		}
		ret = append(ret, r)
	}
	return ret, err
}
