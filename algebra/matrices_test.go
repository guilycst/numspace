package algebra

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFlatMatrix_At(t *testing.T) {
	type fields struct {
		data []float64
		rows int
		cols int
	}
	type args struct {
		i int
		j int
	}

	type testdef struct {
		name   string
		fields fields
		args   args
		want   float64
		err    error
	}

	tests := []testdef{
		{
			name: "Test empty at 0,0 should return 0",
			fields: fields{
				data: []float64{},
				rows: 0,
				cols: 0,
			},
			args: args{
				i: 0,
				j: 0,
			},
			want: 0,
		},
		{
			name: "Test 1x1 at 0,0 should return 1",
			fields: fields{
				data: []float64{1},
				rows: 1,
				cols: 1,
			},
			args: args{
				i: 0,
				j: 0,
			},
			want: 1,
		},
		{
			name: "Test 4x4 at 4,4 should return ErrIndexOutOfRange",
			fields: fields{
				data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
				rows: 4,
				cols: 4,
			},
			args: args{
				i: 4,
				j: 4,
			},
			want: 0,
			err:  ErrorIndexOutOfBounds,
		},
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			tests = append(tests, testdef{
				name: fmt.Sprintf("Test 4x4 at %d,%d should return %d", i, j, i*4+j+1),
				fields: fields{
					data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
					rows: 4,
					cols: 4,
				},
				args: args{
					i: i,
					j: j,
				},
				want: float64(i*4 + j + 1),
			})
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &FlatMatrix{
				data: tt.fields.data,
				rows: tt.fields.rows,
				cols: tt.fields.cols,
			}

			got, err := m.At(tt.args.i, tt.args.j)
			if tt.err != nil && err != tt.err {
				t.Errorf("FlatMatrix.At() = %v, want %v", err, tt.err)
			}
			if got != tt.want {
				t.Errorf("FlatMatrix.At() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlatMatrix_Empty(t *testing.T) {
	type fields struct {
		data []float64
		rows int
		cols int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Test empty matrix should return true",
			fields: fields{
				data: []float64{},
				rows: 0,
				cols: 0,
			},
			want: true,
		},
		{
			name: "Test non-empty 1x1 matrix should return false",
			fields: fields{
				data: []float64{1},
				rows: 1,
				cols: 1,
			},
			want: false,
		},
		{
			name: "Test non-empty 2x2 matrix should return false",
			fields: fields{
				data: []float64{1, 2, 3, 4},
				rows: 2,
				cols: 2,
			},
			want: false,
		},
		{
			name: "Test non-empty 3x3 matrix should return false",
			fields: fields{
				data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
				rows: 3,
				cols: 3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &FlatMatrix{
				data: tt.fields.data,
				rows: tt.fields.rows,
				cols: tt.fields.cols,
			}
			if got := m.Empty(); got != tt.want {
				t.Errorf("FlatMatrix.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlatMatrix_Add(t *testing.T) {
	type fields struct {
		data []float64
		rows int
		cols int
	}
	type args struct {
		other Matrix
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Matrix
		wantErr bool
	}{
		{
			name: "Test adding two 2x2 matrices",
			fields: fields{
				data: []float64{1, 2, 3, 4},
				rows: 2,
				cols: 2,
			},
			args: args{
				other: &FlatMatrix{
					data: []float64{5, 6, 7, 8},
					rows: 2,
					cols: 2,
				},
			},
			want: &FlatMatrix{
				data: []float64{6, 8, 10, 12},
				rows: 2,
				cols: 2,
			},
			wantErr: false,
		},
		{
			name: "Test adding two 3x3 matrices",
			fields: fields{
				data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
				rows: 3,
				cols: 3,
			},
			args: args{
				other: &FlatMatrix{
					data: []float64{9, 8, 7, 6, 5, 4, 3, 2, 1},
					rows: 3,
					cols: 3,
				},
			},
			want: &FlatMatrix{
				data: []float64{10, 10, 10, 10, 10, 10, 10, 10, 10},
				rows: 3,
				cols: 3,
			},
			wantErr: false,
		},
		{
			name: "Test adding matrices with different dimensions should return error",
			fields: fields{
				data: []float64{1, 2, 3, 4},
				rows: 2,
				cols: 2,
			},
			args: args{
				other: &FlatMatrix{
					data: []float64{1, 2, 3},
					rows: 1,
					cols: 3,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test adding nil matrix should return error",
			fields: fields{
				data: []float64{1, 2, 3, 4},
				rows: 2,
				cols: 2,
			},
			args: args{
				other: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &FlatMatrix{
				data: tt.fields.data,
				rows: tt.fields.rows,
				cols: tt.fields.cols,
			}
			got, err := m.Add(tt.args.other)
			if (err != nil) != tt.wantErr {
				t.Errorf("FlatMatrix.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlatMatrix.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlatMatrix_Sub(t *testing.T) {
	type fields struct {
		data []float64
		rows int
		cols int
	}
	type args struct {
		other Matrix
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Matrix
		wantErr bool
	}{
		{
			name: "Test subtracting two 2x2 matrices",
			fields: fields{
				data: []float64{5, 6, 7, 8},
				rows: 2,
				cols: 2,
			},
			args: args{
				other: &FlatMatrix{
					data: []float64{1, 2, 3, 4},
					rows: 2,
					cols: 2,
				},
			},
			want: &FlatMatrix{
				data: []float64{4, 4, 4, 4},
				rows: 2,
				cols: 2,
			},
			wantErr: false,
		},
		{
			name: "Test subtracting two 3x3 matrices",
			fields: fields{
				data: []float64{9, 8, 7, 6, 5, 4, 3, 2, 1},
				rows: 3,
				cols: 3,
			},
			args: args{
				other: &FlatMatrix{
					data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
					rows: 3,
					cols: 3,
				},
			},
			want: &FlatMatrix{
				data: []float64{8, 6, 4, 2, 0, -2, -4, -6, -8},
				rows: 3,
				cols: 3,
			},
			wantErr: false,
		},
		{
			name: "Test subtracting matrices with different dimensions should return error",
			fields: fields{
				data: []float64{1, 2, 3, 4},
				rows: 2,
				cols: 2,
			},
			args: args{
				other: &FlatMatrix{
					data: []float64{1, 2, 3},
					rows: 1,
					cols: 3,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test subtracting nil matrix should return error",
			fields: fields{
				data: []float64{1, 2, 3, 4},
				rows: 2,
				cols: 2,
			},
			args: args{
				other: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &FlatMatrix{
				data: tt.fields.data,
				rows: tt.fields.rows,
				cols: tt.fields.cols,
			}
			got, err := m.Sub(tt.args.other)
			if (err != nil) != tt.wantErr {
				t.Errorf("FlatMatrix.Sub() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlatMatrix.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlatMatrix_ScalarMul(t *testing.T) {
	type fields struct {
		data []float64
		rows int
		cols int
	}
	type args struct {
		scalar float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Matrix
		wantErr bool
	}{
		{
			name: "Test scalar multiplication of 2x2 matrix by 2",
			fields: fields{
				data: []float64{1, 2, 3, 4},
				rows: 2,
				cols: 2,
			},
			args: args{
				scalar: 2,
			},
			want: &FlatMatrix{
				data: []float64{2, 4, 6, 8},
				rows: 2,
				cols: 2,
			},
			wantErr: false,
		},
		{
			name: "Test scalar multiplication of 3x3 matrix by 0.5",
			fields: fields{
				data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
				rows: 3,
				cols: 3,
			},
			args: args{
				scalar: 0.5,
			},
			want: &FlatMatrix{
				data: []float64{0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5},
				rows: 3,
				cols: 3,
			},
			wantErr: false,
		},
		{
			name: "Test scalar multiplication of 2x2 matrix by 0",
			fields: fields{
				data: []float64{1, 2, 3, 4},
				rows: 2,
				cols: 2,
			},
			args: args{
				scalar: 0,
			},
			want: &FlatMatrix{
				data: []float64{0, 0, 0, 0},
				rows: 2,
				cols: 2,
			},
			wantErr: false,
		},
		{
			name: "Test scalar multiplication of empty matrix",
			fields: fields{
				data: []float64{},
				rows: 0,
				cols: 0,
			},
			args: args{
				scalar: 5,
			},
			want: &FlatMatrix{
				data: []float64{},
				rows: 0,
				cols: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &FlatMatrix{
				data: tt.fields.data,
				rows: tt.fields.rows,
				cols: tt.fields.cols,
			}
			got, err := m.ScalarMul(tt.args.scalar)
			if (err != nil) != tt.wantErr {
				t.Errorf("FlatMatrix.ScalarMul() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlatMatrix.ScalarMul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlatMatrix_Mul(t *testing.T) {
	type fields struct {
		data []float64
		rows int
		cols int
	}
	type args struct {
		other Matrix
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Matrix
		wantErr bool
	}{
		{
			name: "Test multiplying two 2x2 matrices",
			fields: fields{
				data: []float64{1, 2, 3, 4},
				rows: 2,
				cols: 2,
			},
			args: args{
				other: &FlatMatrix{
					data: []float64{5, 6, 7, 8},
					rows: 2,
					cols: 2,
				},
			},
			want: &FlatMatrix{
				data: []float64{19, 22, 43, 50},
				rows: 2,
				cols: 2,
			},
			wantErr: false,
		},
		{
			name: "Test multiplying two 3x3 matrices",
			fields: fields{
				data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
				rows: 3,
				cols: 3,
			},
			args: args{
				other: &FlatMatrix{
					data: []float64{9, 8, 7, 6, 5, 4, 3, 2, 1},
					rows: 3,
					cols: 3,
				},
			},
			want: &FlatMatrix{
				data: []float64{30, 24, 18, 84, 69, 54, 138, 114, 90},
				rows: 3,
				cols: 3,
			},
			wantErr: false,
		},
		{
			name: "Test multiplying matrices with incompatible dimensions should return error",
			fields: fields{
				data: []float64{1, 2, 3, 4},
				rows: 2,
				cols: 2,
			},
			args: args{
				other: &FlatMatrix{
					data: []float64{1, 2, 3},
					rows: 1,
					cols: 3,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test multiplying with nil matrix should return error",
			fields: fields{
				data: []float64{1, 2, 3, 4},
				rows: 2,
				cols: 2,
			},
			args: args{
				other: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test multiplying empty matrices",
			fields: fields{
				data: []float64{},
				rows: 0,
				cols: 0,
			},
			args: args{
				other: &FlatMatrix{
					data: []float64{},
					rows: 0,
					cols: 0,
				},
			},
			want: &FlatMatrix{
				data: []float64{},
				rows: 0,
				cols: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &FlatMatrix{
				data: tt.fields.data,
				rows: tt.fields.rows,
				cols: tt.fields.cols,
			}
			got, err := m.Mul(tt.args.other)
			if (err != nil) != tt.wantErr {
				t.Errorf("FlatMatrix.Mul() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlatMatrix.Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlatMatrix_Transpose(t *testing.T) {
	type fields struct {
		data []float64
		rows int
		cols int
	}
	tests := []struct {
		name   string
		fields fields
		want   Matrix
	}{
		{
			name: "Test transpose of 2x2 matrix",
			fields: fields{
				data: []float64{1, 2, 3, 4},
				rows: 2,
				cols: 2,
			},
			want: &FlatMatrix{
				data: []float64{1, 3, 2, 4},
				rows: 2,
				cols: 2,
			},
		},
		{
			name: "Test transpose of 3x3 matrix",
			fields: fields{
				data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
				rows: 3,
				cols: 3,
			},
			want: &FlatMatrix{
				data: []float64{1, 4, 7, 2, 5, 8, 3, 6, 9},
				rows: 3,
				cols: 3,
			},
		},
		{
			name: "Test transpose of 2x3 matrix",
			fields: fields{
				data: []float64{1, 2, 3, 4, 5, 6},
				rows: 2,
				cols: 3,
			},
			want: &FlatMatrix{
				data: []float64{1, 4, 2, 5, 3, 6},
				rows: 3,
				cols: 2,
			},
		},
		{
			name: "Test transpose of 3x2 matrix",
			fields: fields{
				data: []float64{1, 2, 3, 4, 5, 6},
				rows: 3,
				cols: 2,
			},
			want: &FlatMatrix{
				data: []float64{1, 3, 5, 2, 4, 6},
				rows: 2,
				cols: 3,
			},
		},
		{
			name: "Test transpose of empty matrix",
			fields: fields{
				data: []float64{},
				rows: 0,
				cols: 0,
			},
			want: &FlatMatrix{
				data: []float64{},
				rows: 0,
				cols: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &FlatMatrix{
				data: tt.fields.data,
				rows: tt.fields.rows,
				cols: tt.fields.cols,
			}
			if got := m.Transpose(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlatMatrix.Transpose() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlatMatrix_CompareDimensions(t *testing.T) {
	type fields struct {
		data []float64
		rows int
		cols int
	}
	type args struct {
		other Matrix
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Test comparing dimensions of two 2x2 matrices should return true",
			fields: fields{
				data: []float64{1, 2, 3, 4},
				rows: 2,
				cols: 2,
			},
			args: args{
				other: &FlatMatrix{
					data: []float64{5, 6, 7, 8},
					rows: 2,
					cols: 2,
				},
			},
			want: true,
		},
		{
			name: "Test comparing dimensions of 2x2 and 3x3 matrices should return false",
			fields: fields{
				data: []float64{1, 2, 3, 4},
				rows: 2,
				cols: 2,
			},
			args: args{
				other: &FlatMatrix{
					data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
					rows: 3,
					cols: 3,
				},
			},
			want: false,
		},
		{
			name: "Test comparing dimensions with nil matrix should return false",
			fields: fields{
				data: []float64{1, 2, 3, 4},
				rows: 2,
				cols: 2,
			},
			args: args{
				other: nil,
			},
			want: false,
		},
		{
			name: "Test comparing dimensions of two 3x3 matrices should return true",
			fields: fields{
				data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
				rows: 3,
				cols: 3,
			},
			args: args{
				other: &FlatMatrix{
					data: []float64{9, 8, 7, 6, 5, 4, 3, 2, 1},
					rows: 3,
					cols: 3,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &FlatMatrix{
				data: tt.fields.data,
				rows: tt.fields.rows,
				cols: tt.fields.cols,
			}
			if got := m.CompareDimensions(tt.args.other); got != tt.want {
				t.Errorf("FlatMatrix.CompareDimensions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMatrix(t *testing.T) {
	type args struct {
		data [][]float64
	}
	tests := []struct {
		name    string
		args    args
		want    Matrix
		wantErr bool
	}{
		{
			name: "Test creating 2x2 matrix",
			args: args{
				data: [][]float64{
					{1, 2},
					{3, 4},
				},
			},
			want: &FlatMatrix{
				data: []float64{1, 2, 3, 4},
				rows: 2,
				cols: 2,
			},
			wantErr: false,
		},
		{
			name: "Test creating 3x3 matrix",
			args: args{
				data: [][]float64{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: &FlatMatrix{
				data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
				rows: 3,
				cols: 3,
			},
			wantErr: false,
		},
		{
			name: "Test creating non-rectangular matrix should return error",
			args: args{
				data: [][]float64{
					{1, 2},
					{3, 4, 5},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test creating empty matrix",
			args: args{
				data: [][]float64{},
			},
			want: &FlatMatrix{
				data: []float64{},
				rows: 0,
				cols: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMatrix(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMatrix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMatrixFlat(t *testing.T) {
	type args struct {
		data []float64
		rows int
		cols int
	}
	tests := []struct {
		name    string
		args    args
		want    Matrix
		wantErr bool
	}{
		{
			name: "Test creating 2x2 matrix",
			args: args{
				data: []float64{1, 2, 3, 4},
				rows: 2,
				cols: 2,
			},
			want: &FlatMatrix{
				data: []float64{1, 2, 3, 4},
				rows: 2,
				cols: 2,
			},
			wantErr: false,
		},
		{
			name: "Test creating 3x3 matrix",
			args: args{
				data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
				rows: 3,
				cols: 3,
			},
			want: &FlatMatrix{
				data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
				rows: 3,
				cols: 3,
			},
			wantErr: false,
		},
		{
			name: "Test creating matrix with incorrect data length should return error",
			args: args{
				data: []float64{1, 2, 3},
				rows: 2,
				cols: 2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test creating empty matrix",
			args: args{
				data: []float64{},
				rows: 0,
				cols: 0,
			},
			want: &FlatMatrix{
				data: []float64{},
				rows: 0,
				cols: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMatrixFlat(tt.args.data, tt.args.rows, tt.args.cols)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMatrixFlat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMatrixFlat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_retangular(t *testing.T) {
	type args struct {
		data [][]float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test rectangular matrix 2x2 should return true",
			args: args{
				data: [][]float64{
					{1, 2},
					{3, 4},
				},
			},
			want: true,
		},
		{
			name: "Test rectangular matrix 3x3 should return true",
			args: args{
				data: [][]float64{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: true,
		},
		{
			name: "Test non-rectangular matrix should return false",
			args: args{
				data: [][]float64{
					{1, 2},
					{3, 4, 5},
				},
			},
			want: false,
		},
		{
			name: "Test empty matrix should return true",
			args: args{
				data: [][]float64{},
			},
			want: true,
		},
		{
			name: "Test single row matrix should return true",
			args: args{
				data: [][]float64{
					{1, 2, 3, 4},
				},
			},
			want: true,
		},
		{
			name: "Test single column matrix should return true",
			args: args{
				data: [][]float64{
					{1},
					{2},
					{3},
					{4},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRectangular(tt.args.data); got != tt.want {
				t.Errorf("retangular() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_flatten(t *testing.T) {
	type args struct {
		data [][]float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "Test flatten 2x2 matrix",
			args: args{
				data: [][]float64{
					{1, 2},
					{3, 4},
				},
			},
			want: []float64{1, 2, 3, 4},
		},
		{
			name: "Test flatten 3x3 matrix",
			args: args{
				data: [][]float64{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "Test flatten non-rectangular matrix",
			args: args{
				data: [][]float64{
					{1, 2},
					{3, 4, 5},
				},
			},
			want: []float64{1, 2, 3, 4, 5},
		},
		{
			name: "Test flatten empty matrix",
			args: args{
				data: [][]float64{},
			},
			want: []float64{},
		},
		{
			name: "Test flatten single row matrix",
			args: args{
				data: [][]float64{
					{1, 2, 3, 4},
				},
			},
			want: []float64{1, 2, 3, 4},
		},
		{
			name: "Test flatten single column matrix",
			args: args{
				data: [][]float64{
					{1},
					{2},
					{3},
					{4},
				},
			},
			want: []float64{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flatten(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flatten() = %v, want %v", got, tt.want)
			}
		})
	}
}
