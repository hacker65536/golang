https://gobyexample.com/timeouts

```go
package main

import (
    "fmt"
    "time"
)

func main() {

    c1 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c1 <- "result 1"
    }()

    select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(1 * time.Second):
        fmt.Println("timeout 1")
    }

    c2 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "result 2"
    }()
    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(3 * time.Second):
        fmt.Println("timeout 2")
    }
}
```


```console
$ go run timeouts.go
timeout 1
result 2
```


https://golang.org/pkg/time/#After


```go
package main

import (
    "fmt"
    "time"
)

func main() {
    t := <-time.After(1 * time.Second)
    fmt.Println("timed out", t.Format("2006-01-02 03:04:05"))
}
```

```console
$ time go run time.go
timed out 2019-11-18 06:43:08

real    0m1.270s
user    0m0.143s
sys     0m0.047s
```
