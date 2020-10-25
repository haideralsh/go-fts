package main

import (
	"log"
	"regexp"
	"time"
)

func search(docs []document, term string) []document {
	// expression to match (?i) -> case insensitive search \b -> matches the
	// word boundary meaning if we search for `cat`, then `category` is not
	// matched.
	re := regexp.MustCompile(`(?i)\b` + term + `\b`)

	// Documents to return
	var r []document

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

	docs, err := loadDocuments("data/example.json")

	if err != nil {
		log.Fatal("An error occured while loading documents", err)
	}

	log.Printf("Loaded %d document(s) in %v", len(docs), time.Since(start))

	start = time.Now()

	idx := make(index)
	idx.add(docs)

	log.Printf("Indexed %d document(s) in %v", len(docs), time.Since(start))

	ids := idx.indexOf("the one you are")

	// Printing the value for now since go doesn't allow unused variables
	log.Print(ids)

}
