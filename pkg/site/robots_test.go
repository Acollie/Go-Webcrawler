package site

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchRobots(t *testing.T) {

	t.Run("TestFetchRobots", func(t *testing.T) {

		_, err := canVisitURL("https://__invalid_url")
		require.Error(t, err)
	})
	t.Run("TestFetchRobots", func(t *testing.T) {

		resp, err := canVisitURL("https://www.google.com/search/about/")
		require.NoError(t, err)
		require.Equal(t, resp, true)
	})
	t.Run("TestFetchRobots", func(t *testing.T) {
		resp, err := canVisitURL("https://example.com")
		require.NoError(t, err)
		require.NotNil(t, resp)

		require.Equal(t, true, resp)
	})

}

func TestCanVisitURL(t *testing.T) {
	t.Run("ReturnsErrorWhenURLIsInvalid", func(t *testing.T) {
		_, err := canVisitURL("::invalid_url")
		require.Error(t, err)
	})

	t.Run("ReturnsTrueWhenRobotsTxtNotFound", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusNotFound)
		}))
		defer server.Close()

		canVisit, err := canVisitURL(server.URL)
		require.NoError(t, err)
		require.True(t, canVisit)
	})

	t.Run("ReturnsFalseWhenPathIsDisallowed", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.Write([]byte("User-agent: *\nDisallow: /"))
		}))
		defer server.Close()

		canVisit, err := canVisitURL(server.URL)
		require.NoError(t, err)
		require.False(t, canVisit)
	})

	t.Run("ReturnsTrueWhenPathIsAllowed", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.Write([]byte("User-agent: *\nDisallow:"))
		}))
		defer server.Close()

		canVisit, err := canVisitURL(server.URL)
		require.NoError(t, err)
		require.True(t, canVisit)
	})
}
