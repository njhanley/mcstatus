package main

import (
	"fmt"
	"os"
	"time"
)

func fatal(v ...interface{}) {
	fmt.Fprintln(os.Stderr, v...)
	os.Exit(1)
}

func main() {
	status, err := queryMinecraft(os.Args[1], 5*time.Second)
	if err != nil {
		fatal(err)
	}
	fmt.Println(status)
}
