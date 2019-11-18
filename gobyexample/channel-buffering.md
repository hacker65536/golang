
```go
package main

import "fmt"

func main() {
    messages := make(chan string, 2)
    messages <- "buffered"
    messages <- "channel"

    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
```

```console
$ go run channel-buffering.go
buffered
channel
```



wrong implementation
```go
package main

import "fmt"

func main() {
        messages := make(chan string, 2)
        messages <- "buffered"
        messages <- "channel"
        messages <- "input3rd"

        fmt.Println(<-messages)
        fmt.Println(<-messages)
}
```


```console
$ go run channel-buffering.go
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        channel-buffering.go:9 +0x8d
exit status 2
```


```go
package main

import "fmt"

func main() {
        messages := make(chan string, 2)
        messages <- "buffered"
        messages <- "channel"

        fmt.Println(<-messages)
        messages <- "input3rd"
        fmt.Println(<-messages)
        fmt.Println(<-messages)
}
```

```console
$ go run channel-buffering.go
buffered
channel
input3rd
```
