package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// cleanFile accepts a string representing a path
// and converts \r\n into \n
func cleanFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	wFile, err := os.Create(fileName + ".clean")
	if err != nil {
		panic(err)
	}
	defer wFile.Close()
	w := bufio.NewWriter(wFile)

	data := make([]byte, 128)
	for {
		data = data[:cap(data)]
		n, err := file.Read(data)
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		data = data[:n]
		var indexesToDelete []int
		for i, b := range data {
			if b == '\r' {
				if i+1 < len(data) {
					if data[i+1] == '\n' {
						indexesToDelete = append(indexesToDelete, i)
					}
				}
				data[i] = '\n'
			}
		}
		data = removeFromSlice(data, indexesToDelete)
		w.Write(data)
	}
	w.Flush()
}

// removeFromSlice accepts a byte slice and an integer slice
// and deletes each index from the integer slice from the
// byte slice
func removeFromSlice(data []byte, indexes []int) []byte {
	for _, i := range indexes {
		data = append(data[:i], data[i+1:]...)
	}
	return data
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please specify one or more files as input")
		os.Exit(1)
	}

	for _, fileName := range args[1:] {
		cleanFile(fileName)
	}
}
