package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var (
	password string
	info     string
)

func main() {
	// readStats("file.dat")
	// readWholeFile("file.dat")
	// readByLine("file.dat")
	// readByWord("file.dat")
	// readByByte("file.dat", 8)
	readConfig("test.cfg")
}

func readConfig(filename string) {
	file, oErr := os.Open(filename)
	if oErr != nil {
		fmt.Println(oErr.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		raw := strings.Split(scanner.Text(), "=")
		key := raw[0]
		value := raw[1]

		if key == "password" {
			password = value
		} else if key == "test" {
			info = value
		}
	}
	fmt.Println(password)
	fmt.Println(info)

}

func readByByte(filename string, size uint8) {
	file, oErr := os.Open(filename)
	if oErr != nil {
		fmt.Println(oErr.Error())
		return
	}
	defer file.Close()

	buf := make([]byte, size)

	for {
		totalRead, rErr := file.Read(buf)
		if rErr != nil {
			if rErr != io.EOF {
				fmt.Println(rErr.Error())
			}
			break
		}

		fmt.Println(string(buf[:totalRead]))
	}

}

func readByWord(filename string) {
	file, oErr := os.Open(filename)
	if oErr != nil {
		fmt.Println(oErr.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func readByLine(filename string) {
	file, oErr := os.Open(filename)
	if oErr != nil {
		fmt.Println(oErr.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func readWholeFile(filename string) {
	content, rfErr := ioutil.ReadFile(filename)
	if rfErr != nil {
		fmt.Println(rfErr.Error())
		return
	}

	fmt.Println(string(content))
}

func readStats(filename string) {
	file, oErr := os.Open(filename)
	if oErr != nil {
		fmt.Println(oErr.Error())
		return
	}
	defer file.Close()

	stats, sErr := file.Stat()
	if sErr != nil {
		fmt.Println(sErr.Error())
		return
	}
	fmt.Printf("File Name: %s\n", stats.Name())
	fmt.Printf("Time Modified: %s\n", stats.ModTime().Format("15:04:03"))
}
