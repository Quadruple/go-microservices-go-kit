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

func getDimensions(data [][]float64) (int, int) {
	return len(data), len(data[0])
}
