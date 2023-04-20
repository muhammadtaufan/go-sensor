package main

import (
	"github.com/muhammadtaufan/go-sensor/internal/delivery"
	"github.com/muhammadtaufan/go-sensor/internal/usecase"
)

const (
	address = "0.0.0.0:50051"
)

func main() {
	serverConn, grpcClient := delivery.NewGRPCClient(address)
	defer serverConn.Close()

	postSensorData := usecase.NewPostSensorData(grpcClient)

	frequencyUpdater := usecase.NewFrequencyUpdater()

	apiServer := delivery.NewAPIServer()
	apiServer.RegisterFrequencyUpdater(frequencyUpdater)

	frequencyUpdater.Start(postSensorData)
	apiServer.Start()
}
