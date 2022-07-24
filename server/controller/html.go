package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type Html struct {
}

func (html *Html) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("r.URL: %v\n", r.URL)
	f, err := os.Open(fmt.Sprint("../server/html", r.URL))
	if err != nil {
		fmt.Fprint(w, err)
	} else {
		io.Copy(w, f)
	}
}

// func (html *Html) IndexJs(w http.ResponseWriter, r *http.Request) {
// 	f, err := os.Open("../server/html/index.js")
// 	if err != nil {
// 		fmt.Fprint(w, err)
// 	} else {
// 		io.Copy(w, f)
// 	}
// }
