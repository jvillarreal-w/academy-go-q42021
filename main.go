package main

import (
	"fmt"

	"github.com/jvillarreal-w/academy-go-q42021/infrastructure/router"
	"github.com/jvillarreal-w/academy-go-q42021/registry"
	"github.com/labstack/echo"
)

func main() {
	r := registry.NewRegistry()

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Listening at http://localhost:8000")
	if err := e.Start(":8000"); err != nil {
		fmt.Printf("Unable to start server: %v", err)
	}
}
