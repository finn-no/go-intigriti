package v2

const (
// programURI = "/company/v2/programs"
)

type Program struct {
	ID            string `json:"id"`
	Handle        string `json:"handle"`
	CompanyID     string `json:"companyId"`
	CompanyHandle string `json:"companyHandle"`
	LogoURL       string `json:"logoUrl"`
	Name          string `json:"name"`
	Status        struct {
		ID    int    `json:"id"`
		Value string `json:"value"`
	} `json:"status"`
	ConfidentialityLevel struct {
		ID    int    `json:"id"`
		Value string `json:"value"`
	} `json:"confidentialityLevel"`
	WebLinks struct {
		Details string `json:"details"`
	} `json:"webLinks"`
	Type struct {
		ID    int    `json:"id"`
		Value string `json:"value"`
	} `json:"type"`
}
