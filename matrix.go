package math3d

import (
	"fmt"
	"runtime"
)

type Matrix struct {
	values [16]float32
}

var indexes = [4][4]int {
	[4]int{ 0,  1,  2,  3  },
	[4]int{ 4,  5,  6,  7  },
	[4]int{ 8,  9,  10, 11 },
	[4]int{ 12, 13, 14, 15 },
}

var parallelism int

func init() {
	num := runtime.NumCPU()
	
	switch {
		case num > 4 :
			parallelism = 4
		case num < 1 :
			parallelism = 1
		default :
			parallelism = num
	}
}

func NewMatrix(values [16]float32) *Matrix {
	r := new(Matrix)
	r.values = values
	return r
}

func Identity() *Matrix {
	m := new(Matrix)
	m.values = [16]float32 {
			1.0, 0.0, 0.0, 0.0,
			0.0, 1.0, 0.0, 0.0,
			0.0, 0.0, 1.0, 0.0,
			0.0, 0.0, 0.0, 1.0,
		}
	return m
}

func (m *Matrix) Print() {
	fmt.Printf("[")
	for i := 0; i < 4; i++ {
		fmt.Printf("\n\t")
		for j := 0; j < 4; j++ {
			if j > 0 {
				fmt.Printf(", ")
			}
			fmt.Printf("%4.2f", m.values[indexes[i][j]])
		}
	}
	fmt.Printf("\n]\n")
}

func (m *Matrix) SetValues(values [16]float32) {
	m.values = values
}

func MultiplyMatrices(n ... *Matrix) *Matrix {
	r := n[0]
	for i := 1; i < len(n); i++ {
		r = r.Multiply(n[i])
	}
	return r
}

func (m *Matrix) NaiveMultiply(n *Matrix) *Matrix {
	
	r := new(Matrix)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			r.values[indexes[i][j]] = m.values[indexes[i][0]] * n.values[indexes[0][j]] + 
										m.values[indexes[i][1]] * n.values[indexes[1][j]] + 
										m.values[indexes[i][2]] * n.values[indexes[2][j]] +
										m.values[indexes[i][3]] * n.values[indexes[3][j]]
		
		}
	}
	
	return r
}

func parallelHelper(c chan int, i int, m *Matrix, n *Matrix, r *Matrix) {
	for j := 0; j < 4; j++ {
		r.values[indexes[i][j]] = m.values[indexes[i][0]] * n.values[indexes[0][j]] + 
									m.values[indexes[i][1]] * n.values[indexes[1][j]] + 
									m.values[indexes[i][2]] * n.values[indexes[2][j]] +
									m.values[indexes[i][3]] * n.values[indexes[3][j]]
	}
	if c != nil {
		c <- 1
	}
}

func (m *Matrix) ParallelNaiveMultiply(n *Matrix) *Matrix {
	r := new(Matrix)
	c := make(chan int, parallelism)
	
	for i := 0; i < parallelism; i++ {
		go parallelHelper(c, i, m, n, r)
	}
	
	for i := parallelism; i < 4; i++ {
		parallelHelper(nil, i, m, n, r)
	}	
	
	for i := 0; i < parallelism; i++ {
		<-c
	}
	
	return r
}

func (m *Matrix) Multiply(n *Matrix) *Matrix {
	if n == nil {
		return m
	}
	return m.ParallelNaiveMultiply(n)
}