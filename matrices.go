package junglemath

import "errors"

// Some logic was translated with permission from
// https://github.com/mike123ike/Matrix/blob/9897c9bdb6c73563f8ec55c7ad37119d1eb764db/matrix.py

/*
Logic (trimmed):

class Matrix:
    
    def multiply(self, other):
        if not isinstance(other, Matrix):
            result = Matrix(self.rows, self.columns)
            for row in range(self.rows):
                for num in range(self.columns):
                    result.matrix[row][num] = self.matrix[row][num]*other
            return result

        elif self.columns == other.rows:
            result = Matrix(self.rows, other.columns)
            for i in range(self.rows):
                for j in range(other.columns):
                    for k in range(other.rows):
                        result.matrix[i][j] += self.matrix[i][k] * other.matrix[k][j]
            return result

    def transpose(self):
        result = Matrix(self.columns, self.rows)
        for i in range(self.rows):
            for j in range(self.columns):
                result.matrix[j][i] = self.matrix[i][j]
        return result

    def get_cofactor(self, row, column):
        copy = self.duplicate()
        for i in range(copy.rows):
            copy.matrix[i].pop(column)
        copy.matrix.pop(row)
        sign = (-1)**(row+column)
        cofactor = sign * copy.get_determinant()
        del copy
        return cofactor

    def get_determinant(self):
        matrix = self.matrix
        self.update()
        if self.columns == self.rows == 1:
            return self.matrix[0][0]
        elif self.columns == self.rows == 2:
            determinant = (matrix[0][0]*matrix[1][1]) - (matrix[1][0]*matrix[0][1])
            return determinant
        elif self.columns == self.rows:
            total = 0
            for i in range(self.columns):
                a = matrix[0][i]
                total += a * self.get_cofactor(0, i)
            return total

    def get_adjoint(self):
        cofactor_matrix = Matrix(self.rows, self.columns)
        for i in range(self.rows):
            for j in range(self.columns):
                cofactor_matrix.matrix[i][j] = self.get_cofactor(i, j)
        adjoint_matrix = cofactor_matrix.duplicate().transpose()
        return adjoint_matrix

    def invert(self):
        if self.rows == self.columns:
            determinant = self.get_determinant()
            if determinant != 0:
                inverse = self.get_adjoint().multiply(1/determinant)
                inverse.update()
                return inverse

    def exponent(self, power):
        if self.rows == self.columns:
            result = self.duplicate()
            for i in range(power-1):
                result = self.multiply(result)
            return result
*/

func (m *Matrix) SetVal(row, column int, val float64) {
	m.Matrix[column][row] = val
}

func AddMatrices(m1, m2 Matrix) (Matrix, error) {
    // Adds two matrices

    if equal, err := m1.EqualDims(m2); !equal {
        if err != nil {
            return Matrix{}, err
        }
		return Matrix{}, errors.New("error: matrices not of equal size")
    }
    res := Matrix{Matrix: make([][]float64, 0)}
    for i, row := range m1.Matrix {
        for x, val := range row {
            res.Matrix[i][x] = val + m2.Matrix[i][x]
        }
    }
    return res, nil
}

func MultiplyMatrices(m1, m2 Matrix) (Matrix, error) {
    // Multiplies two matrices

    if equal, err := m1.EqualRows(m2); !equal {
        if err != nil {
            return Matrix{}, err
        }
        return Matrix{}, errors.New("error: matrices do not have equal rows")
    }
    res := Matrix{Matrix: make([][]float64, 0)}
    if len(m2.Matrix[0]) > len(m1.Matrix[0]) {
        temp := m1
        m1 = m2
        m2 = temp
    }
    return res, nil
}

func (m *Matrix) EqualRows(other Matrix) (bool, error) {
	// Check if 2 matrices have the same amount of rows
	// Used to check if matrices can be multiplied, etc.
	// Returns false and an error if a matrix is not IsValid()

	if !m.IsValid() || !other.IsValid() {
        return false, errors.New("error: invalid matrix")
    }

    if len(m.Matrix) != len(other.Matrix) {
        return false, nil
    }

    return true, nil 
}

func (m *Matrix) EqualCols(other Matrix) (bool, error) {
    // Check if 2 matrices have the same amount of columns
	// Used to check if matrices can be added, multiplied, etc.
	// Returns false and an error if a matrix is not IsValid()

	if !m.IsValid() || !other.IsValid() {
        return false, errors.New("error: invalid matrix")
    }

    if len(m.Matrix[0]) != len(other.Matrix[0]) {
        return false, nil
    }

    return true, nil 
}

func (m *Matrix) EqualDims(other Matrix) (bool, error) {
    // Check if 2 matrices have the same dimensions
	// Used to check if matrices can be added, multiplied, etc.
	// Returns false and an error if a matrix is not IsValid()

    if !m.IsValid() || !other.IsValid() {
        return false, errors.New("error: invalid matrix")
    }

    equalCols, err := m.EqualCols(other)
    if err != nil {
        return false, err
    }
    if !equalCols {
        return false, nil
    }

    equalRows, err := m.EqualRows(other)
    if err != nil {
        return false, err
    }
    if !equalRows {
        return false, nil
    }

    return true, nil 
}

func (m *Matrix) IsValid() bool {
	// Checks if a matrix is valid
	// Matrix is invalid if not all rows have the same amount of values

	prev := len(m.Matrix[0])
	for _, val := range(m.Matrix) {
		if prev != len(val) { return false }
	}
    return true
}

func (m *Matrix) Add(other Matrix) error {
    // Method to add a matrices
    // Sets the matrix this is called on to the result

	sum, err := AddMatrices(*m, other)
	if err == nil {
		*m = sum
	}
	return err
}