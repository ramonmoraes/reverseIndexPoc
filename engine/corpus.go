package engine

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

type Corpus map[string]document

const engineDocumentSuffix = "-dart-engine.txt"

var engineDocumentsPath = filepath.Join(engineFileSystemPath, "documents")

func createCorpus() Corpus {
	cor := make(Corpus)
	createFolder(engineDocumentsPath)
	return cor
}

func (cor *Corpus) saveDocument(doc document) string {
	docPath := filepath.Join(engineDocumentsPath, doc.identifier+engineDocumentSuffix)
	err := ioutil.WriteFile(docPath, []byte(doc.text), 0644)
	if err != nil {
		log.Fatal(err)
	}
	return docPath
}

func (cor *Corpus) loadDocument(path string) (document, error) {
	return document{}, nil
}
