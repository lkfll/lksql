package lksql

import (
	"github.com/lkfll/lksql/analyze"
	basesql "github.com/lkfll/lksql/baseSql"
	"github.com/lkfll/lksql/repository"
)

// type Factory struct {
// 	// Container // 仓库容器
// }

// // repository容器
// type Container struct {
// 	Entitys map[string]*repository.EntityRepository
// 	Dto     map[string]*repository.DtoRepository
// }

// 实体仓库工厂
func EntityFactory(obj Type, handlers ...analyze.Handler) *repository.EntityRepository {
	at := analyze.EntityAnalyzeStart(obj, handlers...)
	bs := basesql.EntitySqlStart(at)
	return repository.EntityStart(repository.NewRepository(obj, at.GetSelectField()), bs)
}

// 默认工厂 增加，修改 删除主键字段
func EntityDefaultFacory(obj Type) *repository.EntityRepository {
	return EntityFactory(obj, analyze.DefaultAnalyzeTypeHandler)
}

// 工厂方法 创建数据库连接对象
// 创建的 Repository 修改 增加字段包括id谨慎使用 建议使用DefaultFacory
func DtoFactory(obj Type, handlers ...analyze.Handler) *repository.DtoRepository {
	at := analyze.DtoAnalyzeStart(obj, handlers...)
	bs := basesql.DtoStart(at)
	return repository.DtoStart(repository.NewRepository(obj, at.GetSelectField()), bs)
}

// 默认工厂 增加，修改 删除主键字段
func DtoDefaultFacory(obj Type) *repository.DtoRepository {
	return DtoFactory(obj, analyze.DefaultAnalyzeTypeHandler)
}
