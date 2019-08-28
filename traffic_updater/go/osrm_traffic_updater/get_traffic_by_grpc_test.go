package main

import (
	"log"
	"testing"
	"time"

	"github.com/Telenav/osrm-backend/traffic_updater/go/grpc/proxy"
)

func BenchmarkGetAllTrafficDataByGRPC(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping test in short mode.")
	}

	for i := 0; i < b.N; i++ {
		trafficData, err := getTrafficFlowsIncidentsByGRPC(flags.trafficProxyFlags, nil)
		if err != nil {
			b.Error(err)
		}
		quickViewFlows(trafficData.FlowResponses, 10)         //quick view first 10 lines
		quickViewIncidents(trafficData.IncidentResponses, 10) //quick view first 10 lines
	}
}

func TestGetTrafficDataForWaysByGRPC(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	var wayIds []int64
	wayIds = append(wayIds, 829733412, 104489539)

	trafficData, err := getTrafficFlowsIncidentsByGRPC(flags.trafficProxyFlags, wayIds)
	if err != nil {
		t.Error(err)
	}
	quickViewFlows(trafficData.FlowResponses, 10)         //quick view first 10 lines
	quickViewIncidents(trafficData.IncidentResponses, 10) //quick view first 10 lines
}

func TestGetDeltaTrafficDataByGRPCStreaming(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	trafficDataChan := make(chan proxy.TrafficResponse)

	go func() {
		err := getDeltaTrafficFlowsIncidentsByGRPCStreaming(flags.trafficProxyFlags, trafficDataChan)
		if err != nil {
			t.Errorf("getDeltaTrafficFlowsIncidentsByGRPCStreaming failed, err: %v", err)
		}
	}()

	startTime := time.Now()
	statisticsInterval := 120 //120 seconds
	var totalFlowsCount, maxFlowsCount, minFlowsCount int64
	var totalIncidentsCount, maxIncidentsCount, minIncidentsCount int64
	var recvCount int
	for trafficData := range trafficDataChan {
		recvCount++

		currFlowsCount := int64(len(trafficData.FlowResponses))
		totalFlowsCount += currFlowsCount
		if currFlowsCount > maxFlowsCount {
			maxFlowsCount = currFlowsCount
		}
		if minFlowsCount == 0 || currFlowsCount < minFlowsCount {
			minFlowsCount = currFlowsCount
		}

		currIncidentsCount := int64(len(trafficData.IncidentResponses))
		totalIncidentsCount += currFlowsCount
		if currIncidentsCount > maxIncidentsCount {
			maxIncidentsCount = currIncidentsCount
		}
		if minIncidentsCount == 0 || currIncidentsCount < minIncidentsCount {
			minIncidentsCount = currIncidentsCount
		}

		if time.Now().Sub(startTime).Seconds() >= float64(statisticsInterval) {
			log.Printf("received flows from grpc streaming in %f seconds, recv count %d, total got flows count: %d, max per recv: %d, min per recv: %d\n",
				time.Now().Sub(startTime).Seconds(), recvCount, totalFlowsCount, maxFlowsCount, minFlowsCount)
			log.Printf("received incidents from grpc streaming in %f seconds, recv count %d, total got incidents count: %d, max per recv: %d, min per recv: %d\n",
				time.Now().Sub(startTime).Seconds(), recvCount, totalIncidentsCount, maxIncidentsCount, minIncidentsCount)

			quickViewFlows(trafficData.FlowResponses, 5)
			quickViewIncidents(trafficData.IncidentResponses, 5)

			recvCount = 0
			totalFlowsCount = 0
			maxFlowsCount = 0
			minFlowsCount = 0
			totalIncidentsCount = 0
			maxIncidentsCount = 0
			minIncidentsCount = 0
			startTime = time.Now()
		}
	}
}
