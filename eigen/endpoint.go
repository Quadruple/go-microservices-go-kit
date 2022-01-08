package eigen

import (
	"context"
	"go-microservices/matrix"

	"github.com/go-kit/kit/endpoint"
)

type eigenValueRequest struct {
	Matrix [][]float64 `json:"matrix"`
}

type eigenValueResponse struct {
	EigenValues [][]float64 `json:"eigens"`
	Err         string      `json:"err"`
}

func MakeEigenValueEndpoint(eigenValueService EigenValueService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(eigenValueRequest)
		v, err := eigenValueService.GetEigenValues(matrix.ConvertSliceToMatrix(req.Matrix))
		if err != "Successful." {
			return eigenValueResponse{v, err}, nil
		}
		return eigenValueResponse{v, err}, nil
	}
}
