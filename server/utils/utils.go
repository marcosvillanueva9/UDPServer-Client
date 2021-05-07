package utils

import (
	"fmt"
	"net"
	"os"
	"os/exec"
)

func ShowBoard(board [][]int, posX, posY int) {
	cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
	fmt.Println(board)
	fmt.Println("############")
	for _, y := range board {
		fmt.Println("#          #", y)
	}
	fmt.Println("############")
}
    
func GetUDPConn(protocol string, port string) (*net.UDPConn, error){
	UDPAddr, err := net.ResolveUDPAddr(protocol, port)
	if err != nil {
		return nil, err
	}

	conn, err := net.ListenUDP(protocol, UDPAddr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}