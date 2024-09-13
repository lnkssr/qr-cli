package main

import (
	"fmt"
	"log"
	"os"

	"github.com/skip2/go-qrcode"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: qr-cli <text for qr-code>")
		return
	}

	input := os.Args[1]

	qrCode, err := qrcode.New(input, qrcode.Medium)
	if err != nil {
		log.Fatal(err)
	}

	qrString := qrCode.ToString(false)

	fmt.Println(qrString)
}
