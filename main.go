package main

import "fmt"

var version = "0.9.0" // x-release-please-version

func main() {
	fmt.Printf("workflow-test/v2 version: %s\n", version)
	fmt.Println("second feature")
}
