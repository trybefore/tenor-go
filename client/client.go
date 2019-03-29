package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Endpoint string

const (
	SearchEndpoint           Endpoint = "search"
	CategoriesEndpoint       Endpoint = "categories"
	SearchSuggestionEndpoint Endpoint = "search_suggestions"
	AutoCompleteEndpoint     Endpoint = "autocomplete"
	TrendingTermsEndpoint    Endpoint = "trending_terms"
	RegisterShareEndpoint    Endpoint = "registershare"
	GIFSEndpoint             Endpoint = "gifs"
	RandomEndpoint           Endpoint = "random"
	tenorURL                          = "https://api.tenor.com/v1/%s?key=%s&"
)

// TenorClient client holding
type TenorClient struct {
	apiKey string
	client *http.Client
}

// NewTenorClient returns a new tenorclient with the provided api key.
func NewTenorClient(apiKey string) *TenorClient {
	return &TenorClient{
		apiKey: apiKey,
		client: &http.Client{
			Timeout: time.Second * 15,
		},
	}
}

// Search searches the endpoint using the provided query.
// you do not need provide a key.
func (t *TenorClient) Search(endpoint Endpoint, query map[string]interface{}) *TenorResponse {
	return t.search(endpoint, query)
}

func (t *TenorClient) search(endpoint Endpoint, query map[string]interface{}) *TenorResponse {
	response := TenorResponse{}

	url := t.formURL(endpoint, query)
	resp, err := t.client.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &response)

	return &response
}

func (t *TenorClient) formURL(endpoint Endpoint, query map[string]interface{}) string {
	u := fmt.Sprintf(tenorURL, string(endpoint), t.apiKey)

	queries := []string{}

	for k, v := range query {
		value := formatQuery(fmt.Sprintf("%v", v))
		queries = append(queries, fmt.Sprintf("%s=%v", formatQuery(k), value))
	}

	q := strings.Join(queries, "&")

	u = u + q

	return u
}

func formatQuery(u string) string {
	return url.QueryEscape(u)
}
