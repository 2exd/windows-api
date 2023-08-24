package main

import (
	"fmt"
	"github.com/2exd/windows-api/funcs"
	"github.com/atotto/clipboard"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
)

var done = make(chan struct{})
var interrupt = make(chan os.Signal, 1)

func main() {
	log.SetFlags(0)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: viper.GetString("addr"), Path: "/pic"}
	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	go receiveCode()
	go funcs.GetScreen()

	for {
		select {
		case img := <-funcs.ImgChan:
			data := funcs.EncodeToBytes(img)
			err := conn.WriteMessage(websocket.BinaryMessage, data)
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-done:
			return
		case <-interrupt:
			log.Println("interrupt")
			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

func receiveCode() {
	u := url.URL{Scheme: "ws", Host: viper.GetString("addr"), Path: "/bc"}
	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
			clipboard.WriteAll(string(message))
		}
	}()
	for {
		select {
		case <-done:
			return
		}
	}
}
func init() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./config.yaml") // 注意:如果使用相对路径，则是以 main.go 为当前位置与配置文件之间的路径
	err := viper.ReadInConfig()          // 查找并读取配置文件
	if err != nil {                      // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
