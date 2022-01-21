// Program fix-newlines replaces \r\n with \n in text files.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"unicode/utf8"
)

func main() {
	log.SetFlags(0)
	flag.Parse()
	if err := run(flag.Args()); err != nil {
		log.Fatal(err)
	}
}

func run(names []string) error {
	if len(names) == 0 {
		return fmt.Errorf(usageText)
	}
	for _, name := range names {
		if err := fixFile(name); err != nil {
			return err
		}
	}
	return nil
}

func fixFile(path string) error {
	body, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if !bytes.Contains(body, []byte("\r\n")) {
		return nil
	}
	if !utf8.Valid(body) {
		return fmt.Errorf("%s is not a valid utf-8 file", path)
	}
	body = bytes.ReplaceAll(body, []byte("\r\n"), []byte("\n"))
	return os.WriteFile(path, body, 0666)
}

const usageText = "usage: fix-newlines file1.txt ..."

func init() {
	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), usageText)
	}
}
