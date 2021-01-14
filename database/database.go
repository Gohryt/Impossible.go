package database

import gjson "github.com/Gohryt/Impossible.go/json"

type (
	//Configuration is struct wich contains database connection configuration
	Configuration struct {
		Host     string `json:"Host"`
		Port     int    `json:"Port"`
		Name     string `json:"Name"`
		User     string `json:"User"`
		Password string `json:"Password"`
	}
)

//FromMemory function reads json from byte array
func (c *Configuration) FromMemory(data []byte, errorHandler func(*error)) {
	gjson.FromMemory(data, c, errorHandler)
	return
}

//FromFile function reads json from file
func (c *Configuration) FromFile(filePath string, errorHandler func(*error)) {
	gjson.FromFile(filePath, c, errorHandler)
	return
}

//ToMemory function returns json as byte array
func (c *Configuration) ToMemory(errorHandler func(*error)) (data []byte) {
	data = gjson.ToMemory(*c, errorHandler)
	return
}

//ToFile function saves json to file
func (c *Configuration) ToFile(filePath string, errorHandler func(*error)) {
	gjson.ToFile(*c, filePath, errorHandler)
	return
}
