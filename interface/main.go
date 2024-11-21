package main

import "fmt"

type XEncoder struct {

}

func (xEncoder *XEncoder) encoder(password string) {
	fmt.Println("It's encoded by XEncoder", password)
}

func (XEncoder *XEncoder) decoder(password string){
	fmt.Println("It's decoded by Xencoder", password)
}

func main()  {

	var xEncoder *XEncoder = &XEncoder{}
	xEncoder.encoder("123456")
	xEncoder.decoder("123456")
}
