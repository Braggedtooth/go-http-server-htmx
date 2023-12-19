package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/template"
)

func main() {
	app := pocketbase.New()

	// serves static files from the provided public dir (if exists)
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		return nil
	})
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		registry := template.NewRegistry()
		e.Router.GET("/hello/:name", func(c echo.Context) error {
			name := c.PathParam("name")
			records, err := app.Dao().FindRecordsByExpr("greetings")
			record := records[rand.Intn(len(records))]

			if err != nil {
				fmt.Println("The Error is", err)
				return err
			}
			message := fmt.Sprintf(record.GetString("greeting"), name)

			html, err := registry.LoadFiles(
				"views/hello.html",
			).Render(map[string]any{
				"Message": message,
			})

			if err != nil {
				// or redirect to a dedicated 404 HTML page
				return apis.NewNotFoundError("", err)
			}

			return c.HTML(http.StatusOK, html)

		} /* optional middlewares */)

		return nil
	})
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
