package main

import (
	"net/http"

	"github.com/mayyamark/golang-course/13.project-improved/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
