package spec

import (
	"github.com/gavv/httpexpect"
	"github.com/shiv3/go-integration-test/handle"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJSON(t *testing.T) {
	t.Parallel()
	e := getExpectServer(t)

	e.GET("/json").
		Expect().
		Status(http.StatusOK).
		JSON().Object().
			ContainsKey("test").ValueEqual("test",1).
			ContainsKey("test2").ValueEqual("test2","2").
			Value("test3").Array().Elements(1,2,3)
}

func TestJSONSchema(t *testing.T) {
	t.Parallel()
	h := handle.Handler()
	server := httptest.NewServer(h)
	defer server.Close()

	schema := `{
		"type": "object",
		"required": [
	    	"test",
	    	"test2",
	    	"test3"
	  	],
		"properties": {
			"test": {
				"type": "integer"
			},
			"test2": {
				"type": "string"
			},
			"test3": {
				"type": "array"
			}
		}
	}`
	e := httpexpect.New(t, server.URL)

	e.GET("/json").
		Expect().
		Status(http.StatusOK).
		JSON().
		Schema(schema)
}
