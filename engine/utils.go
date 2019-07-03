package engine

import (
	"os"
)

func createFolder(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
    os.Mkdir(path, os.ModePerm)
	}
}