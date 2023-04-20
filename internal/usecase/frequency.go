package usecase

import (
	"log"
	"sync"
	"time"

	"github.com/muhammadtaufan/go-sensor/internal/domain"
)

type FrequencyUpdater struct {
	ticker    *time.Ticker
	mutex     sync.Mutex
	newTicker chan struct{}
}

func NewFrequencyUpdater() domain.FrequencyUpdater {
	return &FrequencyUpdater{
		newTicker: make(chan struct{}, 1),
	}
}

func (fu *FrequencyUpdater) Start(sensorDataSender domain.PostSensorData) {
	fu.setFrequency(1000)

	go func() {
		for {
			select {
			case <-fu.newTicker:
				log.Println("Received newTicker signal")
				continue
			case <-fu.ticker.C:
				log.Println("Received tick from ticker")
				sensorDataSender.SendData()
			}
		}
	}()
}

func (fu *FrequencyUpdater) UpdateFrequency(millis int) {
	fu.setFrequency(millis)
	fu.newTicker <- struct{}{}
}

func (fu *FrequencyUpdater) setFrequency(milliseconds int) {
	fu.mutex.Lock()
	defer fu.mutex.Unlock()

	if fu.ticker != nil {
		fu.ticker.Stop()
	}
	fu.ticker = time.NewTicker(time.Duration(milliseconds) * time.Millisecond)
}
