package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/nochso/tocenize"
)

var VERSION = "?"

func main() {
	log.SetFlags(0)
	flag.Usage = func() {
		fmt.Println("tocenize [options] FILE...")
		fmt.Println()
		flag.PrintDefaults()
	}
	job := tocenize.Job{}
	flag.IntVar(&job.MinDepth, "min", 1, "minimum depth")
	flag.IntVar(&job.MaxDepth, "max", 99, "maximum depth")
	flag.BoolVar(&job.Diff, "d", false, "print full diff to stdout")
	flag.BoolVar(&job.Print, "p", false, "print full result to stdout")
	flag.BoolVar(&job.ExistingOnly, "e", false, "update only existing TOC (no insert)")
	flag.BoolVar(&tocenize.Verbose, "v", false, "verbose output")
	showVersion := flag.Bool("V", false, "print version")
	flag.Parse()

	if *showVersion {
		fmt.Println(VERSION)
		os.Exit(0)
	}
	if flag.NArg() == 0 {
		log.Println("too few arguments")
		flag.Usage()
		os.Exit(2)
	}

	for _, arg := range flag.Args() {
		paths, err := filepath.Glob(arg)
		if err != nil {
			log.Println(err)
			continue
		}

		for _, path := range paths {
			log.SetPrefix(path + ": ")
			doc, err := tocenize.NewDocument(path)
			if err != nil {
				log.Println(err)
				continue
			}
			toc := tocenize.NewTOC(doc, job)
			_, err = doc.Update(toc, job)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
}
