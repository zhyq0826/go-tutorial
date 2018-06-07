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

	// ComputerForm for json data
	ComputerForm struct {
		CPU       uint      `json:"cpu"`
		RAM       uint      `json:"ram"`
		PrivateIP string    `json:"private_ip"`
		PublicIP  string    `json:"public_ip"`
		CreatedAt time.Time `json:"created_at,omitempty"`
		UpdatedAt time.Time `json:"updated_at,omitempty"`
	}
)
