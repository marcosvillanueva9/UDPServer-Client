package main

import (
	"fmt"
	"github.com/marcosvillanueva9/UDPServer-Client/server/utils"
)

const (
	PORT = ":8080"			// Puerto donde se va a levantar el servidor UDP
	PROTOCOL = "udp4"		// Nombre del protocolo
)

var (
	playerPosX = 0
	playerPosY = 0
)

func main() {

	conn, err := utils.GetUDPConn(PROTOCOL, PORT)
	if err != nil {
		fmt.Println("Hubo un error creando la conexion UPD:", err)
		return
	}

	defer conn.Close()

	buffer := make([]byte, 1024)	// Tamaño del buffer (bytes)
	addrs := make(map[string][2]int)

	board := [][]string{
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	}

	for {
		utils.ShowBoard(board, playerPosX, playerPosY)

		n, addr, err := conn.ReadFromUDP(buffer)	// Lectura del UDP
		if n < 1 {
			fmt.Println("Hubo un error leyendo el mensaje")
			return
		}
		mensaje := string(buffer[0:n-2])	// Capturar mensaje del buffer

		if _, ok := addrs[addr.String()]; !ok {
			addrs[addr.String()] = [2]int{0,0}
		}

		if mensaje == "Exit" {
			fmt.Println("Usuario", addrs[addr.String()], "desconectado")
			fmt.Println(addrs[addr.String()])
			delete(addrs, addr.String())
		}

		switch {
			case mensaje == "a":
				if playerPosX != 0 {
					playerPosX--
				}
			case mensaje == "d":
				if playerPosX != 9 {
					playerPosX++
				}
			case mensaje == "w":
				if playerPosY != 0 {
					playerPosY--
				}
			case mensaje == "s":
				if playerPosY != 9 {
					playerPosY++
				}
			default: 
		}

		//fmt.Println("Clientes conectados:", len(addrs))

		data := []byte(fmt.Sprint("Posicione de", addrs[addr.String()], "posx", playerPosX, "posy", playerPosY, "tu mensaje:", mensaje))	// Mensaje al cliente
		_, err = conn.WriteToUDP(data, addr)	// Enviar mensaje al cliente
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}