package sqlRelational

import (
	"fmt"
	"webcrawler/pkg/site"
)

func (c *SqlDB) AddWebsite(website site.Website) error {
	queryString := fmt.Sprintf("INSERT INTO website (baseurl, promancevalue) VALUES ('%s', %f)", website.Url, website.ProminenceValue)
	_, err := c.Client.Exec(queryString)

	return err
}

func (c *SqlDB) AddPage(page site.Page) error {
	queryString := fmt.Sprintf("INSERT INTO page (pageurl, title, body,baseurl) VALUES ('%s', '%s', '%s', '%s')", page.Url, page.Title, page.Body, page.BaseURL)
	_, err := c.Client.Exec(queryString)

	return err
}
