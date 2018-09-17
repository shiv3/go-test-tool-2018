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
	requestValue := "2"
	e.GET("/json").
		WithQuery("v",requestValue).
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		ContainsKey("test").ValueEqual("test", 1).
		ContainsKey("test2").ValueEqual("test2", requestValue).
		ContainsKey("test3").Value("test3").Array().Elements(1, 2, 3)

	e.POST("/json").
		Expect().
		Status(http.StatusMethodNotAllowed)
}

func TestJSONSchema(t *testing.T) {
	t.Parallel()
	h := handle.Handler()
	server := httptest.NewServer(h)
	defer server.Close()
	e := httpexpect.New(t, server.URL)
	schema := `{
		"type": "object",
		"required": ["test","test2","test3"],
		"properties": {
			"test": {
				"type": "integer",
				"maximum": 10
			},
			"test2": {
				"type": "string",
				"pattern": "^test.*$"
			},
			"test3": {
				"type": "array",
				"maxItems": 3
			}
		}
	}`

	e.GET("/json").
		WithQuery("v","test").
		Expect().
		Status(http.StatusOK).
		JSON().
		Schema(schema)
}
