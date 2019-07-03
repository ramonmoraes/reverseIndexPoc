package engine

import (
	"log"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

const engineFileSystemPath = ".reverseindex"
var engineSolidStateIndex = filepath.Join(engineFileSystemPath, "config.json")

type ReverseIndex map[token][]documentPath

type Engine struct {
	ReverseIndex ReverseIndex
	Corpus Corpus
}

func CreateEngine() Engine {
	eng := Engine{
		ReverseIndex: ReverseIndex{},
		Corpus: createCorpus(),
	}
	eng.loadIndexes()
	return eng
}

func (eng Engine) loadIndexes() {
	createFolder(engineFileSystemPath)
	file, err := ioutil.ReadFile(engineSolidStateIndex)
	fmt.Println("Loading indexes...")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal([]byte(file), &eng.ReverseIndex)
	if err != nil {
		log.Fatal(err)
	}
}

func (eng Engine) DumpIndex() {
	file, err := json.MarshalIndent(eng.ReverseIndex, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	
	err = ioutil.WriteFile(engineSolidStateIndex, file, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (eng Engine) Index(input string) {
	tokens := tokenizer(input)
	document := createDocument(input, eng.Corpus)

	for _, token := range tokens {
		eng.ReverseIndex[token] = append(eng.ReverseIndex[token], document.path)
	}
}

func (eng Engine) Search(input string) []documentPath {
	tokens := tokenizer(input)

	paths := []documentPath{}
	for _, token := range(tokens) {
		paths = append(paths, eng.ReverseIndex[token]...)
	}
	return mergePaths(paths)
}

func mergePaths(paths []documentPath) []documentPath {
	return paths
}
