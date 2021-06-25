package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type kankaResource struct {
	name                           string
	subEndpoint                    string
	paginationLinksResourceDecoder func(body string) (error, PaginationLinks)
}

var kResources = []kankaResource{
	{
		"characters", "/characters",
		func(body string) (error, PaginationLinks) {
			chars := Characters{}
			err := json.NewDecoder(strings.NewReader(string(body))).Decode(&chars)
			if err != nil {
				return fmt.Errorf("couldn't json decode body: %s", err), Characters{}
			}

			return nil, chars
		},
	},
	// {"character", "/characters/209208"}, // not needed, same infos in characters
}

func main() {
	log.Println("starting kanka exporter")

	kankaApiToken := os.Getenv("KANKA_API_TOKEN")
	if kankaApiToken == "" {
		log.Println("no KANKA_API_TOKEN provided, stopping execution.")
	}

	kankaCampaignId := os.Getenv("KANKA_CAMPAIGN_ID")
	if kankaCampaignId == "" {
		log.Println("no KANKA_CAMPAIGN_ID provided, stopping execution.")
	}
	log.Printf("using campiagn id: %s \n", kankaCampaignId)

	cwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(cwd) // for example /home/user

	exportDir := cwd + "/generated/exports"

	err = os.MkdirAll(exportDir, 0755)
	if err != nil {
		log.Fatalf("couldn't create directory: %s", err)
	}

	//fetch data
	client := &http.Client{
		// CheckRedirect: redirectPolicyFunc,
	}

	for _, kr := range kResources {
		requestUrl := fmt.Sprintf("https://kanka.io/api/1.0/campaigns/%s", kankaCampaignId)
		requestUrl += kr.subEndpoint

		req, err := http.NewRequest("GET", requestUrl, nil)
		if err != nil {
			log.Fatalf("couldn't create new http request: %s", err)
		}

		var bearer = "Bearer " + kankaApiToken

		req.Header.Add("Authorization", bearer)
		req.Header.Add("Content-type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("couldn't do http request: %s", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("couldn't read body: %s", err)
		}

		exportDirForResource := exportDir + "/" + kr.name

		err = os.MkdirAll(exportDirForResource, 0755)
		if err != nil {
			log.Fatalf("couldn't create directory: %s", err)
		}

		exportFile := exportDirForResource + "/page-1"

		err = ioutil.WriteFile(exportFile, []byte(string(body)), 0644)
		if err != nil {
			log.Fatalf("couldn't write to file: %s", err)
		}

		err, pglResource := kr.paginationLinksResourceDecoder(string(body))
		if err != nil {
			log.Fatalf("couldn't json decode body: %s", err)
			return
		}

		log.Println("klResource.KankaLinks: ", pglResource.PaginationLinks())
	}

	log.Println("finishing kanka exporter")
}
