package main

import "github.com/aleccunningham/pt/cmd"

// Build number and versions injected at compile time
var (
	Version = "unknown"
	Build   = "unknown"
)

func main() {
	cmd.Execute(Version, Build)
}
