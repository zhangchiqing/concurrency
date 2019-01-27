package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// ▸ go run main.go
// [2019-01-27T16:12:44Z] [run start]
// [2019-01-27T16:12:44Z] [getA start]
// [2019-01-27T16:12:47Z] ...[getA end. a, err: A <nil>]
// [2019-01-27T16:12:47Z] ...[getBWithA start. a: A]
// [2019-01-27T16:12:50Z] ......[getBWithA end. b, err: 3 <nil>]
// [2019-01-27T16:12:50Z] ......[getCWithAB start. a, b: A 3]
// [2019-01-27T16:12:53Z] .........[getCWithAB end. c, err: C <nil>]
// [2019-01-27T16:12:53Z] .........[run end C <nil>]
func main() {
	log := start()
	log("run start")
	ns, err := run(log)
	log("run end", ns, err)
}

func run(log logger) (string, error) {
	log("getA start")
	a, err := getA()
	log("getA end. a, err:", a, err)

	if err != nil {
		return "", err
	}

	log("getBWithA start. a:", a)
	b, err := getBWithA(a)
	log("getBWithA end. b, err:", b, err)

	if err != nil {
		return "", err
	}

	log("getCWithAB start. a, b:", a, b)
	c, err := getCWithAB(a, b)
	log("getCWithAB end. c, err:", c, err)
	if err != nil {
		return "", err
	}

	return c, nil
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

func getA() (string, error) {
	wait(3)
	return "A", nil
}

func getBWithA(a string) (int, error) {
	wait(3)
	return 3, nil
}

func getCWithAB(a string, b int) (string, error) {
	wait(3)
	return "C", nil
}

func wait(sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
}
