package main

import (
  "bufio"
  "fmt"
  "os"
  "os/exec"
  "strconv"
  "./api"
  "./structs"
)

func ListStreams(streams []structs.FormattedStream) {
  fmt.Println("++++++ STREAMS ++++++")
  for i, stream := range streams {
    fmt.Printf("%d) %s - %s\n", i + 1, stream.Name, stream.Title)
  }
}

func OpenStream(stream structs.FormattedStream) {
  to_exec := "livestreamer"
  args := []string{stream.Url, "best", "-np 'omxplayer -o hdmi'"}
  cmd := exec.Command(to_exec, args...)
  output, _ := cmd.CombinedOutput()
  fmt.Printf("==> Output: %s\n", string(output))

  os.Exit(1)
}

func CSGOStreams() {
  limit := 20
  offset := 0
  var streams []structs.FormattedStream

  for true {
    streams = api.GetStreams(limit, offset)
    fmt.Println("Enter 'm' to view more streams.")
    ListStreams(streams)
    reader := bufio.NewReader(os.Stdin)
    choice, _ := reader.ReadString('\n')

    if choice == "q\n" {
      os.Exit(1)
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

func FollowedStreams() {
  streams := api.GetFollows()

  if len(streams) == 0 {
    fmt.Println("Sorry, none of your followed streamers are online right now.")
    Menu()
  }

  ListStreams(streams)
  reader := bufio.NewReader(os.Stdin)
  choice, _ := reader.ReadString('\n')

  if choice == "q\n" {
    os.Exit(1)
  } else {
    i, _ := strconv.Atoi(choice)
    length := len(streams)

    if length == 0 {
      fmt.Println("You're not following any streams.")
      os.Exit(1)
    }

    if i <= length && i >= 0 {
      OpenStream(streams[i])
    }
  }
}

func Menu() {
  fmt.Println("RPi Twitch Main Menu")
  fmt.Println("Enter 'f' to view your followed channels.")
  fmt.Println("Enter 'c' to view CSGO streams.")
  fmt.Println("Press 'q' at any time to quit.")

  reader := bufio.NewReader(os.Stdin)
  choice, _ := reader.ReadString('\n')

  switch choice {
    case "f\n": FollowedStreams()
    case "c\n": CSGOStreams()
  }
}

func main() {
  Menu()
}
