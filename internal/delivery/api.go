package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/muhammadtaufan/go-sensor/config"
	"github.com/muhammadtaufan/go-sensor/internal/domain"
)

type APIServer struct {
	server *echo.Echo
	cfg    *config.Config
}

type FrequencyRequest struct {
	Frequency int `json:"frequency"`
}

func NewAPIServer(cfg *config.Config) *APIServer {
	return &APIServer{
		server: echo.New(),
		cfg:    cfg,
	}
}

func (aps *APIServer) Start() {
	aps.server.Logger.Fatal(aps.server.Start(aps.cfg.APIerverAddress))
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
