package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// CompanyAddress ...
type CompanyAddress struct {
	AddressLine1 string
	AddressLine2 string
	CareOf       string
	Country      string
	Locality     string
	POBox        string
	PostalCode   string
	Region       string
}

// CompanySearch ...
type CompanySearch struct {
	api *API
}

func (csr *CompanySearch) search(query string) (*CompanySearchResults, error) {
	r := &CompanySearchResults{}
	var params = QueryParams{"q": query}

	resp, err := csr.api.getResponse("companies/search", params)
	if err != nil {
		return nil, errors.New("xxxx")
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, dataErr := ioutil.ReadAll(resp.Body)
	if dataErr != nil {
		return nil, errors.New("yyyyy")
	}
	fmt.Println("response Body:", string(body))

	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
