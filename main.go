package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Request struct {
	Context     string `json:"context"`
	State       string `json:"state"`
	Description string `json:"description"`
}

func main() {
	token, sha := os.Getenv("GITHUB_TOKEN"), os.Getenv("SHA")
	if len(token) == 0 || len(sha) == 0 {
		log.Fatalf("environment is empty\ttoken: %s, sha:%s", token, sha)
	}

	if err := getStatusV2(token, sha); err != nil {
		log.Fatalf("failed to get status: %v", err)
	}
	fmt.Println("success getStatusV2")

	if err := getStatus(token, sha); err != nil {
		log.Fatalf("failed to get status: %v", err)
	}
	fmt.Println("success getStatus")

	if err := updateStatus(token, sha); err != nil {
		log.Fatalf("failed to update: %v", err)
	}
	fmt.Println("success updateStatus")
}

func getStatusV2(token, sha string) error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	statuses, _, err := client.Repositories.ListStatuses(ctx, "funapy-sandbox", "actions-sandbox", sha, &github.ListOptions{})
	if err != nil {
		return err
	}

	for _, status := range statuses {
		fmt.Printf("%#v\n", status)
	}

	return nil
}

func getStatus(token, sha string) error {
	// url := "https://api.github.com/repos/funapy-sandbox/actions-sandbox/commits/" + sha + "/status/"
	url := "https://api.github.com/repos/funapy-sandbox/actions-sandbox/statuses/" + sha

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", "token "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read all: %w", err)
	}

	var buf bytes.Buffer
	if err := json.Indent(&buf, b, "", "  "); err != nil {
		return err
	}
	fmt.Println(buf.String())
	return nil
}

func updateStatus(token, sha string) error {
	d, err := json.Marshal(Request{
		Context:     "lint",
		State:       "success",
		Description: "lint passed",
	})
	if err != nil {
		return fmt.Errorf("failed to marshal: %w", err)
	}
	url := "https://api.github.com/repos/funapy-sandbox/actions-sandbox/statuses/" + sha

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(d))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", "token "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read body: %w", err)
	}
	return nil
}
