package svd

import (
	"go-microservices/matrix"
)

type SvdService interface {
	GetSingularValueDecomposition([][]float64) ([][]float64, error)
}

type svdService struct{}

type SvdServiceMiddleware func(SvdService) SvdService

func (svdService) GetSingularValueDecomposition(data [][]float64) ([][]float64, error) {
	return matrix.ConvertSliceToMatrix(data).GetSingularValueDecomposition()
}

func NewService() SvdService {
	return &svdService{}
}
