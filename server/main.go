package main

import (
        "fmt"
        "math/rand"
        "net"
        "time"
)

const (
	PORT = ":8080"			// Puerto donde se va a levantar el servidor UDP
	PROTOCOL = "udp4"		// Nombre del protocolo
)

func main() {

	UDPAddr, err := net.ResolveUDPAddr(PROTOCOL, PORT)
	if err != nil {
		fmt.Println("Hubo un error creando el UDP Address:", err)
		return
	}

	conn, err := net.ListenUDP(PROTOCOL, UDPAddr)
	if err != nil {
		fmt.Println("Hubo un error creando la conexion UPD:", err)
		return
	}

	defer conn.Close()
	buffer := make([]byte, 1024)	// Tamaño del buffer (bytes)
	rand.Seed(time.Now().Unix())

	for {
		n, addr, err := conn.ReadFromUDP(buffer)	// Lectura del UDP

		mensaje := string(buffer[0:n-1])	// Capturar mensaje del buffer
		fmt.Println("Mensaje", mensaje)

		data := []byte("Recibido!")	// Mensaje al cliente
		_, err = conn.WriteToUDP(data, addr)	// Enviar mensaje al cliente
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
    
