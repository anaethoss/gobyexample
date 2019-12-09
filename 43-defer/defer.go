package main

import (
	"fmt"
	"os"
	"time"
)

func createFile(p string) *os.File {
	fmt.Println("Creating file")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("Writing to file")
	time.Sleep(time.Second)
	fmt.Fprintln(f, "Writing with this data")
}

func closeFile(f *os.File) {
	fmt.Println("Closing the file")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	f := createFile("./defer.txt")
	defer closeFile(f)
	writeFile(f)
}
