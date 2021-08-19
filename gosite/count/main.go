//count подсчитвает количество вхождений
//каждой строки в файле

package main

import (
  "fmt"
  "github/gosite/datafile"
  "log"
)

func main () {
  //загружает файл votes, разбиравет в нем строки и сохр в массив
  lines, err := datafile.GetStrings("votes.txt")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(lines)
}
