package helper

import (
	"fmt"
	"testing"
)

func TestBcd2Str(t *testing.T) {
	Bcd2Str([]byte("skehen"))
}

func TestBcd2Bytes(t *testing.T) {
	fmt.Println(string([]byte{107, 24, 30}))
	//fmt.Println([]byte("[B@610455d6"))
	//b := Bcd2Bytes([]byte("skehen"))
	//fmt.Println(b)
	//fmt.Println(string(b))
}
