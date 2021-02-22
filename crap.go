package crap

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var basepath string = "./alerts/"

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
