package api

import (
	"fmt"
	"net/http"
	"strings"
)

// API ...
type API struct {
	base    string
	version string
}

type QueryParams map[string]interface{}

// Create ...
func Create() *API {
	api := &API{
		base:    "https://api.opencorporates.com",
		version: "v0.4",
	}

	return api
}

func (api *API) constructURL(path string) string {
	p := strings.ToLower(path)
	if strings.HasPrefix(p, "http") || strings.HasPrefix(p, "https") {
		return path
	}

	url := fmt.Sprintf("%s/%s/%s", api.base, api.version, path)

	return url
}

func (api *API) prepareRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", api.constructURL(url), nil)
	if err != nil {
		return nil, fmt.Errorf("Prepare request: %q", err)
	}

	return req, nil
}

func (api *API) getResponse(url string, params QueryParams) (*http.Response, error) {
	req, err := api.prepareRequest(url)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v.(string))
	}
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Accept", "application/json")

	fmt.Println(req.URL)

	client := &http.Client{}
	r, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Get response do: %q", err)
	}

	return r, nil
}

// SearchCompany ...
func (api *API) SearchCompany(query string) (*CompanySearchResults, error) {
	search := &CompanySearch{
		api: api,
	}

	results, err := search.search(query)
	if err != nil {
		return nil, err
	}

	return results, nil
}
