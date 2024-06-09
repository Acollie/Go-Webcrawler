package openSearchx

func (h OpenSearchHandler) DeleteIndex() error {
	_, err := h.OpenSearchClient.Indices.Delete([]string{h.index})
	if err != nil {
		return err
	}
	return nil

}

func (h OpenSearchHandler) DeleteDocument(doc Document) error {
	_, err := h.OpenSearchClient.Delete(h.index, doc.Url)
	if err != nil {
		return err
	}
	return nil
}
