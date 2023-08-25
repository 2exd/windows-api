package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strconv"
	"time"
	"windows-api/funcs"
)

// 全局客户端map
var clients = make(map[*websocket.Conn]bool)
var lastMod time.Time

// 全局广播消息通道
var broadcast = make(chan []byte)
var upgrader = websocket.Upgrader{} // use default options

func pic(w http.ResponseWriter, r *http.Request) {
	// 升级 HTTP 连接为 WebSocket 连接
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	// 连接成功时将客户端添加到全局 map
	clients[c] = true
	defer func() {
		// 连接关闭时从全局map删除客户端
		delete(clients, c)
		c.Close()
	}()

	// defer client.Close()

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		rgba := funcs.DecodeToRGBA(message)
		funcs.SaveScreen(&rgba)
		// log.Printf("recv: %s", message)
		// err = client.WriteMessage(mt, message)
		// if err != nil {
		// 	log.Println("write:", err)
		// 	break
		// }
	}
}

// WebSocket服务器处理程序
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// 升级HTTP连接为WebSocket连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// 连接成功时将客户端添加到全局 map
	clients[conn] = true
	defer func() {
		// 连接关闭时从全局map删除客户端
		delete(clients, conn)
		conn.Close()
	}()

	if n, err := strconv.ParseInt(r.FormValue("lastMod"), 16, 64); err == nil {
		lastMod = time.Unix(0, n)
	}

	/*	ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()*/

	for {
		select {
		case data := <-funcs.FileChange:
			// 将消息发送到广播通道，以便它可以被广播到所有客户端
			log.Println("Message: ", string(data))
			broadcast <- data
			// case t := <-ticker.C:
			// 	// 将消息发送到广播通道，以便它可以被广播到所有客户端
			// 	log.Println("Message: ", "message"+t.String())
			// 	broadcast <- []byte("message" + t.String())
		}
	}

}
func main() {
	var err error
	log.SetFlags(0)
	http.HandleFunc("/pic", pic)
	http.HandleFunc("/bc", handleWebSocket)
	go func() {
		err := http.ListenAndServe(viper.GetString("ip")+":"+viper.GetString("port"), nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	// 不断处理广播通道上的消息，并向所有客户端发送
	for {
		select {
		case <-ticker.C:
			lastMod, err = funcs.ReadFileIfModified(lastMod, viper.GetString("file"))
			if err != nil {
				log.Println(err)
			}
		case message := <-broadcast:
			for client := range clients {
				err := client.WriteMessage(websocket.TextMessage, message)
				if err != nil {
					log.Println(err)
					client.Close()
					delete(clients, client)
				}
			}
		}
	}
}

func init() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./server.yaml") // 注意:如果使用相对路径，则是以 main.go 为当前位置与配置文件之间的路径
	err := viper.ReadInConfig()          // 查找并读取配置文件
	if err != nil {                      // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
