package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"time"
)

// ▸ go run main.go
// [2019-01-23T18:22:51Z] [run start]
// [2019-01-23T18:22:51Z] [getB start]
// [2019-01-23T18:22:51Z] [getA start]
// [2019-01-23T18:22:52Z] .[getB end: 3 <nil>]
// [2019-01-23T18:22:54Z] ...[getA end: A <nil>]
// [2019-01-23T18:22:54Z] ...[getCWithAB start A 3]
// [2019-01-23T18:22:56Z] .....[getCWithAB end  failed]
// [2019-01-23T18:22:56Z] .....[run end]
func main() {
	log := start()
	log("run start")
	v, err := run(log)
	log("run end:", v, err)
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

type promiseString struct {
	Value string
	Err   error
}

type promiseInt struct {
	Value int
	Err   error
}

func run(log logger) (string, error) {
	aP := make(chan *promiseString)
	bP := make(chan *promiseInt)

	go func(log logger, aP chan<- *promiseString) {
		log("getA start")
		a, err := getA()
		log("getA end:", a, err)
		aP <- &promiseString{a, err}
	}(log, aP)

	go func(log logger, bP chan<- *promiseInt) {
		log("getB start")
		b, err := getB()
		log("getB end:", b, err)
		bP <- &promiseInt{b, err}
	}(log, bP)

	aV := <-aP
	if aV == nil {
		return "", errors.New("aV is nil")
	}

	// Be careful here. Forgetting to handle the error still compiles.
	if aV.Err != nil {
		return "", aV.Err
	}

	bV := <-bP
	if bV.Err != nil {
		return "", bV.Err
	}

	log("getCWithAB start aV:", aV.Value, aV.Err, "bV:", bV.Value, bV.Err)
	c, err := getCWithAB(aV.Value, bV.Value)
	log("getCWithAB end", c, err)
	return c, err
}

func getA() (string, error) {
	wait(3)
	return "A", nil
}

func getB() (int, error) {
	wait(1)
	return 3, nil
}

func getCWithAB(a string, b int) (string, error) {
	wait(2)
	// return "C", nil
	return "", errors.New("failed")
}

func wait(sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
}

func now() string {
	return time.Now().UTC().Format(time.RFC3339)
}
