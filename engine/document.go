package engine

import (
	"crypto/sha1"
	"fmt"
)

type document struct {
	text       string
	identifier string
	path       string
}

func createDocument(text string, corpus Corpus) document {
	doc := document{text: text, identifier: hash(text)}
	doc.path = corpus.saveDocument(doc)
	return doc
}

func hash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	sum := h.Sum(nil)
	return fmt.Sprintf("%x", sum)
}
