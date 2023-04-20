package usecase

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/muhammadtaufan/go-sensor/config"
	"github.com/muhammadtaufan/go-sensor/internal/domain"
	utils "github.com/muhammadtaufan/go-sensor/pkg"
	sensor "github.com/muhammadtaufan/go-sensor/proto"
)

type PostSensorData struct {
	grpcClient sensor.SensorServiceClient
	cfg        *config.Config
}

func NewPostSensorData(grpcClient sensor.SensorServiceClient, cfg *config.Config) domain.PostSensorData {
	return &PostSensorData{
		grpcClient: grpcClient,
		cfg:        cfg,
	}
}

func (ps *PostSensorData) SendData() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data := &sensor.SensorData{
		SensorValue: utils.GetRandomFloat(),
		SensorType:  ps.cfg.SensorType,
		ID1:         utils.GetRandomAlphabets(),
		ID2:         int32(rand.Intn(10-1+1) + 1),
		Timestamp:   time.Now().Unix(),
	}

	_, err := ps.grpcClient.SendSensorData(ctx, data)
	if err != nil {
		log.Printf("Could not send data: %v", err)
	}

	log.Println("Data sent")
}
