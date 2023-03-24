package main
import (
  "fmt"
  "example.com/ghttp"
)

func main() {
  message := ghttp.MultiRequest()
  fmt.Println(message)
}
