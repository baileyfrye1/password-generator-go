package main

import (
	"flag"
	"fmt"

	"golang.design/x/clipboard"
)

func main() {
	var (
		length     int
		special    bool
		passphrase bool
		words      int
		separator  string
	)

	RegisterFlags(&length, &special, &passphrase, &words, &separator)
	flag.Parse()

	if err := clipboard.Init(); err != nil {
		fmt.Println("Error initializing clipboard: ", err)
		return
	}

	if passphrase {
		HandlePassphrase(words, separator)
	} else {
		HandlePassword(special, length)
	}
}
