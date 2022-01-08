package multiplication

import (
	"go-microservices/matrix"
)

type MultiplicationService interface {
	GetMultipliedMatrix([][]float64) ([][]float64, error)
}

type multiplicationService struct{}

type MultiplicationServiceMiddleware func(MultiplicationService) MultiplicationService

func (multiplicationService) GetMultipliedMatrix(data [][]float64) ([][]float64, error) {
	return matrix.ConvertSliceToMatrix(data).GetSingularValueDecomposition()
}

func NewService() MultiplicationService {
	return &multiplicationService{}
}
