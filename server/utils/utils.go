package utils

import (
	"fmt"
	"runtime"
	"net"
	"os"
	"os/exec"
)

var clear map[string]func()

func init() {
	clear = make(map[string]func()) //Initialize it
    clear["linux"] = func() { 
        cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        //cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
		cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}

func ShowBoard(board [][]string, posX, posY int) {

	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    } else { //unsupported platform
        panic("Your platform is unsupported! I can't clear terminal screen :(")
    }

	fmt.Println("############")
	for i, y := range board {
		if posY == i {
			fmt.Print("#")
			for j := range y {
				if posX == j {
					fmt.Print("O")
					continue
				}
				fmt.Print(" ")
			}
			fmt.Println("#")
		}
		fmt.Println("#          #")
	}
	fmt.Println("############")
	fmt.Println("Player position x:", posX, "y:", posY)
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