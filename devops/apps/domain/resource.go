package domain

type (
	domainForm struct {
		Name    string `json:"name"`
		URL     string `json:"url"`
		Private uint   `json:"private"`
	}
)
