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
var conflunece ConfluneceDoc
var now string = time.Now().Format("2006-01-02 15:04:05")

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
	var conf conf
	for _, g := range groups {
		if g.IsDir() {
			fmt.Println("Group:" + g.Name())
			alerts, _ := ioutil.ReadDir(path.Join(basepath, g.Name()))
			for _, a := range alerts {
				if a.IsDir() {
					metapath := path.Join(basepath, g.Name(), a.Name()) + "/.meta.yml"
					var domain string = conf.getConfluenceDomain()
					url := fmt.Sprintf("%v/wiki/rest/api/content/%d?expand=version.number", domain, meta.getConfDocID(metapath))
					fmt.Println(url)
				}
			}
		}
	}
	// conflunece.update()

	// var domain = conf.getConfluenceDomain
	// var docID = meta.getConfDocID
	// fmt.Printf("%v/wiki/rest/api/content/%d?expand=version.number", conf.getConfluenceDomain(), meta.getConfDocID())

	// url := "https://tafmobile.atlassian.net/wiki/rest/api/content/2388721698?expand=version.number,body.storage,space"

	x0 := conflunece.getVersionbyID(2388721698)
	var ver = x0
	fmt.Printf("version %d", ver)

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
