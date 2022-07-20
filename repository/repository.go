// 数据库操作对象
// TODO 对参数进行修改，防止sql注入攻击
package repository

import (
	"github.com/lkfll/lksql/analyze"
)

func NewRepository(typ interface{}, selectField []*analyze.Field) *Repository {
	var ret Repository
	ret.Type = typ
	ret.SelectField = selectField
	return &ret
}

type Repository struct {
	Type        interface{}      // 原对象
	SelectField []*analyze.Field // 查询字段
}
