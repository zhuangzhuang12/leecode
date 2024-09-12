package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

// CheckRedisProtocol checks if the given data is a complete Redis protocol message.
func CheckRedisProtocol(data []byte) (bool, error) {
	reader := bufio.NewReader(bytes.NewReader(data))

	for {
		// Read the first byte to determine the type
		b, err := reader.ReadByte()
		if err != nil {
			return false, nil // not enough data, might be incomplete
		}

		switch b {
		case '+', '-', ':': // Simple Strings, Errors, Integers
			_, err := reader.ReadString('\n')
			if err != nil {
				return false, nil // incomplete data
			}

		case '$': // Bulk Strings
			lenStr, err := reader.ReadString('\n')
			if err != nil {
				return false, nil // incomplete data
			}

			length, err := strconv.Atoi(lenStr[:len(lenStr)-2])
			if err != nil || length < -1 {
				return false, errors.New("invalid bulk string length")
			}

			if length >= 0 {
				buf := make([]byte, length+2) // +2 for \r\n
				_, err := reader.Read(buf)
				if err != nil {
					return false, nil // incomplete data
				}
			}

		case '*': // Arrays
			lenStr, err := reader.ReadString('\n')
			if err != nil {
				return false, nil // incomplete data
			}

			count, err := strconv.Atoi(lenStr[:len(lenStr)-2])
			if err != nil || count < -1 {
				return false, errors.New("invalid array length")
			}

			for i := 0; i < count; i++ {
				complete, err := CheckRedisProtocol(data[len(data)-reader.Buffered():])
				if err != nil || !complete {
					return false, nil // incomplete data
				}
			}

		default:
			return false, errors.New("unknown type")
		}

		// If we've read all data, return true
		if reader.Buffered() == 0 {
			return true, nil
		}
	}
}

func main() {
	// Example data
	data := []byte("*2\r\n$3\r\nSET\r\n$5\r\nmykey\r\n")

	complete, err := CheckRedisProtocol(data)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else if complete {
		fmt.Println("The message is complete")
	} else {
		fmt.Println("The message is incomplete")
	}
}
