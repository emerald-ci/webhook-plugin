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

type Input struct {
	Config struct {
		Url string `json:"url"`
	} `json:"config"`
}

func main() {
	inputContent, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	input := &Input{}
	err = json.Unmarshal(inputContent, &input)
	if err != nil {
		log.Fatal(err)
	}

	var rawJson map[string]*json.RawMessage
	err = json.Unmarshal(inputContent, &rawJson)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rawJson)

	postJson, err := json.Marshal(rawJson["job"])
	if err != nil {
		log.Fatal(err)
	}

	_, err = http.Post(input.Config.Url, "application/json", bytes.NewBuffer(postJson))
	if err != nil {
		log.Fatal(err)
	}
}
