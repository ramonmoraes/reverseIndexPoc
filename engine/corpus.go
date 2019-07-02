package engine

import (
	"path/filepath"
)

const engineDocumentsPath = filepath.Join(engineFileSystemPath, "documents")
type Corpus map[documentPath]document

func (cor *Corpus) saveDocument(doc document) error {
	fmt.Println(engineDocumentsPath)
	return nil
}

func (cor *Corpus) loadDocument(path documentPath) (document, error) {
	return document{}, nil
}

