package spec

import (
	"github.com/gavv/httpexpect"
	"net/http"
	"testing"
)

func TestExampleCom(t *testing.T) {
	e := httpexpect.New(t, "http://example.com")
	e.POST("/").
		Expect().
		Status(http.StatusOK)
}
