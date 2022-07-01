package analyze

import "reflect"

// 结构体成员的构成参数,对应数据库字段
type Field struct {
	Kind      reflect.Kind // 数据类型（种类）
	TableName string       // 所属表名字
	FieldName string       // 数据库对应字段名字
}

// 主键
type KeyField struct {
	KeyName string // 主键名字 主表
	Field
}

// 连接字段
type JoinField struct {
	JoinSql          string // 连接sql语句
	ForeignFieldName string // 外键 字段名字
	ForeignTableName string // 外键 字段名字

	Field
}

// TODO 扩展
