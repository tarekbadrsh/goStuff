package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type user struct {
	ID       string `json:"Id"`
	UserID   string `json:"UserId"`
	UserName string `json:"UserName"`
}

type userfile struct {
	Nextpageid string `json:"nextpageid"`
	UsersList  []user `json:"UsersList"`
}

func addUserTweetToFile(newuser *user) error {
	path := "data.json"
	myuserFile := &userfile{}

	data, err := ioutil.ReadFile(path)
	if os.IsNotExist(err) {
		fmt.Println("not exist")
	} else {
		if len(data) > 0 {
			err = json.Unmarshal(data, myuserFile)
			if err != nil {
				panic(err)
			}
		}
	}

	if len(myuserFile.UsersList) >= 5 {
		unixNow := time.Now().Unix()
		newpath := fmt.Sprintf("%d.json", unixNow)
		err := os.Rename(path, newpath)
		if err != nil {
			panic(err)
		}
		myuserFile.Nextpageid = newpath
		myuserFile.UsersList = nil
	}

	myuserFile.UsersList = append(myuserFile.UsersList, *newuser)

	data, err = json.Marshal(myuserFile)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(path, data, 0666)
	if err != nil {
		panic(err)
	}
	return nil
}

func main() {

	newuser := &user{
		ID:       "123",
		UserID:   "456",
		UserName: "foo",
	}
	addUserTweetToFile(newuser)
}
