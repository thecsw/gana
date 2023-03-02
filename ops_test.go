package gana

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	assert.Equal(t, 1, Min(1, 2))
	assert.Equal(t, 1, Min(1, 2))
	assert.Equal(t, 1, Min(1, 1))
	assert.Equal(t, 1.0, Min(1.0, 2.0))
	assert.Equal(t, 1.0, Min(1.0, 2.0))
	assert.Equal(t, 1.0, Min(1.0, 1.0))
	assert.Equal(t, "a", Min("b", "a"))
	assert.Equal(t, "a", Min("b", "a"))
	assert.Equal(t, "a", Min("a", "a"))
}

func TestMax(t *testing.T) {
	assert.Equal(t, 2, Max(1, 2), "2, 1")
	assert.Equal(t, 2, Max(2, 1), "2, 1")
	assert.Equal(t, 1, Max(1, 1), "1, 1")
	assert.Equal(t, 2.0, Max(1.0, 2.0), "2.0, 1.0")
	assert.Equal(t, 2.0, Max(2.0, 1.0), "2.0, 1.0")
	assert.Equal(t, 1.0, Max(1.0, 1.0), "1.0, 1.0")
	assert.Equal(t, "b", Max("a", "b"), "b, a")
	assert.Equal(t, "b", Max("b", "a"), "b, a")
	assert.Equal(t, "a", Max("a", "a"), "a, a")
}

func TestMinv(t *testing.T) {
	assert.Equal(t, ZeroValue[int](), Minv[int](), "empty minv")
	assert.Equal(t, 1, Minv(1, 2, 3), "1, 2, 3")
	assert.Equal(t, 1, Minv(2, 1, 3), "2, 1, 3")
	assert.Equal(t, 1, Minv(3, 2, 1), "3, 2, 1")
	assert.Equal(t, 1.0, Minv(1.0, 2.0, 3.0), "1.0, 2.0, 3.0")
	assert.Equal(t, 1.0, Minv(2.0, 1.0, 3.0), "2.0, 1.0, 3.0")
	assert.Equal(t, 1.0, Minv(3.0, 2.0, 1.0), "3.0, 2.0, 1.0")
	assert.Equal(t, "a", Minv("a", "b", "c"), `"a", "b", "c"`)
	assert.Equal(t, "a", Minv("b", "a", "c"), `"b", "a", "c"`)
	assert.Equal(t, "a", Minv("c", "b", "a"), `"c", "b", "a"`)
}

func TestOperators(t *testing.T) {
	assert.Equal(t, true, Less(1, 2), "less")
	assert.Equal(t, true, Greater(2, 1), "greater")
	assert.Equal(t, true, Equals(1, 1), "equals")
	assert.Equal(t, false, Equals(1, 2), "not equals")
	assert.Equal(t, true, EqualsClosure(1)(1), "equals closure")
	assert.Equal(t, false, EqualsClosure(1)(2), "not equals closure")
}

func TestMinf(t *testing.T) {
	now := time.Date(2000, 06, 14, 14, 14, 0, 0, time.Local)
	plusSecond := now.Add(time.Second)
	minusHour := now.Add(-time.Hour)
	plusDay := now.Add(time.Hour * 24)
	minusYear := now.Add(-time.Hour * 24 * 365)

	UnixLess := func(a, b time.Time) bool { return a.UnixNano() < b.UnixNano() }
	DurationLess := func(a, b time.Time) bool { return a.Sub(b) < 0 }

	assert.Equal(t, ZeroValue[int](), Minv[int](), "empty minv")
	assert.Equal(t, minusYear, Minf(UnixLess, now, plusSecond, minusHour, plusDay, minusYear), "time one")
	assert.Equal(t, minusYear, Minf(DurationLess, now, plusSecond, minusHour, plusDay, minusYear), "time two")
}

func TestMaxv(t *testing.T) {
	assert.Equal(t, ZeroValue[int](), Maxv[int](), "empty maxv")
	assert.Equal(t, 3, Maxv(1, 2, 3), "1, 2, 3")
	assert.Equal(t, 3, Maxv(2, 1, 3), "2, 1, 3")
	assert.Equal(t, 3, Maxv(3, 2, 1), "3, 2, 1")
	assert.Equal(t, 3.0, Maxv(1.0, 2.0, 3.0), "1.0, 2.0, 3.0")
	assert.Equal(t, 3.0, Maxv(2.0, 1.0, 3.0), "2.0, 1.0, 3.0")
	assert.Equal(t, 3.0, Maxv(3.0, 2.0, 1.0), "3.0, 2.0, 1.0")
	assert.Equal(t, "c", Maxv("a", "b", "c"), `"a", "b", "c"`)
	assert.Equal(t, "c", Maxv("b", "a", "c"), `"b", "a", "c"`)
	assert.Equal(t, "c", Maxv("c", "b", "a"), `"c", "b", "a"`)
}

