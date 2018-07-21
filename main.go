package main

import (
  "os"
  "flag"
  "fmt"
  "bufio"
  "io/ioutil"
  "net/url"
)

func isURL(s string) bool {
  if _, err := url.ParseRequestURI(s); err != nil {
    return false
  }
  return true
}

func main() {
  savePath := flag.String("save", "../../saved_music.txt", "Location of place to save music")
  output := flag.Bool("o", false, "Output list instead of writing to it")
  flag.Parse()

  if *output {
    data, err := ioutil.ReadFile(*savePath)
    if err != nil {
      panic(err)
    }
    fmt.Println(string(data))
  } else {
    var file *os.File
    if _, err := os.Stat(*savePath); os.IsNotExist(err) {
      file, err = os.Create(*savePath)
      if err != nil {
        panic(err)
      }
    } else {
      file, err = os.OpenFile(*savePath, os.O_APPEND | os.O_WRONLY, 0666)
      if err != nil {
        panic(err)
      }
    }

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    in := scanner.Bytes()
    if !isURL(string(in)) {
      fmt.Println("Not a URL, exitting...")
      return
    }
    _, err := file.WriteString(string(in) + "\n")
    if err != nil {
      panic(err)
    }
    fmt.Printf("Wrote: %s \n", string(in))
  }
}
