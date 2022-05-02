package main

/*
	Go's basic types are

	- bool
		Default value is "false"
	- string
		Default value is ""
	- int	int8	int16	int32	int64
	- uint	uint8	uint16	uint32	uint64	uintptr
		Default value is 0
	- byte // alias for uint8
		Default value is 0
	- rune // alias for int32
		   // represent a Unicode point
		Default value is 0
	- float32	float64
		Default value is 0
	- complex64		complex128
		Default value is "(0 + 0i)"


	constant "value" overflows "type"
*/

import "fmt"

func main() {
	// Basic initialization
	var isBool bool
	var thisString string
	var signedInteger int32
	var unsignedInt uint32
	var byteNumber byte
	var runeNumber rune
	var floatNumber float32
	var complexNumber complex64

	isBool = true
	thisString = "This is a string"
	signedInteger = -123
	unsignedInt = 456
	byteNumber = 255
	runeNumber = 987
	floatNumber = 123.123

	fmt.Printf("Type: %T Value: %v\n", isBool, isBool)
	fmt.Printf("Type: %T Value: %v\n", thisString, thisString)
	fmt.Printf("Type: %T Value: %v\n", signedInteger, signedInteger)
	fmt.Printf("Type: %T Value: %v\n", unsignedInt, unsignedInt)
	fmt.Printf("Type: %T Value: %v\n", byteNumber, byteNumber)
	fmt.Printf("Type: %T Value: %v\n", runeNumber, runeNumber)
	fmt.Printf("Type: %T Value: %v\n", floatNumber, floatNumber)
	fmt.Printf("Type: %T Value: %v\n", complexNumber, complexNumber)

	convertedInt := int32(floatNumber)
	fmt.Printf("Type: %T Value: %v\n", convertedInt, convertedInt)

}
