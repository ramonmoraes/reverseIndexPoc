package engine

import (
	"hash/fnv"
	"fmt"
)

type document struct {
	text string
	identifier string
	path documentPath
}

type documentPath string;

func createDocument(text string, corpus Corpus) document {
	doc := document{ text: text, identifier: hash(text)  }
	corpus.saveDocument(doc)
	return doc;
}

func hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return fmt.Sprint(h)
}