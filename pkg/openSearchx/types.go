package openSearchx

import (
	"github.com/opensearch-project/opensearch-go/v2"
	"log"
	"webcrawler/pkg/site"
)

type Document struct {
	Url       string `json:"url"`
	DateAdded int64  `json:"date_added"`
	Expires   int64  `json:"expires"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Relevance int    `json:"relevance"`
}

type SearchResponse struct {
	Hits struct {
		Hits []struct {
			Source site.Page `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

type OpenSearchHandler struct {
	Url              string
	OpenSearchClient *opensearch.Client
	index            string
}

func New(url string, index string) *OpenSearchHandler {
	// Create an opensearch client and use the request-signer
	client, err := opensearch.NewClient(opensearch.Config{
		Addresses: []string{url},
	})
	if err != nil {
		log.Fatal("client creation err", err)
	}
	return &OpenSearchHandler{
		Url:              url,
		index:            index,
		OpenSearchClient: client,
	}
}
