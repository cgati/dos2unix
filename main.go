package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

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
		for i, b := range data {
			if b == '\r' {
				data[i] = '\n'
			}
		}
		w.Write(data)
	}
	w.Flush()
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
