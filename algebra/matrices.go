package algebra

import (
	"errors"
)

var ErrNotRectangular = errors.New("matrix is not rectangular")
var ErrNilMatrix = errors.New("matrix is nil")
var ErrInvalidDimensions = errors.New("invalid dimensions")
var ErrorIndexOutOfBounds = errors.New("index out of range")

type Matrix interface {
	Rows() int
	Cols() int
	At(i, j int) (float64, error)
	MustAt(i, j int) float64
	Empty() bool
	Add(Matrix) (Matrix, error)
	Sub(Matrix) (Matrix, error)
	CompareDimensions(Matrix) bool
}

type FlatMatrix struct {
	data []float64
	rows int
	cols int
}

func (m *FlatMatrix) Rows() int {
	return m.rows
}

func (m *FlatMatrix) Cols() int {
	return m.cols
}

func (m *FlatMatrix) At(i, j int) (float64, error) {
	if m.rows-i <= 0 || m.cols-j <= 0 {
		return 0, ErrorIndexOutOfBounds
	}
	return m.data[i*m.cols+j], nil
}

func (m *FlatMatrix) Empty() bool {
	return m.rows == 0
}

func (m *FlatMatrix) Add(other Matrix) (Matrix, error) {
	if other == nil {
		return nil, ErrNilMatrix
	}

	if !m.CompareDimensions(other) {
		return nil, ErrInvalidDimensions
	}

	result := make([]float64, 0, m.rows*m.cols)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			result = append(result, m.MustAt(i, j)+other.MustAt(i, j))
		}
	}

	return NewMatrixFlat(result, m.rows, m.cols)
}

func (m *FlatMatrix) Sub(other Matrix) (Matrix, error) {
	if other == nil {
		return nil, ErrNilMatrix
	}

	if !m.CompareDimensions(other) {
		return nil, ErrInvalidDimensions
	}

	result := make([]float64, 0, m.rows*m.cols)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			result = append(result, m.MustAt(i, j)-other.MustAt(i, j))
		}
	}

	return NewMatrixFlat(result, m.rows, m.cols)
}

func (m *FlatMatrix) CompareDimensions(other Matrix) bool {
	if other == nil {
		return false
	}

	return m.Rows() == other.Rows() && m.Cols() == other.Cols()
}

func (m *FlatMatrix) MustAt(i, j int) (v float64) {
	var err error
	if v, err = m.At(i, j); err != nil {
		panic(err)
	}
	return
}

func NewMatrix(data [][]float64) (*FlatMatrix, error) {
	if !retangular(data) {
		return nil, ErrNotRectangular
	}

	if len(data) == 0 {
		return NewMatrixFlat([]float64{}, 0, 0)
	}

	return NewMatrixFlat(flatten(data), len(data), len(data[0]))
}

func NewMatrixFlat(data []float64, rows, cols int) (*FlatMatrix, error) {
	if len(data) != rows*cols {
		return nil, ErrNotRectangular
	}
	return &FlatMatrix{
		data: data,
		rows: rows,
		cols: cols,
	}, nil
}

func retangular(data [][]float64) bool {
	rows := len(data)
	if rows == 0 {
		return true
	}

	for _, row := range data {
		if len(row) != len(data[0]) {
			return false
		}
	}

	return true
}

func flatten(data [][]float64) []float64 {
	l := len(data)
	if l == 0 {
		return []float64{}
	}

	flat := make([]float64, 0, l*len(data[0]))
	for _, row := range data {
		flat = append(flat, row...)
	}
	return flat
}
