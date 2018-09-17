package spec

import (
	"net/http"
	"testing"
)

func TestRuok(t *testing.T) {
	t.Parallel()
	e := getExpectServer(t)

	e.GET("/ruok").
		Expect().
		Status(http.StatusOK).
		Text().Equal("imok")

	e.POST("/ruok").
		Expect().
		Status(http.StatusMethodNotAllowed)
}
