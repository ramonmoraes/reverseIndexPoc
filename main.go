package main

import (
	"fmt"
	engine "github.com/ramonmoraes/reverseIndexPoc/engine"
)

func main() {
	eng := engine.CreateEngine()
	input := "Hello darkness my old friend"
	eng.Index(input)
	fmt.Println("Dumping engine index")
	defer eng.DumpIndex()
}
