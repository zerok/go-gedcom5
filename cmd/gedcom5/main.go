package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/pflag"
	gedcom5 "gitlab.com/zerok/go-gedcom5"
)

func main() {
	pflag.Parse()
	for _, a := range pflag.Args() {
		fmt.Printf("# %s\n", a)
		fp, err := os.Open(a)
		if err != nil {
			log.Fatalf("Failed to open %s: %s", a, err.Error())
		}
		var file gedcom5.File
		if err := gedcom5.NewDecoder(fp).Decode(&file); err != nil {
			log.Fatalf("Failed to decode %s: %s", a, err.Error())
		}
		fmt.Printf("Lines decoded: %d\n", len(file.Lines))
	}
}
