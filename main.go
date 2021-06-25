package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type kankaResource struct {
	name                    string
	subEndpoint             string
	metaInfoResourceDecoder func(body string) (MetaInfo, error)
}

var kResources = []kankaResource{
	{
		"characters", "/characters",
		func(body string) (MetaInfo, error) {
			chars := Characters{}
			err := json.NewDecoder(strings.NewReader(string(body))).Decode(&chars)
			if err != nil {
				return Characters{}, fmt.Errorf("couldn't json decode body: %s", err)
			}

			return chars, nil
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

	for _, kr := range kResources {
		exportDirForResource := exportDir + "/" + kr.name

		err = os.MkdirAll(exportDirForResource, 0755)
		if err != nil {
			log.Fatalf("couldn't create directory: %s", err)
		}

		requestUrl := fmt.Sprintf("https://kanka.io/api/1.0/campaigns/%s", kankaCampaignId)
		requestUrl += kr.subEndpoint
		requestUrlPageOne := requestUrl + "?page=1"

		body, err := fetchAndReadBody(requestUrlPageOne, kankaApiToken)
		if err != nil {
			log.Fatalf("couldn't fetch and read body: %s", err)
		}

		exportFile := exportDirForResource + "/page-1"

		err = ioutil.WriteFile(exportFile, body, 0644)
		if err != nil {
			log.Fatalf("couldn't write to file: %s", err)
		}

		metaInfoResource, err := kr.metaInfoResourceDecoder(string(body))
		if err != nil {
			log.Fatalf("couldn't json decode body: %s", err)
			return
		}

		log.Println("metaInfoResource.PaginationLinks: ", metaInfoResource.PaginationLinks())
		log.Println("metaInfoResource.MetaInfo: ", metaInfoResource.MetaInfo())

		lastPage := metaInfoResource.MetaInfo().LastPage
		if lastPage == 1 {
			continue
		}
		for i := 2; i <= metaInfoResource.MetaInfo().LastPage; i++ {
			requestUrlForPage := requestUrl + "?page=" + strconv.Itoa(i)
			body, err := fetchAndReadBody(requestUrlForPage, kankaApiToken)
			if err != nil {
				log.Fatalf("couldn't fetch and read body: %s", err)
			}

			exportFile := exportDirForResource + "/page-" + strconv.Itoa(i)

			err = ioutil.WriteFile(exportFile, body, 0644)
			if err != nil {
				log.Fatalf("couldn't write to file: %s", err)
			}
		}
	}

	log.Println("finishing kanka exporter")
}

func fetchAndReadBody(requestUrl string, kankaApiToken string) ([]byte, error) {
	client := &http.Client{
		// CheckRedirect: redirectPolicyFunc,
	}

	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return []byte{}, fmt.Errorf("couldn't create new http request: %s", err)
	}

	var bearer = "Bearer " + kankaApiToken

	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, fmt.Errorf("couldn't do http request: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("couldn't read body: %s", err)

	}

	return body, nil
}
