package errors

import (
  "fmt"
  "os"
)

func HandleError(err error, msg string) {
  if err != nil {
    fmt.Println(msg)
    os.Exit(1)
  }
}
