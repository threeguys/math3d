package math3d

import (
	"fmt"
	"math"
	"unsafe"
)

type Matrix struct {
	values [16]float32
}

var indexes = [4][4]int {
	[4]int{ 0, 4,  8, 12 },
	[4]int{ 1, 5,  9, 13 },
	[4]int{ 2, 6, 10, 14 },
	[4]int{ 3, 7, 11, 15 },
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

func Perspective(fovy float32, aspect float32, zNear float32, zFar float32) *Matrix {
	r := DegreesToRadians64(float64(fovy) * 0.5) 
	scale := float32(1.0 / math.Tan(r))
	
	result := new(Matrix)
	result.values[indexes[0][0]] = scale / aspect
	result.values[indexes[1][1]] = scale
	
	result.values[indexes[2][3]] = -1.0
	
	result.values[indexes[2][2]] = (zFar + zNear) / (zNear - zFar)
	result.values[indexes[3][2]] = (2.0 * zFar * zNear) / (zNear - zFar)
	
	
	result.Print()
	return result
}

func Translation(v *Vector) *Matrix {
	r := Identity()
	r.values[12] = v.values[0]
	r.values[13] = v.values[1]
	r.values[14] = v.values[2]
	return r
}

func LookAt(eye *Vector, center *Vector, up *Vector) *Matrix {
	f := center.Subtract(eye)
	f = f.Normalize()
	
	u := up.Normalize()
	
	s := f.Cross(u)
	s = s.Normalize()
	
	u = s.Cross(f)
	u = u.Normalize()

	ret := Matrix{}
	ret.values[indexes[0][0]] = s.values[0]
	ret.values[indexes[0][1]] = s.values[1]
	ret.values[indexes[0][2]] = s.values[2]
	
	ret.values[indexes[1][0]] = u.values[0]
	ret.values[indexes[1][1]] = u.values[1]
	ret.values[indexes[1][2]] = u.values[2]
	
	ret.values[indexes[2][0]] = -f.values[0]
	ret.values[indexes[2][1]] = -f.values[1]
	ret.values[indexes[2][2]] = -f.values[2]

	translate := Translation(NewVector([3]float32{ -eye.values[0], -eye.values[1], -eye.values[2] }))
	return ret.Multiply(translate)
}

func (m *Matrix) Print() {
	fmt.Printf("[")
	for i := 0; i < 4; i++ {
		fmt.Printf("\n\t")
		for j := 0; j < 4; j++ {
			if j > 0 {
				fmt.Printf(", ")
			}
			fmt.Printf("%4.3f", m.values[indexes[i][j]])
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

func (m *Matrix) Pointer() unsafe.Pointer {
	return unsafe.Pointer(&m.values[0])
}

func NaiveMultiply(m, n, p *Matrix) *Matrix {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			p.values[indexes[i][j]] = 
					m.values[indexes[i][0]] * n.values[indexes[0][j]] +  
					m.values[indexes[i][1]] * n.values[indexes[1][j]] + 
					m.values[indexes[i][2]] * n.values[indexes[2][j]] + 
					m.values[indexes[i][3]] * n.values[indexes[3][j]]
		}
	}
	return p
}

func unrolledMultiply(m1, m2, mat *Matrix) *Matrix {
	mat.values[0] = m1.values[0] * m2.values[0] + m1.values[4] * m2.values[1] + m1.values[8] * m2.values[2] + m1.values[12] * m2.values[3]
	mat.values[1] = m1.values[1] * m2.values[0] + m1.values[5] * m2.values[1] + m1.values[9] * m2.values[2] + m1.values[13] * m2.values[3];
	mat.values[2] = m1.values[2] * m2.values[0] + m1.values[6] * m2.values[1] + m1.values[10] * m2.values[2] + m1.values[14] * m2.values[3];
	mat.values[3] = m1.values[3] * m2.values[0] + m1.values[7] * m2.values[1] + m1.values[11] * m2.values[2] + m1.values[15] * m2.values[3];

	mat.values[4] = m1.values[0] * m2.values[4] + m1.values[4] * m2.values[5] + m1.values[8] * m2.values[6] + m1.values[12] * m2.values[7];
	mat.values[5] = m1.values[1] * m2.values[4] + m1.values[5] * m2.values[5] + m1.values[9] * m2.values[6] + m1.values[13] * m2.values[7];
	mat.values[6] = m1.values[2] * m2.values[4] + m1.values[6] * m2.values[5] + m1.values[10] * m2.values[6] + m1.values[14] * m2.values[7];
	mat.values[7] = m1.values[3] * m2.values[4] + m1.values[7] * m2.values[5] + m1.values[11] * m2.values[6] + m1.values[15] * m2.values[7];

	mat.values[8] = m1.values[0] * m2.values[8] + m1.values[4] * m2.values[9] + m1.values[8] * m2.values[10] + m1.values[12] * m2.values[11];
	mat.values[9] = m1.values[1] * m2.values[8] + m1.values[5] * m2.values[9] + m1.values[9] * m2.values[10] + m1.values[13] * m2.values[11];
	mat.values[10] = m1.values[2] * m2.values[8] + m1.values[6] * m2.values[9] + m1.values[10] * m2.values[10] + m1.values[14] * m2.values[11];
	mat.values[11] = m1.values[3] * m2.values[8] + m1.values[7] * m2.values[9] + m1.values[11] * m2.values[10] + m1.values[15] * m2.values[11];

	mat.values[12] = m1.values[0] * m2.values[12] + m1.values[4] * m2.values[13] + m1.values[8] * m2.values[14] + m1.values[12] * m2.values[15];
	mat.values[13] = m1.values[1] * m2.values[12] + m1.values[5] * m2.values[13] + m1.values[9] * m2.values[14] + m1.values[13] * m2.values[15];
	mat.values[14] = m1.values[2] * m2.values[12] + m1.values[6] * m2.values[13] + m1.values[10] * m2.values[14] + m1.values[14] * m2.values[15];
	mat.values[15] = m1.values[3] * m2.values[12] + m1.values[7] * m2.values[13] + m1.values[11] * m2.values[14] + m1.values[15] * m2.values[15];
	return mat
}

func (m *Matrix) Multiply(n *Matrix) *Matrix {
	if n == nil {
		return m
	}
	return m.MultiplyP(n, new(Matrix))
}

func (m *Matrix) MultiplyP(n *Matrix, p *Matrix) *Matrix {
	if n != nil {
		unrolledMultiply(m, n, p)
		//NaiveMultiply(m, n, p)
	}
	
	return p
}

