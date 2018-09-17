package spec

import (
	"flag"
	"github.com/gavv/httpexpect"
	"github.com/shiv3/go-integration-test/handle"
	"net/http/httptest"
	"os"
	"testing"
)

var RemoteURL *string

func getExpectServer(t *testing.T) *httpexpect.Expect {
	var server *httptest.Server
	if *RemoteURL == "" {
		server = httptest.NewServer(handle.Handler())
		RemoteURL = &server.URL
	}
	return httpexpect.New(t, *RemoteURL)
}

func TestMain(m *testing.M) {
	RemoteURL = flag.String("url", "", "remote url")
	os.Exit(m.Run())
}