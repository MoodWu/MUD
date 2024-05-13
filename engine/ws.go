package engine

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"log"

	"github.com/gorilla/websocket"
)

func InitWebSocket() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/", func(c *gin.Context) {
		wsHandler(c.Writer, c.Request)
	})

	router.Run(":5250")
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()

	for {
		// 读取消息
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// 处理消息
		log.Printf("收到消息: %s", msg)

		// 回复消息
		err = conn.WriteMessage(websocket.TextMessage, []byte("已收到消息！"))
		if err != nil {
			log.Println(err)
			return
		}
	}
}
