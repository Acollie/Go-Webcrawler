package site

import (
	"fmt"
	"github.com/temoto/robotstxt"
	"net/http"
	"net/url"
	"time"
	"webcrawler/pkg/spider"
)

func FetchRobots(baseUrl string) (bool, error) {
	return canVisitURL(baseUrl)
}

func canVisitURL(targetURL string) (bool, error) {
	parsedURL, err := url.Parse(targetURL)
	if err != nil {
		return false, err
	}

	client := &http.Client{
		Timeout: time.Second * 10, // Set timeout to 10 seconds
	}

	resp, err := client.Get(fmt.Sprintf("%s://%s/robots.txt", parsedURL.Scheme, parsedURL.Host))
	if err != nil {
		return false, err
	}
	if resp.StatusCode == http.StatusNotFound {
		return true, nil
	}

	defer resp.Body.Close()

	robots, err := robotstxt.FromResponse(resp)
	if err != nil {
		return false, err
	}
	group := robots.FindGroup(spider.UserAgent)

	return group.Test(parsedURL.Path), nil
}
