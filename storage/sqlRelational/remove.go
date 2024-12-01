package sqlRelational

import (
	"context"
	"webcrawler/site"
)

func (db *SqlDB) RemovePage(ctx context.Context, website site.Page) error {
	return nil
}

func (db *SqlDB) RemoveWebsite(ctx context.Context, website site.Website) error {
	return nil
}
