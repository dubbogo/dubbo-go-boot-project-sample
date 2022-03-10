package service

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/config"
	"fmt"
	"github.com/dubbogo/dubbo-go-boot-project-sample/dto"
)

var ApiPingService = &ApiPingConsumer{}

type ApiPingProvider struct {
}

func (*ApiPingProvider) MethodMapper() map[string]string {
	return map[string]string{
		"Ping": "ping",
	}
}

func (*ApiPingProvider) Ping(ctx context.Context, req *dto.ApiPingRequestDTO) (*dto.ApiPingResponseDTO, error) {
	if req == nil {
		return nil, fmt.Errorf("Invalid parameter")
	}
	return &dto.ApiPingResponseDTO{
		Result: ApiPing(),
	}, nil
}

func (*ApiPingProvider) Reference() string {
	return "ApiPingProvider"
}

type ApiPingConsumer struct {
	Ping func(ctx context.Context, req *dto.ApiPingRequestDTO) (*dto.ApiPingResponseDTO, error) `dubbo:"ping"`
}

func (*ApiPingConsumer) Reference() string {
	return "ApiPingConsumer"
}

func init() {
	config.SetProviderService(&ApiPingProvider{})
	config.SetConsumerService(ApiPingService)
}
