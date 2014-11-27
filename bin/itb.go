package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sgmac/itbuk"
)

var (
	usage string = `usage: itb -p [npag] <topic>
-p Number of pages per search, defaults to one.
`
	pages = flag.Int("p", 1, "")
)

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s", usage)
		os.Exit(1)
	}
	flag.Parse()

	if len(os.Args) < 2 {
		flag.Usage()
	}
	topic := flag.Args()[0]

	b, err := itbuk.Search(topic, *pages)

	if err != nil {
		log.Fatalf("Error:%s", err)
	}

	for _, book := range b {
		fmt.Println(book)
	}

}
