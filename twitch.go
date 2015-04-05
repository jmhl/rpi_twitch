package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
)

func HandleError(err error) {
  if err != nil {
    fmt.Println("Something went wrong, please try again in a few minutes.")
    os.Exit(1)
  }
}

func GetJSON(url string) Streams {
  res, err := http.Get(url)

  HandleError(err)
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  HandleError(err)

  return DecodeResponse(body)
}

type Streams struct {
  Streams []Stream_info
}

type Stream_info struct {
  Channel Channel_info
}

type Channel_info struct {
  Name string
  Url string
}

func DecodeResponse(body []byte) Streams {
  var data Streams
  err := json.Unmarshal(body, &data)
  HandleError(err);

  return data
}

func ListStreams(streams Streams) {
  for i, stream_info := range streams.Streams {
    fmt.Println(i)
    fmt.Println(stream_info.Channel.Name)
  }
  
}

func main() {
  url := "https://api.twitch.tv/kraken/streams?game=counter-strike:%20Global%20Offensive"
  json := GetJSON(url)
  fmt.Println(json.Streams)

  for i, stream_info := range json.Streams {
    fmt.Println(i)
    fmt.Println(stream_info.Channel.Name)
  }
}
