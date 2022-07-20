package analyze

import "reflect"

// 结构体成员的构成参数,对应数据库字段
type Field struct {
	Kind      reflect.Kind // 数据类型（种类）
	TableName string       // 所属表名字
	Name      string       // 属性名（原名）
	FieldName string       // 数据库对应字段名字
}
