package exec

import (
	"database/sql"
	"reflect"

	"github.com/lkfll/lksql/analyze"
)

// 构造函数
func NewScan(obj interface{}, fields ...*analyze.Field) *Scan {
	var ret Scan
	ret.RefType = reflect.TypeOf(obj)
	ret.Fields = fields
	return &ret
}

// 扫描对象
type Scan struct {
	Fields  []*analyze.Field // 字段名字数组
	RefType reflect.Type     // 反射类型
}

// TODO 优化 使用copy替代反射New
// 扫描成string切片,扫描完成后会关闭sql.rows
func (scan *Scan) ScanHandle(r *sql.Rows) ([]interface{}, error) {
	ret := make([]interface{}, 0)
	for r.Next() { // 每一行
		reflectStruct := reflect.New(scan.RefType).Elem() // 通过反射创建一个对象
		fieldValues := make([]interface{}, 0)
		for _, v := range scan.Fields { // 每个字段
			// 获取字段的各个指针
			fieldValues = append(fieldValues, reflectStruct.FieldByName(v.Name).Addr().Interface())
		}
		r.Scan(fieldValues...)
		ret = append(ret, reflectStruct.Interface())
	}
	return ret, nil
}
