package headers

import (
	"net/http"

	gjson "github.com/Gohryt/Impossible.go/json"
)

type (
	Headers map[string]string
)

//FromMemory function reads json from byte array
func (h *Headers) FromMemory(data []byte, errorHandler func(*error)) {
	gjson.FromMemory(data, h, errorHandler)
	return
}

//FromFile function reads json from file
func (h *Headers) FromFile(filePath string, errorHandler func(*error)) {
	gjson.FromFile(filePath, h, errorHandler)
	return
}

//ToMemory function returns json as byte array
func (h *Headers) ToMemory(errorHandler func(*error)) (data []byte) {
	data = gjson.ToMemory(*h, errorHandler)
	return
}

//ToFile function saves json to file
func (h *Headers) ToFile(filePath string, errorHandler func(*error)) {
	gjson.ToFile(*h, filePath, errorHandler)
	return
}

//ForRequest function adds headers to http request
func (h *Headers) ForRequest(request *http.Request) {
	var (
		key   string
		value string
	)
	for key, value = range *h {
		request.Header.Set(key, value)
	}
	return
}

//ForResponse function adds headers to http response
func (h *Headers) ForResponse(response http.ResponseWriter) {
	var (
		key   string
		value string
	)
	for key, value = range *h {
		response.Header().Set(key, value)
	}
	return
}
