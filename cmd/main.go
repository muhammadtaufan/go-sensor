package main

import (
	"github.com/muhammadtaufan/go-sensor/config"
	"github.com/muhammadtaufan/go-sensor/internal/delivery"
	"github.com/muhammadtaufan/go-sensor/internal/usecase"
)

func main() {
	appConfig := config.LoadConfig()

	serverConn, grpcClient := delivery.NewGRPCClient(appConfig.GRPCServerAddress)
	defer serverConn.Close()

	postSensorData := usecase.NewPostSensorData(grpcClient, appConfig)

	frequencyUpdater := usecase.NewFrequencyUpdater()

	apiServer := delivery.NewAPIServer()
	apiServer.RegisterFrequencyUpdater(frequencyUpdater)

	frequencyUpdater.Start(postSensorData)
	apiServer.Start()
}
