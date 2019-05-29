package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the staircase function below.
func staircase(n int32) {

  for i := int32(1) ; i <= n ; i++ {
    for j:=n-i ; j>0 ; j-- {
      fmt.Printf(" ")
    }
    for k:=int32(0) ; k<i ; k++ {
      fmt.Printf("#")
    }
    fmt.Printf("\n")
  }
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    // fmt.Println("how many steps for the staircase?")
    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    staircase(n)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
