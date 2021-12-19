package eigen

import "go-microservices/matrix"

type EigenValueService interface {
	GetEigenValues(*matrix.MatrixHolder) ([]float64, error)
}

type eigenValueService struct{}

func (eigenValueService) GetEigenValues(matrixHolder *matrix.MatrixHolder) ([]float64, error) {
	return matrixHolder.CalculateEigenValues(), nil
}

func NewService() EigenValueService {
	return &eigenValueService{}
}
