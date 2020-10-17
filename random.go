package gorand

import (
	"math/rand"
	"time"
)

//RandIntn generates new seed and returns random integer within range
func RandIntn(stop int) int {
	x1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.New(x1)
	return y1.Intn(stop)
}

//RandRange generates a seed and returns a random element from the specified range
func RandRange(start int, stop int) int {

	x1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.New(x1)
	return y1.Intn(start-stop) + stop
}

//RandIntc generates a seed and returns a random element from an integer slice
func RandIntc(lst []int) int {
	x1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.New(x1)
	truelength := len(lst) - 1
	return lst[y1.Intn(truelength)]
}

//RandStringc generates a seed and returns a random element from a string slice
func RandStringc(lst []string) string {
	x1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.New(x1)
	truelength := len(lst) - 1
	return lst[y1.Intn(truelength)]
}

//RandFloat64c generates a seed and returns a random element from an float64 slice
func RandFloat64c(lst []float64) float64 {
	x1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.New(x1)
	truelength := len(lst) - 1
	return lst[y1.Intn(truelength)]
}

//RandFloat32c generates a seed and returns a random element from an float32 slice
func RandFloat32c(lst []float32) float32 {
	x1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.New(x1)
	truelength := len(lst) - 1
	return lst[y1.Intn(truelength)]
}
