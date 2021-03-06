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

	for _, g := range groups {
		if g.IsDir() {
			fmt.Println("Group:" + g.Name())
			alerts, _ := ioutil.ReadDir(path.Join(basepath, g.Name()))

			for _, a := range alerts {
				if a.IsDir() {
					fmt.Println(a.Name())
					metapath := path.Join(basepath, g.Name(), a.Name()) + "/.meta.yml"
					// var confDocID = meta.getConfDocID(metapath)
					file, _ := os.Stat(metapath)
					if meta.getModTime(metapath) < file.ModTime().Format("2006-01-02 15:04:05") {
						Logger("WARN", "Manual changes of meta "+metapath+" file detected")
						Logger("WARN", meta.getModTime(metapath)+" < "+file.ModTime().Format("2006-01-02 15:04:05"))

						break
					}
					Logger("WARN", meta.getModTime(metapath)+" < "+file.ModTime().Format("2006-01-02 15:04:05"))
					
					if meta.getVersion(metapath) != conflunece.getVersionbyID(meta.getConfDocID(metapath)) {
						meta.updateVersion(metapath, conflunece.getVersionbyID(meta.getConfDocID(metapath)))
						meta.setModTime(metapath)
					}
					if meta.getAlertname(metapath) != a.Name() {
						meta.setAlertName(metapath, a.Name())
						meta.setModTime(metapath)
					}
					if meta.getAlertGroupName(metapath) != g.Name() {
						meta.setAlertGroupName(metapath, g.Name())
						meta.setModTime(metapath)
					}
					
					fmt.Println(conflunece.getBodybyID(meta.getConfDocID(metapath)))
					josnpath := path.Join(basepath, g.Name(), a.Name()) + "/.conflunece.json"
					ioutil.WriteFile(josnpath, []byte(conflunece.getJsonbyID(meta.getConfDocID(metapath))), 0777)
					
				}
			}
		}
	}

	x0 := conflunece.getVersionbyID(2388721698)
	var ver = x0
	fmt.Printf("version %d", ver)
}
