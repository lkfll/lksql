package serverutils

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// 传入指针
func BodyScanJson(body io.ReadCloser, ScanObj interface{}) error {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, ScanObj)
}
