# A mini, toy, full-text search engine built with go for learning purposes

> Inspired by and based on: [Let's build a Full-Text Search engine](https://artem.krylysov.com/blog/2020/07/28/lets-build-a-full-text-search-engine/)

Loads documents and processes them to create a reverse index.

Todos:

- [x] Take raw text
- [x] Tokenize
- [x] Normalize & filter
- [x] Search in tokens
- [x] Make the input JSON instead of XML
- [x] Return the documents that match instead of their indexes
- [x] Accept the file path as a command line argument
- [ ] Change the file path from an argument to a flag
- [ ] Support wildcards
- [ ] Extend boolean queries to support OR and NOT
- [ ] Sort results by relevance
