package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
  "os/user"
  "path/filepath"
)

func isURL(s string) bool {
	if _, err := url.ParseRequestURI(s); err != nil {
		return false
	}
	return true
}

func main() {
  usr, err := user.Current()
  if err != nil {
    panic(err)
  }
	savePath := flag.String("save", filepath.Join(usr.HomeDir, "saved_music.txt"), "Location of place to save music")
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
      fmt.Printf("Created saved_music.txt in %s \n", *savePath)
			file, err = os.Create(*savePath)
			if err != nil {
				panic(err)
			}
		} else {
			file, err = os.OpenFile(*savePath, os.O_APPEND|os.O_WRONLY, 0666)
			if err != nil {
				panic(err)
			}
		}

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			in := scanner.Text()
			if !isURL(in) {
				fmt.Println("Not a URL, skipping...")
				continue
			}
			_, err := file.WriteString(in + "\n")
			if err != nil {
				panic(err)
			}
			fmt.Printf("Wrote: %s \n", in)
		}
	}
}
