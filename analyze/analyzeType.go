// 通过start 获取Type 分析对象 交与basesql层
package analyze

import (
	"fmt"
	"reflect"
)

// 表名和字段切片结构体
type TableAndFields struct {
	TableName string
	Fields    []*Field
}

// 分析对象，通过反射分析出来的对象
// 表字段map集合
type AnalyzeType struct {
	Type // 原对象

	NumField int // 字段个数

	TableNames []string // 数据库表名 切片

	Key  *KeyField   // TODO 主键字段  扩展为切片 多个表主键
	Join []JoinField // 连接字段集合

	AllField    []*Field // 所有表的字段总和
	InsertField []*Field // 增加字段集合
	UpdateField []*Field // 修改字段集合
	SelectField []*Field // 查询字段集合

	// TODO 扩展
}

// 处理分析对象 函数类型
type AnalyzeTypeHandler func(*AnalyzeType)

// 默认的处理分析对象 函数类型
func DefaultAnalyzeTypeHandler(analyzeType *AnalyzeType) {
	// 删除 插入字段 id主键
	for i, v := range analyzeType.InsertField {
		if v == &analyzeType.Key.Field {
			analyzeType.InsertField = append(analyzeType.InsertField[0:i], analyzeType.InsertField[i+1:]...)
		}
	}
	// 修改字段
	for i, v := range analyzeType.UpdateField {
		if v == &analyzeType.Key.Field {
			analyzeType.UpdateField = append(analyzeType.UpdateField[0:i], analyzeType.UpdateField[i+1:]...)
		}
	}
}

// 开始分析
func Start(obj Type, analyzeTypeHandlers ...AnalyzeTypeHandler) (*AnalyzeType, error) {
	ret, err := analyze(obj)
	if err != nil {
		return nil, err
	}
	for _, ath := range analyzeTypeHandlers { // 外部传入的分析处理函数 调用
		ath(ret)
	}

	return ret, nil
}

// 分析obj 获得 一个不完整的分析对象(*AnalyzeType)
func analyze(obj Type) (*AnalyzeType, error) {
	var ret AnalyzeType // 初始化
	ret.Type = obj      //原对象
	ret.Key = nil
	ret.Join = make([]JoinField, 0)
	ret.AllField = make([]*Field, 0)   // 初始化 AllField
	ret.TableNames = make([]string, 0) // 初始化 tableNames
	typ := reflect.TypeOf(obj)         // 分析传入对象
	ret.NumField = typ.NumField()      // 对象字段个数

	// 分析字段 （成员变量）
	for i := 0; i < ret.NumField; i++ { // 遍历成员变量
		field := typ.Field(i)
		ret.AllField = append(ret.AllField, ret.analyzeField(&field)) // 分析字段 （成员变量）
	}

	// copy 查询 增加 删除字段
	ret.InsertField = make([]*Field, len(ret.AllField))
	ret.UpdateField = make([]*Field, len(ret.AllField))
	ret.SelectField = make([]*Field, len(ret.AllField))

	copy(ret.InsertField, ret.AllField)
	copy(ret.UpdateField, ret.AllField)
	copy(ret.SelectField, ret.AllField)

	// 设置tableNames
	tableNames := make(map[string]string)
	for _, f := range ret.AllField {
		tableNames[f.TableName] = ""
	}
	for k, v := range tableNames {
		if v == "" { // 没什么意义 取消警告
			ret.TableNames = append(ret.TableNames, k)
		}
	}

	return &ret, nil
}

// func 字段名字到sql字段名字
// UserName 变为 user_name
func fieldNameToSqlField(name string) string {
	ret := ""
	for i, v := range name {
		if 'A' <= v && v <= 'Z' { // 是大小写字母
			v += 32
			if i == 0 {
				ret = fmt.Sprint(ret, string(v)) // 是第一个字符
			} else {
				ret = fmt.Sprint(ret, "_", string(v))
			}
			continue
		}
		ret = fmt.Sprint(ret, string(v))
	}
	return ret
}

// 分析成员
func (analyze *AnalyzeType) analyzeField(field *reflect.StructField) *Field {
	// 内联函数
	setRet := func(ret *Field) {
		ret.Kind = field.Type.Kind()                        // 数据类型（种类）
		if name, ok := field.Tag.Lookup(FieldNameTag); ok { // 写明啦标签
			ret.FieldName = name // 字段名
		} else { // 没有标签 默认使用字段名字 大写改下划线
			ret.FieldName = fieldNameToSqlField(field.Name)
		}
		if name, ok := field.Tag.Lookup(TableNameTag); ok { // 是否省略啦表名
			ret.TableName = name // 表名
		} else {
			ret.TableName = analyze.Key.TableName // 省略表名 使用主键的表名
		}
	}
	// 主键
	if value, ok := field.Tag.Lookup(KeyTag); ok {
		var ret KeyField    // 生成主键字段对象
		ret.KeyName = value // 设置主键名字
		setRet(&ret.Field)
		analyze.Key = &ret // 分析对象获得主键字段
		return &ret.Field
	}
	// 连接语句
	if value, ok := field.Tag.Lookup(JoinTag); ok {
		var ret JoinField   // 生成连接字段对象
		ret.JoinSql = value // 设置连接sql
		// TODO 设置外键表名
		// TODO 设置外键字段名字
		setRet(&ret.Field)
		analyze.Join = append(analyze.Join, ret) // 分析对象获得连接字段
		// return 向下继续执行，
		return &ret.Field
	}
	// TODO 扩展
	// if value, ok := field.Tag.Lookup(""); ok {
	// 	var field
	// }

	// 字段 放在最下边 找不到特殊标签才会执行 比如主键标签...
	var ret Field
	setRet(&ret)
	return &ret
}
