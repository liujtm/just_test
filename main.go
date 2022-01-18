package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    resp, err := http.Get("http://httpbin99999.org/get")
    if err != nil {
//        fmt.Println(err)
	fmt.Println(resp.StatusCode)
        return
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
    fmt.Println(resp.StatusCode)
    if resp.StatusCode == 200 {
        fmt.Println("ok")
    }
}
