package inverse

import (
	"go-microservices/matrix"
)

type InverseService interface {
	GetInverse(*matrix.MatrixHolder) ([][]float64, error)
}

type inverseService struct{}

func (inverseService) GetInverse(matrix *matrix.MatrixHolder) ([][]float64, error) {
	return matrix.GetInverse()
}

func NewService() InverseService {
	return &inverseService{}
}
