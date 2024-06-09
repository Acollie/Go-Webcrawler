package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"net/url"
)

type IgnoreList struct {
	Website    map[string]bool
	Relational bool
}
type rawIgnoreList struct {
	Websites   []string `yaml:"websites"`
	Relational bool     `yaml:"relational"`
}

func Fetch() *IgnoreList {
	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	raw := &rawIgnoreList{}
	err = yaml.Unmarshal(yamlFile, raw)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	conf := &IgnoreList{
		Website:    make(map[string]bool),
		Relational: raw.Relational,
	}

	for _, website := range raw.Websites {
		// Extract host from the url
		parsedUrl, err := url.Parse(website)
		if err != nil {
			log.Printf("Failed to parse URL from yaml file: %v, err: %v", website, err)
			continue
		}
		conf.Website[parsedUrl.Host] = true
	}
	return conf
}

func (i *IgnoreList) Ignore(urlStr string) bool {

	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		log.Printf("Url parse error %v ", err)
		return false
	}

	_, isIgnored := i.Website[parsedURL.Host]
	return isIgnored
}
