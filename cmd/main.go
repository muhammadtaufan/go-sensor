package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	utils "github.com/muhammadtaufan/go-sensor/pkg"
	sensor "github.com/muhammadtaufan/go-sensor/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "0.0.0.0:50051"
)

func main() {
	serverConn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect to gRPC server: %v", err)
	}
	defer serverConn.Close()

	c := sensor.NewSensorServiceClient(serverConn)

	// use ticker for the frequency to send data
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		sendData(c)
	}
}

func sendData(c sensor.SensorServiceClient) {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	data := &sensor.SensorData{
		SensorValue: utils.GetRandomFloat(),
		SensorType:  "Temperature",
		ID1:         utils.GetRandomAlphabets(),
		ID2:         int32(rand.Intn(10-1+1) + 1),
		Timestamp:   time.Now().Unix(),
	}

	_, err := c.SendSensorData(ctx, data)
	if err != nil {
		log.Fatalf("Could not send data: %v", err)
	}

	log.Println("Data sent")
}
