package math3d

import (
	"testing"
	"math"
	"fmt"
)

func compare(v1, v2 float32, prec int) int {
	pow := math.Pow(10, float64(prec))
    i1 := math.Trunc(float64(v1) * pow)
	i2 := math.Trunc(float64(v2) * pow)
	
	fmt.Printf("Comparing {%1.7f, %1.7f}\n", i1, i2)

    if i1 < i2 {
        return -1
    } else if i1 > i2{
        return 1
    }

    return 0
}

func compareArray(t *testing.T, values, expected [3]float32, prec int) {
	for i := 0; i < 3; i++ {
		value := compare(expected[i], values[i], prec)
		if value != 0 {
			t.Error( fmt.Sprintf("DOT[%d]: Expected %3.6f but found %3.6f ", i, expected[i], values[i]) )
		}
	}
}

func Test_Add(t *testing.T) {
	u := &Vector{ [3]float32{ 1.0, 2.0, 3.0 } }
	v := &Vector{ [3]float32{ 5.0, 6.0, 7.0 } }
	j := u.Add(v)
	
	expected := [3]float32{ 6.0, 8.0, 10.0 }
	compareArray(t, j.values, expected, 5)
}

func Test_Dot(t *testing.T) {

	u := Vector{ [3]float32{ 1.0,  2.0  } }
	v := Vector{ [3]float32{ 10.0, 20.0 } }
	f := u.Dot(&v)
	
	if f != 50.0 {
		t.Error("Dot product test 1 failed!")
	} else {
		t.Log("Dot product test 1 passed")
	}

	u = Vector{ [3]float32{ 1.0,  2.0,  3.0  } }
	v = Vector{ [3]float32{ 10.0, 20.0, 30.0 } }
	
	f = u.Dot(&v)
	
	if f != 140.0 {
		t.Error("Dot product test 2 failed!")
	} else {
		t.Log("Dot product test 2 passed")
	}
	
}

func Test_Normalize(t *testing.T) {
	
	u := Vector{ [3]float32{ 1.0, 2.0, 3.0 } }
	expected := [3]float32{ 0.267261241912424, 0.534522483824849, 0.801783725737273 }
	v := u.Normalize()
	compareArray(t, v.values, expected, 3)
	
	u = Vector{ [3]float32{ 40.0, 50.0, 60.0 } }
	expected = [3]float32{ 0.455842305838552, 0.56980288229819, 0.683763458757828 }
	v = u.Normalize()
	compareArray(t, v.values, expected, 3)
	
	u = Vector{ [3]float32{ 700.0, 800.0, 900.0 } }
	expected = [3]float32{ 0.502570711032417, 0.57436652689419, 0.646162342755964 }
	v = u.Normalize()
	compareArray(t, v.values, expected, 3)
	
	u = Vector{}
	expected = [3]float32{}
	v = u.Normalize()
	compareArray(t, v.values, expected, 3)
	
}

func Test_CrossProduct(t *testing.T) {
	u := &Vector{ [3]float32{ 1.0, 2.0, 3.0 } }
	v := &Vector{ [3]float32{ 3.0, 4.0, 5.0 } }
	
	j := u.Cross(v)
	expected := [3]float32{ -2.0, 4.0, -2.0 }
	compareArray(t, j.values, expected, 3)
	
	u = &Vector{ [3]float32{ 10.0, 20.0, 30.0 } }
	v = &Vector{ [3]float32{ 300.0, 400.0, 512.0 } }
	
	j = u.Cross(v)
	expected = [3]float32{ -1760.0, 3880.0, -2000.0 }
	compareArray(t, j.values, expected, 3)	
}

