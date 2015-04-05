package structs

// Reusable JSON response shapes
type Channel_info struct {
  Game string
  Name string
  Status string
  Url string
}

// JSON response shape for /channels
type Streams struct {
  Streams []Stream_info
}

type Stream_info struct {
  Channel Channel_info
}

// JSON response shape for /user/{:username}/follows/channels
type Follows struct {
  Follows []Follow_info
}

type Follow_info struct {
  Channel Channel_info
}

// Internal API shapes
type FormattedStream struct {
  Game string
  Name string
  Title string
  Url string
}
