package main

import (
	"log"
	"regexp"
	"time"
)

func search(docs []Document, term string) []Document {
	// expression to match (?i) -> case insensitive search \b -> matches the
	// word boundary meaning if we search for `cat`, then `category` is not
	// matched.
	re := regexp.MustCompile(`(?i)\b` + term + `\b`)

	// Documents to return
	var r []Document

	// we loop through each Document to see if the `Text` field contains the
	// term
	for _, doc := range docs {
		if re.MatchString(doc.Text) {
			r = append(r, doc)
		}
	}
	return r
}

func main() {
	start := time.Now()

	docs, err := LoadDocuments("data/example.json")
	if err != nil {
		log.Fatal("An error occured while loading documents", err)
	}

	log.Printf("loaded %d document(s) in %v", len(docs), time.Since(start))
}
