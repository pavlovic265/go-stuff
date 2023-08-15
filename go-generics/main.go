package gogenerics

import "fmt"

func addInts(list []int) int {
	var sum int
	for _, item := range list {
		sum += item
	}

	return sum
}

func addFloats(list []float32) float32 {
	var sum float32
	for _, item := range list {
		sum += item
	}

	return sum
}

type NumberType interface {
	int | float32 | float64 | uint16 | uint32 | uint64
}

func addList[T NumberType](list []T) T {
	var sum T
	for _, item := range list {
		sum += item
	}

	return sum

}

// func printMe(thing interface{}) {
// 	fmt.Println(thing)
// }

func printMe(thing any) {
	fmt.Println(thing)
}

func main() {

	fmt.Printf("Sum of ints: %d\n", addInts([]int{1, 2, 3, 4, 5}))
	fmt.Printf("Sum of floats: %.2f\n", addFloats([]float32{1.2, 2.3, 3.5, 4.1, 5.3}))

	fmt.Println("Generic bellow ----------------")
	fmt.Printf("Sum of ints: %d\n", addList([]int{1, 2, 3, 4, 5}))
	fmt.Printf("Sum of floats: %.2f\n", addList([]float32{1.2, 2.3, 3.5, 4.1, 5.3}))

	printMe(1.0)
	printMe("hi")

}
