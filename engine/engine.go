package engine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const engineFileSystemPath = ".reverseindex"

var engineSolidStateIndex = filepath.Join(engineFileSystemPath, "config.json")

type ReverseIndex map[token][]string

type Engine struct {
	ReverseIndex ReverseIndex
	Corpus       Corpus
}

func CreateEngine() Engine {
	createFolder(engineFileSystemPath)
	eng := Engine{
		ReverseIndex: ReverseIndex{},
		Corpus:       createCorpus(),
	}
	eng.loadIndexes()
	return eng
}

func (eng Engine) loadIndexes() {
	file, err := ioutil.ReadFile(engineSolidStateIndex)
	fmt.Println("[Loading indexes...]")

	if os.IsNotExist(err) {
		fmt.Println("[Indexes not found]")
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
	fmt.Println("[Indexing]")
	tokens := tokenizer(input)
	document := createDocument(input, eng.Corpus)

	for _, token := range tokens {
		documents := eng.ReverseIndex[token]

		contains := false
		for _, path := range documents {
			contains = strings.Compare(document.path, path) == 0
			if contains {
				break
			}
		}

		if !contains {
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
