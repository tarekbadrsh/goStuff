package configuration

import (
	"encoding/json"
	"os"
)

// ParseJSONConfiguration : to parse config json file in Config object.
// path : file should be json file.
// Cofing : result object
func ParseJSONConfiguration(path string, Config interface{}) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)

	err = decoder.Decode(&Config)
	file.Close()
	if err != nil {
		file.Close()
		panic(err)
	}
}
