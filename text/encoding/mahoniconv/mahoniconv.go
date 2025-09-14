package main

import (
	"flag"
	"io"
	"log"
	"os"

	"gitee.com/quant1x/gox/text/encoding"
)

// An iconv workalike using mahonia.
var from = flag.String("f", "utf-8", "source character set")
var to = flag.String("t", "utf-8", "destination character set")

func main() {
	flag.Parse()

	var r io.Reader = os.Stdin
	var w io.Writer = os.Stdout

	if *from != "utf-8" {
		decode := encoding.NewDecoder(*from)
		if decode == nil {
			log.Fatalf("Could not create decoder for %s", *from)
		}
		r = decode.NewReader(r)
	}

	if *to != "utf-8" {
		encode := encoding.NewEncoder(*to)
		if encode == nil {
			log.Fatalf("Could not create decoder for %s", *to)
		}
		w = encode.NewWriter(w)
	}

	io.Copy(w, r)
}
