package inverse

import (
	"context"
	"go-microservices/matrix"

	"github.com/go-kit/kit/endpoint"
)

type inverseRequest struct {
	Matrix [][]float64 `json:"matrix"`
}

type inverseResponse struct {
	InvertedMatrix float64 `json:"invertedMatrix"`
	Err            string  `json:"err,omitempty"`
}

func MakeInverseEndpoint(svc InverseService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(inverseRequest)
		v, err := svc.GetInverse(matrix.ConvertSliceToMatrix(req.Matrix))
		if err != nil {
			return inverseResponse{v, err.Error()}, nil
		}
		return inverseResponse{v, ""}, nil
	}
}
