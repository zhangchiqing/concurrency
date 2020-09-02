你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
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

### Pulling asynchronously and then extract synchronously
This example shows how to making first calls (pull) asynchronously and then run the second call (extract) synchronously.

The behavior is similar to how docker pull and extract images, where it pulls different layers asynchronously and then extracted one after the other in order.

```
▸ go run main.go
[2019-03-01T22:53:39Z] [run start]
[2019-03-01T22:53:39Z] [pulling...: E]
[2019-03-01T22:53:39Z] [pulling...: B]
[2019-03-01T22:53:39Z] [pulling...: D]
[2019-03-01T22:53:39Z] [pulling...: C]
[2019-03-01T22:53:39Z] [pulling...: A]
[2019-03-01T22:53:40Z] .[pulled: A 1 <nil>]
[2019-03-01T22:53:40Z] .[extracting...: A 1]
[2019-03-01T22:53:41Z] ..[pulled: E 2 <nil>]
[2019-03-01T22:53:41Z] ..[extracted: A 1]
[2019-03-01T22:53:42Z] ...[pulled: B 3 <nil>]
[2019-03-01T22:53:42Z] ...[extracting...: B 3]
[2019-03-01T22:53:43Z] ....[extracted: B 3]
[2019-03-01T22:53:44Z] .....[pulled: D 5 <nil>]
[2019-03-01T22:53:44Z] .....[pulled: C 5 <nil>]
[2019-03-01T22:53:44Z] .....[extracting...: C 5]
[2019-03-01T22:53:46Z] .......[extracted: C 5]
[2019-03-01T22:53:46Z] .......[extracting...: D 5]
[2019-03-01T22:53:48Z] .........[extracted: D 5]
[2019-03-01T22:53:48Z] .........[extracting...: E 2]
[2019-03-01T22:53:51Z] ............[extracted: E 2]
[2019-03-01T22:53:51Z] ............[run end [1 3 5 5 2] <nil>]
```
