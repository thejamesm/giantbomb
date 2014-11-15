package giantbomb

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	ResourceTypeGame      = "game"
	ResourceTypeFranchise = "franchise"
	ResourceTypeCharacter = "character"
	ResourceTypeConcept   = "concept"
	ResourceTypeObject    = "object"
	ResourceTypeLocation  = "location"
	ResourceTypePerson    = "person"
	ResourceTypeCompany   = "company"
	ResourceTypeVideo     = "video"
)

var (
	Host = "https://www.giantbomb.com/api/"
	Key  = "" // Your API key
)

type Response struct {
	Error                string
	Limit                int
	Offset               int
	NumberOfPageResults  int
	NumberOfTotalResults int
	Results              []Result
}

type Result struct{}

// Pass empty string to resourceID if you don't need to specify it.
func getBaseURL(resourceType string, resourceID string) string {
	start := Host + resourceType + "/"
	if resourceID != "" {
		start += resourceID + "/"
	}
	return start + "?format=json&api_key=" + Key
}

func Search(query string, limit int, page int, resources []string) ([]byte, error) {
	url := getBaseURL("search", "") +
		"&query=\"" + query + "\"" +
		"&limit=" + strconv.Itoa(limit) +
		"&page=" + strconv.Itoa(page) +
		"&resources=" + strings.Join(resources, ",")
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Request to %s failed (%s)!", url, resp.Status))
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	// TODO: Implement JSON parsing.
	return body, nil
}
