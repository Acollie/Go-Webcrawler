package sqlRelational

import (
	"fmt"
	"reflect"
	"webcrawler/pkg/site"
)

func (c *SqlDB) UpdateWebsite(page site.Page, website site.Website) error {

	websiteDB, err := c.FetchWebsite(page.BaseURL)
	if err != nil {
		return err
	}
	if reflect.DeepEqual(websiteDB, &site.Website{}) && err == nil {
		println("websiteNot found")
		err = c.AddWebsite(website)
		if err != nil {
			return err
		}
		return nil
	}
	websiteDB.ProminenceValue += 1
	queryString := fmt.Sprintf("UPDATE website SET promancevalue = promancevalue + 1 WHERE baseurl = '%s' ", website.Url)
	_, err = c.Client.Query(queryString)

	return err
}

func (c *SqlDB) UpdatePage(page site.Page) error {

	return nil
}
