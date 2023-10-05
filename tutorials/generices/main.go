package main

import "log"

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func main() {

	ints := map[string]int64{"a": 12, "b": 15, "c": 26}
	floats := map[string]float64{"a": 12.2, "b": 15.7, "c": 26.9}
	log.Printf("ints=%d and floats=%d", SumInts(ints), SumFloats(floats))
	log.Println("sumsFloats:", SumFloats(floats))

}
