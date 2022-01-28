package test

import (
	"fmt"
	"reflect"
	"testing"
)

func Int32ToBytes(digit int32) []byte {
	nn := reflect.TypeOf(digit).Size()
	res := make([]byte, int(nn))
	// fmt.Println(int(nn))
	for i := 0; i < int(nn); i++ {
		// fmt.Printf("%d, %d\n", digit, digit|0xff)
		res[i] = uint8(digit & 0xff)
		digit = digit >> 8
	}
	return res
}

func BytesToInt32(byts []byte) int32 {
	var digit int32 = 0
	for i := 0; i < len(byts); i++ {
		fmt.Printf("%d, %d\n", digit, int32(byts[i]))
		digit = digit | int32(byts[i])<<(i*8)
	}
	return digit
}

func TestIntBytesConvert(t *testing.T) {

	var num int32 = 96
	fmt.Println(Int32ToBytes(num))
	fmt.Println(BytesToInt32(Int32ToBytes(num)))
}
