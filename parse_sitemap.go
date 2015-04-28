package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "encoding/xml"
  "strings"
  )

type Urlset struct {
  XMLName xml.Name `xml:"urlset"`
  Urls []Url `xml:"url"`
}

type Url struct {
    Loc string `xml:"loc"`
    ParsedUrl string
}

func (s Url) String() string {
        return fmt.Sprintf("\t Loc : %s \n", s.Loc)
}

func main () {
  filename := "sitemap.xml"
  baseUrl := "http://BASE_URL/#"
  xmlFile, err := os.Open(filename)
  if err != nil {
    fmt.Println("Error opening file :", err )
    return
  }
  defer xmlFile.Close()

  XMLdata, _ := ioutil.ReadAll(xmlFile)

  var us Urlset
  xml.Unmarshal(XMLdata, &us)
  //fmt.Println(u)

  for _,u := range us.Urls{
    s := strings.Split(u.Loc, "/")
    if len(s) <=3 {
      continue
    }
    u.ParsedUrl = baseUrl + s[3] + "/" + s[len(s) - 1]
    fmt.Println(u.ParsedUrl)
  }
}
