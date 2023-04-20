package delivery

import (
	"log"

	sensor "github.com/muhammadtaufan/go-sensor/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGRPCClient(address string) (*grpc.ClientConn, sensor.SensorServiceClient) {
	serverConn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect to gRPC server: %v", err)
	}

	return serverConn, sensor.NewSensorServiceClient(serverConn)
}
