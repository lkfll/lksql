package analyze

// 操作抽象接口，对象信息

// analyze接受对象
type Type interface {
	ToMap() map[string]interface{} // 获得map
}
