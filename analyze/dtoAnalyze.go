// 通过start 获取Type 分析对象 交与basesql层
package analyze

import "reflect"

// 分析对象，通过反射分析出来的对象
func (analyze *DtoAnalyze) GetJsonSqls() []string {
	return analyze.JoinSqls
}
func (analyze *DtoAnalyze) GetAllTableName() []string {
	ret := make([]string, 0)
	ret = append(ret, analyze.TableName)
	ret = append(ret, analyze.JoinTableNames...)
	return ret
}

var _ Analyze = (*DtoAnalyze)(nil) // 接口校验
type DtoAnalyze struct {
	*EntityAnalyze

	JoinTableNames []string // 连接表名
	JoinSqls       []string // 连接字段集合
}

// 开始分析
func DtoAnalyzeStart(obj Type, handlers ...Handler) Analyze {
	ret := NewDtoAnalyze(obj)
	for _, ath := range handlers { // 外部传入的分析处理函数 调用
		ath(ret)
	}
	return ret
}

// DtoAnalyze 构造函数
func NewDtoAnalyze(obj Type) *DtoAnalyze {
	var ret DtoAnalyze // 初始化
	ret.EntityAnalyze = NewEntityAnalyze(obj)

	ret.JoinSqls = make([]string, 0)
	typ := reflect.TypeOf(obj)

	for i := 0; i < typ.NumField(); i++ { // 获取jsonsqls
		if str, ok := typ.FieldByIndex([]int{i}).Tag.Lookup(JoinTag); ok {
			ret.JoinSqls = append(ret.JoinSqls, str)
		}
		if str, ok := typ.FieldByIndex([]int{i}).Tag.Lookup(TableNameTag); ok {
			if str != ret.TableName { // 不等于主表名
				ret.JoinTableNames = append(ret.JoinTableNames, str)
			}
		}
	}

	return &ret
}
