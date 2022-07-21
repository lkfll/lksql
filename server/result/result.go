package result

import "encoding/json"

type Result struct {
	Code int
	Data interface{}
}

func Ok(obj interface{}) string {
	b, _ := json.Marshal(Result{Code: 0, Data: obj})
	return string(b)
}

func Err(err error) string {
	b, _ := json.Marshal(Result{Code: -1, Data: err})
	return string(b)
}
