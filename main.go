// +build darwin

package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

var (
	flagOutput = flag.String("o", "", "output File")
)

func main() {

	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Please enter exact one url-string. // as args[0]")
		os.Exit(1)
	}
	url := args[0]

	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error while reading the url: %s\nError: %v", url, err)
	}
	defer res.Body.Close()
	var w io.Writer = os.Stdout

	if *flagOutput != "" {
		err := os.MkdirAll(filepath.Dir(*flagOutput), 0755)
		if err != nil {
			fmt.Printf("Error while creating directories: %v", err)
			os.Exit(1)
		}
		f, err := os.OpenFile(*flagOutput, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			fmt.Printf("Error while creating the file from: %s\n%#v", *flagOutput, err)
		}
		defer f.Close()
		w = f
	}

	io.Copy(w, res.Body)
}
