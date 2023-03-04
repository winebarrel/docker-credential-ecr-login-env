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
		fmt.Println(`{"ServerURL":"ServerURL","Username":"Username","Secret":"Secret"}`)
	case "erase":
		// nothing to do
	case "list":
		fmt.Println(`{"ServerURL":"Username"}`)
	case "version":
		fmt.Println("name (package) version")
	}
}
