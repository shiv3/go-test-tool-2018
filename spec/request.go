package spec

import "github.com/gavv/httpexpect"

func Request(e httpexpect.Expect, path string) *httpexpect.Response {
	return e.GET(path).
		Expect()
}
