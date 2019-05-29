package main

import (
	"bufio"
	"fmt"
	"io"
	// "os"
	// "strconv"
	"strings"
)

func roundGrade(g int32) int32 {
	if g>37 && g%5 > 2 {
		return g + 5 - g%5
	}
	return g
}

/*
 * Complete the gradingStudents function below.
 */
func gradingStudents(grades []int32) []int32 {

	for i := 0; i < len(grades); i++ {
		grades[i] = roundGrade(grades[i])
	}
	return grades
}

func main() {
	fmt.Printf("%v\n", gradingStudents([]int32{73, 67, 38, 33}))
	fmt.Printf("%v\n", gradingStudents([]int32{75, 67, 40, 33}))
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
