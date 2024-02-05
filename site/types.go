package site

import (
	"github.com/anaskhan96/soup"
)

type pageI interface {
	Fetch(url string) (*soup.Root, error)
	Save(website Page) error
}
type Page struct {
	Url     string `dynamodbav:"PageURL"`
	Title   string `dynamodbav:"title"`
	Body    string `dynamodbav:"body"`
	BaseURL string `dynamodbav:"BaseURL"`
}

type Website struct {
	Url             string   `dynamodbav:"BaseURL"`
	Links           []string `dynamodbav:"links"`
	ProminenceValue float64  `dynamodbav:"promanceValue"`
}
