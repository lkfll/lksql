package analyze

// 操作抽象接口，对象信息

// analyze接受对象
type Type interface {
	ResultScan
	ToMap() map[string]interface{} // 获得map
}

// 返回值扫描接口
type ResultScan interface {
	ArrayOf(...string) interface{} //
}
