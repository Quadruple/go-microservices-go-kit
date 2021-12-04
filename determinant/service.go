package determinant

import (
	"errors"
	"go-microservices/matrix"
)

var invalidMatrixError = errors.New("Determinant of that matrix cannot be found since it is not a square matrix.")

type DeterminantService interface {
	GetDeterminant(*matrix.MatrixHolder) (float64, error)
}

func (determinantService) GetDeterminant(matrix *matrix.MatrixHolder) (float64, error) {
	return matrix.CalculateDeterminant(), nil
}

type determinantService struct{}

func NewService() DeterminantService {
	return &determinantService{}
}
