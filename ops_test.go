package gana

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	assert.Equal(t, Min(1, 2), 1)
	assert.Equal(t, Min(2, 1), 1)
	assert.Equal(t, Min(1, 1), 1)
	assert.Equal(t, Min(1.0, 2.0), 1.0)
	assert.Equal(t, Min(2.0, 1.0), 1.0)
	assert.Equal(t, Min(1.0, 1.0), 1.0)
	assert.Equal(t, Min("a", "b"), "a")
	assert.Equal(t, Min("b", "a"), "a")
	assert.Equal(t, Min("a", "a"), "a")
}

func TestMax(t *testing.T) {
	assert.Equal(t, Max(1, 2), 2, "1, 2")
	assert.Equal(t, Max(2, 1), 2, "2, 1")
	assert.Equal(t, Max(1, 1), 1, "1, 1")
	assert.Equal(t, Max(1.0, 2.0), 2.0, "1.0, 2.0")
	assert.Equal(t, Max(2.0, 1.0), 2.0, "2.0, 1.0")
	assert.Equal(t, Max(1.0, 1.0), 1.0, "1.0, 1.0")
	assert.Equal(t, Max("a", "b"), "b", "a, b")
	assert.Equal(t, Max("b", "a"), "b", "b, a")
	assert.Equal(t, Max("a", "a"), "a", "a, a")
}

func TestMinv(t *testing.T) {
	assert.Equal(t, Minv[int](), ZeroValue[int](), "empty minv")
	assert.Equal(t, Minv(1, 2, 3), 1, "1, 2, 3")
	assert.Equal(t, Minv(2, 1, 3), 1, "2, 1, 3")
	assert.Equal(t, Minv(3, 2, 1), 1, "3, 2, 1")
	assert.Equal(t, Minv(1.0, 2.0, 3.0), 1.0, "1.0, 2.0, 3.0")
	assert.Equal(t, Minv(2.0, 1.0, 3.0), 1.0, "2.0, 1.0, 3.0")
	assert.Equal(t, Minv(3.0, 2.0, 1.0), 1.0, "3.0, 2.0, 1.0")
	assert.Equal(t, Minv("a", "b", "c"), "a", `"a", "b", "c"`)
	assert.Equal(t, Minv("b", "a", "c"), "a", `"b", "a", "c"`)
	assert.Equal(t, Minv("c", "b", "a"), "a", `"c", "b", "a"`)
}

func TestMaxv(t *testing.T) {
	assert.Equal(t, Maxv[int](), ZeroValue[int](), "empty maxv")
	assert.Equal(t, Maxv(1, 2, 3), 3, "1, 2, 3")
	assert.Equal(t, Maxv(2, 1, 3), 3, "2, 1, 3")
	assert.Equal(t, Maxv(3, 2, 1), 3, "3, 2, 1")
	assert.Equal(t, Maxv(1.0, 2.0, 3.0), 3.0, "1.0, 2.0, 3.0")
	assert.Equal(t, Maxv(2.0, 1.0, 3.0), 3.0, "2.0, 1.0, 3.0")
	assert.Equal(t, Maxv(3.0, 2.0, 1.0), 3.0, "3.0, 2.0, 1.0")
	assert.Equal(t, Maxv("a", "b", "c"), "c", `"a", "b", "c"`)
	assert.Equal(t, Maxv("b", "a", "c"), "c", `"b", "a", "c"`)
	assert.Equal(t, Maxv("c", "b", "a"), "c", `"c", "b", "a"`)
}

