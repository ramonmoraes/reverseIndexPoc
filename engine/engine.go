package engine

const engineFileSystemPath = ".reverseindex"

type ReverseIndex map[token][]documentPath

type Engine struct {
	ReverseIndex ReverseIndex
	Corpus Corpus
}

func CreateEngine() Engine {
	eng := Engine{
		ReverseIndex: make(ReverseIndex),
		Corpus: createCorpus(),
	}
	eng.loadFS()
	return eng
}

func (eng Engine) loadFS() {
	createFolder(engineFileSystemPath)
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