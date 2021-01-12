package regexp

import (
	gjson "github.com/Gohryt/Impossible.go/json"
	"regexp"
)

type (
	Expression struct {
		Compiled *regexp.Regexp
		Real     struct {
			String string `json:"String"`
			With   string `json:"With"`
		}
	}
)

func (r *Expression) Replace(target *string) {
	*target = r.Compiled.ReplaceAllString(*target, r.Real.With)
	return
}

func (r *Expression) FromMemory(data []byte, errorHandler func(*error)) {
	gjson.FromMemory(data, &r.Real, errorHandler)
	r.Compiled = regexp.MustCompile(r.Real.String)
	return
}

func (r *Expression) FromFile(filePath string, errorHandler func(*error)) {
	gjson.FromFile(filePath, &r.Real, errorHandler)
	r.Compiled = regexp.MustCompile(r.Real.String)
	return
}

func (r *Expression) ToMemory(errorHandler func(*error)) (data []byte) {
	data = gjson.ToMemory(r.Real, errorHandler)
	return
}

func (r *Expression) ToFile(filePath string, errorHandler func(*error)) {
	gjson.ToFile(r.Real, filePath, errorHandler)
	return
}
