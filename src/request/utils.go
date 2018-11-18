package request

import (
	"net/http"
	"net/http/httputil"
	"fmt"
)

func DisplayRequest(r *http.Request) {
	re, err := httputil.DumpRequest(r, true); if err == nil {
		fmt.Println(string(re))
	}
}

func DisplayResponse(resp *http.Response) {
	re, err := httputil.DumpResponse(resp, true); if err == nil {
		fmt.Println(string(re))
	}
}