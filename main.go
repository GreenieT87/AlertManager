package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
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

func askForConfirmation(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}

func init() {
	Logger("####", "##################")
	Logger("INIT", "Starting AlertManager")
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
	Logger("INIT", "setup up done.")
	Logger("####", "##################")

}

func main() {
	groups, _ := ioutil.ReadDir(basepath)
	os.Rename("./rules.yml","./rules.bak")
	// creating rules.yml 
	r, err := os.OpenFile("./rules.yml", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Close()
	if _, err := r.WriteString("groups:\n"); err != nil {
		log.Fatal(err)
	}

	for _, g := range groups {
		if g.IsDir() {
			Logger("INFO", "Groupname: "+g.Name())
			alerts, _ := ioutil.ReadDir(path.Join(basepath, g.Name()))

			f, err := os.OpenFile(path.Join(basepath, g.Name())+"/tmp_rules.yml", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatalln(err)
			}
			defer f.Close()
			if _, err := f.WriteString("- name: " + g.Name() + "\n  rules:\n"); err != nil {
				log.Fatal(err)
			}

			for _, a := range alerts {
				if a.IsDir() {
					Logger("INFO", "Alertname: "+g.Name()+"/"+a.Name())
					metapath := path.Join(basepath, g.Name(), a.Name()) + "/.meta.yml"
					rulepath := path.Join(basepath, g.Name(), a.Name()) + "/rule.yaml"

					// appending alertrules to tmp group file
					if _, err := f.WriteString("  - alert: " + a.Name() + "\n"); err != nil {
						log.Fatal(err)
					}
					rule, err := ioutil.ReadFile(rulepath)
					if err != nil {
						fmt.Println(err)
						break
					}
					if _, err := f.Write(rule); err != nil {
						log.Fatal(err)
					}
					if _, err := f.WriteString("\n"); err != nil {
						log.Fatal(err)
					}

					// checking for manual metafilechanges
					file, err := os.Stat(metapath)
					if err != nil {
						log.Fatal(err)
					}
					if meta.getModTime(metapath) != "" && meta.getModTime(metapath) < file.ModTime().Format("2006-01-02 15:04") {
						Logger("WARN", "ModTime: "+meta.getModTime(metapath)+" < OS_ModTime: "+file.ModTime().Format("2006-01-02 15:04"))
						Logger("WARN", "Manual changes of meta "+metapath+" file detected")
						Logger("WARN", "Changes of metafile should only be done by this tool. Manual changes can be done with the proper precautions.")
						Logger("WARN", "Do you wanna proceed with this alert. No just skips this one. [y/n]")
						ask := askForConfirmation("")
						if !ask {
							break
						}
					}

					// update metafile
					if meta.getVersion(metapath) != conflunece.getVersionbyID(meta.getConfDocID(metapath)) ||
						meta.getAlertname(metapath) != a.Name() ||
						meta.getAlertGroupName(metapath) != g.Name() {
						Logger("INFO", "updating metafile")
						meta.updateVersion(metapath, conflunece.getVersionbyID(meta.getConfDocID(metapath)))
						meta.setAlertName(metapath, a.Name())
						meta.setAlertGroupName(metapath, g.Name())
					}

					// Write conflunece json file
					josnpath := path.Join(basepath, g.Name(), a.Name()) + "/.conflunece.json"
					ioutil.WriteFile(josnpath, []byte(conflunece.getJsonbyID(meta.getConfDocID(metapath))), 0666)
				}
			}

			// writing master rules file r
			tmpgrprule, err := ioutil.ReadFile(path.Join(basepath, g.Name()) + "/tmp_rules.yml")
			if err != nil {
				fmt.Println(err)
			}
			if _, err := r.Write(tmpgrprule); err != nil {
				log.Fatal(err)
			}

			// removing temp files
			os.Remove(path.Join(basepath, g.Name()) + "/tmp_rules.yml")
			Logger("INFO", "Group Done!")
		}
	}

	x0 := conflunece.getVersionbyID(2388721698)
	var ver = x0
	fmt.Printf("version %d", ver)
}
