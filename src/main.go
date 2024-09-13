package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/skip2/go-qrcode"
)

func main() {
	consoleFlag := flag.Bool("c", false, "Print QR code to console")
	fileFlag := flag.String("f", "", "File path to save QR code image")

	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Usage: qr-cli [-c | -f <file>] <text for qr-code>")
		return
	}

	input := flag.Args()[0]

	qrCode, err := qrcode.New(input, qrcode.Medium)
	if err != nil {
		log.Fatal(err)
	}

	if *fileFlag != "" {
		err = qrCode.WriteFile(256, *fileFlag)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("QR code saved to file: %s\n", *fileFlag)
	} else if *consoleFlag {
		qrString := qrCode.ToString(false)
		fmt.Println(qrString)
	} else {
		fmt.Println("No valid flag provided. Use -c to print to console or -f <file> to save to a file.")
	}
}
