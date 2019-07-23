package main

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/Telenav/osrm-backend/traffic_updater/go/grpc/proxy"
	"google.golang.org/grpc"
)

func BenchmarkGetAllFlowsByGRPC(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping test in short mode.")
	}

	for i := 0; i < b.N; i++ {
		flows, err := getAllFlowsByGRPC(flags.trafficProxyFlags)
		if err != nil {
			b.Error(err)
		}
		quickViewFlows(flows, 10) //quick view first 10 lines
	}
}

func BenchmarkGetSingleFlowEachTimeByGRPC(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping test in short mode.")
	}

	// make RPC client
	targetServer := flags.trafficProxyFlags.ip + ":" + strconv.Itoa(flags.trafficProxyFlags.port)
	fmt.Println("connect traffic proxy " + targetServer)
	conn, err := grpc.Dial(targetServer, grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxMsgSize)))
	if err != nil {
		b.Errorf("fail to dial: %v", err)
	}
	defer conn.Close()

	// prepare context
	ctx, cancel := context.WithTimeout(context.Background(), proxyConnectionTimeout)
	defer cancel()

	// new proxy client
	client := proxy.NewTrafficProxyClient(conn)

	// test
	startTime := time.Now()
	defer func() {
		fmt.Printf("Processing for getting traffic flow by GRPC takes %f seconds, loop count %d\n",
			time.Now().Sub(startTime).Seconds(), b.N)
	}()

	var req proxy.TrafficRequest
	req.TrafficSource = new(proxy.TrafficSource)
	req.TrafficSource.Region = flags.trafficProxyFlags.region
	req.TrafficSource.TrafficProvider = flags.trafficProxyFlags.trafficProvider
	req.TrafficSource.MapProvider = flags.trafficProxyFlags.mapProvider
	way := new(proxy.TrafficRequest_WayId)
	way.WayId = 753683232
	req.WayIdFields = way
	for i := 0; i < b.N; i++ {
		resp, err := client.GetFlows(ctx, &req)
		if err != nil {
			b.Errorf("GetFlows failed, err: %v.\n", err)
		}
		if i == 0 { // print once
			fmt.Println(resp.GetFlow())
		}
	}
}

func TestGetFlowsByGRPCStreaming(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	flowsChan := make(chan []*proxy.Flow)

	startTime := time.Now()
	statisticsInterval := 120 //120 seconds
	var totalFlowsCount, maxFlowsCount, minFlowsCount int64
	var recvCount int
	for flows := range flowsChan {
		recvCount++

		currFlowsCount := int64(len(flows))
		totalFlowsCount += currFlowsCount
		if currFlowsCount > maxFlowsCount {
			maxFlowsCount = currFlowsCount
		}
		if currFlowsCount < minFlowsCount {
			minFlowsCount = currFlowsCount
		}

		if time.Now().Sub(startTime).Seconds() >= float64(statisticsInterval) {
			fmt.Printf("received flows from grpc streaming in %f seconds, recv count %d, total got flows count: %d, max per recv: %d, min per recv: %d\n",
				time.Now().Sub(startTime).Seconds(), recvCount, totalFlowsCount, maxFlowsCount, minFlowsCount)
			quickViewFlows(flows, 5)

			recvCount = 0
			totalFlowsCount = 0
			maxFlowsCount = 0
			minFlowsCount = 0
			startTime = time.Now()
		}
	}

	go func() {
		err := getFlowsByGRPCStreaming(flags.trafficProxyFlags, flowsChan)
		if err != nil {
			t.Errorf("getFlowsByGRPCStreaming failed, err: %v", err)
		}
	}()
}