func TestMinMaxv(t *testing.T) {
	x, z := MinMaxv[int]()
	assert.Equal(t, x, ZeroValue[int](), "empty MinMaxv, min")
	assert.Equal(t, x, ZeroValue[int](), "empty MinMaxv, max")
	a, b := MinMaxv(1, 2, 3)
	assert.Equal(t, a, 1, "1, 2, 3")
	assert.Equal(t, b, 3, "1, 2, 3")
	a, b = MinMaxv(2, 1, 3)
	assert.Equal(t, a, 1, "2, 1, 3")
	assert.Equal(t, b, 3, "2, 1, 3")
	a, b = MinMaxv(3, 2, 1)
	assert.Equal(t, a, 1, "3, 2, 1")
	assert.Equal(t, b, 3, "3, 2, 1")
	c, d := MinMaxv(1.0, 2.0, 3.0)
	assert.Equal(t, c, 1.0, "1.0, 2.0, 3.0")
	assert.Equal(t, d, 3.0, "1.0, 2.0, 3.0")
	c, d = MinMaxv(2.0, 1.0, 3.0)
	assert.Equal(t, c, 1.0, "2.0, 1.0, 3.0")
	assert.Equal(t, d, 3.0, "2.0, 1.0, 3.0")
	c, d = MinMaxv(3.0, 2.0, 1.0)
	assert.Equal(t, c, 1.0, "3.0, 2.0, 1.0")
	assert.Equal(t, d, 3.0, "3.0, 2.0, 1.0")
	e, f := MinMaxv("a", "b", "c")
	assert.Equal(t, e, "a", `"a", "b", "c"`)
	assert.Equal(t, f, "c", `"a", "b", "c"`)
	e, f = MinMaxv("b", "a", "c")
	assert.Equal(t, e, "a", `"b", "a", "c"`)
	assert.Equal(t, f, "c", `"b", "a", "c"`)
	e, f = MinMaxv("c", "b", "a")
	assert.Equal(t, e, "a", `"c", "b", "a"`)
	assert.Equal(t, f, "c", `"c", "b", "a"`)
}

func TestFirst(t *testing.T) {
	assert.Equal(t, First([]int{1, 2, 3}), 1, "1, 2, 3")
	assert.Equal(t, First([]int{}), 0, "[]int{}")
	assert.Equal(t, First([]float64{1.0, 2.0, 3.0}), 1.0, "1.0, 2.0, 3.0")
	assert.Equal(t, First([]float64{}), 0.0, "[]float64{}")
	assert.Equal(t, First([]string{"a", "b", "c"}), "a", `"a", "b", "c"`)
	assert.Equal(t, First([]string{}), "", "[]string{}")
}

func TestLast(t *testing.T) {
	assert.Equal(t, Last([]int{1, 2, 3}), 3, "1, 2, 3")
	assert.Equal(t, Last([]int{}), 0, "[]int{}")
	assert.Equal(t, Last([]float64{1.0, 2.0, 3.0}), 3.0, "1.0, 2.0, 3.0")
	assert.Equal(t, Last([]float64{}), 0.0, "[]float64{}")
	assert.Equal(t, Last([]string{"a", "b", "c"}), "c", `"a", "b", "c"`)
	assert.Equal(t, Last([]string{}), "", "[]string{}")
}

func TestZeroValue(t *testing.T) {
	assert.Equal(t, ZeroValue[int](), 0, "int")
	assert.Equal(t, ZeroValue[float64](), 0.0, "float64")
	assert.Equal(t, ZeroValue[string](), "", "string")
}

func TestMap(t *testing.T) {
	assert.Equal(t, Map(func(x int) int { return x * 2 }, []int{1, 2, 3}), []int{2, 4, 6}, "1, 2, 3")
	assert.Equal(t, Map(func(x float64) float64 { return x * 2 }, []float64{1.0, 2.0, 3.0}), []float64{2.0, 4.0, 6.0}, "1.0, 2.0, 3.0")
	assert.Equal(t, Map(func(x string) string { return x + x }, []string{"a", "b", "c"}), []string{"aa", "bb", "cc"}, `"a", "b", "c"`)
	assert.Equal(t, Map(func(x int) int { return x * 2 }, []int{}), []int{}, "[]int{}")
	assert.Equal(t, Map(func(x float64) float64 { return x * 2 }, []float64{}), []float64{}, "[]float64{}")
	assert.Equal(t, Map(func(x string) string { return x + x }, []string{}), []string{}, "[]string{}")
}

