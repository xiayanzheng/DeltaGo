package dapr

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

func IntToString(data int) string {
	return strconv.Itoa(data)
}

func StringToInt(data string) int {
	Int, err := strconv.Atoi(data)
	if err != nil {
		log.Fatal(err)
	}
	return Int
}

func GetValType(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println("type:", reflect.TypeOf(arg))
			//fmt.Println(arg, "is an unknown type.")
		}
	}
}
