package configuration

import (
	"encoding/json"
	"os"
)

// JSON : to parse config json file in Config object.
// path : file should be json file.
// Cofing : result object
func JSON(path string, config interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(config)
	if err != nil {
		return err
	}
	return nil
}
