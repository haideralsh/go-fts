package main

import (
	"flag"
	"github.com/haideralsh/go-fts/pkg/document"
	"github.com/haideralsh/go-fts/pkg/index"
	"log"
	"path/filepath"
	"regexp"
	"time"
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
	f := flag.String("path", "data/example.txt", "Path for the file to index")
	q := flag.String("query", "the one you are", "Search query")
	flag.Parse()

	path, err := filepath.Abs(*f)
	if err != nil {
		log.Fatalf("Invalid path %s", f)
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

	ids := idx.IndexOf(*q)
	if len(ids) == 0 {
		log.Fatalf("No results found for %s", *q)
	}

	for _, id := range ids {
		log.Print(docs[id])
	}
}
