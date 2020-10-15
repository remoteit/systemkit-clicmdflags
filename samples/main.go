package main

import (
	"log"
	"os"

	"github.com/remoteit/systemkit-clicmdflags/samples/cmdflags"
)

func main() {
	os.Args = append(os.Args, "device", "help")
	if err := cmdflags.Execute(); err != nil {
		log.Fatal(err)
	}
}
