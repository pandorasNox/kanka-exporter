package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

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

	//fetch data
	client := &http.Client{
		// CheckRedirect: redirectPolicyFunc,
	}

	requestUrl := fmt.Sprintf("https://kanka.io/api/1.0/campaigns/%s/characters?page=1", kankaCampaignId)
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("couldn't read body: %s", err)
	}
	_ = body

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

	exportFile := exportDir + "/characters-page-1"

	err = ioutil.WriteFile(exportFile, []byte(string(body)), 0644)
	if err != nil {
		log.Fatalf("couldn't write to file: %s", err)
	}

	log.Println("finishing kanka exporter")
}
