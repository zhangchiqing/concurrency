package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

// ▸ go run main.go
// [2019-01-27T15:59:44Z] [run start]
// [2019-01-27T15:59:44Z] [getNWithC start: E]
// [2019-01-27T15:59:44Z] [getNWithC start: B]
// [2019-01-27T15:59:44Z] [getNWithC start: C]
// [2019-01-27T15:59:44Z] [getNWithC start: A]
// [2019-01-27T15:59:44Z] [getNWithC start: D]
// [2019-01-27T15:59:46Z] ..[getNWithC end: C 2 <nil>]
// [2019-01-27T15:59:46Z] ..[getNWithC end: D 2 <nil>]
// [2019-01-27T15:59:46Z] ..[getNWithC end: B 2 <nil>]
// [2019-01-27T15:59:47Z] ...[getNWithC end: A 3 <nil>]
// [2019-01-27T15:59:49Z] .....[getNWithC end: E 5 <nil>]
// [2019-01-27T15:59:49Z] .....[run end [3 2 2 2 5] <nil>]
func main() {
	log := start()
	log("run start")
	ns, err := run(log)
	log("run end", ns, err)
}

type logger = func(msgs ...interface{})

func start() logger {
	begin := time.Now()
	return func(msgs ...interface{}) {
		cur := time.Now()
		secsFloat64 := cur.Sub(begin).Seconds()
		secs := int(math.Round(secsFloat64))
		dots := strings.Repeat(".", secs)
		fmt.Println(fmt.Sprintf("[%v] %v%v", now(), dots, msgs))
	}
}

func now() string {
	return time.Now().UTC().Format(time.RFC3339)
}

type promiseInt struct {
	Value int
	Err   error
}

type promiseIntChan = chan *promiseInt

func run(log logger) ([]int, error) {
	cs := []string{"A", "B", "C", "D", "E"}

	nPs := [](chan *promiseInt){}
	for _, c := range cs {
		nP := make(chan *promiseInt)
		// Be careful: `c` has to pass into the closure, otherwise it would aways be
		// the last c in cs
		go func(c string, nP chan<- *promiseInt) {
			log("getNWithC start:", c)
			n, err := getNWithC(c)
			log("getNWithC end:", c, n, err)
			nP <- &promiseInt{n, err}
		}(c, nP)
		nPs = append(nPs, nP)
	}

	ms := []int{}
	for i, nP := range nPs {
		mV := <-nP
		if mV == nil {
			return nil, fmt.Errorf("mV is nil, i:%v", i)
		}

		if mV.Err != nil {
			return nil, mV.Err
		}
		ms = append(ms, mV.Value)
	}

	return ms, nil
}

func getNWithC(c string) (int, error) {
	// 1 <= sec <= 5
	sec := randN(5) + 1
	time.Sleep(time.Duration(sec) * time.Second)
	return sec, nil
}

func randN(n int) int {
	s := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(s)
	return gen.Intn(n)
}
