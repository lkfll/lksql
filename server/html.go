package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type Html struct {
}

func (html *Html) Index(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("../server/index.html")
	if err != nil {
		fmt.Fprint(w, err)
	} else {
		io.Copy(w, f)
	}
}