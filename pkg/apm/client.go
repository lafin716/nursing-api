package apm

import (
	"github.com/pinpoint-apm/pinpoint-go-agent"
)

type PinpointClient interface {
	GetAgent() (pinpoint.Agent, error)
}

type PinpointConfig struct {
	AppName       string
	AgentId       string
	CollectorHost string
	CollectorPort int
}

type pinpointClient struct {
	config *PinpointConfig
}

func NewPinpointClient(config *PinpointConfig) PinpointClient {
	return &pinpointClient{
		config: config,
	}
}

func (m pinpointClient) GetAgent() (pinpoint.Agent, error) {
	opts := []pinpoint.ConfigOption{
		pinpoint.WithAppName(m.config.AppName),
		pinpoint.WithAgentId(m.config.AgentId),
		pinpoint.WithCollectorHost(m.config.CollectorHost),
		pinpoint.WithCollectorAgentPort(m.config.CollectorPort),
	}
	cfg, _ := pinpoint.NewConfig(opts...)
	return pinpoint.NewAgent(cfg)
}
