package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Request struct {
	Context     string `json:"context"`
	State       string `json:"state"`
	Description string `json:"description"`
}

func main() {
	//
	token, sha := os.Getenv("GITHUB_TOKEN"), os.Getenv("SHA")
	if len(token) == 0 || len(sha) == 0 {
		log.Fatalf("environment is empty\ttoken: %s, sha:%s", token, sha)
	}

	d, err := json.Marshal(Request{
		Context:     "lint",
		State:       "success",
		Description: "lint passed",
	})
	if err != nil {
		log.Fatalf("failed to marshal: %s", err.Error())
	}

	url := "https://api.github.com/repos/funapy-sandbox/actions-sandbox/statuses/" + sha

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(d))
	if err != nil {
		log.Fatalf("failed to create request: %s", err.Error())
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", "token "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("failed to send request: %s", err.Error())
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read body: %s", err.Error())
	}
}
