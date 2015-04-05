package twitch_request

import (
  "encoding/json"
  "io/ioutil"
  "net/http"
  "../errors"
  "../structs"
)

const msg = "Something went wrong, please try again in a few minutes."

func GetJSON(url string) []byte {
  res, err := http.Get(url)
  errors.HandleError(err, msg)
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  errors.HandleError(err, msg)

  return body
}

func DecodeStreamsResponse(body []byte) structs.Streams {
  var data structs.Streams
  err := json.Unmarshal(body, &data)
  errors.HandleError(err, msg);

  return data
}

func DecodeFollowsResponse(body []byte) structs.Follows {
  var data structs.Follows
  err := json.Unmarshal(body, &data)
  errors.HandleError(err, msg);

  return data
}
