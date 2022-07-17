// 数据库操作对象
// TODO 对参数进行修改，防止sql注入攻击
package repository

import (
	"github.com/lkfll/lksql/analyze"
	basesql "github.com/lkfll/lksql/baseSql"
)

// 数据库操作对象
// 交与用户进行操作
// var _ common.Repositoryer = (*Repository)(nil)
type DtoRepository struct {
	AnalyzeType *analyze.AnalyzeType // 分析对象
	BaseSQL     *basesql.DtoSql      // 生成的基础sql语句
	// where
	// limmit'
	// 等
}

// 实体类仓库
type EntityRepository struct {
	AnalyzeType *analyze.AnalyzeType // 分析对象
	EntitySql   *basesql.EntitySql   // 生成的基础sql语句
	// where
	// limmit'
	// 等
}

//
func DtoStart(analyze *analyze.AnalyzeType, baseSql *basesql.DtoSql) *DtoRepository {
	var ret DtoRepository
	ret.AnalyzeType = analyze
	ret.BaseSQL = baseSql
	return &ret
}

//
func EntityStart(analyze *analyze.AnalyzeType, entitySql *basesql.EntitySql) *EntityRepository {
	var ret EntityRepository
	ret.AnalyzeType = analyze
	ret.EntitySql = entitySql
	return &ret
}

// TODO select other
