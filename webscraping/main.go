package main

import (
	"net/http"

	"github.com/UnderTechnologies/go-micrho/micrho"
	micrho_http "github.com/UnderTechnologies/go-micrho/micrho/http"
)

func main() {
	httpServer := micrho_http.NewDefaultServer("8080")
	app := micrho.NewDefaultApp(micrho.WithServer(httpServer))

	httpServer.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Hello, World!"))
	})

	app.Start()
}
