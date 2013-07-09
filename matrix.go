package math3d

import "fmt"

type Matrix struct {
	values [16]float32
}

var indexes = [4][4]int {
	[4]int{ 0,  1,  2,  3  },
	[4]int{ 4,  5,  6,  7  },
	[4]int{ 8,  9,  10, 11 },
	[4]int{ 12, 13, 14, 15 },
}

/*
var parallelism int

var functions = [4]func(m, n, p *Matrix) {
	NaiveMultiply,
	ParallelMultiply2,
	ParallelMultiply3,
	ParallelMultiply4,
}

func init() {
	num := runtime.NumCPU()
	fmt.Printf("Num CPU: %d\n", num)
	
	switch {
		case num > 4 :
			parallelism = 4
		case num < 1 :
			parallelism = 1
		default :
			parallelism = num
	}
	
	parallelism = 4
}

func SetParallelism(p int) {
	if p <= 0 {
		parallelism = 1
	
	} else if p > 4 {
		parallelism = 4

	} else {
		parallelism = p
	}
}

func GetParallelism() int {
	return parallelism
}
*/

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

func NaiveMultiply(m, n, p *Matrix) *Matrix {
	for i := 0; i < 4; i++ {
		calculateRow(i, m, n, p)
	}
	return p
}

func calculateRow(i int, m *Matrix, n *Matrix, r *Matrix) {
	for j := 0; j < 4; j++ {
		r.values[indexes[i][j]] = m.values[indexes[i][0]] * n.values[indexes[0][j]] + 
									m.values[indexes[i][1]] * n.values[indexes[1][j]] + 
									m.values[indexes[i][2]] * n.values[indexes[2][j]] +
									m.values[indexes[i][3]] * n.values[indexes[3][j]]
	}
}

/*
func parallelHelper(c chan int, i int, m *Matrix, n *Matrix, r *Matrix) {
	calculateRow(i, m, n, r)
	if c != nil {
		c <- 1
	}
}

func ParallelMultiply2(m, n, p *Matrix) {
	c := make(chan int, 2)
	go parallelHelper(c, 0, m, n, p)
	calculateRow(1, m, n, p)
	go parallelHelper(c, 2, m, n, p)
	calculateRow(3, m, n, p)
	<-c
	<-c
}

func ParallelMultiply3(m, n, p *Matrix) {
	c := make(chan int, 2)
	go parallelHelper(c, 0, m, n, p)
	go parallelHelper(c, 1, m, n, p)
	calculateRow(2, m, n, p)
	calculateRow(3, m, n, p)
	<-c
	<-c
}

func ParallelMultiply4(m, n, p *Matrix) {
	c := make(chan int, 3)
	go parallelHelper(c, 0, m, n, p)
	go parallelHelper(c, 1, m, n, p)
	go parallelHelper(c, 2, m, n, p)
	calculateRow(3, m, n, p)
	<-c
	<-c
	<-c
}
*/

func (m *Matrix) Multiply(n *Matrix) *Matrix {
	if n == nil {
		return m
	}
	return m.MultiplyP(n, new(Matrix))
}

func (m *Matrix) MultiplyP(n *Matrix, p *Matrix) *Matrix {
	if n != nil {
		NaiveMultiply(m, n, p)
	}
	
	return p
}