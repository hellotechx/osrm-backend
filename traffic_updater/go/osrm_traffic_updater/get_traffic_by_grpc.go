package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/Telenav/osrm-backend/traffic_updater/go/grpc/proxy"
	"google.golang.org/grpc"
)

const (
	proxyConnectionTimeout = 60 * time.Second
	maxMsgSize             = 1024 * 1024 * 1024
)

func quickViewFlows(flows []*proxy.Flow, viewCount int) {
	for i := 0; i < viewCount && i < len(flows); i++ {
		fmt.Println(flows[i])
	}
}

func getAllFlowsByGRPC(f trafficProxyFlags) ([]*proxy.Flow, error) {

	startTime := time.Now()
	defer func() {
		fmt.Printf("Processing time for getting traffic flows takes %f seconds\n", time.Now().Sub(startTime).Seconds())
	}()

	// make RPC client
	targetServer := f.ip + ":" + strconv.Itoa(f.port)
	fmt.Println("connect traffic proxy " + targetServer)
	conn, err := grpc.Dial(targetServer, grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxMsgSize)))
	if err != nil {
		return nil, fmt.Errorf("fail to dial: %v", err)
	}
	defer conn.Close()

	// prepare context
	ctx, cancel := context.WithTimeout(context.Background(), proxyConnectionTimeout)
	defer cancel()

	// new proxy client
	client := proxy.NewTrafficProxyClient(conn)

	// get flows
	fmt.Println("getting flows")
	var req proxy.TrafficRequest
	req.TrafficSource = new(proxy.TrafficSource)
	req.TrafficSource.Region = f.region
	req.TrafficSource.TrafficProvider = f.trafficProvider
	req.TrafficSource.MapProvider = f.mapProvider
	ways := new(proxy.TrafficRequest_All)
	ways.All = true
	req.WayIdFields = ways
	resp, err := client.GetFlows(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("GetFlows failed, err: %v.\n", err)
	}
	fmt.Printf("GetFlows succeed, code: %d, msg: %s, got flows count: %d\n",
		resp.GetCode(), resp.GetMsg(), len(resp.GetFlows().Flows))

	return resp.GetFlows().Flows, nil
}
