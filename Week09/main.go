package main

import (
	"bufio"
	"net"
	"fmt"
)
//用 Go 实现一个 tcp server ，用两个 goroutine 读写 conn，两个 goroutine 通过 chan 可以传递 message，能够正确退出
type Message struct {
	MsgChan chan string
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Printf("listen error:%v\n", err)
		return
	}
	fmt.Println("程序启动端口：8000")
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept error:%v\n", err)
			continue
		}
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	defer conn.Close()

	rd := bufio.NewReader(conn)

	msg := &Message{make(chan string, 8)}
	go sendMsg(conn, msg.MsgChan)

	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			fmt.Printf("read error:%v\n", err)
			return
		}

		msg.MsgChan <- string(line)
	}
}

func sendMsg(conn net.Conn, ch <-chan string) {
	wr := bufio.NewWriter(conn)

	for msg := range ch {
		//将读到的信息，发出去
		wr.WriteString(msg)
		wr.Flush()
	}
}
