package analyze

import "reflect"

func (analyze *EntityAnalyze) GetJsonSqls() []string {
	return make([]string, 0)
}
func (analyze *EntityAnalyze) GetInsertField() []*Field {
	return analyze.InsertField
}
func (analyze *EntityAnalyze) GetUpdateField() []*Field {
	return analyze.UpdateField
}
func (analyze *EntityAnalyze) GetSelectField() []*Field {
	return analyze.SelectField
}
func (analyze *EntityAnalyze) GetType() *Type {
	return &analyze.Type
}
func (analyze *EntityAnalyze) GetTableName() string {
	return analyze.TableName
}
func (analyze *EntityAnalyze) GetAllTableName() []string {
	ret := make([]string, 0)
	ret = append(ret, analyze.TableName)
	return ret
}
func (analyze *EntityAnalyze) SetInsertField(field []*Field) {
	analyze.InsertField = field
}
func (analyze *EntityAnalyze) SetUpdateField(field []*Field) {
	analyze.UpdateField = field
}
func (analyze *EntityAnalyze) SetSelectField(field []*Field) {
	analyze.SelectField = field
}

var _ Analyze = (*EntityAnalyze)(nil) // 接口校验
type EntityAnalyze struct {
	Type // 原对象

	TableName string // 表名

	KeyField *Field // 主键字段

	AllField    []*Field // 字段
	InsertField []*Field // 增加字段集合
	UpdateField []*Field // 修改字段集合
	SelectField []*Field // 查询字段集合
}

// 开始分析
func EntityAnalyzeStart(obj Type, handlers ...Handler) Analyze {
	ret := NewEntityAnalyze(obj)
	for _, ath := range handlers { // 外部传入的分析处理函数 调用
		ath(ret)
	}

	return ret
}

// 分析obj 获得 一个不完整的分析对象(*AnalyzeType)
func NewEntityAnalyze(obj Type) *EntityAnalyze {
	var ret EntityAnalyze            // 创建EntityAnalyze
	ret.Type = obj                   //原对象
	ret.AllField = make([]*Field, 0) // 初始化 Fields

	typ := reflect.TypeOf(obj) // 反射type
	// 设置主键
	keyField := typ.Field(0)                                // 约定第一个字段是主键,table 标签是表名称，没有标签默认用类名下划线格式
	if value, ok := keyField.Tag.Lookup(TableNameTag); ok { // 设置所属表名称
		ret.TableName = value
	} else {
		ret.TableName = fieldNameToSqlField(typ.Name())
	}
	ret.KeyField = ret.CreateField(&keyField) // 创建主键字段

	// 获取所有字段信息 （成员变量）
	ret.AllField = append(ret.AllField, ret.KeyField)
	for i := 1; i < typ.NumField(); i++ { // 遍历成员变量
		field := typ.Field(i)
		ret.AllField = append(ret.AllField, ret.CreateField(&field)) // 创建字段 （成员变量）
	}

	// copy 查询 增加 删除字段
	ret.InsertField = make([]*Field, len(ret.AllField))
	ret.UpdateField = make([]*Field, len(ret.AllField))
	ret.SelectField = make([]*Field, len(ret.AllField))

	// 拷贝增改查操作的字段信息
	copy(ret.InsertField, ret.AllField)
	copy(ret.UpdateField, ret.AllField)
	copy(ret.SelectField, ret.AllField)

	return &ret
}

// 创建一个字段对象
func (analyze *EntityAnalyze) CreateField(field *reflect.StructField) *Field {
	var ret Field
	ret.Kind = field.Type.Kind() // 数据类型（种类）
	ret.Name = field.Name
	if name, ok := field.Tag.Lookup(FieldNameTag); ok { // 写明啦标签
		ret.FieldName = name // 字段名
	} else { // 没有标签 默认使用字段名字 大写改下划线
		ret.FieldName = fieldNameToSqlField(field.Name)
	}
	if name, ok := field.Tag.Lookup(TableNameTag); ok { // 是否省略啦表名
		ret.TableName = name // 表名
	} else {
		ret.TableName = analyze.TableName // 省略表名 使用主键的表名
	}
	return &ret
}
