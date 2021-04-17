package chapter3

import (
	"math/rand"
	"time"
)

// PermuteInPlace permutes given array
func PermuteInPlace(tab []int) {
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range tab {
		randomIndex := r1.Intn(len(tab))
		tab[i], tab[randomIndex] = tab[randomIndex], tab[i]
	}
}
