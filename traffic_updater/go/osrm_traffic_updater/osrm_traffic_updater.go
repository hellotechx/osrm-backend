package main

import (
	"flag"
	"fmt"
	"time"
)

type trafficProxyFlags struct {
	port            int
	ip              string
	region          string
	trafficProvider string
	mapProvider     string
}

var flags struct {
	trafficProxyFlags trafficProxyFlags
	testMode          string
	mappingFile       string
	csvFile           string
}

func init() {
	flag.IntVar(&flags.trafficProxyFlags.port, "p", 6666, "traffic proxy listening port")
	flag.StringVar(&flags.trafficProxyFlags.ip, "c", "127.0.0.1", "traffic proxy ip address")
	flag.StringVar(&flags.trafficProxyFlags.region, "region", "na", "region")
	flag.StringVar(&flags.trafficProxyFlags.trafficProvider, "traffic", "", "traffic data provider")
	flag.StringVar(&flags.trafficProxyFlags.mapProvider, "map", "", "map data provider")
	flag.StringVar(&flags.testMode, "testmode", "", "test mode, e.g. grpc_get_all, ut_grpc_get_way")
	flag.StringVar(&flags.mappingFile, "m", "wayid2nodeids.csv", "OSRM way id to node ids mapping table")
	flag.StringVar(&flags.csvFile, "f", "traffic.csv", "OSRM traffic csv file")
}

const TASKNUM = 128
const CACHEDOBJECTS = 4000000

func main() {
	flag.Parse()

	startTime := time.Now()
	defer func() {
		endTime := time.Now()
		fmt.Printf("Total processing time %f seconds\n", endTime.Sub(startTime).Seconds())
	}()

	if flags.testMode != "" {

		flows, err := getAllFlowsByGRPC(flags.trafficProxyFlags)
		if err != nil {
			fmt.Println(err)
			return
		}
		for i := 0; i < 10 && i < len(flows); i++ {
			fmt.Println(flows[i])
		}

		// if test mode, always return after test run
		return
	}

	isFlowDoneChan := make(chan bool, 1)
	wayid2speed := make(map[int64]int)
	go getTrafficFlow(flags.trafficProxyFlags.ip, flags.trafficProxyFlags.port, wayid2speed, isFlowDoneChan)

	var sources [TASKNUM]chan string
	for i := range sources {
		sources[i] = make(chan string, CACHEDOBJECTS)
	}
	go loadWay2NodeidsTable(flags.mappingFile, sources)

	isFlowDone := wait4PreConditions(isFlowDoneChan)
	if isFlowDone {
		var ds dumperStatistic
		ds.Init(TASKNUM)
		dumpSpeedTable4Customize(wayid2speed, sources, flags.csvFile, &ds)
		ds.Output()
	}
}

func wait4PreConditions(flowChan <-chan bool) bool {
	var isFlowDone bool
loop:
	for {
		select {
		case f := <-flowChan:
			if !f {
				fmt.Printf("[ERROR] Communication with traffic server failed.\n")
				break loop
			} else {
				isFlowDone = true
				break loop
			}
		}
	}
	return isFlowDone
}
