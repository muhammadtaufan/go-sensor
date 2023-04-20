package domain

type FrequencyUpdater interface {
	UpdateFrequency(millis int)
	Start(sensorDataSender PostSensorData)
}
