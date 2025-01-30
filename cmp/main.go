package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"sort"
)

func main() {
	trans := cmp.Transformer("Sort", func(in []int) []int {
		out := append([]int(nil), in...)
		sort.Ints(out)
		return out
	})
	x := struct {
		Ints []int
	}{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}
	y := struct {
		Ints []int
	}{[]int{2, 8, 0, 9, 6, 1, 4, 7, 3, 5}}
	z := struct {
		Ints []int
	}{[]int{0, 0, 1, 2, 3, 4, 5, 6, 7, 8}}
	fmt.Println(cmp.Equal(x, y, trans))
	fmt.Println(cmp.Equal(y, z, trans))
	fmt.Println(cmp.Equal(z, x, trans))

}
