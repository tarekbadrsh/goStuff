package configuration_test

import (
	"testing"

	"github.com/tarekbadrshalaan/goStuff/configuration"
)

//!+test
//go test -v
func TestJSON(t *testing.T) {
	type Objconf struct {
		LogPath   string
		StopAfter int64
		Emails    []string
	}
	config := &Objconf{}
	err := configuration.JSON("JSON_test.json", config)
	if err != nil {
		t.Errorf("TestJSON\nerror:configuration.JSON:%v", err)
	}
	expectedconfig := &Objconf{
		LogPath:   "logpath",
		StopAfter: 60,
		Emails:    []string{"foo@gmail.com", "bar@gmail.com"},
	}
	if !(expectedconfig.LogPath == config.LogPath &&
		expectedconfig.StopAfter == config.StopAfter &&
		expectedconfig.Emails[0] == config.Emails[0] &&
		expectedconfig.Emails[1] == config.Emails[1]) {

		t.Errorf("TestJSON\nparsed object not as expected\nexpected:%v\nparsed:%v", expectedconfig, config)
	}
}

//!-tests

//!+bench
//go test -v  -bench=.
func BenchmarkJSON(b *testing.B) {
	type Objconf struct {
		LogPath   string
		StopAfter int64
		Emails    []string
	}
	for index := 0; index < b.N; index++ {
		config := &Objconf{}

		err := configuration.JSON("JSON_test.json", config)
		if err != nil {
			b.Errorf("BenchmarkJSON\nerror:configuration.JSON:%v", err)
		}
	}
}

//!-bench
