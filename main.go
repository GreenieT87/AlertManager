package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

// TODO: put path vars in config.yml
var configpath string = "./config.yml"
var logpath string = "./STUFF/"
var basepath string = "./alerts/"
var meta metaData
var config conf
var now string = time.Now().Format("2006-01-02 15:04:05")

// MetaData holds meta informatiob about an alert. For internal processing only. Data gets persisted in .meta.yml

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
	_, erre := os.Stat(basepath)
	if os.IsNotExist(erre) {
		fmt.Println(basepath + " missing")
		os.Exit(1)
	}
	fmt.Println("All good")

}

func main() {
	groups, _ := ioutil.ReadDir(basepath)
	for _, g := range groups {
		if g.IsDir() {
			fmt.Println("Group:" + g.Name())
			alerts, _ := ioutil.ReadDir(path.Join(basepath, g.Name()))
			for _, a := range alerts {
				if a.IsDir() {
					metapath := path.Join(basepath, g.Name(), a.Name()) + "/.meta.yml"
					if meta.getCreated(metapath) == "" {
						meta.setCreated(metapath)
					}
					meta.setModTime(metapath)
					fmt.Println(path.Join("Alert: " + a.Name()))
					fmt.Printf("Version: %v\n", meta.getVersion(path.Join(basepath, g.Name(), a.Name())+"/.meta.yml"))
				}
			}
		}
	}

	// Logger("ERROR", "BLA")
	// fmt.Println(meta.getVersion("STUFF/meta.yml"))
	// fmt.Println(config.getConfluenceAPIKey())
	// Logger("INFO", "TEST")
	// meta.updateVersion("STUFF/meta.yml", 1)
	// fmt.Println(meta.getVersion("STUFF/meta.yml"))
	// file, _ := os.Stat("STUFF/meta.yml")
	// meta.setModTime("STUFF/meta.yml")
	// if meta.getModTime("STUFF/meta.yml") < file.ModTime().Format("2006-01-02 15:04:05") {
	// Logger("WARN", "Manual changes of meta file detected")
	// }
	// fmt.Println(file.ModTime().Format("2006-01-02 15:04:05"))
}
