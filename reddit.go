package reddit

import (
  //"io"
  "fmt"
  "encoding/json"
  "log"
  "errors"
  "net/http"
  //"os"
)

type Item struct {
  Title string
  URL string
  Comments int `json:"num_comments"`
}

func (i Item) String() string {
  com := ""
  switch i.Comments {
    case 0: 
      //nothing
  case 1:
      com = " (1 comment)"
  default:
      com = fmt.Sprintf(" (%d comments)", i.Comments)
  }
  return fmt.Sprintf("%s%s\n%s",i.Title, com, i.URL)
}

type Response struct {
  Data struct {
    Children []struct {
      Data Item
    }
  }
}

func Get(reddit string) ([]Item, error) {
  url := fmt.Sprintf("http://reddit.com/r/%s.json", reddit)
  resp,err := http.Get(url)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()
  if resp.StatusCode != http.StatusOK {
     return nil, errors.New(resp.Status)
  }
  r := new(Response)
  err = json.NewDecoder(resp.Body).Decode(r)
  if err != nil {
    return nil, err
  }

  items := make([]Item, len(r.Data.Children))
  for i,child := range r.Data.Children {
    items[i] = child.Data
  }
  return items, nil

}



func main_second() {
  items,err := Get("golang")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("START")
  for _, item := range items {
    fmt.Println(item)
  }


}

func main_first() {
  resp,err := http.Get("http://reddit.com/r/golang.json")
  if err != nil {
    log.Fatal(err)
  }

  if resp.StatusCode != http.StatusOK {
     log.Fatal(resp.Status)
  }

  //_, err = io.Copy(os.Stdout, resp.Body)
  r := new(Response)
  err = json.NewDecoder(resp.Body).Decode(r)
  if err != nil {
    log.Fatal(err)
  }

  for _,child := range r.Data.Children {
    fmt.Println(child.Data.Title)
  }

}
