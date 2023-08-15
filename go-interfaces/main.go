package main

import "fmt"

type Printer interface {
	printInfo()
}

type Book struct {
	Title string
	Price float32
}

func (b Book) printInfo() {
	fmt.Printf("Book: %s Price: %.2f\n", b.Title, b.Price)
}

type Drink struct {
	Name  string
	Price float32
}

func (d Drink) printInfo() {
	fmt.Printf("Drink: %s Price: %.2f\n", d.Name, d.Price)
}

type Toy struct {
	Name  string
	Price float32
}

func (t Toy) printInfo() {
	fmt.Printf("Toy: %s Price: %.2f\n", t.Name, t.Price)
}

func main() {

	gunslinger := Book{
		Title: "The gunslinger",
		Price: 4.75,
	}

	coffee := Drink{
		Name:  "Coffee",
		Price: 2.25,
	}

	rubixCube := Toy{
		Name:  "Rubix Cube",
		Price: 5.0,
	}

	gunslinger.printInfo()
	coffee.printInfo()
	rubixCube.printInfo()

	info := []Printer{gunslinger, coffee, rubixCube}

	info[0].printInfo()
	info[1].printInfo()
	info[2].printInfo()

}
