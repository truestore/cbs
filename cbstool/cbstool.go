package main

import (
	"cbs/netflow"
	"cbs/writer"

	"compress/gzip"
	"encoding/csv"
	"os"
	"fmt"
)

const BS = 256 * 1024

func main() {
	ifile, err := os.Open(os.Args[1])
	if err != nil {
		return
	}

	igz, err := gzip.NewReader(ifile)
	if err != nil {
		return
	}

	icsv := csv.NewReader(igz)

	ww := writer.NewUint32Writer(BS)

	for {
		rec, err := icsv.Read()
		if err != nil {
			break
		}

		rv5, err := netflow.ParseCsv(rec)
		if err != nil {
			break
		}

		ww.Append(rv5.DstAddr)
		if ww.Len() >= BS {
			h, v, err := ww.Flush()
			fmt.Println(h, len(v), err)
		}

	}
}

