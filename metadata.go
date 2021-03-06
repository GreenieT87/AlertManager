package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type metaData struct {
	Version   int
	Alertname string
	Groupname string
	ModTime   string
	Created   string
	ConfDocID int
}

func getMetaData(filename string) (*metaData, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			Logger("ERROR", "no Meta "+filename+" file")
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
func (m metaData) getVersion(metapath string) (version int) {
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
func (m metaData) getCreated(metapath string) (created string) {
	md, err := getMetaData(metapath)
	if err != nil {
		log.Fatal(err)
	}
	return md.Created
}
func (m metaData) getConfDocID(metapath string) (confdocid int) {
	md, err := getMetaData(metapath)
	if err != nil {
		log.Fatal(err)
	}
	return md.ConfDocID
}

// updateVersion can be used to set a version number
// or by giving `0` as version it autoincrements
func (m metaData) updateVersion(filename string, version int) {
	if version == 0 {
		version = meta.getVersion(filename) + 1
	}
	data := metaData{
		Version:   version,
		Alertname: meta.getAlertname(filename),
		Groupname: meta.getAlertGroupName(filename),
		ModTime:   meta.getModTime(filename),
		Created:   meta.getCreated(filename),
		ConfDocID: meta.getConfDocID(filename),
	}
	file, _ := yaml.Marshal(data)
	_ = ioutil.WriteFile(filename, file, 0666)
}
func (m metaData) setAlertName(filename string, alertname string) {
	data := metaData{
		Version:   meta.getVersion(filename),
		Alertname: alertname,
		Groupname: meta.getAlertGroupName(filename),
		ModTime:   meta.getModTime(filename),
		Created:   meta.getCreated(filename),
		ConfDocID: meta.getConfDocID(filename),
	}
	file, _ := yaml.Marshal(data)
	_ = ioutil.WriteFile(filename, file, 0666)
}
func (m metaData) setAlertGroupName(filename string, alertgroupname string) {
	data := metaData{
		Version:   meta.getVersion(filename),
		Alertname: meta.getAlertname(filename),
		Groupname: alertgroupname,
		ModTime:   meta.getModTime(filename),
		Created:   meta.getCreated(filename),
		ConfDocID: meta.getConfDocID(filename),
	}
	file, _ := yaml.Marshal(data)
	_ = ioutil.WriteFile(filename, file, 0666)
}
func (m metaData) setModTime(filename string) {
	data := metaData{
		Version:   meta.getVersion(filename),
		Alertname: meta.getAlertname(filename),
		Groupname: meta.getAlertGroupName(filename),
		ModTime:   now,
		Created:   meta.getCreated(filename),
		ConfDocID: meta.getConfDocID(filename),
	}
	file, _ := yaml.Marshal(data)
	_ = ioutil.WriteFile(filename, file, 0666)
}
func (m metaData) setCreated(filename string) {
	data := metaData{
		Version:   meta.getVersion(filename),
		Alertname: meta.getAlertname(filename),
		Groupname: meta.getAlertGroupName(filename),
		ModTime:   now,
		Created:   now,
		ConfDocID: meta.getConfDocID(filename),
	}
	file, _ := yaml.Marshal(data)
	_ = ioutil.WriteFile(filename, file, 0666)
}
func (m metaData) setConfDocID(filename string, confdocid int) {
	data := metaData{
		Version:   meta.getVersion(filename),
		Alertname: meta.getAlertname(filename),
		Groupname: meta.getAlertGroupName(filename),
		ModTime:   time.Now().Format("2006-01-02 15:04:05"),
		Created:   meta.getCreated(filename),
		ConfDocID: confdocid,
	}
	file, _ := yaml.Marshal(data)
	_ = ioutil.WriteFile(filename, file, 0666)
}
