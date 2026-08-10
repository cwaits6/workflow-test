package main

import "fmt"

var version = "0.10.0" // x-release-please-version

func main() {
	fmt.Printf("workflow-test/v2 version: %s\n", version)
	fmt.Println("second feature (patched)")
}

func unused() string { return "sticky release-as probe" }
