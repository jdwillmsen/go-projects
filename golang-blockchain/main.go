package main

import (
	"github.com/jdwillmsen/golang-blockchain/cli"
	"os"
)

// TODO: Migrate project to work with latest version of golang
// TODO: Continue from part 7 - Try to figure out broken send
func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
