package spec

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)
type TestCases struct {
	Cases []Case `yaml:"case"`
}
type Case struct {
	Name string `yaml:"name"`
	Request struct {
		Method string `yaml:"method"`
		Path string `yaml:"path"`
		Body string `yaml:"body"`
		Json map[string]interface{} `yaml:"json"`
	} `yaml:"request"`
	Expect struct{
		Status int `yaml:"status"`
		Body string `yaml:"body"`
		Json map[string]interface{} `yaml:"json"`
	}
} 

func TestJSONFromYaml(t *testing.T) {
	e := getExpectServer(t)

	testCases := TestCases{}
	buf, err := ioutil.ReadFile("../spec_yaml/testcase1.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(buf), &testCases)

	for _,c := range testCases.Cases {
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()
			r := e.Request(c.Request.Method, c.Request.Path)
			if c.Request.Json != nil {
				r = r.WithJSON(c.Request.Json)
			} else {
				r = r.WithBytes([]byte(c.Request.Body))
			}
			expect := r.Expect().
				Status(c.Expect.Status)

			if c.Expect.Json != nil {
				expect.JSON().Equal(c.Expect.Json)
			} else {
				expect.Body().Equal(c.Expect.Body)
			}
		})
	}
}
