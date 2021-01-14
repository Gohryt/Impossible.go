package regexp

import (
	"regexp"

	gjson "github.com/Gohryt/Impossible.go/json"
)

type (
	//Expression is struct wich contains compiled regexp and fields for reading and saving it to file
	Expression struct {
		Compiled *regexp.Regexp
		Real     struct {
			String string `json:"String"`
			With   string `json:"With"`
		}
	}
)

//Replace function replaces string with
func (r *Expression) Replace(target *string) {
	*target = r.Compiled.ReplaceAllString(*target, r.Real.With)
	return
}

//FromMemory function reads json from byte array
func (r *Expression) FromMemory(data []byte, errorHandler func(*error)) {
	gjson.FromMemory(data, &r.Real, errorHandler)
	r.Compiled = regexp.MustCompile(r.Real.String)
	return
}

//FromFile function reads json from file
func (r *Expression) FromFile(filePath string, errorHandler func(*error)) {
	gjson.FromFile(filePath, &r.Real, errorHandler)
	r.Compiled = regexp.MustCompile(r.Real.String)
	return
}

//ToMemory function returns json as byte array
func (r *Expression) ToMemory(errorHandler func(*error)) (data []byte) {
	data = gjson.ToMemory(r.Real, errorHandler)
	return
}

//ToFile function saves json to file
func (r *Expression) ToFile(filePath string, errorHandler func(*error)) {
	gjson.ToFile(r.Real, filePath, errorHandler)
	return
}
