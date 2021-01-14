package templates

type (
	//Description is struct wich contains basic fields for html filling
	Description struct {
		Title       string `json:"Title"`
		Description string `json:"Description"`
		Keywords    string `json:"Keywords"`
		Domain      string `json:"Domain"`
		Type        string `json:"Type"`
		Author      string `json:"Author"`
	}
)
