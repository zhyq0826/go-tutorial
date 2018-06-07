package resources

import "time"

type (
	// DomainForm for json data
	DomainForm struct {
		Name      string    `json:"name"`
		URL       string    `json:"url"`
		Private   uint      `json:"private"`
		CreatedAt time.Time `json:"created_at,omitempty"`
		UpdatedAt time.Time `json:"updated_at,omitempty"`
	}
)
