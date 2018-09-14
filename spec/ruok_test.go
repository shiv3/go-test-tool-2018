package spec

import (
	"net/http"
	"testing"
)


//func TestExampleCom(t *testing.T) {
//	e := httpexpect.New(t,"http://example.com")
//	e.POST("").
//		Expect().
//		Status(http.StatusOK)
//}

func TestRuok(t *testing.T) {
	t.Parallel()
	e := getExpectServer(t)

	e.GET("/ruok").
		Expect().
		Status(http.StatusOK).
		Body().Equal("imok")

	e.POST("/ruok").
		Expect().
		Status(http.StatusMethodNotAllowed)
}

