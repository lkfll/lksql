package lksql

import (
	"github.com/lkfll/lksql/analyze"
	basesql "github.com/lkfll/lksql/baseSql"
	"github.com/lkfll/lksql/repository"
)

// 工厂方法 创建数据库连接对象
// 创建的 Repository 修改 增加字段包括id谨慎使用 建议使用DefaultFacory
func Factory(obj Type, analyzeTypeHandlers ...analyze.AnalyzeTypeHandler) (*repository.Repository, error) {
	at, err := analyze.Start(obj, analyzeTypeHandlers...)
	if err != nil {
		return nil, err
	}
	bs := basesql.Start(at)
	return repository.Start(at, bs), nil
}

// 默认工厂 增加，修改 删除主键字段
func DefaultFacory(obj Type) (*repository.Repository, error) {
	return Factory(obj, analyze.DefaultAnalyzeTypeHandler)
}
