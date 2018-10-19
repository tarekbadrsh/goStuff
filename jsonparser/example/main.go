package main

import (
	"fmt"

	"github.com/tarekbadrshalaan/goStuff/jsonparser"
)

func main() {

	data1 := []byte(`{"json": {"errors": [], "data": {"url": "www.example.com", "count": 0, "id": "123", "name": "foo"}}}`)
	obj, err := jsonparser.JSONParser(data1, "json", "data")
	if err != nil {
		panic(err)
	}
	fmt.Println(obj["name"])

	data2 := []byte(`{"json": {"errors": [], "data": {"url": "www.example.com", "count": 0, "id": "123", "name": "foo"}}}`)
	name2, err := jsonparser.JSONParserstring(data2, "json", "data", "name")
	if err != nil {
		panic(err)
	}
	fmt.Println(name2)
}
