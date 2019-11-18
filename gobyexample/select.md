https://gobyexample.com/select

```go
package main

import (
    "fmt"
    "time"
)

func main() {

    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}
```

```console
$ time go run select.go
received one
received two
0.15user 0.04system 0:02.14elapsed 8%CPU (0avgtext+0avgdata 45960maxresident)k
0inputs+5384outputs (1major+15439minor)pagefaults 0swaps
```


arrange
```go
package main

import (
        "fmt"
        "time"
)

func main() {

        c1 := make(chan string)
        c2 := make(chan string)
        c3 := make(chan string)

        go func() {
                time.Sleep(1 * time.Second)
                c1 <- "one"
        }()

        go func() {
                time.Sleep(2 * time.Second)
                c2 <- "two"
        }()

        go func() {
                time.Sleep(3 * time.Second)
                c3 <- "three"
        }()

        for i := 0; i < 3; i++ {
                select {
                case msg1 := <-c1:
                        fmt.Println("received", msg1)
                case msg2 := <-c2:
                        fmt.Println("received", msg2)
                case msg3 := <-c3:
                        fmt.Println("received", msg3)
                }
        }
}
```

```console
$ time go run select.go
received one
received two
received three

real    0m3.143s
user    0m0.155s
sys     0m0.029s

```
