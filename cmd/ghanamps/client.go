package main

import (
	"encoding/json"
	"fmt"
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

func handleGetMembers(party string) {
	url := baseURL + "/members"
	if party != "" {
		url = baseURL + "/members/" + party
	}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ghanamps: error fetching members: %s\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	payload := serverResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		fmt.Fprintf(os.Stderr, "ghanamps: request failed: %s\n", err)
		os.Exit(1)
	}

	if !payload.Ok {
		fmt.Fprintf(os.Stderr, "ghanamps: request failed: %s\n", payload.Msg)
		os.Exit(1)
	}

	if len(payload.Data.([]interface{})) == 0 {
		fmt.Println("ghanamps: no members found for the given party: " + party)
		os.Exit(0)
	}

	data, err := json.MarshalIndent(payload.Data, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "ghanamps: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(string(data))
}

func handleGetLeadership() {
	url := baseURL + "/leadership"

	resp, err := client.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ghanamps: error fetching leaders: %s\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	payload := serverResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		fmt.Fprintf(os.Stderr, "ghanamps: request failed: %s\n", err)
		os.Exit(1)
	}

	if !payload.Ok {
		fmt.Fprintf(os.Stderr, "ghanamps: request failed: %s\n", payload.Msg)
		os.Exit(1)
	}

	data, err := json.MarshalIndent(payload.Data, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "ghanamps: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(string(data))
}
