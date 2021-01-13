package main

import "fmt"

const (
	mod            = 20201227
	pubKeyDoor     = 8335663
	pubKeyCard     = 8614349
	initialSubject = 7
)

func main() {
	loopSizeDoor := calculateLoopSize(initialSubject, pubKeyDoor)

	fmt.Printf("loop size door: %d\n", loopSizeDoor)
	fmt.Printf("key: %d\n", encryptionKey(pubKeyCard, loopSizeDoor))
}

func transform(subject, value int) int {
	return (value * subject) % mod
}

func calculateLoopSize(subject, publicKey int) int {
	value := 1

	for loopSize := 1; ; loopSize++ {
		value = transform(subject, value)
		if value == publicKey {
			return loopSize
		}
	}
}

func encryptionKey(subject, loopSize int) int {
	var value int = 1
	for i := 0; i < int(loopSize); i++ {
		value = transform(subject, value)
	}

	return value
}