func TestMaxf(t *testing.T) {
	now := time.Date(2000, 06, 14, 14, 14, 0, 0, time.Local)
	plusSecond := now.Add(time.Second)
	minusHour := now.Add(-time.Hour)
	plusDay := now.Add(time.Hour * 24)
	minusYear := now.Add(-time.Hour * 24 * 365)

	UnixLess := func(a, b time.Time) bool { return a.UnixNano() < b.UnixNano() }
	DurationLess := func(a, b time.Time) bool { return a.Sub(b) < 0 }

	assert.Equal(t, ZeroValue[int](), Minv[int](), "empty minv")
	assert.Equal(t, plusDay, Maxf(UnixLess, now, plusSecond, minusHour, plusDay, minusYear), "time one")
	assert.Equal(t, plusDay, Maxf(DurationLess, now, plusSecond, minusHour, plusDay, minusYear), "time two")
}

func TestMinMaxv(t *testing.T) {
	x, z := MinMaxv[int]()
	assert.Equal(t, ZeroValue[int](), x, "empty MinMaxv, min")
	assert.Equal(t, ZeroValue[int](), z, "empty MinMaxv, max")
	a, b := MinMaxv(1, 2, 3)
	assert.Equal(t, 1, a, "1, 2, 3")
	assert.Equal(t, 3, b, "1, 2, 3")
	a, b = MinMaxv(2, 1, 3)
	assert.Equal(t, 1, a, "2, 1, 3")
	assert.Equal(t, 3, b, "2, 1, 3")
	a, b = MinMaxv(3, 2, 1)
	assert.Equal(t, 1, a, "3, 2, 1")
	assert.Equal(t, 3, b, "3, 2, 1")
	c, d := MinMaxv(1.0, 2.0, 3.0)
	assert.Equal(t, 1.0, c, "1.0, 2.0, 3.0")
	assert.Equal(t, 3.0, d, "1.0, 2.0, 3.0")
	c, d = MinMaxv(2.0, 1.0, 3.0)
	assert.Equal(t, 1.0, c, "2.0, 1.0, 3.0")
	assert.Equal(t, 3.0, d, "2.0, 1.0, 3.0")
	c, d = MinMaxv(3.0, 2.0, 1.0)
	assert.Equal(t, 1.0, c, "3.0, 2.0, 1.0")
	assert.Equal(t, 3.0, d, "3.0, 2.0, 1.0")
	e, f := MinMaxv("a", "b", "c")
	assert.Equal(t, "a", e, `"a", "b", "c"`)
	assert.Equal(t, "c", f, `"a", "b", "c"`)
	e, f = MinMaxv("b", "a", "c")
	assert.Equal(t, "a", e, `"b", "a", "c"`)
	assert.Equal(t, "c", f, `"b", "a", "c"`)
	e, f = MinMaxv("c", "b", "a")
	assert.Equal(t, "a", e, `"c", "b", "a"`)
	assert.Equal(t, "c", f, `"c", "b", "a"`)
}

func TestMinMaxf(t *testing.T) {
	now := time.Date(2000, 06, 14, 14, 14, 0, 0, time.Local)
	plusSecond := now.Add(time.Second)
	minusHour := now.Add(-time.Hour)
	plusDay := now.Add(time.Hour * 24)
	minusYear := now.Add(-time.Hour * 24 * 365)

	UnixLess := func(a, b time.Time) bool { return a.UnixNano() < b.UnixNano() }
	DurationLess := func(a, b time.Time) bool { return a.Sub(b) < 0 }

	x, z := MinMaxv[int]()
	assert.Equal(t, ZeroValue[int](), x, "empty MinMaxv, min")
	assert.Equal(t, ZeroValue[int](), z, "empty MinMaxv, max")

	a, b := MinMaxf(UnixLess, now, plusSecond, minusHour, plusDay, minusYear)
	assert.Equal(t, minusYear, a, "time one")
	assert.Equal(t, plusDay, b, "time two")
	c, d := MinMaxf(DurationLess, now, plusSecond, minusHour, plusDay, minusYear)
	assert.Equal(t, minusYear, c, "time one")
	assert.Equal(t, plusDay, d, "time two")
}

func TestZeroValue(t *testing.T) {
	assert.Equal(t, 0, ZeroValue[int](), "int")
	assert.Equal(t, 0.0, ZeroValue[float64](), "float64")
	assert.Equal(t, "", ZeroValue[string](), "string")
}

