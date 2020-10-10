package main

import (
   "encoding/xml"
   "log"
   "os"
   "regexp"
   "strings"
   "time"
   "unicode"
)

type document struct {
   Title string `xml:"title"`
   URL   string `xml:"url"`
   Text  string `xml:"abstract"`
   ID    int
}

func loadDocuments(path string) ([]document, error) {
   // Open the file
   f, err := os.Open(path)

   if err != nil {
      return nil, err
   }
   defer f.Close()

   // Create a new decoder
   dec := xml.NewDecoder(f)

   // Inline struct for an object containing array of documents
   dump := struct {
      Documents []document `xml:"doc"`
   }{}

   // Decode the object
   if err := dec.Decode(&dump); err != nil {
      return nil, err
   }

   for i := range dump.Documents {
      // change their IDs to their index
      dump.Documents[i].ID = i
   }
   return dump.Documents, nil
}

// To build an inverted index
// Take raw text -> tokenize -> normalize & filter -> search in tokens
func tokenize(text string) []string {
   // takes a string and split it according to the rules inside the anonymous
   // function
   return strings.FieldsFunc(text, func(r rune) bool {
      // Split on any character that is not a letter or a number.
      return !unicode.IsLetter(r) && !unicode.IsNumber(r)
   })
}

// Make everything lowercase
func lowercaseFilter(tokens []string) []string {
   // new slice of strings using the length of all tokens
   r := make([]string, len(tokens))

   // take each token and make it lowercase and send it to the save it in the
   // temp array
   for i, token := range tokens {
      r[i] = strings.ToLower(token)
   }
   return r
}

// Remove stop words (commonly used words)
// basically a map of just keys, i.e a set. Since Go doesn't have built-in sets.
var stopwords = map[string]struct{}{
   "a": {}, "and": {}, "be": {}, "have": {}, "i": {},
   "in": {}, "of": {}, "that": {}, "the": {}, "to": {},
}

func stopwordFilter(tokens []string) []string {
   // Same thing, init an empty slice with the length of the tokens
   r := make([]string, 0, len(tokens))

   for _, token := range tokens {
      // in Go when you do myMap["key"] it returns two things
      // the value for that key, and a boolean whether it was
      // found or not.
      if _, ok := stopwords[token]; !ok {
         r = append(r, token)
      }
      // can be written as:
      //_, found := stopwords[token]
      //if !found {
      // r = append(r, token)
      //}
   }
   return r
}

func search(docs []document, term string) []document {
   // expression to match (?i) -> case insensitive search \b -> matches the
   // word boundary meaning if we search for `cat`, then `category` is not
   // matched.
   re := regexp.MustCompile(`(?i)\b` + term + `\b`)

   // documents to return
   var r []document

   // we loop through each document to see if the `Text` field contains the
   // term
   for _, doc := range docs {
      if re.MatchString(doc.Text) {
         r = append(r, doc)
      }
   }
   return r
}

func analyze(text string) []string {
   tokens := tokenize(text)
   tokens = lowercaseFilter(tokens)
   tokens = stopwordFilter(tokens)

   return tokens
}

// the index type is for an object with the structure:
// idx := { "apple": [1, 2], "ball": [2, 3], "cat": [1] }
type index map[string][]int

func (idx index) add(docs []document) {
   // loop over the document (which we decoded from the JSON and added to them
   // their indexes as IDs)
   for _, doc := range docs {
      // loop over the each token that is got normalized and filtered
      for _, token := range analyze(doc.Text) {
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

func (idx index) search(text string) [][]int {
   var result [][]int

   // Filter the the text (lowercase, etc...) before searching
   // then loop over each word
   for _, token := range analyze(text) {
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

func main() {
   start := time.Now()

   docs, err := loadDocuments("path/to/file");
   if err != nil {
      log.Fatal(err)
   }

   log.Printf("loaded %d documents in %v", docs, time.Since(start))
}
