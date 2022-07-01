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

type Repository struct {
	AnalyzeType *analyze.AnalyzeType // 分析对象
	BaseSQL     *basesql.BaseSQL     // 生成的基础sql语句
	// where
	// limmit'
	// 等
}

//
func Start(analyze *analyze.AnalyzeType, baseSql *basesql.BaseSQL) *Repository {
	var ret Repository
	ret.AnalyzeType = analyze
	ret.BaseSQL = baseSql
	return &ret
}

// TODO select other
