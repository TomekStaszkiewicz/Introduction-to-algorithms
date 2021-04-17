package chapter3

import (
	"math"
	"math/rand"
	"sort"
	"time"
)

// SortingPermutation creates a new table which is a permutted tab arrau
// permutation is done using the sorting + priority algorythms
func SortingPermutation(tab []int) []int {
	length := float64(len(tab))
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	newTab := PermutationTable{
		[]PermutationItem{},
	}

	for _, value := range tab {
		newTab.tab = append(newTab.tab, PermutationItem{
			value,
			r1.Intn(int(math.Pow(length, 3))),
		})
	}
	newTab.sort()

	return newTab.getTab()
}

type PermutationItem struct {
	value    int
	priority int
}

type PermutationTable struct {
	tab []PermutationItem
}

func (p PermutationTable) sort() {
	sort.Slice(p.tab, func(a, b int) bool {
		return p.tab[a].priority > p.tab[b].priority
	})
}

func (p PermutationTable) getTab() []int {
	result := []int{}

	for _, val := range p.tab {
		result = append(result, val.value)
	}

	return result
}
