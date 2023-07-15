package main

import "fmt"

func main() {
	// fmt.Println(test(3, "world"))

	// var a, b int = 1, 2
	// passBy(a, &b)
	// fmt.Printf("a %d, b %d\n", a, b)

	result := sum(1, 2, 3, 4)
	fmt.Println("result", result)
	// sum(make([]int, 4, 1, 2, 3, 4))
}

func sum(numbers ...int) int {
	sum := 0
	for n := range numbers {
		sum += n
	}

	return sum
}

func passBy(val1 int, val2 *int) {
	val1 = 10
	*val2 = 20

}

func test(a int, b string) (int, string) {
	return (a + 1), ("Hello " + b)
}
