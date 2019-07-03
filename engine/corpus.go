package engine

import (
	"fmt"
	"path/filepath"
	"io/ioutil"
)

type Corpus map[documentPath]document
var engineDocumentsPath = filepath.Join(engineFileSystemPath, "documents")

func createCorpus() Corpus{
	cor := make(Corpus)
	createFolder(engineDocumentsPath)
	return cor;
}

func (cor *Corpus) saveDocument(doc document) error {
	fmt.Println(engineDocumentsPath)
	docPath := filepath.Join(engineDocumentsPath, doc.identifier + "-dart-engine.txt")
	doc.path = documentPath(docPath)
	return ioutil.WriteFile(docPath, []byte(doc.text), 0644)
}

func (cor *Corpus) loadDocument(path documentPath) (document, error) {
	return document{}, nil
}