func TestFilter(t *testing.T) {
	assert.Equal(t, Filter(func(x int) bool { return x > 2 }, []int{1, 2, 3}), []int{3}, "1, 2, 3")
	assert.Equal(t, Filter(func(x float64) bool { return x > 2.0 }, []float64{1.0, 2.0, 3.0}), []float64{3.0}, "1.0, 2.0, 3.0")
	assert.Equal(t, Filter(func(x string) bool { return x > "b" }, []string{"a", "b", "c"}), []string{"c"}, `"a", "b", "c"`)
	assert.Equal(t, Filter(func(x int) bool { return x > 2 }, []int{}), []int{}, "[]int{}")
	assert.Equal(t, Filter(func(x float64) bool { return x > 2.0 }, []float64{}), []float64{}, "[]float64{}")
	assert.Equal(t, Filter(func(x string) bool { return x > "b" }, []string{}), []string{}, "[]string{}")
}

func TestTake(t *testing.T) {
	assert.Equal(t, Take(2, []int{1, 2, 3}), []int{1, 2}, "1, 2, 3")
	assert.Equal(t, Take(2, []float64{1.0, 2.0, 3.0}), []float64{1.0, 2.0}, "1.0, 2.0, 3.0")
	assert.Equal(t, Take(2, []string{"a", "b", "c"}), []string{"a", "b"}, `"a", "b", "c"`)
	assert.Equal(t, Take(2, []int{}), []int{}, "[]int{}")
	assert.Equal(t, Take(2, []float64{}), []float64{}, "[]float64{}")
	assert.Equal(t, Take(2, []string{}), []string{}, "[]string{}")
}

func TestTail(t *testing.T) {
	assert.Equal(t, Tail(2, []int{1, 2, 3}), []int{2, 3}, "1, 2, 3")
	assert.Equal(t, Tail(2, []float64{1.0, 2.0, 3.0}), []float64{2.0, 3.0}, "1.0, 2.0, 3.0")
	assert.Equal(t, Tail(2, []string{"a", "b", "c"}), []string{"b", "c"}, `"a", "b", "c"`)
	assert.Equal(t, Tail(2, []int{}), []int{}, "[]int{}")
	assert.Equal(t, Tail(2, []float64{}), []float64{}, "[]float64{}")
	assert.Equal(t, Tail(2, []string{}), []string{}, "[]string{}")
}

func TestDrop(t *testing.T) {
	assert.Equal(t, Drop[int, uint](2, []int{1, 2, 3}), []int{3}, "1, 2, 3")
	assert.Equal(t, Drop[float64, uint](2, []float64{1.0, 2.0, 3.0}), []float64{3.0}, "1.0, 2.0, 3.0")
	assert.Equal(t, Drop[string, uint](2, []string{"a", "b", "c"}), []string{"c"}, `"a", "b", "c"`)
	assert.Equal(t, Drop[int, uint](2, []int{}), []int{}, "[]int{}")
	assert.Equal(t, Drop[float64, uint](2, []float64{}), []float64{}, "[]float64{}")
	assert.Equal(t, Drop[string, uint](2, []string{}), []string{}, "[]string{}")
}

func TestSkip(t *testing.T) {
	assert.Equal(t, Skip[int, uint](2, []int{1, 2, 3}), []int{3}, "1, 2, 3")
	assert.Equal(t, Skip[float64, uint](2, []float64{1.0, 2.0, 3.0}), []float64{3.0}, "1.0, 2.0, 3.0")
	assert.Equal(t, Skip[string, uint](2, []string{"a", "b", "c"}), []string{"c"}, `"a", "b", "c"`)
	assert.Equal(t, Skip[int, uint](2, []int{}), []int{}, "[]int{}")
	assert.Equal(t, Skip[float64, uint](2, []float64{}), []float64{}, "[]float64{}")
	assert.Equal(t, Skip[string, uint](2, []string{}), []string{}, "[]string{}")
}

