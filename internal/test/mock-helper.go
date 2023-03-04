package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	key := os.Args[1]

	switch key {
	case "store":
		// nothing to do
	case "get":
		fmt.Printf(
			`{"ServerURL":"%s","Username":"%s","Secret":"%s"}`,
			os.Getenv("TEST_ServerURL"),
			os.Getenv("TEST_Username"),
			os.Getenv("TEST_Secret"),
		)
	case "erase":
		// nothing to do
	case "list":
		fmt.Println(`{"ServerURL":"Username"}`)
	case "version":
		fmt.Println("name (package) version")
	}
}
