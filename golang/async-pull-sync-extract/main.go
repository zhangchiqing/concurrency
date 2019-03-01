package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

// ▸ go run main.go
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

type io struct {
	Input  string
	Output (chan *promiseInt)
}

type promiseIntChan = chan *promiseInt

func run(log logger) ([]int, error) {
	cs := []string{"A", "B", "C", "D", "E"}

	nPs := []*io{}
	for _, c := range cs {
		nP := make(chan *promiseInt)
		// Be careful: `c` has to pass into the closure, otherwise it would aways be
		// the last c in cs
		go func(c string, nP chan<- *promiseInt) {
			log("pulling...:", c)
			n, err := pull(c)
			log("pulled:", c, n, err)
			nP <- &promiseInt{n, err}
		}(c, nP)
		nPs = append(nPs, &io{
			Input:  c,
			Output: nP,
		})
	}

	ms := []int{}
	for _, nP := range nPs {
		io := nP
		mV := <-io.Output
		if mV.Err != nil {
			return nil, mV.Err
		}
		log("extracting...:", io.Input, mV.Value)
		err := extract(mV.Value)
		log("extracted:", io.Input, mV.Value)
		if err != nil {
			return nil, err
		}
		ms = append(ms, mV.Value)
	}

	return ms, nil
}

func pull(c string) (int, error) {
	// 1 <= sec <= 5
	sec := randN(5) + 1
	time.Sleep(time.Duration(sec) * time.Second)
	return sec, nil
}

func extract(c int) error {
	// 1 <= sec <= 3
	sec := randN(3) + 1
	time.Sleep(time.Duration(sec) * time.Second)
	return nil
}

func randN(n int) int {
	s := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(s)
	return gen.Intn(n)
}
