package matrix

import (
	"gonum.org/v1/gonum/mat"
)

type MatrixHolder struct {
	Matrix mat.Matrix
}

type MatrixOperations interface {
	ConvertSliceToMatrix([][]float64) *MatrixHolder
	CalculateDeterminant() float64
	CalculateEigenValues() ([]complex128, error)
	GetInverse() [][]float64
}

func ConvertSliceToMatrix(data [][]float64) *MatrixHolder {
	row, col := getDimensions(data)
	emptyMatrix := mat.NewDense(row, col, nil)
	for index, row := range data {
		emptyMatrix.SetRow(index, row)
	}
	return &MatrixHolder{
		Matrix: emptyMatrix,
	}
}

func (matrixHolder *MatrixHolder) CalculateDeterminant() float64 {
	return mat.Det(matrixHolder.Matrix)
}

func (matrixHolder *MatrixHolder) CalculateEigenValues() ([][]float64, bool) {
	var eigenValues mat.Eigen
	err := eigenValues.Factorize(matrixHolder.Matrix, mat.EigenNone)
	if !err {
		return nil, err
	}
	return convertComplexSliceToFloatSlice(eigenValues.Values(nil)), false
}

func (matrixHolder *MatrixHolder) GetInverse() ([][]float64, error) {
	var invertedMatrix mat.Dense
	err := invertedMatrix.Inverse(matrixHolder.Matrix)
	if err != nil {
		return nil, err
	}
	return convertDenseToFloat(invertedMatrix), nil
}

func getDimensions(data [][]float64) (int, int) {
	return len(data), len(data[0])
}

func convertDenseToFloat(data mat.Dense) [][]float64 {
	row, col := data.Dims()
	convertedSlice := [][]float64{}
	for i := 0; i < row; i++ {
		tempSlice := []float64{}
		for j := 0; j < col; j++ {
			tempSlice = append(tempSlice, data.At(i, j))
		}
		convertedSlice = append(convertedSlice, tempSlice)
	}
	return convertedSlice
}

func convertComplexSliceToFloatSlice(data []complex128) [][]float64 {
	floatSlice := [][]float64{}
	for index := range data {
		floatSlice = append(floatSlice, []float64{real(data[index]), imag(data[index])})
	}
	return floatSlice
}
