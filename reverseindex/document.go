package reverseindex

import (
	"fmt"
)

type document struct {
	text string
	path documentPath
}

type documentPath string;

func createDocument(text string) document {
	doc := document{ text: text }
	doc.save()
	return doc;
}

func (doc *document) save() {
	fmt.Println(doc)
	fmt.Printf("Saving at %s\n", engineFileSystemPath)
}