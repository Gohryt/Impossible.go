package database

import gjson "github.com/Gohryt/Impossible.go/json"

type (
	Configuration struct {
		Host     string `json:"Host"`
		Port     int    `json:"Port"`
		Name     string `json:"Name"`
		User     string `json:"User"`
		Password string `json:"Password"`
	}
)

func (c *Configuration) FromMemory(data []byte, errorHandler func(*error)) {
	gjson.FromMemory(data, c, errorHandler)
	return
}

func (c *Configuration) FromFile(filePath string, errorHandler func(*error)) {
	gjson.FromFile(filePath, c, errorHandler)
	return
}

func (c *Configuration) ToMemory(errorHandler func(*error)) (data []byte) {
	data = gjson.ToMemory(*c, errorHandler)
	return
}

func (c *Configuration) ToFile(filePath string, errorHandler func(*error)) {
	gjson.ToFile(*c, filePath, errorHandler)
	return
}
