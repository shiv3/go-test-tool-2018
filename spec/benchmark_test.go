package spec

import (
	"net/http"
	"testing"
)

func BenchmarkRuok(b *testing.B) {
	var t testing.T
	e := getExpectServer(&t)
	for n := 0; n < b.N; n++ {
		e.GET("/ruok").
			Expect().
			Status(http.StatusOK).
			Body().Equal("imok")

		e.POST("/ruok").
			Expect().
			Status(http.StatusMethodNotAllowed)
	}
}
