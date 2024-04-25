package main

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"go/structures"
)

var logger = logrus.New()
var q = &structures.Queue{}

func intToBytes(num int) []byte {
	converted := make([]byte, 4)
	for i := 0; i < 4; i++ {
		converted[i] = byte(num >> (i * 8) & 0xFF)
	}
	return converted
}

func bytesToInt(bytes []byte) int {
	var num int
	for i := 0; i < 4; i++ {
		num |= int(bytes[i]) << (i * 8)
	}
	return num
}

func main() {
	logger.Infof("Init main program")

	// Exemplo de uso da fila
	// q.Enqueue(intToBytes(5))
	// q.Enqueue(intToBytes(3))
	q.Enqueue([]byte{1, 2, 3})
	q.Enqueue([]byte{4, 5, 6})
	q.Enqueue([]byte{7, 8, 9})

	// Verificando o tamanho da fila
	fmt.Println("Tamanho da fila:", q.GetSize())

	// Removendo elementos da fila (dequeue) e imprimindo-os
	for !q.IsEmpty() {
		item, _ := q.Dequeue()
		fmt.Println("Elemento removido da fila:", item)
	}

}
