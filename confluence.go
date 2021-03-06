package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var cg conf
var key string = cg.getConfluenceAPIKey()
var domain string = cg.getConfluenceDomain()

// ConfluneceDoc holds the current confluence document
type ConfluneceDoc struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Status string `json:"status"`
	Title  string `json:"title"`
	Space  struct {
		ID         int    `json:"id"`
		Key        string `json:"key"`
		Name       string `json:"name"`
		Type       string `json:"type"`
		Status     string `json:"status"`
		Expandable struct {
			Settings    string `json:"settings"`
			Metadata    string `json:"metadata"`
			Operations  string `json:"operations"`
			LookAndFeel string `json:"lookAndFeel"`
			Identifiers string `json:"identifiers"`
			Permissions string `json:"permissions"`
			Icon        string `json:"icon"`
			Description string `json:"description"`
			Theme       string `json:"theme"`
			History     string `json:"history"`
			Homepage    string `json:"homepage"`
		} `json:"_expandable"`
		Links struct {
			Webui string `json:"webui"`
			Self  string `json:"self"`
		} `json:"_links"`
	} `json:"space"`
	Version struct {
		By struct {
			Type           string `json:"type"`
			AccountID      string `json:"accountId"`
			AccountType    string `json:"accountType"`
			Email          string `json:"email"`
			PublicName     string `json:"publicName"`
			ProfilePicture struct {
				Path      string `json:"path"`
				Width     int    `json:"width"`
				Height    int    `json:"height"`
				IsDefault bool   `json:"isDefault"`
			} `json:"profilePicture"`
			DisplayName            string `json:"displayName"`
			IsExternalCollaborator bool   `json:"isExternalCollaborator"`
			Expandable             struct {
				Operations    string `json:"operations"`
				PersonalSpace string `json:"personalSpace"`
			} `json:"_expandable"`
			Links struct {
				Self string `json:"self"`
			} `json:"_links"`
		} `json:"by"`
		When          time.Time `json:"when"`
		FriendlyWhen  string    `json:"friendlyWhen"`
		Message       string    `json:"message"`
		Number        int       `json:"number"`
		MinorEdit     bool      `json:"minorEdit"`
		SyncRev       string    `json:"syncRev"`
		SyncRevSource string    `json:"syncRevSource"`
		ConfRev       string    `json:"confRev"`
		Expandable    struct {
			Collaborators string `json:"collaborators"`
			Content       string `json:"content"`
		} `json:"_expandable"`
		Links struct {
			Self string `json:"self"`
		} `json:"_links"`
	} `json:"version"`
	MacroRenderedOutput struct {
	} `json:"macroRenderedOutput"`
	Body struct {
		Storage struct {
			Value           string        `json:"value"`
			Representation  string        `json:"representation"`
			EmbeddedContent []interface{} `json:"embeddedContent"`
			Expandable      struct {
				Content string `json:"content"`
			} `json:"_expandable"`
		} `json:"storage"`
		Expandable struct {
			Editor              string `json:"editor"`
			AtlasDocFormat      string `json:"atlas_doc_format"`
			View                string `json:"view"`
			ExportView          string `json:"export_view"`
			StyledView          string `json:"styled_view"`
			Dynamic             string `json:"dynamic"`
			Editor2             string `json:"editor2"`
			AnonymousExportView string `json:"anonymous_export_view"`
		} `json:"_expandable"`
	} `json:"body"`
	Extensions struct {
		Position int `json:"position"`
	} `json:"extensions"`
	Expandable struct {
		ChildTypes          string `json:"childTypes"`
		Container           string `json:"container"`
		Metadata            string `json:"metadata"`
		Operations          string `json:"operations"`
		SchedulePublishDate string `json:"schedulePublishDate"`
		Children            string `json:"children"`
		Restrictions        string `json:"restrictions"`
		History             string `json:"history"`
		Ancestors           string `json:"ancestors"`
		Descendants         string `json:"descendants"`
	} `json:"_expandable"`
	Links struct {
		Editui     string `json:"editui"`
		Webui      string `json:"webui"`
		Context    string `json:"context"`
		Self       string `json:"self"`
		Tinyui     string `json:"tinyui"`
		Collection string `json:"collection"`
		Base       string `json:"base"`
	} `json:"_links"`
}

func (c ConfluneceDoc) update() {
	// var domain = conf.getConfluenceDomain()
	// var docID = meta.getConfDocID()

	url := domain+"/wiki/rest/api/content/2388721698?expand=version.number,body.storage,space"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic dC5ncnVlbkB0YWZtb2JpbGUuZGU6dlRzTWJEdW5mSEQwMUdMVHhEQjhGNTNE")
	req.Header.Add("Cookie", "JSESSIONID=BEBEEA2BBE754C2EE6163D7E95A71B6E; atlassian.xsrf.token=B0YG-UIL8-KTON-RQPS_d1725762d46de319d4d08160cafa0cae140fa2ed_lout")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(string(body))
	var cp ConfluneceDoc
	json.Unmarshal([]byte(body), &cp)
	fmt.Printf("version: %v, Body: %s", cp.Version.Number, cp.Body.Storage.Value)
}

func (c ConfluneceDoc) getVersionbyID(ID int) (version int) {

	url := fmt.Sprintf("https://tafmobile.atlassian.net/wiki/rest/api/content/%d?expand=version.number", ID)
	method := "GET"
	// fmt.Println(url) // debug line
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+key) 
	req.Header.Add("Cookie", "JSESSIONID=BEBEEA2BBE754C2EE6163D7E95A71B6E; atlassian.xsrf.token=B0YG-UIL8-KTON-RQPS_d1725762d46de319d4d08160cafa0cae140fa2ed_lout")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(string(body)) // debug line
	var cp ConfluneceDoc
	json.Unmarshal([]byte(body), &cp)
	version = cp.Version.Number
	return version
}
