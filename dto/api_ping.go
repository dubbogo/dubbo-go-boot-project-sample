package dto

import (
	hessian "github.com/apache/dubbo-go-hessian2"
)

type ApiPingRequestDTO struct {
}

func (*ApiPingRequestDTO) JavaClassName() string {
	return "com.chansos.dubbogo.dto.ApiPingRequestDTO"
}

type ApiPingResponseDTO struct {
	Result string `json:"result"`
}

func (*ApiPingResponseDTO) JavaClassName() string {
	return "com.chansos.dubbogo.dto.ApiPingResponseDTO"
}

func init() {
	hessian.RegisterPOJO(&ApiPingRequestDTO{})
	hessian.RegisterPOJO(&ApiPingResponseDTO{})
}
