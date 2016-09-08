package main

import (
        "bufio"
        "fmt"
        "net"
        "net/url"
        "os"
)

func main() {
        fname := os.Args[1]
        file, err := os.Open(fname)
        if err != nil {
                panic(err)
        }
        defer file.Close()

        sc := bufio.NewScanner(file)
        for i := 1; sc.Scan(); i++ {
                if err := sc.Err(); err != nil {
                        break
                }
                u, _ := url.Parse(sc.Text())
                ad, _ := net.LookupIP(u.Host)

                fmt.Println(ad," ",u.Host)
//              fmt.Println(ad)
        }

}
