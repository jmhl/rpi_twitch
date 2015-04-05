package format

import "../structs"

func FormatStreams(json structs.Streams) []structs.FormattedStream {
  var streams []structs.FormattedStream

  for _, stream_info := range json.Streams {
    stream := structs.FormattedStream{
      Game: stream_info.Channel.Game,
      Name: stream_info.Channel.Name,
      Title: stream_info.Channel.Status,
      Url: stream_info.Channel.Url,
    }

    streams = append(streams, stream)
  }

  return streams
}

func FormatFollows(json structs.Follows) []structs.FormattedStream {
  var streams []structs.FormattedStream

  for _, stream_info := range json.Follows {
    stream := structs.FormattedStream{
      Game: stream_info.Channel.Game,
      Name: stream_info.Channel.Name,
      Title: stream_info.Channel.Status,
      Url: stream_info.Channel.Url,
    }

    streams = append(streams, stream)
  }

  return streams
}
