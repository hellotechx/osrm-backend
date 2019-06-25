package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var flags struct {
	idsmapping  string
	mocktraffic string
	output      string
}

func init() {
	flag.StringVar(&flags.idsmapping, "i", "", "Input id mapping file.")
	flag.StringVar(&flags.mocktraffic, "m", "", "Mock traffic file")
	flag.StringVar(&flags.output, "o", "", "Output csv file")
}

func main() {
	flag.Parse()

	if len(flags.idsmapping) == 0 || len(flags.mocktraffic) == 0 || len(flags.output) == 0 {
		fmt.Printf("[ERROR]Input or Output file path is empty.\n")
		return
	}

	wayid2speed := make(map[uint64]int)
	loadMockTraffic(flags.mocktraffic, wayid2speed)

	GenerateSpeedTable(wayid2speed, flags.idsmapping, flags.output)
}

func loadMockTraffic(trafficPath string, wayid2speed map[uint64]int) {
	// load mock traffic file
	mockfile, err := os.Open(trafficPath)
	defer mockfile.Close()
	if err != nil {
		log.Fatal(err)
		fmt.Printf("Open pbf file of %v failed.\n", trafficPath)
		return
	}
	fmt.Printf("Open pbf file of %s succeed.\n", trafficPath)

	csvr := csv.NewReader(mockfile)
	for {
		row, err := csvr.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
				break
			} else {
				fmt.Printf("Error during decoding mock traffic, row = %v\n", err)
				return
			}
		}

		var wayid uint64
		var speed int64
		if wayid, err = strconv.ParseUint(row[0], 10, 64); err != nil {
			fmt.Printf("#Error during decoding wayid, row = %v\n", row)
		}
		if speed, err = strconv.ParseInt(row[1], 10, 32); err != nil {
			fmt.Printf("#Error during decoding speed, row = %v\n", row)
		}

		wayid2speed[wayid] = (int)(speed)
	}
}
