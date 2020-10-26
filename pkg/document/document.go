package document

import (
	"encoding/json"
	"io/ioutil"
)

// Document is the smallest unit we add in the our DB (in memory index)
// We're using an object instead of just the text string in order to add other
// properties if needed
type Document struct {
	ID   int
	Text string `json:"text"`
}

type documents []Document

func Load(path string) ([]Document, error) {
	// Open the file
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	dump := documents{}

	if err := json.Unmarshal(f, &dump); err != nil {
		return nil, err
	}

	for i := range dump {
		// change their IDs to their index
		dump[i].ID = i
	}

	return dump, nil
}
