package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("start single kanka fetch")

	kankaApiToken := os.Getenv("KANKA_API_TOKEN")
	if kankaApiToken == "" {
		log.Println("no KANKA_API_TOKEN provided, stopping execution.")
		return
	}

	kankaCampaignId := os.Getenv("KANKA_CAMPAIGN_ID")
	if kankaCampaignId == "" {
		log.Println("no KANKA_CAMPAIGN_ID provided, stopping execution.")
		return
	}
	log.Printf("using campiagn id: %s \n", kankaCampaignId)

	requestUrl := fmt.Sprintf("https://kanka.io/api/1.0/campaigns/%s", kankaCampaignId)
	requestUrl += "/entities"

	body, err := fetchAndReadBody(requestUrl, kankaApiToken)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("")
	fmt.Println(string(body))
	fmt.Println("")

	log.Println("finish")
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
