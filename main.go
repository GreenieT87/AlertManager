package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

var basepath string = "./alerts/"

type alertGroup struct {
	alertGroupName string
	AlertGroupPath []os.FileInfo
}

type alert struct {
	alertName      string
	alertPath      []os.FileInfo
	alertGroupInfo alertGroup
}

// Logger provides a simple log interface
func Logger(level string, message string) {
	if level == "ERROR" || level == "error" {
		log.Fatalf("%v %v", level, message)
	}
	log.Printf("%v %v", level, message)
}

func printAlertGroups() {
	for _, a := range getAlertGroups() {
		if a.IsDir() {
			fmt.Println(path.Base(a.Name()))
			Logger("INFO", a.Name())
			fmt.Println(a.Name())
		}

	}
}

func getAlertGroups(groups []os.FileInfo) {
	groups, err := ioutil.ReadDir(basepath)
	if err != nil {
		log.Fatal(err)
	}
	if len(groups) == 0 {
		Logger("ERROR", "no alert folders found")
	}
	println(groups)
	//ag.alertGroupName = string(groups)
	//return ag.alertGroupName
}

func printAlerts() {
	for _, a := range getAlerts() {
		if a.IsDir() {
			Logger("INFO", a.Name())
			fmt.Println(path.Base(a.Name()))

		}
	}
}

func getAlerts() (alerts []os.FileInfo) {
	for _, g := range getAlertGroups() {
		if g.IsDir() {
			groupPath := path.Join(basepath, g.Name())
			println(groupPath)
			alerts, err := ioutil.ReadDir(groupPath)
			for _, A := range alerts {
				if A.IsDir() {
					if A.Name() == groupPath {
						println(path.Join(groupPath, string(a.Name())))
					}
				}
			}
			if err != nil {
				log.Fatal(err)
			}
			if len(alerts) == 0 {
				Logger("INFO", "no alert folders found")

			}

		}

	}
	return alerts
}

func createTemplate() {
	basepath = "./alerts/general/template"
	_, err := os.Stat(basepath)
	if _, err := os.Stat(basepath); !os.IsNotExist(err) {
		Logger("ERROR", "Template already exists at "+basepath)
	}
	if err != nil {
		println(err)
		if os.IsExist(err) {
			Logger("ERROR", basepath+" already exists/nexitting...")
		}
		if os.IsNotExist(err) {
			os.MkdirAll(basepath, 0700)
			Logger("INFO", basepath+" created.")
		}
		os.Create(basepath + "/rule.yaml")
		Logger("INFO", "rulefile created.")
		os.Create(basepath + "/readme.md")
		Logger("INFO", "readme created.")
		os.Create(basepath + "/confluence.json")
		Logger("INFO", "confluence page created.")
		if !os.IsNotExist(err) {
			Logger("ERROR", "Something is really wrong")
		}
	}
	Logger("INFO", "Template created at "+basepath)
}

func validateFSstucture() {
	_, err := os.Stat(basepath)
	if err != nil {
		if os.IsNotExist(err) {
			Logger("WARN", "Main Alerts directory does not exist!")
			fmt.Println("Want to create it? [y/(n)]")
			var b []byte = make([]byte, 1)
			os.Stdin.Read(b)
			if string(b) == "y" {
				os.MkdirAll(basepath, 0700)
				validateFSstucture()
			}
			Logger("INFO", "exiting...")
			return
		}
		if !os.IsNotExist(err) {
			Logger("ERROR", "Something is really wrong")
		}
	}
	groups, err := ioutil.ReadDir(basepath)
	if err != nil {
		log.Fatal(err)
	}
	if len(groups) == 0 {
		Logger("INFO", "no alert folders found")
	}
	for _, a := range groups {
		if a.IsDir() {
			Logger("INFO", a.Name())
		}
	}
}

func main() {
	//createTemplate()
	getAlertGroups()
	//	validateFSstucture()
}
