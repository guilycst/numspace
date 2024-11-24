package algebra

import (
	"errors"
)

var (
	// ErrNotRectangular indicates that the input data is not a rectangular matrix.
	ErrNotRectangular = errors.New("matrix is not rectangular")

	// ErrNilMatrix is returned when a nil matrix is provided to an operation.
	ErrNilMatrix = errors.New("matrix is nil")

	// ErrInvalidDimensions indicates a mismatch in matrix dimensions during an operation.
	ErrInvalidDimensions = errors.New("invalid dimensions")

	// ErrorIndexOutOfBounds is returned when an invalid row or column index is accessed.
	ErrorIndexOutOfBounds = errors.New("index out of range")

	// ErrMulDimensions indicates that matrix dimensions are incompatible for multiplication.
	ErrMulDimensions = errors.New("invalid dimensions for multiplication")
)

// Matrix defines a general interface for matrix operations.
// Implementations must ensure correct handling of dimensions and error reporting.
type Matrix interface {
	// Rows returns the number of rows in the matrix.
	Rows() int

	// Cols returns the number of columns in the matrix.
	Cols() int

	// At retrieves the element at row i and column j.
	// Returns an error if the indices are out of bounds.
	At(i, j int) (float64, error)

	// MustAt retrieves the element at row i and column j, panicking if indices are out of bounds.
	MustAt(i, j int) float64

	// Empty checks whether the matrix has zero rows or columns.
	Empty() bool

	// Add performs element-wise addition with another matrix.
	// Returns an error if dimensions do not match.
	Add(Matrix) (Matrix, error)

	// Sub performs element-wise subtraction with another matrix.
	// Returns an error if dimensions do not match.
	Sub(Matrix) (Matrix, error)

	// ScalarMul multiplies each element of the matrix by the given scalar.
	ScalarMul(float64) (Matrix, error)

	// CompareDimensions checks if the dimensions of the current matrix match another matrix.
	CompareDimensions(Matrix) bool

	// Mul performs matrix multiplication with another matrix.
	// Returns an error if dimensions are incompatible.
	Mul(Matrix) (Matrix, error)

	// Transpose returns a new matrix that is the transpose of the current matrix.
	Transpose() Matrix
}

// FlatMatrix is a concrete implementation of the Matrix interface.
// It stores matrix data as a flat slice for efficient memory access and computation.
type FlatMatrix struct {
	// data is the flat slice that stores all matrix elements in row-major order.
	data []float64

	// rows is the number of rows in the matrix.
	rows int

	// cols is the number of columns in the matrix.
	cols int
}

func (m *FlatMatrix) Rows() int {
	return m.rows
}

func (m *FlatMatrix) Cols() int {
	return m.cols
}

func (m *FlatMatrix) At(i, j int) (float64, error) {
	if i < 0 || i >= m.rows || j < 0 || j >= m.cols {
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

	result := make([]float64, m.rows*m.cols)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			result[i*m.cols+j] = m.MustAt(i, j) + other.MustAt(i, j)
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

	result := make([]float64, m.rows*m.cols)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			result[i*m.cols+j] = m.MustAt(i, j) - other.MustAt(i, j)
		}
	}

	return NewMatrixFlat(result, m.rows, m.cols)
}

func (m *FlatMatrix) ScalarMul(scalar float64) (Matrix, error) {
	if m.Empty() {
		return NewMatrixZero(0, 0)
	}

	result := make([]float64, m.rows*m.cols)
	for i := 0; i < len(m.data); i++ {
		result[i] = m.data[i] * scalar
	}
	return NewMatrixFlat(result, m.rows, m.cols)
}

func (m *FlatMatrix) Mul(other Matrix) (Matrix, error) {
	if other == nil {
		return nil, ErrNilMatrix
	}

	if m.Cols() != other.Rows() {
		return nil, ErrMulDimensions
	}

	nm, err := NewMatrixZero(m.Rows(), other.Cols())
	if err != nil {
		panic(err)
	}

	result := nm.(*FlatMatrix)

	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < other.Cols(); j++ {
			var sum float64
			for k := 0; k < m.Cols(); k++ {
				sum += m.MustAt(i, k) * other.MustAt(k, j)
			}
			result.data[i*result.cols+j] = sum
		}
	}

	return result, nil
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

func (m *FlatMatrix) Transpose() Matrix {
	result, err := NewMatrixZero(m.Cols(), m.Rows())
	if err != nil {
		panic(err)
	}

	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			result.(*FlatMatrix).data[j*m.Rows()+i] = m.data[i*m.Cols()+j]
		}
	}
	return result
}

func NewMatrix(data [][]float64) (Matrix, error) {
	if !isRectangular(data) {
		return nil, ErrNotRectangular
	}

	if len(data) == 0 {
		return NewMatrixZero(0, 0)
	}

	return NewMatrixFlat(flatten(data), len(data), len(data[0]))
}

func NewMatrixFlat(data []float64, rows, cols int) (Matrix, error) {
	if err := validateConstructor(len(data), rows, cols); err != nil {
		return nil, err
	}

	return &FlatMatrix{
		data: data,
		rows: rows,
		cols: cols,
	}, nil
}

func NewMatrixZero(rows, cols int) (Matrix, error) {
	if err := validateConstructor(rows*cols, rows, cols); err != nil {
		return nil, err
	}
	return &FlatMatrix{
		data: make([]float64, rows*cols),
		rows: rows,
		cols: cols,
	}, nil
}

func validateConstructor(dataLen int, rows, cols int) error {
	if dataLen > 0 && dataLen != (rows*cols) {
		return ErrInvalidDimensions
	}
	if rows < 0 || cols < 0 {
		return ErrInvalidDimensions
	}
	return nil
}

func isRectangular(data [][]float64) bool {
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
