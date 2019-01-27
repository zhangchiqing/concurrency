# Concurreny Models in different languages

## Golang Concurreny With Channels and Goroutines
### Chaining and block
There is no concurrency in this example. One async call chained after another

```
▸ cd golang/chain
▸ go run main.go
[2019-01-27T16:12:44Z] [run start]
[2019-01-27T16:12:44Z] [getA start]
[2019-01-27T16:12:47Z] ...[getA end. a, err: A <nil>]
[2019-01-27T16:12:47Z] ...[getBWithA start. a: A]
[2019-01-27T16:12:50Z] ......[getBWithA end. b, err: 3 <nil>]
[2019-01-27T16:12:50Z] ......[getCWithAB start. a, b: A 3]
[2019-01-27T16:12:53Z] .........[getCWithAB end. c, err: C <nil>]
[2019-01-27T16:12:53Z] .........[run end C <nil>]
```

### Multiple async calls run concurrently
Async calls running concurrently could have different types of input and output.

In this example, `getA` and `getB` have different types, they run concurrently and `getCWithAB` is not called until both `getA` and `getB` are finished. And the return values of `getA` and `getB` will be passed to `getCWIthAB`.

```
▸ cd golang/batch
▸ go run main.go
[2019-01-27T15:59:19Z] [run start]
[2019-01-27T15:59:19Z] [getB start]
[2019-01-27T15:59:19Z] [getA start]
[2019-01-27T15:59:20Z] .[getB end: 3 <nil>]
[2019-01-27T15:59:22Z] ...[getA end: A <nil>]
[2019-01-27T15:59:22Z] ...[getCWithAB start aV: A <nil> bV: 3 <nil>]
[2019-01-27T15:59:24Z] .....[getCWithAB end  failed]
[2019-01-27T15:59:24Z] .....[run end:  failed]
```

### Making async calls in a batch
This example shows how to make a batch of async calls concurrently with inputs from a list of values, and how to wait and return the list of return values when all the async calls finish.

```
▸ cd golang/all
▸ go run main.go
[2019-01-27T15:59:44Z] [run start]
[2019-01-27T15:59:44Z] [getNWithC start: E]
[2019-01-27T15:59:44Z] [getNWithC start: B]
[2019-01-27T15:59:44Z] [getNWithC start: C]
[2019-01-27T15:59:44Z] [getNWithC start: A]
[2019-01-27T15:59:44Z] [getNWithC start: D]
[2019-01-27T15:59:46Z] ..[getNWithC end: C 2 <nil>]
[2019-01-27T15:59:46Z] ..[getNWithC end: D 2 <nil>]
[2019-01-27T15:59:46Z] ..[getNWithC end: B 2 <nil>]
[2019-01-27T15:59:47Z] ...[getNWithC end: A 3 <nil>]
[2019-01-27T15:59:49Z] .....[getNWithC end: E 5 <nil>]
[2019-01-27T15:59:49Z] .....[run end [3 2 2 2 5] <nil>]
```

A fixed-length queue could be added to throttle the number of async calls running concurrently.

### Timeout
This example shows how to give a timeout for an async call. The implementation can also be used the same for racing two async calls.

```
▸ cd golang/timeout
▸ go run main.go
[2019-01-27T16:01:44Z] [run start]
[2019-01-27T16:01:44Z] [getA start]
[2019-01-27T16:01:47Z] ...[getA end: 0 getA Timeout after 3 secs]
[2019-01-27T16:01:47Z] ...[run end 0 getA Timeout after 3 secs]
```
