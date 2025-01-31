package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"math"
	"sort"
)

type Contact struct {
	Phone string
	Email string
}

type User struct {
	Name    string
	Age     int
	Contact *Contact
}

type UserWithUnExportedField struct {
	Name    string
	Age     int
	contact *Contact
}

func main() {
	// sortCmp()
	// structCmp()
	// UnExportedFieldCmp()
	// FloatCmp()
	// NilCmp()
	// SliceCmp()
	CustomEqual()

}

func sortCmp() {
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

// cmp.Equal可以比较指针指向的内容 ==不行
func structCmp() {
	u1 := User{Name: "lyb", Age: 18}
	u2 := User{Name: "lyb", Age: 18}
	fmt.Println("u1 == u2 ? ", u1 == u2)
	fmt.Println("u1 equals u2? ", cmp.Equal(u1, u2))
	c1 := &Contact{Phone: "123456789", Email: "lyb@example.com"}
	c2 := &Contact{Phone: "123456789", Email: "lyb@example.com"}
	u1.Contact = c1
	u2.Contact = c1
	fmt.Println("u1 == u2 with same pointer? ", u1 == u2)
	fmt.Println("u1 equals u2 with same pointer? ", cmp.Equal(u1, u2))
	u2.Contact = c2
	fmt.Println("u1 == u2 with different pointer? ", u1 == u2)
	fmt.Println("u1 equals u2 with different pointer? ", cmp.Equal(u1, u2))
}

// cmp.Equal默认不比较未导出字段
func UnExportedFieldCmp() {
	c1 := &Contact{Phone: "123456789", Email: "lyb@example.com"}
	c2 := &Contact{Phone: "123456789", Email: "lyb@example.com"}
	u1 := UserWithUnExportedField{Name: "lyb", Age: 18, contact: c1}
	u2 := UserWithUnExportedField{Name: "lyb", Age: 18, contact: c2}
	// cmpopts.IgnoreUnexported(UserWithUnExportedField{}) 表示忽略UserWithExportedField的直接未导出字段
	fmt.Println("u1 equals u2? ", cmp.Equal(u1, u2, cmp.AllowUnexported(UserWithUnExportedField{})))
}

type FloatPair struct {
	X float64
	Y float64
}

func FloatCmp() {
	p1 := FloatPair{X: math.NaN()}
	p2 := FloatPair{X: math.NaN()}
	fmt.Println("p1 equals p2? ", cmp.Equal(p1, p2))
	fmt.Println("p1 equals p2(cmpopts.EquateNaNs)? ", cmp.Equal(p1, p2, cmpopts.EquateNaNs()))
	f1 := 0.1
	f2 := 0.2
	f3 := 0.3
	p3 := FloatPair{X: f1 + f2}
	p4 := FloatPair{X: f3}
	fmt.Println("p3 equals p4? ", cmp.Equal(p3, p4))
	fmt.Println("p3 equals p4(cmpopts.EquateApprox)? ", cmp.Equal(p3, p4, cmpopts.EquateApprox(0.1, 0.001)))
	p5 := FloatPair{X: 0.1 + 0.2}
	p6 := FloatPair{X: 0.3}
	fmt.Println("p5 equals p6? ", cmp.Equal(p5, p6))
}

func NilCmp() {
	var s1 []int
	s2 := make([]int, 0)
	var m1 map[int]int
	m2 := make(map[int]int)
	fmt.Println("s1 equals s2? ", cmp.Equal(s1, s2))
	fmt.Println("m1 equals m2? ", cmp.Equal(m1, m2))
	fmt.Println("s1 equals s2(cmpopts.EquateEmpty()? ", cmp.Equal(s1, s2, cmpopts.EquateEmpty()))
	fmt.Println("m1 equals m2(cmpopts.EquateEmpty()? ", cmp.Equal(m1, m2, cmpopts.EquateEmpty()))
}

func SliceCmp() {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{4, 3, 2, 1}
	fmt.Println("s1 equals s2? ", cmp.Equal(s1, s2))
	fmt.Println("s1 equals s2(cmpopts.SortSlices)? ", cmp.Equal(s1, s2, cmpopts.SortSlices(func(i, j int) bool {
		return i < j
	})))
	m1 := map[int]int{1: 1, 2: 2, 3: 3, 4: 4}
	m2 := map[int]int{4: 4, 3: 3, 2: 2, 1: 1}
	fmt.Println("m1 equals m2? ", cmp.Equal(m1, m2))
	fmt.Println("m1 equals m2(cmpopts.SortMaps)? ", cmp.Equal(m1, m2, cmpopts.SortMaps(func(i, j int) bool {
		return i < j
	})))
}

type NetAddr struct {
	IP   string
	Port int
}

func (a NetAddr) Equal(b NetAddr) bool {
	if b.Port != a.Port {
		return false
	}
	if a.IP != b.IP {
		if a.IP == "127.0.0.1" && b.IP == "localhost" {
			return true
		}
		if a.IP == "localhost" && b.IP == "127.0.0.1" {
			return true
		}
		return false
	}
	return true
}

func CustomEqual() {
	a1 := NetAddr{"127.0.0.1", 5000}
	a2 := NetAddr{"localhost", 5000}
	a3 := NetAddr{"192.168.1.1", 5000}
	fmt.Println("a1 equals a2? ", cmp.Equal(a1, a2, cmp.Comparer(NetAddr.Equal)))
	fmt.Println("a1 equals a3? ", cmp.Equal(a1, a3, cmp.Comparer(NetAddr.Equal)))
}
