package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"golang.org/x/net/websocket"
)

// 单机端口数为 65535 所以单机最多同时连接数为 65535

var origin = "http://127.0.0.1:8080/"

var url = "ws://127.0.0.1:9501/001"

var userNum = 10000

func main() {

	for i := 0; i <= userNum; i++ {
		// 每一微妙新增一个用户
		time.Sleep(1 * time.Millisecond)
		// 用户操作
		// 连接 后等待 2s 发送信息，接收信息， 等待2s 挂断连接
		go func(i int) {
			ws, err := websocket.Dial(url, "", origin)

			if err != nil {
				log.Fatal(err)
			}
			var str = strconv.Itoa(i)
			message := []byte("hello, world!你好" + str)

			fmt.Println("conn:" + str)
			time.Sleep(2 * time.Second)

			_, err = ws.Write(message)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Send: %s\n", message)

			var msg = make([]byte, 512)
			m, err := ws.Read(msg)
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(2 * time.Second)
			fmt.Printf("Receive: %s\n", msg[:m])
			ws.Close() //关闭连接
			//_ = ws
		}(i)
	}

	select {
	// loop 等待操作执行完毕，需手动关闭
	}
}
