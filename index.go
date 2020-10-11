package main

import (
	"encoding/json"
	"io/ioutil"
)

// Document struct
// Using an object instead of just the text string in order to add other properties if needed
type Document struct {
	ID   int
	Text string `json:"title"`
}

// LoadDocuments ...
func LoadDocuments(path string) ([]Document, error) {
	// Open the file
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Inline struct for an object containing array of documents
	dump := struct {
		Documents []Document
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

// the index type is for an object with the structure:
// idx := { "apple": [1, 2], "ball": [2, 3], "cat": [1] }
type index map[string][]int

func (idx index) add(docs []Document) {
	// loop over the Document (which we decoded from the JSON and added to them
	// their indexes as IDs)
	for _, doc := range docs {
		// loop over the each token that is got normalized and filtered
		for _, token := range process(doc.Text) {
			// Get the array for each token from the map: idx["cat"] would
			// return [1]
			ids := idx[token]

			// if the array exists and the last int in the array is not equal to
			// the id of the Document because say we get the token "cat" twice
			// in the same Document, we want to add the ID for
			// document only once.
			if ids != nil && ids[len(ids)-1] == doc.ID {
				// Don't add same ID twice.
				continue
			}

			// append the array that is stored in the map under the "cat" key
			// the current document ID
			idx[token] = append(ids, doc.ID)
		}
	}
}

func (idx index) search(text string) [][]int {
	var result [][]int

	// Filter the the text (lowercase, etc...) before searching
	// then loop over each word
	for _, token := range process(text) {
		// Check if the index contains the word
		if ids, ok := idx[token]; ok {
			// if so add the document id to the results array
			result = append(result, ids)
		}
	}

	// Return the document ids that contain this any word (token) of the
	// passed text
	return result
}
