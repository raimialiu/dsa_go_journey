package main

import (
	"dsa_go_journey/problems/arrays"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	/*
		I begin my DSA journey again today January 13th, 2026. I want to get cracked and become excellent in all ramifications.
		I feel like I dont know anything and I getting wasting away. So I am going to give my brain and myself some work over the course of
		the next four years to become really excellent. So the first the I want to do go crack DSA.
		So I created this repo to track my journey. This in anyway will not have any AI content, it will capture my progress in this journey
		and my DSA journey. I am going to ensure I understand my domain space deeply and properly.

		DELIBERATE INTENTION BRO, VERY IMPORTANT....
	*/
	//fmt.Println(hex("255"))
	//in2dex := 42
	//	response := int2bin(in2dex, "")

	//fmt.Println(fmt.Sprintf("response: %v", response))

	//fmt.Println("Let us now print reverse")
	//reversedValue := bin2decimal(response)
	//fmt.Println(fmt.Sprintf("reversedValue: %v", reversedValue))

	fmt.Println("Two Sum Implementations")
	result := arrays.TwoSums([]int{1, 2, 3, 4, 5, 6}, 9)
	fmt.Println("Result:", result)
}

/*
* We understand that hex is 4 bit(4 binary digit) and that it combines zero and alphabet.
* 0-9 and A -F (A=10,B=11,C=12,D=13,E=14,F=15). So I am going to use the knowledge of this
* to build a simple function that can convert a string to hex equivalent.
Edge Case => what about other alphabets (I'll treat later)
*/

func hex(value interface{}) string {
	// PSEUDOCOES
	// create a local mapping of possible value equivalent in hex
	// split apart
	// for each of item in results of split, do

	var result = ""
	mapping := map[string]interface{}{
		"A": 10,
		"B": 11,
		"C": 12,
		"D": 13,
		"E": 14,
		"F": 15,
	}

	valueType := reflect.TypeOf(value)
	if valueType.Kind() == reflect.String {
		stringValue := value.(string)
		values := strings.Split(stringValue, "")
		for _, char := range values {
			if _, err := strconv.Atoi(char); err != nil {
				charToUpper := strings.ToUpper(char)
				if _, ok := mapping[charToUpper]; ok {
					result += charToUpper
				}
			}

			result += char
		}
	}
	return "0x" + result
}

func bin2decimal(value string) string {
	// create a variable for response
	// split the input by empty space
	// get the length of the string
	// start a loop, for each value of the current iteration
	// do value * power(2, lentgh-1)
	// add everything together, this will be your response

	response := 0
	values := strings.Split(value, "")
	length := len(values)

	for _, char := range values {
		length -= 1
		if integerValue, err := strconv.Atoi(char); err != nil {
			panic(err)
		} else {
			power := math.Pow(2, float64(length))
			currentValue := float64(integerValue) * power
			response += int(currentValue)
		}
	}

	return fmt.Sprintf("%d", response)
}

func int2bin(value int, response string) string {
	// PSEUDO
	// 1. create a response variable, set it to empty string, convert input to string
	// 2. get the length of the string
	// 3. start counting from the end (not needed)
	// 4. start loop(while),
	// 5. divide by 2, keep reminder, append to response variable
	// 6. continue until divisor is  zero

	// 1 Create response variable, set it to empty string, convert input to string
	if value <= 0 {
		values := strings.Split(response, "")
		var actualResponse = ""
		counter := len(values)
		for i := counter; i < 0; i-- {
			actualResponse = actualResponse + values[i]
		}
		return actualResponse
	}

	reminder := value % 2
	response += strconv.Itoa(reminder)
	return int2bin(value/2, response)
}
