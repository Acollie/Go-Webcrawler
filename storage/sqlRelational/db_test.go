package sqlRelational

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/modules/mariadb"
	"log"
	"reflect"
	"testing"
	"webcrawler/site"
)

func setupDB(db *SqlDB) {
	if _, err := db.Client.Exec(`GRANT ALL PRIVILEGES ON website.* TO 'root'@'localhost'; FLUSH PRIVILEGES;`); err != nil {
		panic(err)
	}
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
	ctx := context.Background()
	mariaDBContainer, err := mariadb.Run(ctx, "mariadb:11.0.3")
	if err != nil {
		log.Fatalf("Failed to run MariaDB container: %v", err)
	}
	defer func() {
		_ = mariaDBContainer.Terminate(ctx) // Ensure container is terminated
	}()

	// Get the port of the running container
	portStr, err := mariaDBContainer.MappedPort(ctx, "3306")
	if err != nil {
		log.Fatalf("Failed to get port endpoint: %v", err)
	}

	// Parse port from string to int if needed
	var port int
	_, err = fmt.Sscanf(string(portStr), "%d", &port)
	if err != nil {
		log.Fatalf("Failed to parse port: %v", err)
	}

	host := "localhost"
	password := "test"
	user := "test"
	dbName := "test"

	sqlDB, err := New(host, password, user, dbName, port)
	if err != nil {
		log.Fatalf("Failed to initialize SQL DB: %v", err)
	}
	setupDB(sqlDB)
	defer teardown(sqlDB)

	t.Run("AddWebsite", func(t *testing.T) {
		website := site.Website{
			Url:             "http://www.google.com",
			ProminenceValue: 1,
		}

		err := sqlDB.AddWebsite(ctx, website)
		require.NoError(t, err)

		websiteReturned, err := sqlDB.FetchWebsite(ctx, website.Url)
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

		err := sqlDB.AddPage(ctx, page)
		require.NoError(t, err)

		pageReturned, err := sqlDB.FetchPage(ctx, page.Url)
		require.NoError(t, err)
		require.Equal(t, &page, pageReturned)
	})

	t.Run("Fetch page empty", func(t *testing.T) {
		pageReturn, err := sqlDB.FetchPage(ctx, "does not exist")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(pageReturn, &site.Page{}), true)
	})

	t.Run("Fetch website empty", func(t *testing.T) {
		pageReturn, err := sqlDB.FetchWebsite(ctx, "does not exist website")
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
		err := sqlDB.UpdateWebsite(ctx, page, website)
		require.NoError(t, err)
		returnDB, err := sqlDB.FetchWebsite(ctx, website.Url)
		require.NoError(t, err)
		require.Equal(t, returnDB.ProminenceValue, website.ProminenceValue+1)

	})
}
