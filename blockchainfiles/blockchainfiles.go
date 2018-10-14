package blockchainfiles

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/pkg/errors"
)

// DataFile : structure of file:
// contain Nextpageid: refer to next page path and Recoreds : list of records.
type DataFile struct {
	Nextpageid string        `json:"nextpageid"`
	Recoreds   []interface{} `json:"Recoreds"`
}

func archiveFile(basefileName string) (string, error) {
	unixNow := time.Now().Unix()
	newpath := fmt.Sprintf("%d.json", unixNow)
	counter := 0
	for {
		if _, err := os.Stat(newpath); os.IsNotExist(err) {
			break
		}
		newpath = fmt.Sprintf("%d_%d.json", unixNow, counter)
		counter++
	}

	err := os.Rename(basefileName, newpath)
	if err != nil {
		return "", errors.Wrapf(err, "could not rename basefileName:%v to newpath:%v", basefileName, newpath)
	}
	return newpath, nil
}

// GetdataFile : return object of Datafile from file path
// it will return empty object if file not exist.
func GetdataFile(path string) (*DataFile, error) {
	dataFile := &DataFile{}
	file, err := os.Open(path)
	if os.IsNotExist(err) {
		return dataFile, nil
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(dataFile)
	if err != nil {
		return nil, errors.Wrap(err, "could not Decode dataFile")
	}
	return dataFile, nil
}

// AddRecord : append new recored to file
// If the file does not exist, AddRecord creates it with permissions perm.
func AddRecord(basefileName string, newrecord interface{}, maxfileRecored int) error {
	datafile, err := GetdataFile(basefileName)
	if err != nil {
		return errors.Wrap(err, "could not getdataFile")
	}
	if len(datafile.Recoreds) >= maxfileRecored {
		newpath, err := archiveFile(basefileName)
		if err != nil {
			return errors.Wrapf(err, "could not archive file %v", basefileName)
		}
		datafile.Nextpageid = newpath
		datafile.Recoreds = nil
	}

	datafile.Recoreds = append(datafile.Recoreds, newrecord)

	marshaldata, err := json.Marshal(datafile)
	if err != nil {
		return errors.Wrapf(err, "could not Marshal dataFile %v", datafile)
	}

	err = ioutil.WriteFile(basefileName, marshaldata, 0666)
	if err != nil {
		return errors.Wrapf(err, "could not WriteFile %v", basefileName)
	}
	return nil
}
