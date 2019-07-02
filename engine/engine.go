package engine

const engineFileSystemPath = "./\\.reverseindex"

type ReverseIndex map[token][]documentPath
type Corpus map[documentPath]document

type Engine struct {
	ReverseIndex ReverseIndex
	Corpus Corpus
}

func (eng Engine) index(input string) {
	tokens := tokenizer(input)
	document := createDocument(input)

	for _, token := range tokens {
		eng.ReverseIndex[token] = append(eng.ReverseIndex[token], document.path)
	}
}

func (eng Engine) search(input string) []documentPath {
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