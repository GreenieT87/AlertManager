package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

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
