package math3d

import (
	"fmt"
	"math"
)

const (
	x = 0
	y = 1
	z = 2
)

var positiveX = [3]float32{ 1.0, 0.0, 0.0 }
var positiveY = [3]float32{ 0.0, 1.0, 0.0 }
var positiveZ = [3]float32{ 0.0, 0.0, 1.0 }

var negativeX = [3]float32{ -1.0,  0.0,  0.0 }
var negativeY = [3]float32{  0.0, -1.0,  0.0 }
var negativeZ = [3]float32{  0.0,  0.0, -1.0 }

// A Vector consists of three components, this is implemented using a 3 element float32 array
// for ease of use with OpenGL
type Vector struct {
	values [3]float32
}

// Creates a new instance of Vector, specifying the components 
func NewVector(values [3]float32) *Vector {
	u := new(Vector)
	u.values = values
	return u
}

// Sets the component values for the vector in array form
func (u *Vector) SetValues(values [3]float32) {
	u.values = values
}

// Sets the individual components for the vector
func (u *Vector) SetComponents(x, y, z float32) {
	u.values[0] = x
	u.values[1] = y
	u.values[2] = z
}

// Sets all components of the vector to zero 
func (u *Vector) Zero() {
	u.SetComponents(0.0, 0.0, 0.0)
}

// Returns the X component of the vector
func (u *Vector) X() float32 {
	return u.values[x]
}

// Returns the Y component of the vector
func (u *Vector) Y() float32 {
	return u.values[y]
}

// Returns the Z component of the vector
func (u *Vector) Z() float32 {
	return u.values[z]
}

// Returns the length of the vector
func (u *Vector) Length() float32 {
	return float32(math.Sqrt(float64((u.Dot(u)))))
}

// Adds v to the target vector, returns a new instance of Vector containing the sum
func (u *Vector) Add(v *Vector) *Vector {
	return u.AddP(v, new(Vector))
}

// Adds v to the target vector and places the result into p, returns p
func (u *Vector) AddP(v, p *Vector) *Vector {
	p.values[x] = u.values[x] + v.values[x]
	p.values[y] = u.values[y] + v.values[y]
	p.values[z] = u.values[z] + v.values[z]
	return p
}

// Subtracts v from the target Vector and returns a new instance of Vector containing the subtraction
func (u *Vector) Subtract(v *Vector) *Vector {
	return u.SubtractP(v, new(Vector))
}

// Subtracts v from the target Vector and places the subtraction value into p, returns p
func (u *Vector) SubtractP(v, p *Vector) *Vector {
	p.values[x] = u.values[x] - v.values[x]
	p.values[y] = u.values[y] - v.values[y]
	p.values[z] = u.values[z] - v.values[z]
	return p
}

// Returns a new instance of Vector where the return value is the normalized form of the target Vector
func (u *Vector) Normalize() *Vector {
	return u.NormalizeP(new(Vector))
}

// Calculates the normalized form of the target vector and places the answer into p, returns p
func (u *Vector) NormalizeP(p *Vector) *Vector {
	len := u.Length()
	
	if len > 0 {	
		p.values[x] = u.values[x] / len
		p.values[y] = u.values[y] / len
		p.values[z] = u.values[z] / len
	}
	
	return p
}

// Calculates the Dot product of the target vector and v (u . v)
func (u *Vector) Dot(v *Vector) float32 {
	return u.values[x] * v.values[x] +
			u.values[y] * v.values[y] +
			u.values[z] * v.values[z]
}

// Calculates the Cross product of the target vector and v (u x v) and returns a new instance of Vector with the calculated value
func (u *Vector) Cross(v *Vector) *Vector {
	return u.CrossP(v, new(Vector))
}

// Calculates the Cross product of the target vector and v (u x v) and places the calculated value into p, returns p
func (u *Vector) CrossP(v *Vector, p *Vector) *Vector {
	// x.y * y.z - y.y * x.z,
	p.values[x] = u.values[y] * v.values[z] - v.values[y] * u.values[z]
	
	// x.z * y.x - y.z * x.x,
	p.values[y] = u.values[z] * v.values[x] - v.values[z] * u.values[x]
	
	// x.x * y.y - y.x * x.y
	p.values[z] = u.values[x] * v.values[y] - v.values[x] * u.values[y]
	
	return p
}

// Multiplies the target vector by the Matrix m and returns a new instance of Vector with the multiplied value
func (u *Vector) Multiply(m *Matrix) *Vector {
	return u.MultiplyP(m, new(Vector))
}

// Multiplies the target vector by the Matrix m and places the answer into p, returns p
func (u *Vector) MultiplyP(m *Matrix, p *Vector) *Vector {
	
	p.values[x] = u.values[x] * m.values[indexes[0][0]] + 
					u.values[y] * m.values[indexes[0][1]] +
					u.values[z] * m.values[indexes[0][2]] + 
					m.values[indexes[0][3]];
					
    p.values[y] = u.values[x] * m.values[indexes[1][0]] + 
					u.values[y] * m.values[indexes[1][1]] + 
					u.values[z] * m.values[indexes[1][2]] + 
					m.values[indexes[1][3]];
					
    p.values[z] = u.values[x] * m.values[indexes[2][0]] + 
					u.values[y] * m.values[indexes[2][1]] + 
					u.values[z] * m.values[indexes[2][2]] + 
					m.values[indexes[2][3]];
	
	return p
}

// Scales the target vector and returns a new instance of Vector containing the scaled vector
func (u *Vector) Scale(factor float32) *Vector {
	return u.ScaleP(factor, new(Vector))
}

// Scales the target vector and places the answer into p, returns p
func (u *Vector) ScaleP(factor float32, p *Vector) *Vector {
	p.values[x] = u.values[x] * factor
	p.values[y] = u.values[y] * factor
	p.values[z] = u.values[z] * factor
	return p
}

func (u *Vector) Print() {
	fmt.Printf("%4.3f", u.values)
}

