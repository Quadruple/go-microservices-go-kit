package eigen

import "go-microservices/matrix"

type EigenValueService interface {
	GetEigenValues(*matrix.MatrixHolder) ([][]float64, bool)
}

type eigenValueService struct{}

func (eigenValueService) GetEigenValues(matrixHolder *matrix.MatrixHolder) ([][]float64, bool) {
	return matrixHolder.CalculateEigenValues()
}

func NewService() EigenValueService {
	return &eigenValueService{}
}
