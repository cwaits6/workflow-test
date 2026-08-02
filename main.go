package main

import "fmt"

var version = "0.3.0" // x-release-please-version

func main() {
	fmt.Printf("workflow-test version: %s\n", version)
	fmt.Println("release-please validation: feature one")
}
