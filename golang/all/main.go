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
	run(log)
	log("run end")
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

type promiseInt struct {
	int
	error
}

func run(log logger) {
	// using anonymouse structs and fields
	aP := make(chan *struct {
		string
		error
	})

	// using struct
	bP := make(chan *promiseInt)

	// using anonymouse struct and closure
	go func() {
		log("getA start")
		a, err := getA()
		log("getA end:", a, err)
		aP <- &struct {
			string
			error
		}{a, err}
	}()

	// using struct and passing in values
	go func(log logger, bP chan<- *promiseInt) {
		log("getB start")
		b, err := getB()
		log("getB end:", b, err)
		bP <- &promiseInt{b, err}
	}(log, bP)

	aV := <-aP
	// Be careful here. Forgetting to handle the error still compiles.
	if aV.error != nil {
		return
	}

	bV := <-bP
	if bV.error != nil {
		return
	}

	log("getCWithAB start", aV.string, bV.int)
	c, err := getCWithAB(aV.string, bV.int)
	log("getCWithAB end", c, err)

	if err != nil {
		return
	}
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
