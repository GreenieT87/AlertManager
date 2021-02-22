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
	AlertGroupPath string
}

type alert struct {
	alertName      string
	alertPath      string
	alertGroupInfo alertGroup
}

// Logger provides a simple log interface
func Logger(level string, message string) {
	if level == "ERROR" || level == "error" {
		log.Fatalf("%v %v", level, message)
	}
	log.Printf("%v %v", level, message)
}

func (ag alertGroup) printAlertGroup() {
	Logger("INFO", string(ag.alertGroupName))
}

func (al alert) printAlert() {
	// Logger("INFO", "Alertname: "+al.alertName+" Path: "+al.alertPath+" Groupname: "+al.alertGroupInfo.alertGroupName)
	fmt.Printf(`
	Alertname:      %v 
	AlertPath:      %v
	AlertGroupName: %v
	AlertGroupPath: %v
	`, al.alertName, al.alertPath, al.alertGroupInfo.alertGroupName, al.alertGroupInfo.AlertGroupPath)
}

// func (agi alertGroup) getAlertGroups() (ag alert) {
// 	groups, err := ioutil.ReadDir(basepath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if len(groups) == 0 {
// 		Logger("ERROR", "no alert folders found")
// 	}
// 	for _, g := range groups {
// 		// println(g.Name())
// 		agi := alertGroup{alertGroupName: g.Name()}
// 	}
// 	return ag
// }

func printAlerts() {
	groups, err := ioutil.ReadDir(basepath)
	if err != nil {
		log.Fatal(err)
	}
	if len(groups) == 0 {
		Logger("ERROR", "no alert folders found")
	}
	for _, g := range groups {
		// println(g.Name())
		ag := &alertGroup{
			alertGroupName: g.Name(),
		}
		// ag.printAlertGroup()

		alerts, err := ioutil.ReadDir(basepath + ag.alertGroupName)
		if err != nil {
			log.Fatal(err)
		}
		if len(alerts) == 0 {
			Logger("ERROR", "no alert folders found")
		}
		for _, a := range alerts {
			// println(g.Name())
			Al := &alert{
				alertName: a.Name(),
				alertPath: path.Join(basepath, ag.alertGroupName, a.Name()),
				alertGroupInfo: alertGroup{
					alertGroupName: ag.alertGroupName,
					AlertGroupPath: path.Join(basepath, ag.alertGroupName),
				},
			}
			Al.printAlert()
		}
	}
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
	printAlerts()
	al.AlertPath
}
