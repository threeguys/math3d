package math3d

import (
	"math"
)

const (
	x = 0
	y = 1
	z = 2
)

type Vector struct {
	values [3]float32
}

func NewVector(values [3]float32) *Vector {
	v := new(Vector)
	v.values = values
	return v
}

func (u *Vector) X() float32 {
	return u.values[x]
}

func (u *Vector) Y() float32 {
	return u.values[y]
}

func (u *Vector) Z() float32 {
	return u.values[z]
}

func (u *Vector) Length() float32 {
	return float32(math.Sqrt(float64((u.Dot(u)))))
}

func (u *Vector) Add(v *Vector) *Vector {
	j := new(Vector)
	j.values[x] = u.values[x] + v.values[x]
	j.values[y] = u.values[y] + v.values[y]
	j.values[z] = u.values[z] + v.values[z]
	return j
}

func (u *Vector) Normalize() *Vector {
	v := new(Vector)
	len := u.Length()
	
	if len > 0 {	
		v.values[x] = u.values[x] / len
		v.values[y] = u.values[y] / len
		v.values[z] = u.values[z] / len
	}
	
	return v
}

func (u *Vector) Dot(v *Vector) float32 {
	return u.values[x] * v.values[x] +
			u.values[y] * v.values[y] +
			u.values[z] * v.values[z]
}

func (u *Vector) Cross(v *Vector) *Vector {
	ret := new(Vector)
	
	// x.y * y.z - y.y * x.z,
	ret.values[x] = u.values[y] * v.values[z] - v.values[y] * u.values[z]
	
	// x.z * y.x - y.z * x.x,
	ret.values[y] = u.values[z] * v.values[x] - v.values[z] * u.values[x]
	
	// x.x * y.y - y.x * x.y
	ret.values[z] = u.values[x] * v.values[y] - v.values[x] * u.values[y]
	
	return ret
}
