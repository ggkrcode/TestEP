package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/epstatus", getEPStatus).Methods("POST")
	
	
	log.Fatal(http.ListenAndServe("0.0.0.0:" + os.Getenv("PORT"), r))
}

func getEPStatus(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Received request to get device details")
	var request WebHookRequest 
	_ = json.NewDecoder(r.Body).Decode(&request)
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
	fmt.Println(err)
	var speech = ""
	var displayText = ""
	
	intentName := request.Result.Metadata.IntentName
	
	if intentName == "list" {
		speech = "New List.20"
		displayText = "New List.20"
	} else if intentName == "status" {
		speech = "All New"
		displayText = "All New"
	}

	hookResp := WebHookResp {
		speech,
		displayText,
		
	}

	json.NewEncoder(w).Encode(hookResp)
}


