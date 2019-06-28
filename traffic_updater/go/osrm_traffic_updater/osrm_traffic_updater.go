package main

import (
	"context"
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/Telenav/osrm-backend/traffic_updater/go/gen-go/proxy"
	"github.com/apache/thrift/lib/go/thrift"
)

var flags struct {
	port          int
	ip            string
	mappingFile   string
	csvFile       string
	highPrecision bool
}

func init() {
	flag.IntVar(&flags.port, "p", 6666, "traffic proxy listening port")
	flag.StringVar(&flags.ip, "c", "127.0.0.1", "traffic proxy ip address")
	flag.StringVar(&flags.mappingFile, "m", "wayid2nodeids.csv", "OSRM way id to node ids mapping table")
	flag.StringVar(&flags.csvFile, "f", "traffic.csv", "OSRM traffic csv file")
	flag.BoolVar(&flags.highPrecision, "d", false, "use high precision speeds, i.e. decimal")
}

func flows2map(flows []*proxy.Flow, m map[uint64]int) {
	for _, flow := range flows {
		wayid := (uint64)(flow.WayId)
		m[wayid] = int(flow.Speed)
	}
}

func main() {
	flag.Parse()

	var transport thrift.TTransport
	var err error

	// make socket
	targetServer := flags.ip + ":" + strconv.Itoa(flags.port)
	fmt.Println("connect traffic proxy " + targetServer)
	transport, err = thrift.NewTSocket(targetServer)
	if err != nil {
		fmt.Println("Error opening socket:", err)
		return
	}

	// Buffering
	transport, err = thrift.NewTFramedTransportFactoryMaxLength(thrift.NewTTransportFactory(), 1024*1024*1024).GetTransport(transport)
	if err != nil {
		fmt.Println("Error get transport:", err)
		return
	}
	defer transport.Close()
	if err := transport.Open(); err != nil {
		fmt.Println("Error opening transport:", err)
		return
	}

	// protocol encoder&decoder
	protocol := thrift.NewTCompactProtocolFactory().GetProtocol(transport)

	// create proxy client
	client := proxy.NewProxyServiceClient(thrift.NewTStandardClient(protocol, protocol))

	// get flows
	breforeGetFlowsTime := time.Now()
	fmt.Println("getting flows")
	var defaultCtx = context.Background()
	flows, err := client.GetAllFlows(defaultCtx)
	if err != nil {
		fmt.Println("get flows failed:", err)
		return
	}
	afterGotFlowTime := time.Now()
	fmt.Printf("got flows count: %d, time used: %f seconds\n", len(flows), afterGotFlowTime.Sub(breforeGetFlowsTime).Seconds())

	wayid2speed := make(map[uint64]int)
	flows2map(flows, wayid2speed)
	afterFlows2MapTime := time.Now()
	fmt.Printf("flows2map time used: %f seconds\n", afterFlows2MapTime.Sub(afterGotFlowTime).Seconds())

	generateSpeedTable(wayid2speed, flags.mappingFile, flags.csvFile)
	afterGenerateSpeedTableTime := time.Now()
	fmt.Printf("generateSpeedTable time used: %f seconds\n", afterGenerateSpeedTableTime.Sub(afterFlows2MapTime).Seconds())
}
