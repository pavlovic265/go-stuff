package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	// createFile("createFile.txt")
	// writeBytes("writeBytes.txt", []byte("Hello I am a byte slice"))
	// writeByLine("writeByLine.txt", []string{"line1", "line2", "line3"})
	append("writeByLine.txt", "we are appending data")
}

func writeByLine(filename string, lines []string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	for _, val := range lines {
		_, fErr := fmt.Fprintln(file, val)
		if err != nil {
			fmt.Println(fErr.Error())
			return
		}
	}
}

func writeBytes(filename string, data []byte) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	size, wErr := file.Write(data)
	if err != nil {
		fmt.Println(wErr.Error())
		return
	}

	fmt.Printf("Wrote %d byte to the file\n", size)
}

func createFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	_, wErr := file.WriteString("new file")
	if err != nil {
		fmt.Println(wErr.Error())
		return
	}
}

func append(filename, data string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, fs.ModeAppend)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	_, fErr := fmt.Fprintln(file, data)
	if err != nil {
		fmt.Println(fErr.Error())
		return
	}

}
