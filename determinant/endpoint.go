package determinant

import (
	"context"
	"go-microservices/matrix"

	"github.com/go-kit/kit/endpoint"
)

type determinantRequest struct {
	Matrix [][]float64 `json:"matrix"`
}

type determinantResponse struct {
	Determinant float64 `json:"determinant"`
	Err         string  `json:"err,omitempty"`
}

func MakeUppercaseEndpoint(svc DeterminantService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(determinantRequest)
		v, err := svc.GetDeterminant(matrix.ConvertSliceToMatrix(req.Matrix))
		if err != nil {
			return determinantResponse{v, err.Error()}, nil
		}
		return determinantResponse{v, ""}, nil
	}
}
