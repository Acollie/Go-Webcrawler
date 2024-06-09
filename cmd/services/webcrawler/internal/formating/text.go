package formating

import (
	"github.com/anaskhan96/soup"
	"net/url"
	"path"
	"strings"
	"webcrawler/pkg/queue"
)

const (
	MaxDepth = 8
)

func GetLinks(fetchingURL string, body string) ([]string, error) {
	var links []string
	doc := soup.HTMLParse(body)
	for _, link := range doc.FindAll("a") {
		link := resolveURL(fetchingURL, link.Attrs()["href"])
		links = append(links, link)
	}

	links = removeLargeWebSites(links)
	links = removeAnchors(links)
	links = removeDuplicates(links)
	links = removeMailTo(links)
	links = removeDepthLinks(links)

	return links, nil
}

func ResolveLinkToQueueMessage(links []string) []queue.Message {
	return convertLinksToQueueMessage(links)
}

func convertLinksToQueueMessage(links []string) []queue.Message {
	messages := []queue.Message{}
	for _, link := range links {
		message := queue.Message{Url: link}
		messages = append(messages, message)
	}
	return messages
}

func resolveURL(baseURL, relURL string) string {
	if strings.Contains(relURL, "http") {
		return relURL
	}
	u, err := url.Parse(baseURL)
	if err != nil {
		return ""
	}
	u.Path = path.Join(u.Path, relURL)
	return u.String()
}

func removeDuplicates(links []string) []string {
	encountered := map[string]bool{}
	var result []string

	for v := range links {
		if !encountered[links[v]] {
			encountered[links[v]] = true
			result = append(result, links[v])
		}
	}
	return result
}

func removeDepthLinks(links []string) []string {
	var result []string
	for _, link := range links {
		if testDepthLink(link, MaxDepth) {
			result = append(result, link)
		}

	}
	return result
}

func testDepthLink(link string, maxDepth int) bool {
	res := strings.Split(link, "/")
	return len(res) <= maxDepth+3
}

func removeAnchors(links []string) []string {
	var result []string
	for _, link := range links {
		u, err := url.Parse(link)
		if err != nil {
			return nil
		}
		u.Fragment = ""
		result = append(result, u.String())
	}
	return result
}

func removeMailTo(links []string) []string {
	var result []string
	for _, link := range links {
		u, err := url.Parse(link)
		if err != nil {
			return nil
		}
		if u.Scheme == "mailto" {
			continue
		}
		result = append(result, u.String())
	}
	return result
}

func removeLargeWebSites(links []string) []string {
	largeWebsites := map[string]struct{}{
		"facebook.com": {}, "twitter.com": {}, "instagram.com": {}, "youtube.com": {},
		"linkedin.com": {}, "pinterest.com": {}, "tumblr.com": {}, "reddit.com": {},
		"snapchat.com": {}, "whatsapp.com": {}, "quora.com": {}, "flickr.com": {},
		"vimeo.com": {}, "medium.com": {}, "vk.com": {}, "soundcloud.com": {},
	}

	var result []string
	for i, link := range links {
		url, _ := url.Parse(link)
		if _, ok := largeWebsites[url.Hostname()]; ok {
			continue
		}
		result = append(result, links[i])
		if len(result) == 100 {
			break
		}
	}
	return result
}
