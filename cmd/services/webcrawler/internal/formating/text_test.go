package formating

import (
	"github.com/stretchr/testify/require"
	"testing"
	"webcrawler/pkg/site"
)

func TestFetchLinks(t *testing.T) {
	link := "http://www.alexcollie.com/"

	t.Run("Test basic", func(t *testing.T) {
		_, resp, err := site.NewPage(link)
		links, err := GetLinks(link, resp)
		require.NoError(t, err)
		require.Equal(t, len(links) > 0, true)

	})

}

func TestLinkResolve(t *testing.T) {
	url := resolveURL("http://www.example.com/", "test")
	require.Equal(t, url, "http://www.example.com/test", "Resolving relative URL failed")
	url = resolveURL("http://www.example.com/", "/test")
	require.Equal(t, url, "http://www.example.com/test", "Resolving absolute URL failed")

	// Test with another base URL
	url = resolveURL("https://www.anotherexample.com/", "/sample")
	require.Equal(t, url, "https://www.anotherexample.com/sample", "Resolving absolute URL with https base URL failed")

	// Test with empty path
	url = resolveURL("https://www.anotherexample.com/", "")
	require.Equal(t, url, "https://www.anotherexample.com/", "Resolving with empty path failed")

	// Test with path that has a leading slash and subdirectory
	url = resolveURL("https://www.anotherexample.com", "/a/b/c")
	require.Equal(t, url, "https://www.anotherexample.com/a/b/c", "Resolving path with subdirectories failed")

	// Test with path without a leading slash and with subdirectory
	url = resolveURL("https://www.anotherexample.com", "a/b/c")
	require.Equal(t, url, "https://www.anotherexample.com/a/b/c", "Resolving path without leading slash and with subdirectories failed")

}

func TestRemoveDepthLinks(t *testing.T) {
	tests := []struct {
		url      string
		expected bool
	}{
		{"http://www.example.com/", true},
		{"http://www.example.com/test", true},
		{"http://www.example.com/test/test", true},
		{"http://www.example.com/test/test/test", true},
		{"http://www.example.com/test/test/test/test", true},
		{"http://www.example.com/test/test/test/test/test", true},
		{"http://www.example.com/test/test/test/test/test/test", false},
	}
	t.Run("Test basic", func(t *testing.T) {
		for _, test := range tests {
			result := testDepthLink(test.url, 5)
			require.Equalf(t, result, test.expected, "Removing depth links failed %s", test.url)
		}

	})
}

func TestRemoveDuplicates(t *testing.T) {
	links := []string{"http://www.example.com/", "http://www.example.com/", "http://www.example.com/test", "http://www.example.com/test"}
	links = removeDuplicates(links)
	require.Equal(t, len(links), 2, "Removing duplicates failed")
}

func TestRemoveAnchors(t *testing.T) {
	links := []string{"http://www.example.com/#", "http://www.example.com/test#"}
	links = removeAnchors(links)
	require.Equal(t, len(links), 2, "Removing anchors failed")
}
