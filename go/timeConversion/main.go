package main

import (
    "fmt"
    "strconv"
)

/*
 * Complete the timeConversion function below.
 */
func timeConversion(s string) string {
  hours,_ := strconv.Atoi(s[0:2])

  if "PM" == s[8:10] && hours != 12 {
    hours+=12
  }
  if hours == 12 && "AM" == s[8:10] {
    hours=0
  }
  return fmt.Sprintf("%02d%s", hours,s[2:8])
}

func main() {
  fmt.Printf("07:05:45PM = %s\n", timeConversion("07:05:45PM"))
  fmt.Printf("07:05:45AM = %s\n", timeConversion("07:05:45AM"))
  fmt.Printf("12:05:45AM = %s\n", timeConversion("12:05:45AM"))
  fmt.Printf("12:05:45PM = %s\n", timeConversion("12:05:45PM"))

}

// func main() {
//     reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

//     outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
//     checkError(err)

//     defer outputFile.Close()

//     writer := bufio.NewWriterSize(outputFile, 1024 * 1024)

//     s := readLine(reader)

//     result := timeConversion(s)

//     fmt.Fprintf(writer, "%s\n", result)

//     writer.Flush()
// }

// func readLine(reader *bufio.Reader) string {
//     str, _, err := reader.ReadLine()
//     if err == io.EOF {
//         return ""
//     }

//     return strings.TrimRight(string(str), "\r\n")
// }

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
