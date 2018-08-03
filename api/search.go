package api

// CompanySearchResults ...
type CompanySearchResults struct {
	Version string              `json:"api_version"`
	Results CompanySearchResult `json:"results"`
}

// CompanySearchResult ...
type CompanySearchResult struct {
	Companies  []CompanyList `json:"companies"`
	Page       int           `json:"page"`
	PerPage    int           `json:"per_page"`
	TotalCount int           `json:"total_count"`
	TotalPages int           `json:"total_pages"`
}

// CompanyList ...
type CompanyList struct {
	Company CompanyResult `json:"company"`
}

// CompanyResult ...
type CompanyResult struct {
	CompanyNumber     string                   `json:"company_number"`
	CompanyType       string                   `json:"company_type"`
	CurrentStatus     string                   `json:"current_status"`
	DissolutionDate   string                   `json:"dissolution_date"`
	Inactive          bool                     `json:"bool"`
	IncorporationDate string                   `json:"incorporation_date"`
	JurisdictionCode  string                   `json:"jurisdiction_code"`
	Name              string                   `json:"name"`
	OpenCorporatesURL string                   `json:"opencorporates_url"`
	RegisteredAddress CompanyRegisteredAddress `json:"registered_address"`
	RegistryURL       string                   `json:"registry_url"`
}

// CompanyRegisteredAddress ...
type CompanyRegisteredAddress struct {
	Country       string `json:"country"`
	Locality      string `json:"locality"`
	PostalCode    string `json:"postal_code"`
	Region        string `json:"region"`
	StreetAddress string `json:"street_address"`
}
