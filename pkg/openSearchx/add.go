package openSearchx

import (
	"encoding/json"
	"strings"
)

func (h OpenSearchHandler) AddDocument(doc Document) error {
	docBytes, err := json.Marshal(doc)
	if err != nil {
		return err
	}
	_, err = h.OpenSearchClient.Index(h.index, strings.NewReader(string(docBytes)))
	if err != nil {
		return err
	}
	return nil
}
func (h OpenSearchHandler) CreateIndex() error {
	_, err := h.OpenSearchClient.Indices.Create(h.index)
	if err != nil {
		return err
	}
	return nil
}
