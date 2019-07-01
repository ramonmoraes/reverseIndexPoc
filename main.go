package main

import (
	"fmt"
	"strings"
)

type reverseIndex map[index][]*document

type document struct {
	value string
}

type index struct {
	key string
}

func main() {
	db := make(reverseIndex)
	input := "Hello darkness my old friend"
	
	inputDocument := document{value: input}
	inputIndex := getIndexesFromString(inputDocument)

	db.saveIndex(inputIndex, inputDocument)
	fmt.Println(db)
}

func (db reverseIndex) saveIndex(indexes []index, doc document) {
	for _, index := range indexes {
		db[index] = append(db[index], &doc)
	}
}

func getIndexesFromString(input document) []index {
	output := []index{}
	spacelessInput := strings.Split(input.value, " ")
	for _, word := range spacelessInput {
		if len(word) > 3 {
			newIndex := index{key: word}
			output = append(output, newIndex)
		}
	}

	return output
}
