package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/ben181231/algolia-upload/config"
)

func main() {
	conf, err := config.ReadFromEnv()
	if err != nil {
		log.Fatalf("Init Error: %s", err)
		return
	}

	log.Print("Reading data from stdin...")
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Data Reading Error: %s", err)
		return
	}

	dataSize := len(data)
	log.Printf("Got %d bytes of data.", dataSize)

	if dataSize == 0 {
		log.Print("Bye.")
		return
	}

	items := []json.RawMessage{}
	if err := json.Unmarshal(data, &items); err != nil {
		log.Fatalf("Unmarshal Data Error: %s", err)
		return
	}

	itemCount := len(items)
	log.Printf("Got %d items", itemCount)

	if itemCount == 0 {
		log.Print("Bye.")
		return
	}

	client := search.NewClient(conf.AppID, conf.AdminKey)
	index := client.InitIndex(conf.IndexName)

	if _, err := index.SaveObjects(items); err != nil {
		log.Fatalf("Error Uploading Items: %s", err)
		return
	}
}
