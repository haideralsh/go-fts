package index

import (
	"github.com/haideralsh/go-fts/pkg/document"
	"github.com/haideralsh/go-fts/pkg/process"
)

// the Index type is for an object with the structure:
// idx := { "apple": [1, 2], "ball": [2, 3], "cat": [1] }
type Index map[string][]int

func (idx Index) Add(docs []document.Document) {
	// loop over the document (which we decoded from the JSON and added to them
	// their indexes as IDs)
	for _, doc := range docs {
		// loop over the each token that is got normalized and filtered
		for _, token := range process.Process(doc.Text) {
			// Get the array for each token from the map: idx["cat"] would
			// return [1]
			ids := idx[token]

			// if the array exists and the last int in the array is not equal to
			// the id of the document because say we get the token "cat" twice
			// in the same document, we want to add the ID for
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

func (idx Index) IndexOf(text string) []int {
	var result [][]int

	// Filter the the text (lowercase, etc...) before searching
	// then loop over each word
	for _, token := range process.Process(text) {
		// Check if the index contains the word
		if ids, ok := idx[token]; ok {
			// if so add the document id to the results array
			result = append(result, ids)
		}
	}

	// Return the document ids that contain this word (token)
	return flatten(result)
}

func flatten(arr [][]int) []int {
	// Make a new hashmap that we will use to insure we only return unique ids
	// using a hashmap so that our id look up is O(1)
	flat := make(map[int]int)

	// for each array in the big array
	for _, inner := range arr {
		// For each id in the inner array
		for _, id := range inner {
			if _, ok := flat[id]; !ok {
				flat[id] = id
			}
		}
	}

	var result []int

	for _, n := range flat {
		result = append(result, n)
	}

	return result
}
