package main

import (
	"fmt"
	"github.com/marcosvillanueva9/UDPServer-Client/server/utils"
)

const (
	PORT = ":8080"			// Puerto donde se va a levantar el servidor UDP
	PROTOCOL = "udp4"		// Nombre del protocolo
)

func main() {

	conn, err := utils.GetUDPConn(PROTOCOL, PORT)
	if err != nil {
		fmt.Println("Hubo un error creando la conexion UPD:", err)
		return
	}

	defer conn.Close()

	buffer := make([]byte, 1024)	// Tamaño del buffer (bytes)
	addrs := make(map[string]string)
	playerPosX := 0
	playerPosY := 0
	board := make([][]int, 10, 10)

	for {
		utils.ShowBoard(board, playerPosX, playerPosY)

		n, addr, err := conn.ReadFromUDP(buffer)	// Lectura del UDP
		if n < 1 {
			fmt.Println("Hubo un error leyendo el mensaje")
			return
		}
		mensaje := string(buffer[0:n-1])	// Capturar mensaje del buffer

		if _, ok := addrs[addr.String()]; !ok {
			addrs[addr.String()] = mensaje
		}

		if mensaje == "Exit" {
			fmt.Println("Usuario", addrs[addr.String()], "desconectado")
			fmt.Println(addrs[addr.String()])
			delete(addrs, addr.String())
		}

		//fmt.Println("Clientes conectados:", len(addrs))

		data := []byte(fmt.Sprint("Bienvenido", addrs[addr.String()]))	// Mensaje al cliente
		_, err = conn.WriteToUDP(data, addr)	// Enviar mensaje al cliente
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}