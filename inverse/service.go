package inverse

import (
	"errors"
	"go-microservices/matrix"
)

var invalidMatrixError = errors.New("Determinant of that matrix cannot be found since it is not a square matrix.")

type InverseService interface {
	GetInverse(*matrix.MatrixHolder) (float64, error)
}

type inverseService struct{}

func (inverseService) GetInverse(matrix *matrix.MatrixHolder) (float64, error) {
	return matrix.CalculateDeterminant(), nil
}

func NewService() InverseService {
	return &inverseService{}
}
