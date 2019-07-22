package models

type People struct {
	Craft string `json:"craft"`
	Name  string `json:name`
}

type AstrosResponse struct {
	Peoples []People
	Number  int    `json:"number"`
	Message string `json:"message"`
}

type RESTAPI interface {
	Get() (int, error)
	fetch() ([]byte, error)
}

type Entry struct {
	ID           int    `json: "id,omitemtpy"`
	FirstName    string `json: "first_name,omitempty"`
	LastName     string `json: "last_name,omitempty"`
	EmailAddress string `json: "email_address,omitempty"`
}
