package json

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func FromMemory(data []byte, target interface{}, errorHandler func(*error)) {
	if len(data) > 6 {
		err := json.Unmarshal(data, target)
		errorHandler(&err)
	}
	return
}

func FromFile(filePath string, target interface{}, errorHandler func(*error)) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	errorHandler(&err)
	data, err := ioutil.ReadAll(file)
	errorHandler(&err)
	if len(data) > 6 {
		err = json.Unmarshal(data, target)
		errorHandler(&err)
	}
	err = file.Close()
	errorHandler(&err)
	return
}

func ToMemory(target interface{}, errorHandler func(*error)) (data []byte) {
	data, err := json.Marshal(target)
	errorHandler(&err)
	return
}

func ToFile(target interface{}, filePath string, errorHandler func(*error)) {
	data, err := json.Marshal(target)
	errorHandler(&err)
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	errorHandler(&err)
	if file != nil {
		_, err = file.Write(data)
		errorHandler(&err)
		err = file.Close()
		errorHandler(&err)
	}
	return
}
