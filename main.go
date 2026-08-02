package main

import "fmt"

var version = "dev"

func main() {
	fmt.Printf("workflow-test version: %s\n", version)
	fmt.Println("release-please validation: feature one")
}
