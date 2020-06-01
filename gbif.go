package gbif

import (
  "encoding/json"
  "fmt"
  "bufio"
  "os"
  // "io"
  "io/ioutil"
  "log"
  "net/http"
  "time"
  "strings"
  termutil "github.com/andrew-d/go-termutil"
)

// func Hello() {
//   fmt.Println("Hello, World!")
// }

type dat struct {
  Key int
  Species string
}

func SpeciesName() {
	url := ""

	if termutil.Isatty(os.Stdin.Fd()) {
    fmt.Println("Nothing on STDIN")
  } else {
   	reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
    url = "https://api.gbif.org/v1/species/" + text
  }

  url = strings.TrimRight(url, "\r\n")
	
  spaceClient := http.Client{
    Timeout: time.Second * 2, // Maximum of 2 secs
  }

  req, err := http.NewRequest(http.MethodGet, url, nil)
  if err != nil {
    log.Fatal(err)
  }

  req.Header.Set("User-Agent", "gbif-testing-from-golang")

  res, getErr := spaceClient.Do(req)
  if getErr != nil {
    log.Fatal(getErr)
  }

  body, readErr := ioutil.ReadAll(res.Body)
  if readErr != nil {
    log.Fatal(readErr)
  }

  out := dat{}
  jsonErr := json.Unmarshal(body, &out)
  if jsonErr != nil {
    log.Fatal(jsonErr)
  }

  fmt.Println(out.Species)
}

// func SpeciesId2Occurrences() {
// 	url := ""

// 	if termutil.Isatty(os.Stdin.Fd()) {
//     fmt.Println("Nothing on STDIN")
//   } else {
//    	reader := bufio.NewReader(os.Stdin)
// 		text, _ := reader.ReadString('\n')
//     url = "https://api.gbif.org/v1/species/" + text
//   }

//   url = strings.TrimRight(url, "\r\n")
	
//   spaceClient := http.Client{
//     Timeout: time.Second * 2, // Maximum of 2 secs
//   }

//   req, err := http.NewRequest(http.MethodGet, url, nil)
//   if err != nil {
//     log.Fatal(err)
//   }

//   req.Header.Set("User-Agent", "gbif-testing-from-golang")

//   res, getErr := spaceClient.Do(req)
//   if getErr != nil {
//     log.Fatal(getErr)
//   }

//   body, readErr := ioutil.ReadAll(res.Body)
//   if readErr != nil {
//     log.Fatal(readErr)
//   }

//   out := dat{}
//   jsonErr := json.Unmarshal(body, &out)
//   if jsonErr != nil {
//     log.Fatal(jsonErr)
//   }

//   fmt.Println(out.Species)
// }
