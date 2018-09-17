package spec

import (
	"github.com/gavv/httpexpect"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

type TestCases struct {
	*testing.T
	*httpexpect.Expect
	Cases []Case `yaml:"case"`
}

type Case struct {
	Name    string `yaml:"name"`
	Request struct {
		Method string                 `yaml:"method"`
		Path   string                 `yaml:"path"`
		Body   string                 `yaml:"body"`
		Json   map[string]interface{} `yaml:"json"`
		Query  map[string]interface{} `yaml:"query"`
	} `yaml:"request"`
	Expect struct {
		Status int                    `yaml:"status"`
		Body   string                 `yaml:"body"`
		Json   map[string]interface{} `yaml:"json"`
	}
}

func (c *TestCases) run() {
	for _,test := range c.Cases {
		c.T.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			r := c.Expect.Request(test.Request.Method, test.Request.Path).
				WithQueryObject(test.Request.Query)
			if test.Request.Json != nil {
				r = r.WithJSON(test.Request.Json)
			} else {
				r = r.WithBytes([]byte(test.Request.Body))
			}
			expect := r.Expect().
				Status(test.Expect.Status)

			if test.Expect.Json != nil {
				expect.JSON().Equal(test.Expect.Json)
			} else {
				expect.Body().Equal(test.Expect.Body)
			}
		})
	}
}

func testYaml(t *testing.T,filepath string){
	testCases := TestCases{
		Expect:getExpectServer(t),
		T:t,
	}
	buf, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(buf), &testCases)
	testCases.run()
}

func TestJSONFromYaml(t *testing.T) {
	filepath.Walk("../spec_yaml", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir(){
			testYaml(t,path)
		}
		return nil
	})
}
