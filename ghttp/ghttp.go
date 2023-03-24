package ghttp

import (
  "fmt"
  "github.com/imroc/req/v3"
  "sync"
  "strconv"
)

func MultiRequest() string {

  var wg sync.WaitGroup

  for i :=1; i<=5; i++ {
    wg.Add(1)

    i := 1

    go func() {
      defer wg.Done()
      Request(strconv.Itoa(i))
    }()
      
  }
  wg.Wait()

  return "Done"
}



func Request(name string) string {
  message := fmt.Sprintf("Hi, %v. Welcome!", name)
  req.DevMode()
  req.MustGet("https://httpbin.org/uuid")

  req.EnableForceHTTP1()
  req.MustGet("https://httpbin.org/uuid")
  fmt.Printf(name)
  return message
}
