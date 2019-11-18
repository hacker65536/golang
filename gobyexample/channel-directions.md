https://gobyexample.com/channel-directions


```go
package main

import "fmt"

func ping(pings chan<- string, msg string) {　// send only
    pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {　// pings is recive only, pongs is  send only
    msg := <-pings
    pongs <- msg
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}
```

```console
$ go run channel-directions.go
passed message
```
