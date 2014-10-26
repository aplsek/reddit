package main

import (
  //"io"
  "fmt"
  //"github.com/aplsek/reddit"

  //"os"
)

func main() {
  items, err := reddit.Get("golang")
  if err != nil {
    log.Fatal(err)
  }


  for _,child := range r.Data.Children {
    fmt.Println(child.Data.Title)
  }

}
