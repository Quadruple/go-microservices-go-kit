package svd

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type svdRequest struct {
	Matrix [][]float64 `json:"matrix"`
}

type svdResponse struct {
	SvdMatrix [][]float64 `json:"svdMatrix"`
	Err       string      `json:"err,omitempty"`
}

func MakeSvdEndpoint(svc SvdService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(svdRequest)
		v, err := svc.GetSingularValueDecomposition(req.Matrix)
		if err != nil {
			return svdResponse{v, err.Error()}, nil
		}
		return svdResponse{v, ""}, nil
	}
}
