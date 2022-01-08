package multiplication

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type multiplicationRequest struct {
	Matrix [][]float64 `json:"matrix"`
}

type multiplicationResponse struct {
	MultipliedMatrix [][]float64 `json:"svdMatrix"`
	Err              string      `json:"err,omitempty"`
}

func MakeMultiplicationEndpoint(svc MultiplicationService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(multiplicationRequest)
		v, err := svc.GetMultipliedMatrix(req.Matrix)
		if err != nil {
			return multiplicationResponse{v, err.Error()}, nil
		}
		return multiplicationResponse{v, ""}, nil
	}
}
