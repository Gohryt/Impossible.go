package json

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//FromMemory function reads json from byte array
func FromMemory(data []byte, target interface{}, errorHandler func(*error)) {
	var (
		err error
	)
	if len(data) > 6 {
		err = json.Unmarshal(data, target)
		errorHandler(&err)
	}
	return
}

//FromFile function reads json from file
func FromFile(filePath string, target interface{}, errorHandler func(*error)) {
	var (
		file *os.File
		data []byte
		err  error
	)
	file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	errorHandler(&err)
	data, err = ioutil.ReadAll(file)
	errorHandler(&err)
	if len(data) > 6 {
		err = json.Unmarshal(data, target)
		errorHandler(&err)
	}
	err = file.Close()
	errorHandler(&err)
	return
}

//ToMemory function returns json as byte array
func ToMemory(target interface{}, errorHandler func(*error)) (data []byte) {
	var (
		err error
	)
	data, err = json.Marshal(target)
	errorHandler(&err)
	return
}

//ToFile function saves json to file
func ToFile(target interface{}, filePath string, errorHandler func(*error)) {
	var (
		file *os.File
		data []byte
		err  error
	)
	data, err = json.Marshal(target)
	errorHandler(&err)
	file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	errorHandler(&err)
	if file != nil {
		_, err = file.Write(data)
		errorHandler(&err)
		err = file.Close()
		errorHandler(&err)
	}
	return
}
