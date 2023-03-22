package main

import "C"

import (
	"fmt"
	"os"
	"path/filepath"
	"xpl0itu/godecrypt/src/decrypt"
	"xpl0itu/godecrypt/src/extract"
)

func removeAppDecFiles() error {
	files, err := filepath.Glob("*.app.dec")
	if err != nil {
		return err
	}
	for _, file := range files {
		err = os.Remove(file)
		if err != nil {
			fmt.Printf("Failed to remove file %s: %v\n", file, err)
		}
	}
	return nil
}

func DecryptAndExtract(path *C.char) {
	os.Chdir(C.GoString(path))
	decrypt.Decrypt()
	extract.Extract()
	removeAppDecFiles()
}

func main() {}
