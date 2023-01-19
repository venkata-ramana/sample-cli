package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Process Started")

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	currPath := filepath.Dir(ex)

	os.Rename(filepath.Join(currPath, "sample-cli"), filepath.Join(currPath, "new-sample-cli"))
	fmt.Println("File renamed successfully")
	fmt.Println("Process Ended")
}
