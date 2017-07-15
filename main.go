package main

import (
	"log"

	"github.com/gobuffalo/envy"
	"github.com/saopayne/blogify-web/actions"
)

func main() {
	port := envy.Get("PORT", "3000")
	app := actions.App()
	log.Fatal(app.Start(port))
}
