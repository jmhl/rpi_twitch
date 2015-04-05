package main

import (
  "bufio"
  "fmt"
  "os"
  "os/exec"
  "strings"
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

func ListGames(games []structs.FormattedGame) {
  fmt.Println("++++++ GAMES ++++++")
  for i, game := range games {
    fmt.Printf("%d) %s\n", i + 1, game.Name)
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

func GameStreams(name string) {
  limit := 20
  offset := 0
  var streams []structs.FormattedStream

  for true {
    streams = api.GetStreams(name, limit, offset)
    fmt.Println("Enter 'm' to view more streams.")
    ListStreams(streams)
    reader := bufio.NewReader(os.Stdin)
    choice, _ := reader.ReadString('\n')
    choice = strings.Replace(choice, "\n", "", 1)

    if choice == "q" {
      os.Exit(1)
    } else if choice == "m" {
      offset = limit
      limit = limit + 20
    } else {
      i, _ := strconv.Atoi(choice)
      i = i - 1
      length := len(streams)

      if i <= length && i >= 0 {
        OpenStream(streams[i])
      }
    }
  }
}

func CSGOStreams() {
  limit := 20
  offset := 0
  var streams []structs.FormattedStream

  for true {
    streams = api.GetCSGOStreams(limit, offset)
    fmt.Println("Enter 'm' to view more streams.")
    ListStreams(streams)
    reader := bufio.NewReader(os.Stdin)
    choice, _ := reader.ReadString('\n')
    choice = strings.Replace(choice, "\n", "", 1)

    if choice == "q" {
      os.Exit(1)
    } else if choice == "m" {
      offset = limit
      limit = limit + 20
    } else {
      i, _ := strconv.Atoi(choice)
      i = i - 1
      length := len(streams)

      if i <= length && i >= 0 {
        OpenStream(streams[i])
      }
    }
  }
}

func TopGames() {
  limit := 20
  offset := 0
  var games []structs.FormattedGame

  for true {
    games = api.GetGames(limit, offset)
    fmt.Println("Enter 'm' to view more games.")
    ListGames(games)
    reader := bufio.NewReader(os.Stdin)
    choice, _ := reader.ReadString('\n')
    choice = strings.Replace(choice, "\n", "", 1)

    if choice == "q" {
      os.Exit(1)
    } else if choice == "m" {
      offset = limit
      limit = limit + 20
    } else {
      i, _ := strconv.Atoi(choice)
      i = i - 1
      length := len(games)

      if i <= length && i >= 0 {
	GameStreams(games[i].Name)
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
  choice = strings.Replace(choice, "\n", "", 1)

  if choice == "q" {
    os.Exit(1)
  } else {
    i, _ := strconv.Atoi(choice)
    i = i - 1
    length := len(streams)

    if i <= length && i >= 0 {
      OpenStream(streams[i])
    }
  }
}

func Menu() {
  fmt.Println("RPi Twitch Main Menu")
  fmt.Println("Enter 'f' to view your followed channels.")
  fmt.Println("Enter 'g' to view the different games.")
  fmt.Println("Enter 'go' to view CSGO streams.")
  fmt.Println("Press 'q' at any time to quit.")

  reader := bufio.NewReader(os.Stdin)
  choice, _ := reader.ReadString('\n')
  choice = strings.Replace(choice, "\n", "", 1)

  switch choice {
    case "f"  : FollowedStreams()
    case "g"  : TopGames()
    case "go" : CSGOStreams()
  }
}

func main() {
  Menu()
}
