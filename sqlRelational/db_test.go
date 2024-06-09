package sqlRelational

import (
	"github.com/stretchr/testify/require"
	"os"
	"reflect"
	"testing"
	"webcrawler/pkg/site"
)

func setupDB(db *SqlDB) {
	if _, err := db.Client.Exec(addWebsite); err != nil {
		panic(err)
	}
	if _, err := db.Client.Exec(addPage); err != nil {
		panic(err)
	}
}
func teardown(db *SqlDB) {
	_, err := db.Client.Exec(dropPage)
	if err != nil {
		panic(err)
	}
	_, err = db.Client.Exec(dropWebsite)
	if err != nil {
		panic(err)
	}
}

func TestDB(t *testing.T) {
	if os.Getenv("ENVIRONMENT") != "local" {
		t.Skipf("skipping since not local")
	}

	sqlDB := New("webcrawlertest")
	setupDB(sqlDB)
	defer teardown(sqlDB)

	t.Run("AddWebsite", func(t *testing.T) {
		website := site.Website{
			Url:             "http://www.google.com",
			ProminenceValue: 1,
		}

		err := sqlDB.AddWebsite(website)
		require.NoError(t, err)

		websiteReturned, err := sqlDB.FetchWebsite(website.Url)
		require.NoError(t, err)
		require.Equal(t, &website, websiteReturned)
	})

	t.Run("AddPage", func(t *testing.T) {
		page := site.Page{
			Url:     "http://www.google.com",
			Title:   "Google",
			Body:    "Search Engine",
			BaseURL: "http://www.google.com",
		}

		err := sqlDB.AddPage(page)
		require.NoError(t, err)

		pageReturned, err := sqlDB.FetchPage(page.Url)
		require.NoError(t, err)
		require.Equal(t, &page, pageReturned)
	})

	t.Run("Fetch page empty", func(t *testing.T) {
		pageReturn, err := sqlDB.FetchPage("does not exist")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(pageReturn, &site.Page{}), true)
	})

	t.Run("Fetch website empty", func(t *testing.T) {
		pageReturn, err := sqlDB.FetchWebsite("does not exist website")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(pageReturn, &site.Website{}), true)
	})

	t.Run("update website", func(t *testing.T) {
		page := site.Page{
			Url:     "http://www.google.com",
			Title:   "Google",
			Body:    "Search Engine",
			BaseURL: "http://www.google.com",
		}

		website := site.Website{
			Url:             "http://www.google.com",
			ProminenceValue: 1,
		}
		err := sqlDB.UpdateWebsite(page, website)
		require.NoError(t, err)
		returnDB, err := sqlDB.FetchWebsite(website.Url)
		require.NoError(t, err)
		require.Equal(t, returnDB.ProminenceValue, website.ProminenceValue+1)

	})
}
