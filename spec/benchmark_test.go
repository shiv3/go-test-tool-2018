package spec

import (
	"testing"
)

func bench(b *testing.B, path string) {
	e := getExpectServer(&testing.T{})
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Request(e, path)
	}
}

func BenchmarkRuok(b *testing.B) {
	bench(b, "/ruok")
}

func BenchmarkSlow(b *testing.B) {
	bench(b, "/slow")
}
