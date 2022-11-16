package gana

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenericWorker(t *testing.T) {
	inputs := make(chan int)
	outputs := make(chan string)
	wg := &sync.WaitGroup{}
	sq := func(a int) string {
		return strconv.Itoa(a * a)
	}
	go GenericWorker(inputs, outputs, sq, wg)
	wg.Add(1)
	for i := -1000; i <= 1000; i++ {
		inputs <- i
		assert.Equal(t, sq(i), <-outputs,
			fmt.Sprintf("worker [%d]", i))
	}
	close(inputs)
	close(outputs)
	wg.Wait()
}

func TestGenericWorkers(t *testing.T) {
	inputs := make(chan int)
	sq := func(a int) int {
		return a * a
	}
	double := func(a int) int {
		return a + a
	}
	sub5 := func(a int) int {
		return a - 5
	}
	strc := func(a int) string {
		return strconv.Itoa(a)
	}
	toint := func(a string) int {
		f, err := strconv.ParseInt(a, 10, 64)
		assert.Nil(t, err, fmt.Sprintf("parse int of %s", a))
		return int(f)
	}
	squares := GenericWorkers(inputs, sq, 10, 20)
	doubles := GenericWorkers(squares, double, 20, 40)
	sub5s := GenericWorkers(doubles, sub5, 40, 60)
	strcs := GenericWorkers(sub5s, strc, 60, 5)
	ints := GenericWorkers(strcs, toint, 5, 1)

	results := make([]int, 0, 1001)
	go func() {
		for i := 0; i <= 1000; i++ {
			inputs <- i
		}
		close(inputs)
	}()
	for ii := range ints {
		results = append(results, ii)
	}
	sort.Ints(results)
	for i := 0; i <= 1000; i++ {
		assert.Equal(t, toint(strc(sub5(double(sq(i))))), results[i],
			fmt.Sprintf("Checking value %d", i))
	}
}
