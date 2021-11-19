package main

import (
	"github.com/jvillarreal-w/academy-go-q42021/infrastructure/router"
	"github.com/jvillarreal-w/academy-go-q42021/registry"
	u "github.com/jvillarreal-w/academy-go-q42021/utils"
	"github.com/labstack/echo"
)

func main() {
	u.GeneralLogger.Println("Starting...")
	r := registry.NewRegistry()

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	u.GeneralLogger.Println("Listening at http://localhost:8000")
	if err := e.Start(":8000"); err != nil {
		u.ErrorLogger.Printf("Unable to start server: %v", err)
	}
}
