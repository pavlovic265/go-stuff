package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	content, err := os.ReadFile("file.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
	} else {
		fmt.Println(string(content))
	}

	result, err := divide(7, 0)
	// result, err := divide(7, 0)
	if err != nil {
		fmt.Println("Error divide: ", err.Error())
		// os.Exit(-1)
		log.Fatal(err)
	}
	fmt.Println(result)
}

func divide(x int, y int) (float64, error) {
	// if y == 0 {
	// 	return float64(0), errors.New("can not divide by zero")
	// }
	if y == 0 {
		return float64(0), CustomError{Message: "Can not divide by zero", Code: -1}
	}

	return float64(x) / float64(y), nil
}

type CustomError struct {
	Message string
	Code    int
}

func (c CustomError) Error() string {
	return c.Message + " " + strconv.Itoa(c.Code)
}
