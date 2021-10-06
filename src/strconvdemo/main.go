package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := true
	//i := 65
	f := 123.123

	formatBool := strconv.FormatBool(b)
	fl := strconv.FormatFloat(f, 'f', -1, 64)
	fmt.Println(formatBool)
	fmt.Println(fl)

	bStr := "true"
	fStr := "123.123"
	iStr := "-1234500000000000000000000000"
	parseBool, _ := strconv.ParseBool(bStr)
	parseFloat, _ := strconv.ParseFloat(fStr, 64)
	parseInt, err := strconv.ParseInt(iStr, 10, 0)
	fmt.Println(err)
	fmt.Println(^int(^uint(0) >> 1))
	formatUint := strconv.FormatUint(uint64(^uint(0)>>1), 2)
	fmt.Println(formatUint)
	fmt.Printf("%T %v\n", parseBool, parseBool)
	fmt.Printf("%T %v\n", parseFloat, parseFloat)
	fmt.Printf("%T %v\n", parseInt, parseInt)
}
