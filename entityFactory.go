package lksql

import (
	"github.com/lkfll/lksql/analyze"
	basesql "github.com/lkfll/lksql/baseSql"
	"github.com/lkfll/lksql/repository"
)

// 实体仓库工厂
func EntityFactory(obj Type, analyzeTypeHandlers ...analyze.AnalyzeTypeHandler) (*repository.EntityRepository, error) {
	at, err := analyze.Start(obj, analyzeTypeHandlers...)
	if err != nil {
		return nil, err
	}
	bs := basesql.EntitySqlStart(at)
	return repository.EntityStart(at, bs), nil
}

// 默认工厂 增加，修改 删除主键字段
func EntityDefaultFacory(obj Type) (*repository.EntityRepository, error) {
	return EntityFactory(obj, analyze.DefaultAnalyzeTypeHandler)
}
