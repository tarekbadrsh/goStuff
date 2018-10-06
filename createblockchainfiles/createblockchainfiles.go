package createblockchainfiles

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// DataFile :
type DataFile struct {
	Nextpageid string        `json:"nextpageid"`
	Recoreds   []interface{} `json:"Recoreds"`
}

// AddRecord : appen new recored to file
func AddRecord(newrecord interface{}, maxfileRecored int) error {
	path := "data.json"
	mydataFile := &DataFile{}

	data, err := ioutil.ReadFile(path)
	if os.IsNotExist(err) {
		log.Printf("%v not exist", path)
	} else {
		if len(data) > 0 {
			err = json.Unmarshal(data, mydataFile)
			if err != nil {
				panic(err)
			}
		}
	}

	if len(mydataFile.Recoreds) >= maxfileRecored {
		unixNow := time.Now().Unix()
		newpath := fmt.Sprintf("%d.json", unixNow)
		err := os.Rename(path, newpath)
		if err != nil {
			panic(err)
		}
		mydataFile.Nextpageid = newpath
		mydataFile.Recoreds = nil
	}

	mydataFile.Recoreds = append(mydataFile.Recoreds, newrecord)

	data, err = json.Marshal(mydataFile)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(path, data, 0666)
	if err != nil {
		panic(err)
	}
	return nil
}
