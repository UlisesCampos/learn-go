/*
The function File.Seek() Seek sets the offset for the next Read or Write on file to offset, interpreted according to whence: 0 means relative to the origin of the file, 1 means relative to the current offset, and 2 means relative to the end. It returns the new offset and an error, if any. 
*/
package main

import (
  "os"
  "fmt"
  "log"
)
func main() {
  file, err := os.Create("test.txt")
  if err != nil {
    log.Fatal(err)
  }
  file, _ = os.Open("test.txt")
  defer file.Close()

  // Offset is how many bytes to move
  // Offset can be positive or negative
  var offset int64 = 5

  // Whence is the point of reference for offset
  var whence int = 0
  newPosition, err := file.Seek(offset, whence)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Just moved to 5:", newPosition)

  // Go back 2 bytes from current position
  newPosition, err = file.Seek(-2, 1)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Just moved back two", newPosition)

  // Find the currrent position by getting the return value from Seek after moving 0 bytes
  currentPosition, err := file.Seek(0, 1)
  fmt.Println("Current position", currentPosition)

  // Go to beginning of file
  newPosition, err = file.Seek(0, 0)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Position after seeking 0, 0:", newPosition)
}
