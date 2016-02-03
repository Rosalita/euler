package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
)

func main(){

  var rawGrid []byte
  var err error

  // read the grid to be solved from a text file
  rawGrid, err = ioutil.ReadFile("grid.txt")
  if err != nil {
    panic(err)
  }

  //convert the raw bytes into string format
  stringGrid := string(rawGrid)

  //break the raw string down into a slice containing of rows of numbers
  gridRows := strings.Split(stringGrid, "\n")

  // gridRows on windows will still contain "\r" carriage returns
  // These extra "\r" won't be present on linux because it is a better OS
  // be nice to anyone running this code on windows by looping through trimming off the "\r" white space

 for i := 0; i < len(gridRows); i ++ {
    fmt.Println(gridRows[i])
    gridRows[i] = strings.TrimSpace(gridRows[i])
  }
  checkHorizontal(gridRows)

}

func checkHorizontal(gridRows []string) int64{
  for i := 0; i < len(gridRows); i ++ {
    lineToCheck := strings.Split(gridRows[i], " ")
    fmt.Println(lineToCheck)

    int1, err:= strconv.ParseInt(lineToCheck[0], 10, 64)
    if err != nil {
       panic(err)
    }
    int2, err:= strconv.ParseInt(lineToCheck[1], 10, 64)
    if err != nil {
       panic(err)
    }
    int3, err:= strconv.ParseInt(lineToCheck[2], 10, 64)
    if err != nil {
       panic(err)
    }
    int4, err:= strconv.ParseInt(lineToCheck[3], 10, 64)
    if err != nil {
       panic(err)
    }
    fmt.Println(int1, int2, int3, int4)

  }


  return 1
}
