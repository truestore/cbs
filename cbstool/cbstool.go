package main

import (
	"github.com/truestore/cbs/netflow"
	"github.com/truestore/cbs/writer"

	"compress/gzip"
	"encoding/csv"
	"fmt"
	"os"
)

const BS = 256 * 1024
const LOCALPREFIX = 3569426432
const LOCALMASK = 4294959104

func localIp(x uint32) bool {
	return (x & LOCALMASK) == LOCALPREFIX
}

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
	wd := writer.NewUint16Writer(BS)

	for {
		rec, err := icsv.Read()
		if err != nil {
			break
		}

		rv5, err := netflow.ParseCsv(rec)
		if err != nil {
			break
		}

		if !localIp(rv5.DstAddr) {
			ww.Append(rv5.DstAddr)
		}
		wd.Append(rv5.DstPort)
		if ww.Len() >= BS {
			h, v, err := ww.Flush()
			fmt.Println(h, len(v), err)
			h, v, err = wd.Flush()
			fmt.Println(h, len(v), err)
		}

	}
}
