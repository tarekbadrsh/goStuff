package blockchainfiles_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/tarekbadrshalaan/goStuff/blockchainfiles"
)

func removeAllJsonFiles() error {
	//remove all old json , start test from scratch
	files, err := filepath.Glob(filepath.Join("", "*.json"))
	if err != nil {
		return errors.Wrap(err, "filepath.Glob")
	}
	for _, file := range files {
		err = os.Remove(file)
		if err != nil {
			return errors.Wrap(err, "os.Remove(file)")
		}
	}
	//remove all old json , start test from scratch
	return nil
}

//!+test
//go test -v
func TestAddRecord(t *testing.T) {
	err := removeAllJsonFiles()
	if err != nil {
		t.Errorf("TestAddRecord\nremoveAllJsonFiles:%v", err)
	}

	type user struct {
		ID       int    `json:"Id"`
		UserID   string `json:"UserId"`
		UserName string `json:"UserName"`
	}
	for index := 0; index < 12; index++ {
		newuser := &user{
			ID:       index,
			UserID:   "456",
			UserName: "foo",
		}
		err = blockchainfiles.AddRecord("data.json", newuser, 5)
		if err != nil {
			t.Errorf("TestAddRecord\nerror:AddRecord:%v", err)
		}
	}

	files, err := filepath.Glob(filepath.Join("", "*.json"))
	if err != nil {
		t.Errorf("filepath.Glob:%v", err)
	}
	if len(files) != 3 {
		t.Errorf("filepath.Glob files count not as expected\nexpected:%v\actual:%v", 3, len(files))
	}

	datefile0, err := blockchainfiles.GetdataFile(files[0])
	if err != nil {
		t.Errorf("blockchainfiles.GetdataFile Error:%v", err)
	}
	actual0 := user{}
	mapstructure.Decode(datefile0.Recoreds[0], &actual0)
	expected0 := user{ID: 0, UserID: "456", UserName: "foo"}
	if datefile0.Nextpageid != "" || actual0 != expected0 {
		t.Errorf("datefile[0] is not as expected\nexpected %v\nactual %v", expected0, actual0)
	}

	datefile1, err := blockchainfiles.GetdataFile(files[1])
	if err != nil {
		t.Errorf("blockchainfiles.GetdataFile Error:%v", err)
	}
	actual1 := user{}
	mapstructure.Decode(datefile1.Recoreds[0], &actual1)
	expected1 := user{ID: 5, UserID: "456", UserName: "foo"}
	if datefile1.Nextpageid != files[0] || actual1 != expected1 {
		t.Errorf("datefile[1] is not as expected\nexpected %v\nactual %v", expected1, actual1)
	}

	datefile2, err := blockchainfiles.GetdataFile(files[2])
	if err != nil {
		t.Errorf("blockchainfiles.GetdataFile Error:%v", err)
	}
	actual2 := user{}
	mapstructure.Decode(datefile2.Recoreds[0], &actual2)
	expected2 := user{ID: 10, UserID: "456", UserName: "foo"}
	if datefile2.Nextpageid != files[1] || actual2 != expected2 {
		t.Errorf("datefile[2] is not as expected\nexpected %v\nactual %v", expected2, actual2)
	}

}

//!-tests

//!+bench
//go test -v  -bench=.
func BenchmarkAddRecord(b *testing.B) {

	err := removeAllJsonFiles()
	if err != nil {
		b.Errorf("TestAddRecord\nremoveAllJsonFiles:%v", err)
	}

	type user struct {
		ID       int    `json:"Id"`
		UserID   string `json:"UserId"`
		UserName string `json:"UserName"`
	}
	for index := 0; index < b.N; index++ {
		newuser := &user{
			ID:       index,
			UserID:   "456",
			UserName: "foo",
		}
		err = blockchainfiles.AddRecord("data.json", newuser, 1000)
		if err != nil {
			b.Errorf("BenchmarkAddRecord\nerror:AddRecord:%v", err)
		}
	}
}

//!-bench
