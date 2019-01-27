package main

import (
	"fmt"
	"time"
)

// ▸ go run main.go
// [2019-01-27T16:00:15Z] start
// [2019-01-27T16:00:15Z] getA start
// [2019-01-27T16:00:18Z] getA end
// [2019-01-27T16:00:18Z] getBWithA start
// [2019-01-27T16:00:21Z] getBWithA end
// [2019-01-27T16:00:21Z] getCWithAB start
// [2019-01-27T16:00:24Z] getCWithAB end
// [2019-01-27T16:00:24Z] end: C
func main() {
	log("start")
	a, err := getA()
	if err != nil {
		log(err)
	}

	b, err := getBWithA(a)
	if err != nil {
		log(err)
	}

	c, err := getCWithAB(a, b)
	if err != nil {
		log(err)
	}

	log("end: " + c)
}

func log(msg interface{}) {
	fmt.Println(fmt.Sprintf("[%v] %v", now(), msg))
}

func now() string {
	return time.Now().UTC().Format(time.RFC3339)
}

func getA() (string, error) {
	log("getA start")
	wait(3)
	log("getA end")
	return "A", nil
}

func getBWithA(a string) (int, error) {
	log("getBWithA start")
	wait(3)
	log("getBWithA end")
	return 3, nil
}

func getCWithAB(a string, b int) (string, error) {
	log("getCWithAB start")
	wait(3)
	log("getCWithAB end")
	return "C", nil
}

func wait(sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
}
