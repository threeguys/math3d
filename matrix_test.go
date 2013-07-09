package math3d

import (
	"fmt"
	"testing"
)

func (m *Matrix) equals(n *Matrix) bool {
	for i := 0; i < 16; i++ {
		if m.values[i] != n.values[i] {
			return false
		}
	}
	
	return true
}

func Test_Identity(t *testing.T) {
	m := Identity()
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			value := m.values[indexes[i][j]]
			if i == j && value != 1.0 {
				t.Error(fmt.Sprintf("Expected 1.0 in identity (%d,%d) but value was %f", i, j, value))
			} else if i != j && value != 0.0 {
				t.Error(fmt.Sprintf("Expected 0.0 in identity (%d,%d) but value was %f", i, j, value))
			}
		}
	}
}

func Test_NaiveMultiply(t *testing.T) {
	m := new(Matrix)
	m.values = [16]float32 {
		1,  2,  3,  4,
		5,  6,  7,  8,
		9,  10, 11, 12,
		13, 14, 15, 16,
	}
	m.Print()
	
	n := new(Matrix)
	n.values = [16]float32 {
		17, 18, 19, 20,
		21, 22, 23, 24,
		25, 26, 27, 28,
		29, 30, 31, 32,
	}
	n.Print()
	
	r := m.NaiveMultiply(n)
	r.Print()
	
	vrfy := &Matrix{ [16]float32 {
			 250.0,  260.0,  270.0,  280.0,
			 618.0,  644.0,  670.0,  696.0,
			 986.0, 1028.0, 1070.0, 1112.0,
			1354.0, 1412.0, 1470.0, 1528.0,
		} }
		
	if !vrfy.equals(r) {
		t.Error("naiveMultiply was not correct!")
		fmt.Printf("=========== EXPECTED ===========\n")
		vrfy.Print()
		fmt.Printf("=========== ACTUAL =============\n")
		r.Print()
	}
}

func Test_ParallelNaiveMultiply(t *testing.T) {
	m := new(Matrix)
	m.values = [16]float32 {
		1,  2,  3,  4,
		5,  6,  7,  8,
		9,  10, 11, 12,
		13, 14, 15, 16,
	}
	m.Print()
	
	n := new(Matrix)
	n.values = [16]float32 {
		17, 18, 19, 20,
		21, 22, 23, 24,
		25, 26, 27, 28,
		29, 30, 31, 32,
	}
	n.Print()
	
	r := m.ParallelNaiveMultiply(n)
	r.Print()
	
	vrfy := &Matrix{ [16]float32 {
			 250.0,  260.0,  270.0,  280.0,
			 618.0,  644.0,  670.0,  696.0,
			 986.0, 1028.0, 1070.0, 1112.0,
			1354.0, 1412.0, 1470.0, 1528.0,
		} }
		
	if !vrfy.equals(r) {
		t.Error("naiveMultiply was not correct!")
		fmt.Printf("=========== EXPECTED ===========\n")
		vrfy.Print()
		fmt.Printf("=========== ACTUAL =============\n")
		r.Print()
	}
}

func Test_MultiplyMatrices(t *testing.T) {
	m := new(Matrix)
	m.values = [16]float32 {
		1,  2,  3,  4,
		5,  6,  7,  8,
		9,  10, 11, 12,
		13, 14, 15, 16,
	}
	m.Print()
	
	n := new(Matrix)
	n.values = [16]float32 {
		17, 18, 19, 20,
		21, 22, 23, 24,
		25, 26, 27, 28,
		29, 30, 31, 32,
	}
	n.Print()
	
	r := MultiplyMatrices(m, n)
	r.Print()
	
	vrfy := &Matrix{ [16]float32 {
			 250.0,  260.0,  270.0,  280.0,
			 618.0,  644.0,  670.0,  696.0,
			 986.0, 1028.0, 1070.0, 1112.0,
			1354.0, 1412.0, 1470.0, 1528.0,
		} }
		
	if !vrfy.equals(r) {
		t.Error("naiveMultiply was not correct!")
		fmt.Printf("=========== EXPECTED ===========\n")
		vrfy.Print()
		fmt.Printf("=========== ACTUAL =============\n")
		r.Print()
	}
}
