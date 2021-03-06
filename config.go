package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type conf struct {
	Conflunece struct {
		ConfluenceAPIKey   string `yaml:"confluence_api_key,omitempty"`
		ConfluenceSpaceKey string `yaml:"confluence_space_key,omitempty"`
		ConfluenceDomain   string `yaml:"confluence_domain,omitempty"`
		ModTime            string `yaml:"mod_time,omitempty"`
	} `yaml:"conflunece,omitempty"`
	Grafana struct {
		GrafanaDomain string `yaml:"grafana_domain,omitempty"`
		GrafanaAPIKey string `yaml:"grafama_api_key,omitempty"`
	} `yaml:"grafana,omitempty"`
}

func getConf(filename string) (*conf, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &conf{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}

	return c, nil
}
func (c conf) getConfluenceAPIKey() (confluenceapikey string) {
	cd, err := getConf(configpath)
	if err != nil {
		log.Fatal(err)
	}
	// if cd.Conflunece.ConfluenceAPIKey == "" {
	// Logger("ERROR", "no api key found in config")
	// }
	return cd.Conflunece.ConfluenceAPIKey
}
func (c conf) getConfluenceSpaceKey() (confluencespacekey string) {
	cd, err := getConf(configpath)
	if err != nil {
		log.Fatal(err)
	}
	return cd.Conflunece.ConfluenceSpaceKey
}
func (c conf) getConfluenceDomain() (confluencedomain string) {
	cd, err := getConf(configpath)
	if err != nil {
		log.Fatal(err)
	}
	return cd.Conflunece.ConfluenceDomain
}
func (c conf) getGrafanaDoamin() (grafanadomain string) {
	cd, err := getConf(configpath)
	if err != nil {
		log.Fatal(err)
	}
	return cd.Grafana.GrafanaDomain
}
func (c conf) getGrafanaAPIKey() (grafanaapikey string) {
	cd, err := getConf(configpath)
	if err != nil {
		log.Fatal(err)
	}
	return cd.Grafana.GrafanaAPIKey
}
