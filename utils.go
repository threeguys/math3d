package math3d

import "math"

const (
	piOver180 = math.Pi / 180.0
	piUnder180 = 180.0 / math.Pi
)

func RadiansToDegrees(radians float32) float32 {
	return radians * piUnder180
}

func RadiansToDegrees64(radians float64) float64 {
	return radians * piUnder180
}

func DegreesToRadians(degrees float32) float32 {
	return degrees * piOver180
}

func DegreesToRadians64(degrees float64) float64 {
	return degrees * piOver180
}

func Clamp(x, min, max float32) float32 {
	switch {
		case x < min :
			return min
			
		case x > max :
			return max
			
		default :
			return x
	}
}