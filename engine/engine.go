package engine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
)

const engineFileSystemPath = ".reverseindex"

var engineSolidStateIndex = filepath.Join(engineFileSystemPath, "config.json")

type ReverseIndex map[token][]string

type Engine struct {
	ReverseIndex ReverseIndex
	Corpus       Corpus
}

func CreateEngine() Engine {
	eng := Engine{
		ReverseIndex: ReverseIndex{},
		Corpus:       createCorpus(),
	}
	eng.loadIndexes()
	return eng
}

func (eng Engine) loadIndexes() {
	createFolder(engineFileSystemPath)
	file, err := ioutil.ReadFile(engineSolidStateIndex)
	fmt.Println("Loading indexes...")

	if os.IsNotExist(err) {
		fmt.Println("Indexes not found")
		return
	} else if err != nil {
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
		sort.Strings(eng.ReverseIndex[token])
		if sort.SearchStrings(eng.ReverseIndex[token], string(token)) == 0 {
			eng.ReverseIndex[token] = append(eng.ReverseIndex[token], document.path)
		}
	}
}

func (eng Engine) Search(input string) []string {
	tokens := tokenizer(input)

	paths := []string{}
	for _, token := range tokens {
		paths = append(paths, eng.ReverseIndex[token]...)
	}
	return mergePaths(paths)
}

func mergePaths(paths []string) []string {
	return paths
}
