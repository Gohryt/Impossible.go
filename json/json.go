package json

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//FromMemory function reads json from byte array
func FromMemory(data []byte, target interface{}) (err error) {
	err = json.Unmarshal(data, target)

	return
}

//FromFile function reads json from file
func FromFile(path string, target interface{}) (err error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, target)
	if err != nil {
		return
	}
	err = file.Close()

	return
}

//ToMemory function returns json as byte array
func ToMemory(target interface{}) (data []byte, err error) {
	data, err = json.Marshal(target)

	return
}

//ToFile function saves json to file
func ToFile(target interface{}, path string, errorHandler func(error)) (i int, err error) {
	data, err := json.Marshal(target)
	if err != nil {
		return
	}
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	i, err = file.Write(data)
	if err != nil {
		return
	}
	err = file.Close()

	return
}
