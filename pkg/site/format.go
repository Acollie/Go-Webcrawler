package site

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
	"webcrawler/pkg/queue"
)

const (
	MaxBodyLength = 10000
)

func NewPage(fetchURL string) (Page, string, error) {
	client := &http.Client{
		Timeout: time.Second * 10, // Set timeout to 10 seconds
	}

	resp, err := client.Get(fetchURL)
	if err != nil {
		return Page{}, "", fmt.Errorf("get %s", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return Page{}, "", fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	// response into string
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return Page{}, "", err
	}

	doc := soup.HTMLParse(string(body))
	titleTag := doc.Find("title")
	title := ""
	if titleTag.Error == nil {
		title = titleTag.Text()
	} else {
		title = "title not found"
	}

	baseUrl, _ := url.Parse(fetchURL)
	return Page{
		Url:         fetchURL,
		Title:       title,
		Body:        fetchText(&doc),
		BaseURL:     baseUrl.Hostname(),
		CrawledDate: uint64(time.Now().Unix()),
		Meta:        fetchMeta(&doc),
	}, string(body), nil
}

func fetchMeta(root *soup.Root) map[string]string {
	metaInformation := make(map[string]string)
	meta := root.FindAll("meta")
	for _, value := range meta {
		name := value.Attrs()["name"]
		content := value.Attrs()["content"]
		if name != "" && content != "" {
			metaInformation[name] = content
		}
	}
	if len(metaInformation) == 0 {
		return map[string]string{}
	}
	return metaInformation
}

func NewWebsite(urlFull string, links []queue.Message) Website {
	hostName, _ := url.Parse(urlFull)
	return Website{
		Url:             hostName.Hostname(),
		Links:           resolveMessageIntoLinks(links),
		ProminenceValue: 0,
	}
}

// This function gets all of the text from a soup file
func fetchText(root *soup.Root) string {
	res := ""
	for _, node := range root.FindAll("p") {
		res += node.FullText()
	}
	for _, node := range root.FindAll("h1") {
		res += node.FullText()
	}
	for _, node := range root.FindAll("h2") {
		res += node.FullText()
	}
	for _, node := range root.FindAll("div") {
		res += node.FullText()
	}
	for _, node := range root.FindAll("span") {
		res += node.FullText()
	}

	res = strings.ReplaceAll(res, "\n", "")
	res = strings.ReplaceAll(res, "\t", "")
	res = strings.ReplaceAll(res, "\r", "")
	res = strings.ReplaceAll(res, "  ", " ")
	res = strings.ReplaceAll(res, "\u00A0", " ") // Non-breaking space
	res = strings.ReplaceAll(res, "\f", "")      // Form feed
	res = strings.ReplaceAll(res, "\v", "")      // Vertical tab
	res = strings.ReplaceAll(res, "\u200B", "")  // Zero-width space
	res = strings.TrimSpace(res)

	if len(res) < MaxBodyLength {
		return res
	}
	return res[:MaxBodyLength]
}

func stripHTML(text *string) {
	re := regexp.MustCompile("<[^>]*>")
	*text = re.ReplaceAllString(*text, "")

}

func scriptEmptyLines(text *string) {
	re := regexp.MustCompile(`(?m)^\s*$[\r\n]*|<[^>]*>`)
	*text = re.ReplaceAllString(*text, "")

}

func stripJavascript(text *string) {
	re := regexp.MustCompile("(?i)<script[^>]*>(.*?)</script>")
	*text = re.ReplaceAllString(*text, "")
}

func stripCSS(text *string) {
	re := regexp.MustCompile("(?i)<style[^>]*>(.*?)</style>")
	*text = re.ReplaceAllString(*text, "")
}

func resolveMessageIntoLinks(messages []queue.Message) []string {
	var links []string
	for _, message := range messages {
		links = append(links, message.Url)
	}
	return links
}
