package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	// Write your code here
	var arrayLength = len(arr)
	var positiveCount, negativeCount, zerosCount int

	for _, v := range arr {
		switch {
		case v > 0:
			positiveCount++
		case v < 0:
			negativeCount++
		default:
			zerosCount++
		}
	}
	var proportionPositive float64 = float64(positiveCount) / float64(arrayLength)
	var proportionNegative float64 = float64(negativeCount) / float64(arrayLength)
	var proportionZeros float64 = float64(zerosCount) / float64(arrayLength)

	fmt.Println(proportionPositive)
	fmt.Println(proportionNegative)
	fmt.Println(proportionZeros)
	// fmt.Println(arrayLength)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
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