package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
)

const (
	HOST = "127.0.0.1"	// Local Host
	PORT = ":8080"		// Nombre del puerto
	PROTOCOL = "udp4"		// Nombre del protocolo
)

func main() {

	s, err := net.ResolveUDPAddr(PROTOCOL, HOST+PORT)
	if err != nil {
		fmt.Println("Hubo un error creando el UDP Address:", err)
		return
	}

	c, err := net.DialUDP(PROTOCOL, nil, s)
	if err != nil {
		fmt.Println("Hubo un error creando el Dial UDP:", err)
		return
	}

	fmt.Printf("El servidor UDP es %s\n", c.RemoteAddr().String())	// Solo para tener de Info
	fmt.Println("Tu Address local es", c.LocalAddr().String())		// Solo para tener de Info
	defer c.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingresa tu nombre")
	text, _ := reader.ReadString('\n')
	data := []byte(text)
	_, err = c.Write(data)

	for {
		
		text, _ := reader.ReadString('\n')
		data := []byte(text)
		_, err = c.Write(data)
		if err != nil {
			fmt.Println("Hubo un error en el mensaje:", err)
			return
		}
		
		if text == "Exit\n" {
			return
		}

		buffer := make([]byte, 1024)
		n, _, err := c.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Hubo un error leyendo la respuesta del servidor:", err)
			return
		}

		fmt.Print(string(buffer[0:n]))
	}
}
      
