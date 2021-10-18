package controllers

import (
	"net/http"
	"regexp"
)

// handle the routing of http requests comes into our web server to the correct method that's going to handle that in the model layer

type userController struct {
	userIDPattern *regexp.Regexp // userIDPattern - to check which type of request it's going to work with (check the path)
}

/*
Add a behavior to the userController struct	- here the signature is important

(uc userController) bind the func to the userController type => the function becomes a method
receives Req & Res objects from the net/http package => implements the Handler interface (from net/http package)
*/
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// return a static message	[]byte("message") => type conversion (convert strings to byte slices)
	w.Write([]byte("Hello from the User Controller!"))
}

/*
A constructor function - by convention: starts with new

returns a pointer to the userController
constructing a new userController object
*/
func newUserController() *userController {
	return &userController{ // using addressof operator
		// /users/<number>
		userIDPattern: regexp.MustCompile(`^/users(\d+)/?`),
	}
}
