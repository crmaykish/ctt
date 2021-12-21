package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tarm/serial"
)

const DEFAULT_BAUDRATE = 115200
const MAGIC_BYTE = 0xDE

func main() {
	var port string
	var fileName string
	var baudrate int

	// Parse command line arguments
	flag.StringVar(&port, "p", "", "Serial port")
	flag.IntVar(&baudrate, "b", DEFAULT_BAUDRATE, "Baudrate")
	flag.StringVar(&fileName, "f", "", "File name")
	flag.Parse()

	// Must provide a serial port and filename
	if len(port) == 0 || len(fileName) == 0 {
		fmt.Println("Colin's Transfer Tool")
		fmt.Println("Usage: ctt -p <port> -f <file> -b [baudrate]")
		os.Exit(1)
	}

	// Open the serial port at the given baudrate
	serialConfig := &serial.Config{Name: port, Baud: baudrate}
	serialPort, err := serial.OpenPort(serialConfig)

	if err != nil {
		fmt.Printf("Failed to open serial port: %s\r\n", err)
		os.Exit(1)
	}

	fmt.Printf("Connected to %s at %d baud\r\n", port, baudrate)

	// Read the file for transfer
	file, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Printf("Failed to open file: %s\r\n", err)
		os.Exit(1)
	}

	fmt.Printf("Starting transfer...\r\n")

	// Write the magic start byte
	serialPort.Write([]byte{MAGIC_BYTE})

	// Write the file bytes
	bytesWritten, err := serialPort.Write(file)

	if err != nil {
		fmt.Printf("Failed to transfer bytes: %s\r\n", err)
		os.Exit(1)
	}

	fmt.Printf("Wrote %d bytes\r\n", bytesWritten)

	// Write the magic end bytes
	serialPort.Write([]byte{MAGIC_BYTE, MAGIC_BYTE, MAGIC_BYTE})

	fmt.Println("Done!")
}
