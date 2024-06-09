package site

import (
	"github.com/anaskhan96/soup"
)

type pageI interface {
	Fetch(url string) (*soup.Root, error)
	Save(website Page) error
}
type Page struct {
	Url         string            `dynamodbav:"PageURL" json:"url,omitempty"`
	Title       string            `dynamodbav:"title" json:"title,omitempty"`
	Body        string            `dynamodbav:"body" json:"body,omitempty"`
	BaseURL     string            `dynamodbav:"BaseURL" json:"baseURL,omitempty"`
	Meta        map[string]string `dynamodbav:"-" json:"meta,omitempty"`
	CrawledDate uint64            `dynamodbav:"crawledDate" json:"crawledDate,omitempty"`
	Links       []string          `dynamodbav:"-" `
}

type Website struct {
	Url             string   `dynamodbav:"BaseURL"`
	Links           []string `dynamodbav:"links"`
	ProminenceValue float64  `dynamodbav:"promanceValue"`
}
