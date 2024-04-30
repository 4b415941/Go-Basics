package main

import "fmt"

type XEncoder struct {
}

type YEncoder struct {
}

type IEncoder interface {
	Encode(value string)
	Decode(value string)
}

func (XEncoder *XEncoder) Encode(value string) {
	fmt.Println("Encoded with XEncoder.")
}

func (XEncoder *XEncoder) Decode(value string) {
	fmt.Println("Decoded with XEncoder.")
}

func (YEncoder *YEncoder) Encode(value string) {
	fmt.Println("Encoded with YEncoder.")
}

func (YEncoder *YEncoder) Decode(value string) {
	fmt.Println("Decoded with YEncoder.")
}

func main() {
	/* var xencoder *XEncoder

	xencoder = &XEncoder{}
	xencoder.Encode("12345")
	xencoder.Decode("12345") */

	/* var yencoder *YEncoder

	yencoder = &YEncoder{}
	yencoder.Encode("12345")
	yencoder.Decode("12345") */

	var encoder IEncoder
	encoder = &YEncoder{}

	encoder.Encode("123354")
	encoder.Decode("123354")
}
