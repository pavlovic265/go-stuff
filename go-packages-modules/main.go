package main

import (
	"fmt"

	"example.com/packages/util"
	// "net/http"
	// "github.com/gorilla/mux"
)

func main() {
	greeting := fmt.Sprintf("Hello %s!", "Marko")
	fmt.Println(greeting)

	fmt.Printf("Length of greeting is %d!\n", util.StringLength(greeting))
	fmt.Printf("Greeting util %s!\n", util.GetGreetingUtil())

	// r := mux.NewRouter()

	// http.ListenAndServe(":9000", r)

}
