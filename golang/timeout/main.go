package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"time"
)

// ▸ go run main.go
// [2019-01-26T13:32:20Z] [run start]
// [2019-01-26T13:32:20Z] [getA start]
// [2019-01-26T13:32:22Z] ..[getA end: 3 <nil>]
// [2019-01-26T13:32:22Z] ..[run end 3 <nil>]

// ▸ go run main.go
// [2019-01-26T13:33:16Z] [run start]
// [2019-01-26T13:33:16Z] [getA start]
// [2019-01-26T13:33:18Z] ..[getA end aV (this might or might not print, if print, it might appear before or after 'getA end:***': {3 <nil>}]
// [2019-01-26T13:33:18Z] ..[if timeout, this won't be printed. if not timeout, this might be printed after 'run end ***']
// [2019-01-26T13:33:18Z] ..[getA end: 3 <nil>]
// [2019-01-26T13:33:18Z] ..[run end 3 <nil>]
func main() {
	log := start()
	log("run start")
	n, err := run(log)
	log("run end", n, err)
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

func run(log logger) (int, error) {
	aP := make(chan *promiseInt)
	go func(aP chan *promiseInt) {
		log("getA start")
		a, err := getA()
		aV := &promiseInt{a, err}
		aP <- aV
		log("getA end aV (this might or might not print, if print, it might appear before or after 'getA end:***':", *aV)
		log("if timeout, this won't be printed. if not timeout, this might be printed after 'run end ***'")
	}(aP)

	var aV *promiseInt
	select {
	case aV = <-aP:
	case <-time.After(3 * time.Second):
		aV = &promiseInt{0, errors.New("getA Timeout after 5 secs")}
	}

	// make sure to handle the nil pointer case
	if aV == nil {
		log("getA end:", nil)
		return 0, nil
	}
	log("getA end:", aV.Value, aV.Err)

	if aV.Err != nil {
		return 0, aV.Err
	}
	return aV.Value, nil
}

func getA() (int, error) {
	// use wait(4) to trigger timeout
	// wait(4)
	wait(2)
	return 3, nil
	// return 3, errors.New("getA failed")
}

func wait(sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
}
