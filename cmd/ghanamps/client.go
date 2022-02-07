package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const baseURL = "https://ghanamps.herokuapp.com"

var client = &http.Client{Timeout: time.Second * 10}

type serverResponse struct {
	Ok   bool
	Msg  string
	Data interface{}
}

func prepareRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	req.Header.Set("User-Agent", userAgent)
	return req, err
}

func handleGetMembers(party string) {
	url := baseURL + "/members"
	if party != "" {
		url = baseURL + "/members/" + party
	}

	req, err := prepareRequest("GET", url, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", command, err)
		os.Exit(1)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: error fetching members: %s\n", command, err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	payload := serverResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		fmt.Fprintf(os.Stderr, "%s: request failed: %s\n", command, err)
		os.Exit(1)
	}

	if !payload.Ok {
		fmt.Fprintf(os.Stderr, "%s: request failed: %s\n", command, payload.Msg)
		os.Exit(1)
	}

	if len(payload.Data.([]interface{})) == 0 {
		fmt.Fprintf(os.Stdout, "%s: no members found for the given party: "+party, command)
		os.Exit(0)
	}

	data, err := json.MarshalIndent(payload.Data, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", command, err)
		os.Exit(1)
	}
	fmt.Println(string(data))
}

func handleGetLeadership() {
	url := baseURL + "/leadership"
	req, err := prepareRequest("GET", url, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", command, err)
		os.Exit(1)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: error fetching leaders: %s\n", command, err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	payload := serverResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		fmt.Fprintf(os.Stderr, "%s: request failed: %s\n", command, err)
		os.Exit(1)
	}

	if !payload.Ok {
		fmt.Fprintf(os.Stderr, "%s: request failed: %s\n", command, payload.Msg)
		os.Exit(1)
	}

	data, err := json.MarshalIndent(payload.Data, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", command, err)
		os.Exit(1)
	}
	fmt.Println(string(data))
}
