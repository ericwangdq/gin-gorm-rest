package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadFile(name string) (error, map[string]interface{}) {

	file, err := os.Open(name)
	if err != nil {
		return err, nil
	}

	defer file.Close()
	byteValue, err := ioutil.ReadAll(file)

	var result map[string]interface{}
	err = json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		return err, nil
	}

	return nil, result
}
