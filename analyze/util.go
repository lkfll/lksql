package analyze

import "fmt"

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
