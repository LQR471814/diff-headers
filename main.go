package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/rodaine/table"
)

func main() {
	flag.Parse()
	left := flag.Arg(0)
	right := flag.Arg(1)

	leftFile, err := os.ReadFile(left)
	if err != nil {
		log.Fatal(err)
	}
	rightFile, err := os.ReadFile(right)
	if err != nil {
		log.Fatal(err)
	}

	leftHeaders := ParseHeaders(string(leftFile))
	rightHeaders := ParseHeaders(string(rightFile))

	fmt.Printf("===== %s is missing the following headers from %s\n\n", right, left)
	t1 := table.New("HEADER", "VALUE")
	for _, pair := range leftHeaders.Pairs {
		_, ok := rightHeaders.Get(pair.Key)
		if !ok {
			t1.AddRow(pair.Key, pair.Value)
		}
	}
	t1.Print()

	fmt.Printf("\n===== %s is missing the following headers from %s\n\n", left, right)
	t2 := table.New("HEADER", "VALUE")
	for _, pair := range rightHeaders.Pairs {
		_, ok := leftHeaders.Get(pair.Key)
		if !ok {
			t2.AddRow(pair.Key, pair.Value)
		}
	}
	t2.Print()

	fmt.Printf("\n===== Differences in shared header values\n\n")
	t3 := table.New("HEADER", left, right)
	for _, pair := range leftHeaders.Pairs {
		rightValue, ok := rightHeaders.Get(pair.Key)
		if ok && pair.Value != rightValue {
			t3.AddRow(pair.Key, pair.Value, rightValue)
		}
	}
	t3.Print()
}
