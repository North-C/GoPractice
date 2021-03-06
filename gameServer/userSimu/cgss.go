package main

import (
	"bufio"
	"fmt"
	"gameserver/centerServer"
	"gameserver/ipc"
	"os"
	"strconv"
	"strings"
)

var centerClient *centerServer.CenterClient

func startCenterService() error{
	server := ipc.NewIpcServer(&centerServer.CenterServer{})
	client := ipc.NewIpcClient(server)
	centerClient = &centerServer.CenterClient{client}

	return nil
}

func Help(args []string)int{
	fmt.Println(
		`Command:
				login<username><level><exp>
				logout<username>
				send<message>
				listplayer
				quit(q)
				help(h)
		`)
	return 0
}


func Quit(args []string)int{
	return 1
}

func Logout(args []string) int {
	if len(args) != 2{
		fmt.Println("USAGE:logout <username>")
		return 0
	}
	centerClient.RemovePlayer(args[1])
	return 0
}

func Login(args []string) int{
	if len(args) != 4{
		fmt.Println("USAGE: login<username><level><exp>")
		return 0
	}

	level, err := strconv.Atoi(args[2])
	if err != nil{
		fmt.Println("Invalid Parameter: <level> should be an integer.")
		return 0
	}
	exp, err := strconv.Atoi(args[3])
	if err != nil{
		fmt.Println("Invalid Parameter: <exp> should be an integer.")
		return 0
	}

	player := centerServer.NewPlayer()
	player.Name = args[1]
	player.Level = level
	player.Exp = exp

	err = centerClient.AddPlayer(player)
	if err != nil{
		fmt.Println("Failed to add player", err)
	}
	return 0
}

// 将输入的指令和函数直接对应起来
func GetCommandHandlers() map[string]func(args []string)int{
	return map[string]func(args []string)int{
		"help":       Help,
		"H":          Help,
		"quit":       Quit,
		"q":          Quit,
		"login":      Login,
		"logout":     Logout,
		"listplayer": ListPlayer,
		"send":       Send,
	}
}

func ListPlayer(args []string)int{
	ps, err := centerClient.ListPlayer("")
	if err != nil{
		fmt.Println("Failed.", err)		
	}else{
		for i, v := range ps{
			fmt.Println(i+1, ":", v)
		}
	}
	return 0
}

func Send(args []string)int{
	message := strings.Join(args[1:], " ")

	err := centerClient.Broadcast(message)
	if err != nil{
		fmt.Println("Failed.", err)
	}
	return 0
}


func main() { 
	fmt.Println("Casual Game Server Solution")
	startCenterService()
	Help(nil)
	r := bufio.NewReader(os.Stdin)
	handlers := GetCommandHandlers()
	for{		// 循环读入用户输入
		fmt.Print("Command> ")
		b, _, _ := r.ReadLine()
		line := string(b)		// 读取分解用户参数
		tokens :=  strings.Split(line, " ")

		if handler, ok := handlers[tokens[0]]; ok{
			ret := handler(tokens)
			if ret != 0{
				break
			}
		}else{
			fmt.Println("Unknown command:", tokens[0])
		}
	}
}


