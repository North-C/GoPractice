package centerServer

import "fmt"

// 用户的基本信息
type Player struct{
	Name string
	Level int
	Exp int
	Room int

	mq chan *Message 		//等待收取的消息
}

// 构建新的Player
func NewPlayer() *Player{
	m := make(chan *Message, 1024)
	player := &Player{"", 0, 0, 0, m}

	go func(p *Player){
		for{
			msg := <-p.mq
			fmt.Println(p.Name, "recieved message:", msg.Content)
		}
	}(player)

	return player
}




