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
	return Getkey(objmap, keys...)
}

// Getkey : get key from map string recursively
func Getkey(o map[string]interface{}, keys ...string) (map[string]interface{}, error) {
	if len(keys) > 1 {
		return Getkey(o[keys[0]].(map[string]interface{}), keys[1:]...)
	}
	if val, ok := o[keys[0]].(map[string]interface{}); ok {
		return val, nil
	}
	return nil, fmt.Errorf("Unmarshal failed key:%v | value:%v", keys[0], o[keys[0]])
}

// Getkeystring : get key from map string recursively
func Getkeystring(o map[string]interface{}, keys ...string) (string, error) {
	if len(keys) > 1 {
		if val, ok := o[keys[0]].(map[string]interface{}); ok {
			return Getkeystring(val, keys[1:]...)
		} else if val, ok := o[keys[0]].(string); ok {
			var objmap map[string]interface{}
			if err := json.Unmarshal([]byte(val), &objmap); err == nil {
				return Getkeystring(objmap, keys[1:]...)
			}
		}
	}
	if val, ok := o[keys[0]].(string); ok {
		return val, nil
	}
	return "", fmt.Errorf("Unmarshal failed key:%v | value:%v", keys[0], o[keys[0]])
}
