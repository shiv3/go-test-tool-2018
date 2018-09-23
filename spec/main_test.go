package spec

import (
	"flag"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	RemoteURL = flag.String("url", "", "remote url")
	os.Exit(m.Run())
}
