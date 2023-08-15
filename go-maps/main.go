package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["Marko"] = 32
	fmt.Printf("Marko is %d years old!\n", ages["Marko"])
	ages["Susan"] += 1
	fmt.Printf("Susan is %d years old!\n ", ages["Susan"])

	gpas := map[string]float32{
		"Marko": 3.0,
		"Alice": 3.9,
	}

	fmt.Printf("Marko GPA is %.2f!\n", gpas["Marko"])
	fmt.Printf("Marko GPA is %.2f!\n", gpas["Alice"])

	var visited map[string]bool
	visited = make(map[string]bool)
	visited["A"] = true
	fmt.Printf("Marko GPA is %t!\n", visited["A"])

	fruits := []string{
		"bananas",
		"kiwi",
		"apples",
		"strawberries",
		"tomatoes",
		"bananas",
	}

	fmt.Println("Duplicated: ", duplicates(fruits))
}

func duplicates(words []string) string {
	dupMap := make(map[string]bool)
	for _, item := range words {
		if !dupMap[item] {
			dupMap[item] = true
		} else {
			return item
		}
	}

	return ""
}
