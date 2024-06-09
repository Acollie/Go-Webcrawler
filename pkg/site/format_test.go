package site

import (
	"github.com/anaskhan96/soup"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func loadExampleHTML(fileName string) string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func TestFullPage(t *testing.T) {
	html := loadExampleHTML("example.html")
	root := soup.HTMLParse(html)
	res := fetchText(&root)
	require.LessOrEqual(t, len(res), MaxBodyLength, "Expected empty string, got %s", res)
}

func TestFetchMeta(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]string
	}{
		{
			name:  "Test basic",
			input: `<meta name="description" content="This is a test description">`,
			expected: map[string]string{
				"description": "This is a test description",
			},
		},
		{
			name:  "Test with multiple meta",
			input: `<meta name="description" content="This is a test description"><meta name="author" content="Alex Collie">`,
			expected: map[string]string{
				"description": "This is a test description",
				"author":      "Alex Collie",
			},
		},
		{
			name:     "Test with no meta",
			input:    `<title>This is a test</title>`,
			expected: map[string]string{},
		},
		{
			name:     "Test with no content",
			input:    `<meta name="description">`,
			expected: map[string]string{},
		},
		{
			name:     "Test with no name",
			input:    `<meta content="This is a test description">`,
			expected: map[string]string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root := soup.HTMLParse(test.input)
			res := fetchMeta(&root)
			require.Equal(t, test.expected, res)
		})

	}
}

func TestStripJavascript(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test basic",
			input:    "This is a test <script>console.log('test')</script>",
			expected: "This is a test ",
		},
		{
			name:     "Test with multiple scripts",
			input:    "This is a test <script>console.log('test')</script> <script>console.log('test')</script>",
			expected: "This is a test  ",
		},
		{
			name:     "Test with no script",
			input:    "This is a test",
			expected: "This is a test",
		},
		{
			name:     "Complex ref test",
			input:    `<!DOCTYPE html><html lang="en"><head><meta charSet="utf-8"/><meta name="viewport" content="width=device-width"/><title>Alex Collie&#x27;s personal site</title><meta name="description" content="Alex Collie&#x27;s personal website, a site for a software developer based in London"/><meta name="next-head-count" content="4"/><link rel="preload" href="/_next/static/css/090c144453b7cc79.css" as="style"/><link rel="stylesheet" href="/_next/static/css/090c144453b7cc79.css" data-n-g=""/><link rel="preload" href="/_next/static/css/ae4ed9c503fd1e33.css" as="style"/><link rel="stylesheet" href="/_next/static/css/ae4ed9c503fd1e33.css" data-n-p=""/><noscript data-n-css=""></noscript><script defer="" nomodule="" src="/_next/static/chunks/polyfills-78c92fac7aa8fdd8.js"></script><script src="/_next/static/chunks/webpack-b8f8d6679aaa5f42.js" defer=""></script><script src="/_next/static/chunks/framework-66d32731bdd20e83.js" defer=""></script><script src="/_next/static/chunks/main-d190bcef2284c937.js" defer=""></script>`,
			expected: `<!DOCTYPE html><html lang="en"><head><meta charSet="utf-8"/><meta name="viewport" content="width=device-width"/><title>Alex Collie&#x27;s personal site</title><meta name="description" content="Alex Collie&#x27;s personal website, a site for a software developer based in London"/><meta name="next-head-count" content="4"/><link rel="preload" href="/_next/static/css/090c144453b7cc79.css" as="style"/><link rel="stylesheet" href="/_next/static/css/090c144453b7cc79.css" data-n-g=""/><link rel="preload" href="/_next/static/css/ae4ed9c503fd1e33.css" as="style"/><link rel="stylesheet" href="/_next/static/css/ae4ed9c503fd1e33.css" data-n-p=""/><noscript data-n-css=""></noscript>`,
		},
		{
			name:     "Complex ref test 2",
			input:    `<script>window.env={"AD_SLOT_CLIENT_INJECTOR_REGISTRY":"https://test.com"};</script>`,
			expected: ``,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.input
			stripJavascript(&result)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestRemoveCSS(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test basic",
			input:    "This is a test <style>body{color: red}</style>",
			expected: "This is a test ",
		},
		{
			name:     "Test with multiple styles",
			input:    "This is a test <style>body{color: red}</style> <style>body{color: red}</style>",
			expected: "This is a test  ",
		},
		{
			name:     "Test with no style",
			input:    "This is a test",
			expected: "This is a test",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.input
			stripCSS(&result)
			require.Equal(t, test.expected, result)
		})
	}
}
