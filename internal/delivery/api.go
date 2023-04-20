package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/muhammadtaufan/go-sensor/internal/domain"
)

type APIServer struct {
	server *echo.Echo
}

type FrequencyRequest struct {
	Frequency int `json:"frequency"`
}

func NewAPIServer() *APIServer {
	return &APIServer{
		server: echo.New(),
	}
}

func (aps *APIServer) Start() {
	aps.server.Logger.Fatal(aps.server.Start(":3000"))
}

func (aps *APIServer) RegisterFrequencyUpdater(fu domain.FrequencyUpdater) {
	aps.server.POST("/update-frequency", func(c echo.Context) error {
		req := new(FrequencyRequest)

		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid JSON format")
		}

		fu.UpdateFrequency(req.Frequency)
		return c.JSON(http.StatusOK, "Frequency Updated")
	})
}
