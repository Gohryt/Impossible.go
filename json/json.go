package json

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//FromMemory function reads json from byte array
func FromMemory(data []byte, target interface{}, errorHandler func(error)) {
	err := json.Unmarshal(data, target)
	errorHandler(err)
}

//FromFile function reads json from file
func FromFile(path string, target interface{}, errorHandler func(error)) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	errorHandler(err)
	data, err := ioutil.ReadAll(file)
	errorHandler(err)
	err = json.Unmarshal(data, target)
	errorHandler(err)
	err = file.Close()
	errorHandler(err)
}

//ToMemory function returns json as byte array
func ToMemory(target interface{}, errorHandler func(error)) (data []byte) {
	data, err := json.Marshal(target)
	errorHandler(err)

	return
}

//ToFile function saves json to file
func ToFile(target interface{}, path string, errorHandler func(error)) (i int) {
	data, err := json.Marshal(target)
	errorHandler(err)
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	errorHandler(err)
	i, err = file.Write(data)
	errorHandler(err)
	err = file.Close()
	errorHandler(err)

	return
}
