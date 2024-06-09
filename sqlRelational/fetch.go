package sqlRelational

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"webcrawler/pkg/site"
)

func (c *SqlDB) FetchWebsite(url string) (*site.Website, error) {
	queryString := fmt.Sprintf("SELECT baseurl, promanceValue FROM website where baseurl = '%s' limit 1", url)
	query := c.Client.QueryRow(queryString)
	var website site.Website
	err := query.Scan(&website.Url, &website.ProminenceValue)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("no record")
			return &site.Website{}, nil
		}
		return &site.Website{}, err
	}
	return &website, nil
}

func (c *SqlDB) FetchPage(url string) (*site.Page, error) {
	queryString := fmt.Sprintf("SELECT pageurl, title, body,baseurl from page where pageurl ='%s' limit 1", url)
	query := c.Client.QueryRow(queryString)
	var page site.Page
	err := query.Scan(&page.Url, &page.Title, &page.Body, &page.BaseURL)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("no page")
			return &site.Page{}, nil
		}
		return &site.Page{}, err
	}
	return &page, nil
}
