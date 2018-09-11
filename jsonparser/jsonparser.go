package jsonparser

import (
	"encoding/json"
	"fmt"
)

// JSONParser : convert byte array of data and return path of keys value
// ex. how to use:
// data := []byte(`{"json": {"errors": [], "data": {"url": "www.example.com", "count": 0, "id": "123", "name": "foo"}}}`)
// obj, err := jsonparser.JSONParser(data, "json", "data")
// if err != nil {
// 	panic(err)
// }
// fmt.Println(obj["name"])
func JSONParser(data []byte, keys ...string) (map[string]interface{}, error) {
	var objmap map[string]interface{}
	if err := json.Unmarshal(data, &objmap); err != nil {
		return nil, err
	}
	return getkey(objmap, keys...)
}

func getkey(o map[string]interface{}, keys ...string) (map[string]interface{}, error) {
	if len(keys) > 1 {
		return getkey(o[keys[0]].(map[string]interface{}), keys[1:]...)
	}
	if val, ok := o[keys[0]].(map[string]interface{}); ok {
		return val, nil
	}
	return nil, fmt.Errorf("Unmarshal failed key:%v | value:%v", keys[0], o[keys[0]])
}
