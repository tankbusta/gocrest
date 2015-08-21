package gocrest

import (
	"fmt"
)

type AllianceDetails struct {
	StartDate       string `json:"startDate"`
	NumCorporations int    `json:"corporationsCount"`
	Description     string `json:"description"`
	Url             string `json:"url"`
	ShortName       string `json:"shortName"`
	Name            string `json:"name"`
}

type AlliancePage struct {
	Alliances []struct {
		AlliancePageDetails `json:"href"`
	} `json:"items"`

	PageCount  int `json:"pageCount"`
	TotalCount int `json:"totalCount"`

	currentPage int
}

type AlliancePageDetails struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Ticker string `json:"shortName"`
}

func (s AlliancePageDetails) GetHref() string {
	return fmt.Sprintf("%s/alliances/%d/", PUBLIC_CREST, s.ID)
}

func GetAlliances() (*AlliancePage, error) {
	ap := &AlliancePage{}
	ep := fmt.Sprintf("%s/alliances/", PUBLIC_CREST)
	// Make the request and unmarshal the response...
	if err := c.doGetRequest(ep, ap); err != nil {
		return nil, err
	}
	// First Page...
	if ap.currentPage == 0 {
		ap.currentPage = 1
	}
	return ap, nil
}
