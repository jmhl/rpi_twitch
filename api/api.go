package api

import (
  "fmt"
  "../format"
  "../get_json"
  "../structs"
  "../user"
)

func GetFollows() []structs.FormattedStream {
  url := fmt.Sprintf("https://api.twitch.tv/kraken/users/%s/follows/channels", user.Name())
  res := twitch_request.GetJSON(url)
  json := twitch_request.DecodeFollowsResponse(res)
  streams := format.FormatFollows(json)

  return streams
}

func GetStreams(limit int, offset int) []structs.FormattedStream {
  url := fmt.Sprintf("https://api.twitch.tv/kraken/streams?game=counter-strike:%sGlobal%sOffensive&limit=%d&offset=%d", "%20", "%20", limit, offset)
  res := twitch_request.GetJSON(url)
  json := twitch_request.DecodeStreamsResponse(res)
  streams := format.FormatStreams(json)

  return streams
}
