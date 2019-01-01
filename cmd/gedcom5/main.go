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
		fmt.Printf("Header:\n")
		for _, l := range file.Header.Lines {
			fmt.Printf("  - %s\n", l.String())
		}
		fmt.Printf("Trailer:\n")
		for _, l := range file.Trailer.Lines {
			fmt.Printf("  - %s\n", l.String())
		}
		fmt.Printf("Lines decoded: %d\n", len(file.Lines))
		fmt.Printf("Top-level records: %d\n", len(file.Records))
		indis := make([]*gedcom5.IndividualRecord, 0, 10)
		var families int
		var multimedias int
		var notes int
		var repos int
		var sources int
		var submitters int
		var unknowns int
		for _, rec := range file.Records {
			switch rec.(type) {
			case *gedcom5.IndividualRecord:
				indis = append(indis, rec.(*gedcom5.IndividualRecord))
			case *gedcom5.FamilyRecord:
				families++
			case *gedcom5.MultimediaRecord:
				multimedias++
			case *gedcom5.NoteRecord:
				notes++
			case *gedcom5.RepositoryRecord:
				repos++
			case *gedcom5.SourceRecord:
				sources++
			case *gedcom5.SubmitterRecord:
				submitters++
			default:
				unknowns++
			}
		}
		fmt.Printf(" - Individuals: %d\n", len(indis))
		for _, indi := range indis {
			printIndividual(indi)
		}
		fmt.Printf(" - Families: %d\n", families)
		fmt.Printf(" - Multimedia items: %d\n", multimedias)
		fmt.Printf(" - Notes: %d\n", notes)
		fmt.Printf(" - Repositories: %d\n", repos)
		fmt.Printf(" - Sources: %d\n", sources)
		fmt.Printf(" - Submitters: %d\n", submitters)
		fmt.Printf(" - Unknown records: %d\n", unknowns)
	}
}

func printIndividual(indi *gedcom5.IndividualRecord) {
	pfx := "        "
	fmt.Printf("    - %s\n", indi.String())
	if indi.Sex != "" {
		fmt.Printf("%sSex: %s\n", pfx, indi.Sex)
	}
	if indi.Occupation != "" {
		fmt.Printf("%sOccupation: %s\n", pfx, indi.Occupation)
	}
	if indi.Religion != "" {
		fmt.Printf("%sReligion: %s\n", pfx, indi.Religion)
	}
	if indi.Birth.Date != "" {
		fmt.Printf("%sBirthdate: %s\n", pfx, indi.Birth.Date)
	}
}
