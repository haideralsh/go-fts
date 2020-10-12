package main

import (
	"encoding/json"
	"io/ioutil"
)

// Using an object instead of just the text string in order to add other properties if needed
type document struct {
	ID   int
	Text string `json:"title"`
}

func loadDocuments(path string) ([]document, error) {
	// Open the file
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Inline struct for an object containing array of documents
	dump := struct {
		Documents []document
	}{}

	if err := json.Unmarshal(f, &dump.Documents); err != nil {
		return nil, err
	}

	for i := range dump.Documents {
		// change their IDs to their index
		dump.Documents[i].ID = i
	}

	return dump.Documents, nil
}
