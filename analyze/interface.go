package analyze

// 处理分析对象 函数类型
type Handler func(SetAnalyze)

type Analyze interface {
	GetInsertField() []*Field
	GetUpdateField() []*Field
	GetSelectField() []*Field
	GetJsonSqls() []string

	GetType() *Type
	GetTableName() string
	GetAllTableName() []string
}

type SetAnalyze interface {
	Analyze
	SetInsertField([]*Field)
	SetUpdateField([]*Field)
	SetSelectField([]*Field)
}

// 默认的处理分析对象 函数类型
func DefaultAnalyzeTypeHandler(analyze SetAnalyze) {
	// 删除 插入字段 id主键
	analyze.SetInsertField(analyze.GetInsertField()[1:])
	// 修改字段
	analyze.SetUpdateField(analyze.GetUpdateField()[1:])
}
