package main

import (
	"github.com/tarekbadrshalaan/goStuff/createblockchainfiles"
)

type user struct {
	ID       string `json:"Id"`
	UserID   string `json:"UserId"`
	UserName string `json:"UserName"`
}

type data struct {
	ID        string `json:"Id"`
	Text      string `json:"Text"`
	CreatedAt int64  `json:"CreatedAt"`
}

func main() {
	for index := 0; index < 13; index++ {
		// newuser := &user{
		// 	ID:       "123",
		// 	UserID:   "456",
		// 	UserName: "foo",
		// }
		newdata := &data{
			ID:        "123",
			Text:      "hello",
			CreatedAt: 123456789,
		}

		err := createblockchainfiles.AddRecord("data.json", newdata, 5)
		if err != nil {
			panic(err)
		}
	}
}
