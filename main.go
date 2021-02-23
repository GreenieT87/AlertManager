package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	"gopkg.in/yaml.v2"
)

// TODO: put path vars in config.yml
var configpath string = "./config.yml"
var logpath string = "./STUFF/"
var basepath string = "./alerts/"
var meta metaData
var config conf
var now string = time.Now().Format("2006-01-02 15:04:05")

// MetaData holds meta informatiob about an alert. For internal processing only. Data gets persisted in .meta.yml
type metaData struct {
	Version   int64
	Alertname string
	Groupname string
	ModTime   string
}

func getMetaData(filename string) (*metaData, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			Logger("ERROR", "no Meta file")
		}
		return nil, err
	}
	c := &metaData{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}
	return c, nil
}
func (m metaData) getVersion(metapath string) (version int64) {
	md, err := getMetaData(metapath)
	if err != nil {
		log.Fatal(err)
	}
	return md.Version
}
func (m metaData) getAlertname(metapath string) (alertname string) {
	md, err := getMetaData(metapath)
	if err != nil {
		log.Fatal(err)
	}
	return md.Alertname
}
func (m metaData) getAlertGroupName(metapath string) (alertgroupname string) {
	md, err := getMetaData(metapath)
	if err != nil {
		log.Fatal(err)
	}
	return md.Groupname
}
func (m metaData) getModTime(metapath string) (modtime string) {
	md, err := getMetaData(metapath)
	if err != nil {
		log.Fatal(err)
	}
	return md.ModTime
}

// updateVersion can be used to set a version number
// or by giving `0` as version it autoincrements
func (m metaData) updateVersion(filename string, version int64) {
	if version == 0 {
		version = meta.getVersion(filename) + 1
	}
	data := metaData{
		Version:   version,
		Alertname: meta.getAlertname(filename),
		Groupname: meta.getAlertGroupName(filename),
		ModTime:   now,
	}
	file, _ := yaml.Marshal(data)
	_ = ioutil.WriteFile(filename, file, 0666)
}
func (m metaData) setAlertName(filename string, alertname string) {
	data := metaData{
		Version:   meta.getVersion(filename),
		Alertname: alertname,
		Groupname: meta.getAlertGroupName(filename),
		ModTime:   "",
	}
	file, _ := yaml.Marshal(data)
	_ = ioutil.WriteFile(filename, file, 0666)
}
func (m metaData) setAlertGroupName(filename string, alertgroupname string) {
	data := metaData{
		Version:   meta.getVersion(filename),
		Alertname: "",
		Groupname: "",
		ModTime:   "",
	}
	data.Version = meta.getVersion(filename)
	data.Alertname = meta.getAlertname(filename)
	data.Groupname = alertgroupname
	file, _ := yaml.Marshal(data)
	_ = ioutil.WriteFile(filename, file, 0666)
}

func (m metaData) setModTime(filename string) {
	data := metaData{
		Version:   meta.getVersion(filename),
		Alertname: meta.getAlertname(filename),
		Groupname: meta.getAlertGroupName(filename),
		ModTime:   now,
	}
	file, _ := yaml.Marshal(data)
	_ = ioutil.WriteFile(filename, file, 0666)
}

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

// Logger provides a simple log interface. It logs to file `log.log`,
// on ERROR it exits after printing the provided message
func Logger(level string, message string) {
	var logline string = now + " " + level + " " + message
	fmt.Println(logline)
	if level == "ERROR" || level == "error" {
		os.Exit(1)
	}
	f, err := os.OpenFile(logpath+"log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	if _, err := f.WriteString(logline + "\n"); err != nil {
		log.Fatal(err)
	}
}

func init() {
	_, err := os.Stat(configpath)
	if os.IsNotExist(err) {
		fmt.Println("Config file missing")
		os.Exit(1)
	}
	fmt.Println("All good")
}

func main() {
	groups, _ := ioutil.ReadDir("./alerts")
	for _, g := range groups {
		if g.IsDir() {
			alerts, _ := ioutil.ReadDir(path.Join("./alerts", g.Name()))
			for _, a := range alerts {
				if a.IsDir() {
					fmt.Println(path.Join(g.Name(), a.Name()))
				}
			}
		}
	}

	// Logger("ERROR", "BLA")
	fmt.Println(meta.getVersion("STUFF/meta.yml"))
	// fmt.Println(config.getConfluenceAPIKey())
	// Logger("INFO", "TEST")
	// meta.updateVersion("STUFF/meta.yml", 1)
	// fmt.Println(meta.getVersion("STUFF/meta.yml"))
	file, _ := os.Stat("STUFF/meta.yml")
	meta.setModTime("STUFF/meta.yml")
	if meta.getModTime("STUFF/meta.yml") < file.ModTime().Format("2006-01-02 15:04:05") {
		Logger("WARN", "Manual changes of meta file detected")
	}
	fmt.Println(file.ModTime().Format("2006-01-02 15:04:05"))
}
