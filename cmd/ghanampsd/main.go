// This is the source code for the ghanamps web server called 'ghanampsd'. It has
// one endpoint (the index) which returns a JSON array containing the details of
// all the current members of parliament.
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/yeboahnanaosei/ghanamps"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"ok": true, "msg": "Welcome to the ghanamps API. Please use the /members or /leaders endpoint to get the data"}`))
}

// membersHandler handles the members endpoint. It returns a JSON of the current
// members of parliament.
func membersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	members, err := ghanamps.Members()
	if err != nil {
		log.Println("error fetching members:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"ok": false, "msg": "Request failed. An internal error occured"}`))
		return
	}

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	enc.SetIndent(" ", "  ")

	payload := struct {
		Ok   bool              `json:"ok"`
		Msg  string            `json:"msg"`
		Data []ghanamps.Member `json:"data"`
	}{
		Ok:   true,
		Msg:  "Request successful",
		Data: members,
	}

	err = enc.Encode(payload)
	if err != nil {
		log.Println("error encoding JSON response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"ok": false, "msg": "Request failed. An internal error occured"}`))
		return
	}
}

// leadershipHandler returns handles the leadership endpoint. It returns a JSON
// of the current leaders of parliament.
func leadershipHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	leaders, err := ghanamps.Leaders()
	if err != nil {
		log.Println("error fetching leadership:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"ok": false, "msg": "Request failed. An internal error occured"}`))
		return
	}

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	enc.SetIndent(" ", "  ")

	payload := struct {
		Ok   bool              `json:"ok"`
		Msg  string            `json:"msg"`
		Data []ghanamps.Leader `json:"data"`
	}{
		Ok:   true,
		Msg:  "Request successful",
		Data: leaders,
	}

	err = enc.Encode(payload)
	if err != nil {
		log.Println("error encoding JSON response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"ok": false, "msg": "Request failed. An internal error occured"}`))
		return
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/members", membersHandler)
	http.HandleFunc("/leaders", leadershipHandler)
	log.Println("listening on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