func TestMap(t *testing.T) {
	assert.Equal(t, []int{2, 4, 6}, Map(func(x int) int { return x * 2 }, []int{1, 2, 3}), "1, 2, 3")
	assert.Equal(t, []float64{2.0, 4.0, 6.0}, Map(func(x float64) float64 { return x * 2 }, []float64{1.0, 2.0, 3.0}), "1.0, 2.0, 3.0")
	assert.Equal(t, []string{"aa", "bb", "cc"}, Map(func(x string) string { return x + x }, []string{"a", "b", "c"}), `"a", "b", "c"`)
	assert.Equal(t, []int{}, Map(func(x int) int { return x * 2 }, []int{}), "[]int{}")
	assert.Equal(t, []float64{}, Map(func(x float64) float64 { return x * 2 }, []float64{}), "[]float64{}")
	assert.Equal(t, []string{}, Map(func(x string) string { return x + x }, []string{}), "[]string{}")
}

func TestFirst(t *testing.T) {
	assert.Equal(t, 1, First([]int{1, 2, 3}), "1, 2, 3")
	assert.Equal(t, 0, First([]int{}), "[]int{}")
	assert.Equal(t, 1.0, First([]float64{1.0, 2.0, 3.0}), "1.0, 2.0, 3.0")
	assert.Equal(t, 0.0, First([]float64{}), "[]float64{}")
	assert.Equal(t, "a", First([]string{"a", "b", "c"}), `"a", "b", "c"`)
	assert.Equal(t, "", First([]string{}), "[]string{}")

}

func TestLast(t *testing.T) {
	assert.Equal(t, 0, Last([]int{}), "[]int{}")
	assert.Equal(t, 3.0, Last([]float64{1.0, 2.0, 3.0}), "1.0, 2.0, 3.0")
	assert.Equal(t, 0.0, Last([]float64{}), "[]float64{}")
	assert.Equal(t, "c", Last([]string{"a", "b", "c"}), `"a", "b", "c"`)
	assert.Equal(t, "", Last([]string{}), "[]string{}")
}

func TestFilter(t *testing.T) {
	assert.Equal(t, []int{3}, Filter(func(x int) bool { return x > 2 }, []int{1, 2, 3}), "1, 2, 3")
	assert.Equal(t, []float64{3.0}, Filter(func(x float64) bool { return x > 2.0 }, []float64{1.0, 2.0, 3.0}), "1.0, 2.0, 3.0")
	assert.Equal(t, []string{"c"}, Filter(func(x string) bool { return x > "b" }, []string{"a", "b", "c"}), `"a", "b", "c"`)
	assert.Equal(t, []int{}, Filter(func(x int) bool { return x > 2 }, []int{}), "[]int{}")
	assert.Equal(t, []float64{}, Filter(func(x float64) bool { return x > 2.0 }, []float64{}), "[]float64{}")
	assert.Equal(t, []string{}, Filter(func(x string) bool { return x > "b" }, []string{}), "[]string{}")
}

func TestRemoveIf(t *testing.T) {
	assert.Equal(t, []int{3}, RemoveIf(func(x int) bool { return x <= 2 }, []int{1, 2, 3}), "1, 2, 3")
	assert.Equal(t, []float64{3.0}, RemoveIf(func(x float64) bool { return x <= 2.0 }, []float64{1.0, 2.0, 3.0}), "1.0, 2.0, 3.0")
	assert.Equal(t, []string{"c"}, RemoveIf(func(x string) bool { return x <= "b" }, []string{"a", "b", "c"}), `"a", "b", "c"`)
	assert.Equal(t, []int{}, RemoveIf(func(x int) bool { return x <= 2 }, []int{}), "[]int{}")
	assert.Equal(t, []float64{}, RemoveIf(func(x float64) bool { return x <= 2.0 }, []float64{}), "[]float64{}")
	assert.Equal(t, []string{}, RemoveIf(func(x string) bool { return x <= "b" }, []string{}), "[]string{}")
}

func TestTake(t *testing.T) {
	assert.Equal(t, []int{1, 2}, Take(2, []int{1, 2, 3}), "1, 2, 3")
	assert.Equal(t, []float64{1.0, 2.0}, Take(2, []float64{1.0, 2.0, 3.0}), "1.0, 2.0, 3.0")
	assert.Equal(t, []string{"a", "b"}, Take(2, []string{"a", "b", "c"}), `"a", "b", "c"`)
	assert.Equal(t, []int{}, Take(2, []int{}), "[]int{}")
	assert.Equal(t, []float64{}, Take(2, []float64{}), "[]float64{}")
	assert.Equal(t, []string{}, Take(2, []string{}), "[]string{}")
}