func TestAny(t *testing.T) {
	assert.True(t, Any(2, []int{1, 2, 3}), "1, 2, 3")
	assert.True(t, Any(2.0, []float64{1.0, 2.0, 3.0}), "1.0, 2.0, 3.0")
	assert.True(t, Any("b", []string{"a", "b", "c"}), `"a", "b", "c"`)
	assert.False(t, Any(4, []int{1, 2, 3}), "1, 2, 3")
	assert.False(t, Any(4.0, []float64{1.0, 2.0, 3.0}), "1.0, 2.0, 3.0")
	assert.False(t, Any("d", []string{"a", "b", "c"}), `"a", "b", "c"`)
}

func TestAnyf(t *testing.T) {
	assert.True(t, Anyf(func(x int) bool { return x > 2 }, []int{1, 2, 3}), "1, 2, 3")
	assert.True(t, Anyf(func(x float64) bool { return x > 2.0 }, []float64{1.0, 2.0, 3.0}), "1.0, 2.0, 3.0")
	assert.True(t, Anyf(func(x string) bool { return x > "b" }, []string{"a", "b", "c"}), `"a", "b", "c"`)
	assert.False(t, Anyf(func(x int) bool { return x > 4 }, []int{1, 2, 3}), "1, 2, 3")
	assert.False(t, Anyf(func(x float64) bool { return x > 4.0 }, []float64{1.0, 2.0, 3.0}), "1.0, 2.0, 3.0")
	assert.False(t, Anyf(func(x string) bool { return x > "d" }, []string{"a", "b", "c"}), `"a", "b", "c"`)
}

func TestAll(t *testing.T) {
	assert.False(t, All(2, []int{1, 2, 3}), "1, 2, 3")
	assert.False(t, All(2.0, []float64{1.0, 2.0, 3.0}), "1.0, 2.0, 3.0")
	assert.False(t, All("b", []string{"a", "b", "c"}), `"a", "b", "c"`)
	assert.True(t, All(20, Repeat(20, 10)), "Repeat(20, 10)")
	assert.True(t, All(20.0, Repeat(20.0, 10)), "Repeat(20.0, 10)")
	assert.True(t, All("b", Repeat("b", 10)), `Repeat("b", 10)`)
}

func TestAllf(t *testing.T) {
	assert.False(t, Allf(func(x int) bool { return x > 2 }, []int{1, 2, 3}), "1, 2, 3")
	assert.False(t, Allf(func(x float64) bool { return x > 2.0 }, []float64{1.0, 2.0, 3.0}), "1.0, 2.0, 3.0")
	assert.False(t, Allf(func(x string) bool { return x > "b" }, []string{"a", "b", "c"}), `"a", "b", "c"`)
	assert.True(t, Allf(func(x int) bool { return x > 0 }, Repeat(20, 10)), "Repeat(20, 10)")
	assert.True(t, Allf(func(x float64) bool { return x > 0.0 }, Repeat(20.0, 10)), "Repeat(20.0, 10)")
	assert.True(t, Allf(func(x string) bool { return x > "a" }, Repeat("b", 10)), `Repeat("b", 10)`)
}

func TestRepeat(t *testing.T) {
	assert.Equal(t, Repeat(1, 0), []int{}, "Repeat(1, 0)")
	assert.Equal(t, Repeat(1, 1), []int{1}, "Repeat(1, 1)")
	assert.Equal(t, Repeat(1, 2), []int{1, 1}, "Repeat(1, 2)")
	assert.Equal(t, Repeat(1, 3), []int{1, 1, 1}, "Repeat(1, 3)")
	assert.Equal(t, Repeat(1, 4), []int{1, 1, 1, 1}, "Repeat(1, 4)")
	assert.Equal(t, Repeat(1, 5), []int{1, 1, 1, 1, 1}, "Repeat(1, 5)")
}

func TestGetPointer(t *testing.T) {
	var i int
	i_ptr := &i
	i_ptr2 := &i_ptr
	assert.Equal(t, GetPointer(i), i_ptr, "GetPointer(i)")
	assert.Equal(t, GetPointer(i_ptr), i_ptr2, "GetPointer(&i)")
}
