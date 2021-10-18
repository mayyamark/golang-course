package controllers

import "net/http"

// root routing => go to the correct controller
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)  // the /users is handled by the usersController
	http.Handle("/users/", *uc) // the /users/ is handled by the usersController
}
