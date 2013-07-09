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
	[4]int{ 0,  1,  2,  3  },
	[4]int{ 4,  5,  6,  7  },
	[4]int{ 8,  9,  10, 11 },
	[4]int{ 12, 13, 14, 15 },
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

	r := DegreesToRadians64(float64(fovy) / 2.0)
	deltaZ := zFar - zNear
	s := math.Sin(r)

	if (deltaZ == 0 || s == 0 || aspect == 0) {
		return nil
	}
	
	cotangent := float32(math.Cos(r) / s)
	
	result := Identity()
	result.values[0] = cotangent / aspect
	result.values[5] = cotangent
	result.values[10] = - (zFar + zNear) / deltaZ;
	result.values[11] = -float32(1.0);
	result.values[14] = (-float32(2.0) * zNear * zFar) / deltaZ;
	result.values[15] = 0.0
	result.Print()
	return result
/*
	r := DegreesToRadians64(float64(fovy)) / 2.0
	cotanHalfFovy := float32(1.0 / math.Tan(r))
	
	result := new(Matrix)
	result.values[indexes[0][0]] = cotanHalfFovy / aspect
	result.values[indexes[1][1]] = cotanHalfFovy
	
	result.values[indexes[2][3]] = -1.0
	
	result.values[indexes[2][2]] = (zFar + zNear) / (zNear - zFar)
	result.values[indexes[2][3]] = (2.0 * zFar * zNear) / (zNear - zFar)
	
	
	result.Print()
	//result.values[15] = 0.0
	return result
		*/
	/*
	xymax := float64(zNear) * math.Tan(float64(fovy) * math.Pi / 360.0)
	ymin := -xymax
	xmin := -xymax

	width := xymax - xmin
	height := xymax - ymin

	depth := zFar - zNear
	q := -(zFar + zNear) / depth
	qn := -2 * (zFar * zNear) / depth

	w := 2 * float64(zNear) / width
	w = w / float64(aspect)
	h := 2 * float64(zNear) / height

	m := new(Matrix)
	m.values[0]  = float32(w)
	m.values[1]  = 0
	m.values[2]  = 0
	m.values[3]  = 0

	m.values[4]  = 0
	m.values[5]  = float32(h)
	m.values[6]  = 0
	m.values[7]  = 0

	m.values[8]  = 0
	m.values[9]  = 0
	m.values[10] = q
	m.values[11] = -1

	m.values[12] = 0
	m.values[13] = 0
	m.values[14] = qn
	m.values[15] = 0
	m.Print()
	return m
	*/
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
	ret.values[0] = s.values[0]
	ret.values[4] = s.values[1]
	ret.values[8] = s.values[2]
	
	ret.values[1] = u.values[0]
	ret.values[5] = u.values[1]
	ret.values[9] = u.values[2]
	
	ret.values[2] = -f.values[0]
	ret.values[6] = -f.values[1]
	ret.values[10] = -f.values[2]

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

func (m *Matrix) Pointer() unsafe.Pointer {
	return unsafe.Pointer(&m.values[0])
}

func NaiveMultiply(m, n, p *Matrix) *Matrix {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			p.values[indexes[i][j]] = m.values[indexes[i][0]] * n.values[indexes[0][j]] + 
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
	}
	
	return p
}

