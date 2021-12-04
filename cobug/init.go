package cobug

import "net/http"

var (
	client                *http.Client
	disAutoLoacationClien *http.Client
)

func init() {
	client = &http.Client{}

	disAutoLoacationClien = &http.Client{}
	disAutoLoacationClien.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
}