func TestTail(t *testing.T) {
	assert.Equal(t, []int{2, 3}, Tail(2, []int{1, 2, 3}), "1, 2, 3")
	assert.Equal(t, []float64{2.0, 3.0}, Tail(2, []float64{1.0, 2.0, 3.0}), "1.0, 2.0, 3.0")
	assert.Equal(t, []string{"b", "c"}, Tail(2, []string{"a", "b", "c"}), `"a", "b", "c"`)
	assert.Equal(t, []int{}, Tail(2, []int{}), "[]int{}")
	assert.Equal(t, []float64{}, Tail(2, []float64{}), "[]float64{}")
	assert.Equal(t, []string{}, Tail(2, []string{}), "[]string{}")
}

func TestDrop(t *testing.T) {
	assert.Equal(t, []int{3}, Drop[int, uint](2, []int{1, 2, 3}), "1, 2, 3")
	assert.Equal(t, []float64{3.0}, Drop[float64, uint](2, []float64{1.0, 2.0, 3.0}), "1.0, 2.0, 3.0")
	assert.Equal(t, []string{"c"}, Drop[string, uint](2, []string{"a", "b", "c"}), `"a", "b", "c"`)
	assert.Equal(t, []int{}, Drop[int, uint](2, []int{}), "[]int{}")
	assert.Equal(t, []float64{}, Drop[float64, uint](2, []float64{}), "[]float64{}")
	assert.Equal(t, []string{}, Drop[string, uint](2, []string{}), "[]string{}")
}

func TestSkip(t *testing.T) {
	assert.Equal(t, []int{3}, Skip[int, uint](2, []int{1, 2, 3}), "1, 2, 3")
	assert.Equal(t, []float64{3.0}, Skip[float64, uint](2, []float64{1.0, 2.0, 3.0}), "1.0, 2.0, 3.0")
	assert.Equal(t, []string{"c"}, Skip[string, uint](2, []string{"a", "b", "c"}), `"a", "b", "c"`)
	assert.Equal(t, []int{}, Skip[int, uint](2, []int{}), "[]int{}")
	assert.Equal(t, []float64{}, Skip[float64, uint](2, []float64{}), "[]float64{}")
	assert.Equal(t, []string{}, Skip[string, uint](2, []string{}), "[]string{}")
}

func TestZip(t *testing.T) {
	assert.Equal(t,
		[]Tuple[int, int]{NewTuple(1, 4), NewTuple(2, 5), NewTuple(3, 6)},
		Zip([]int{1, 2, 3}, []int{4, 5, 6}),
		"1, 2, 3, 4, 5, 6")

	assert.Equal(t,
		[]Tuple[float64, float64]{NewTuple(1.0, 4.0), NewTuple(2.0, 5.0), NewTuple(3.0, 6.0)},
		Zip([]float64{1.0, 2.0, 3.0}, []float64{4.0, 5.0, 6.0}),
		"1.0, 2.0, 3.0, 4.0, 5.0, 6.0")

	assert.Equal(t,
		[]Tuple[string, string]{NewTuple("a", "d"), NewTuple("b", "e"), NewTuple("c", "f")},
		Zip([]string{"a", "b", "c"}, []string{"d", "e", "f"}),
		`"a", "b", "c", "d", "e", "f"`)

	assert.Equal(t,
		[]Tuple[int, int]{},
		Zip([]int{}, []int{}),
		"[]int{}, []int{}")
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
	assert.Equal(t, []int{}, Repeat(1, 0), "Repeat(1, 0)")
	assert.Equal(t, []int{1}, Repeat(1, 1), "Repeat(1, 1)")
	assert.Equal(t, []int{1, 1}, Repeat(1, 2), "Repeat(1, 2)")
	assert.Equal(t, []int{1, 1, 1}, Repeat(1, 3), "Repeat(1, 3)")
	assert.Equal(t, []int{1, 1, 1, 1}, Repeat(1, 4), "Repeat(1, 4)")
	assert.Equal(t, []int{1, 1, 1, 1, 1}, Repeat(1, 5), "Repeat(1, 5)")
}

func TestGetPointer(t *testing.T) {
	var i int
	i_ptr := &i
	i_ptr2 := &i_ptr
	assert.Equal(t, GetPointer(i), i_ptr, "GetPointer(i)")
	assert.Equal(t, GetPointer(i_ptr), i_ptr2, "GetPointer(&i)")
}
