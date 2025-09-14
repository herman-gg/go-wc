package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

type Flags struct {
	isForLines bool
	isForBytes bool
}

func main() {
	isForLines := flag.Bool("l", false, "Count lines")
	isForBytes := flag.Bool("b", false, "Count bytes")
	flag.Parse()

	flags := Flags{*isForLines, *isForBytes}

	fmt.Println(count(os.Stdin, flags))
}

func count(r io.Reader, flags Flags) int {
	scanner := bufio.NewScanner(r)

	splitFunc := bufio.ScanWords
	if flags.isForLines {
		splitFunc = bufio.ScanLines
	} else if flags.isForBytes {
		splitFunc = bufio.ScanBytes
	}

	scanner.Split(splitFunc)

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}
