package test

import (
	"github.com/eris-ltd/eris-db/test/fixtures"
	"testing"
)

// Needs to be in a _test.go file to be picked up
func TestWrapper(runner func() int) int {
	ffs := fixtures.NewFileFixtures("Eris-DB")

	defer ffs.RemoveAll()

	err := initGlobalVariables(ffs)

	if err != nil {
		panic(err)
	}

	// start a node
	ready := make(chan error)
	go newNode(ready)
	err = <-ready

	if err != nil {
		panic(err)
	}

	return runner()
}

// This main function exists as a little convenience mechanism for running the
// delve debugger which doesn't work well from go test yet. In due course it can
// be removed, but it's flux between pull requests should be considered
// inconsequential, so feel free to insert your own code if you want to use it
// as an application entry point for delve debugging.
func DebugMain() {
	t := &testing.T{}
	TestWrapper(func() int {
		testNameReg(t, "JSONRPC")
		return 0
	})
}

func Successor(x int) int {
	return x + 1
}
