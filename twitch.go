package main

import (
  "bufio"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
  "strconv"
)

func HandleError(err error, msg string) {
  if err != nil {
    fmt.Println(msg)
    Exit()
  }
}

func Exit() {
  os.Exit(1)
}

func GetJSON(url string) Streams {
  res, err := http.Get(url)
  msg := "Something went wrong, please try again in a few minutes."

  HandleError(err, msg)
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  HandleError(err, msg)

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
  Status string
  Url string
}

func DecodeResponse(body []byte) Streams {
  msg := "Something went wrong, please try again in a few minutes."
  var data Streams
  err := json.Unmarshal(body, &data)
  HandleError(err, msg);

  return data
}

type FormattedStream struct {
  name string
  title string
  url string
}

func FormatStreams(json Streams) []FormattedStream {
  var streams []FormattedStream

  for _, stream_info := range json.Streams {
    stream := FormattedStream{
      name: stream_info.Channel.Name,
      title: stream_info.Channel.Status,
      url: stream_info.Channel.Url,
    }

    streams = append(streams, stream)
  }

  return streams
}

func ListStreams(streams []FormattedStream) {
  fmt.Println("++++++ STREAMS ++++++")
  for i, stream := range streams {
    fmt.Printf("%d) %s - %s\n", i + 1, stream.name, stream.title)
  }
}

func OpenStream(stream FormattedStream) {
  str := fmt.Sprintf("livestreamer %s best -np 'omxplayer -o hdmi'", stream.url)
  os.exec(str)
  Exit()
}

func GetStreams(limit int, offset int) []FormattedStream {
  url := fmt.Sprintf("https://api.twitch.tv/kraken/streams?game=counter-strike:%sGlobal%sOffensive&limit=%d&offset=%d", "%20", "%20", limit, offset)
  json := GetJSON(url)
  streams := FormatStreams(json)

  return streams
}

func Menu() {
  limit := 20
  offset := 0
  var streams []FormattedStream

  for true {
    streams = GetStreams(limit, offset)
    ListStreams(streams)
    reader := bufio.NewReader(os.Stdin)
    choice, _ := reader.ReadString('\n')

    fmt.Println(choice)
    if choice == "q\n" {
      Exit()
    } else if choice == "m\n" {
      offset = limit
      limit = limit + 20
    } else {
      i, _ := strconv.Atoi(choice)
      length := len(streams)

      if i <= length && i >= 0 {
        OpenStream(streams[i])
      }
    }
  }
}

func main() {
  Menu()
}
