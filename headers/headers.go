package headers

import (
	gjson "github.com/Gohryt/Impossible.go/json"
	"net/http"
)

type (
	Headers map[string]string
)

func (h *Headers) FromMemory(data []byte, errorHandler func(*error)) {
	gjson.FromMemory(data, h, errorHandler)
	return
}

func (h *Headers) FromFile(filePath string, errorHandler func(*error)) {
	gjson.FromFile(filePath, h, errorHandler)
	return
}

func (h *Headers) ToMemory(errorHandler func(*error)) (data []byte) {
	data = gjson.ToMemory(*h, errorHandler)
	return
}

func (h *Headers) ToFile(filePath string, errorHandler func(*error)) {
	gjson.ToFile(*h, filePath, errorHandler)
	return
}

func (h *Headers) ForRequest(request *http.Request) {
	for key, value := range *h {
		request.Header.Set(key, value)
	}
}

func (h *Headers) ForResponse(response http.ResponseWriter) {
	for key, value := range *h {
		response.Header().Set(key, value)
	}
}
