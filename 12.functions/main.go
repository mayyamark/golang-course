package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hi")
	port := 3000
	// isStarted := startWebServer(port, 2)
	// fmt.Println(isStarted) // true

	// err := imporvedStartWebServer(port)
	// fmt.Println((err)) // Something went wrong!

	_, err := imporvedStartWebServerWithMultipleRecivedAndReturnedData(port) // ignore the first return value
	fmt.Println(err)                                                         // 3000	Something went wrong!

}

func startWebServer(port, numberOfRetries int) bool {
	// or func startWebServer(port int, numberOfRetries int) { => it's the same, the types are equal
	fmt.Println("Starting server...")
	// Some logic
	fmt.Println("Server started on port: ", port)
	fmt.Println("Number of retries: ", numberOfRetries)

	return true
}

// return an error by using the error extension
func imporvedStartWebServer(port int) error {
	// or func startWebServer(port int, numberOfRetries int) { => it's the same, the types are equal
	fmt.Println("Starting server...")
	// Some logic
	fmt.Println("Server started on port: ", port)

	return errors.New("Something went wrong!")
}

func imporvedStartWebServerWithMultipleRecivedAndReturnedData(port int) (int, error) {
	// or func startWebServer(port int, numberOfRetries int) { => it's the same, the types are equal
	fmt.Println("Starting server...")
	// Some logic
	fmt.Println("Server started on port: ", port)

	return port, errors.New("Something went wrong!")
}
