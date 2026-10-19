// Command go-infra is the infrastructure shell: one process hosting the members compiled in below, each on its own port with
// its own runtime (be-sdk-go's shell launcher). Which members it hosts in a deployment is BrickKit's choice
// (BRICKKIT_SERVED_MEMBERS_CONFIG); the launcher refuses members that are not compiled in or not at the declared
// version.
package main

import "github.com/brickKit/be-sdk-go/v3/shell"

// ID and Version are the shell's own identity; component.yaml says the same.
const (
	ID      = "be/go-infra"
	Version = "3.0.14"
)

func main() { shell.Main(shell.Identity{ID: ID, Version: Version}, members()...) }
