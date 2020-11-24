package main

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/haideralsh/go-fts/pkg/document"
	"github.com/haideralsh/go-fts/pkg/index"
)

func search(docs []document.Document, term string) []document.Document {
	// expression to match (?i) -> case insensitive search \b -> matches the
	// word boundary meaning if we search for `cat`, then `category` is not
	// matched.
	re := regexp.MustCompile(`(?i)\b` + term + `\b`)

	// Documents to return
	var r []document.Document

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
	if len(os.Args) < 2 {
		log.Fatal("Please provide the data file path as an argument.")
	}

	r := os.Args[1]
	path, err := filepath.Abs(r)
	if err != nil {
		log.Fatalf("Invalid path %s", r)
	}

	start := time.Now()
	docs, err := document.Load(path)
	if err != nil {
		log.Fatal("An error occured while loading documents. ", err)
	}
	log.Printf("Loaded %d document(s) in %v", len(docs), time.Since(start))

	start = time.Now()
	idx := make(index.Index)
	idx.Add(docs)
	log.Printf("Indexed %d document(s) in %v", len(docs), time.Since(start))

	ids := idx.IndexOf("the one you are")
	for _, id := range ids {
		log.Print(docs[id])
	}
}
