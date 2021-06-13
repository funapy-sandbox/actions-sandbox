package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	token, sha := os.Getenv("GITHUB_TOKEN"), os.Getenv("GITHUB_SHA")
	if len(token) == 0 || len(sha) == 0 {
		log.Fatalf("environment is empty\ttoken: %s, sha:%s", token, sha)
	}

	fmt.Println(sha)

	d, err := json.Marshal(Request{
		Context:     "lint",
		State:       "success",
		Description: "lint passed",
	})
	if err != nil {
		log.Fatalf("failed to marshal: %s", err.Error())
	}

	log.Println(string(d))

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

	log.Printf("status code: %d\n", resp.StatusCode)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read body: %s", err.Error())
	}
	log.Println(string(b))
}
