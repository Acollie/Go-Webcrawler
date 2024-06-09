package openSearchx

import (
	"encoding/json"
	"github.com/opensearch-project/opensearch-go/v2/opensearchapi"
	"strings"
	"webcrawler/pkg/site"
)

func (h OpenSearchHandler) Get(doc site.Page) error {
	_, err := h.OpenSearchClient.Get(h.index, doc.Url)
	if err != nil {
		return err
	}
	return nil
}

func (h OpenSearchHandler) GetIndex() error {
	_, err := h.OpenSearchClient.Indices.Get([]string{h.index})
	if err != nil {
		return err
	}
	return nil
}

func (h OpenSearchHandler) Search(query string) ([]site.Page, error) {
	res, err := h.OpenSearchClient.Search(
		func(req *opensearchapi.SearchRequest) {
			req.Index = []string{h.index}
			req.Body = strings.NewReader(query)
		},
	)
	if err != nil {
		return nil, err
	}

	var searchRequest SearchResponse

	if err := json.NewDecoder(res.Body).Decode(&searchRequest); err != nil {
		return nil, err
	}

	var result []site.Page
	for _, hit := range searchRequest.Hits.Hits {
		result = append(result, hit.Source)
	}

	return result, nil
}
