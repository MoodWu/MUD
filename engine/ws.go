package engine

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gorilla/websocket"
)

func InitWebSocket() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//将 ./static 目录下的文件直接映射到根路径 /
	router.StaticFS("/client", http.Dir("./client"))

	// // 处理HTTP请求,
	// router.Any("/", func(c *gin.Context) {
	// 	if c.Request.Header.Get("Upgrade") == "websocket" {
	// 		log.Println("upgrade")
	// 		// c.Next() // 继续后续的处理，交给WebSocket处理器
	// 		wsHandler(c.Writer, c.Request)
	// 	} else {
	// 		c.File("mud.html") // 返回client.html文件

	// 	}
	// })

	router.GET("/ws", func(c *gin.Context) {
		if c.Request.Header.Get("Upgrade") == "websocket" {
			log.Println("upgrade")
			// c.Next() // 继续后续的处理，交给WebSocket处理器
			wsHandler(c.Writer, c.Request)
		}
	})

	// // GET 与 Any针对同一个url，Any必须在Get之前
	// router.GET("/", func(c *gin.Context) {
	// 	wsHandler(c.Writer, c.Request)
	// })
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
		log.Println("upgrade error.", err)
		return
	}

	defer conn.Close()

	ws := WebSocket{Conn: conn}

	err = ws.sendWelcomeMsg()
	if err != nil {
		log.Println("send welcome msg error.", err)
		return
	}

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
		err = ws.process(msg)
		// err = conn.WriteMessage(websocket.TextMessage, []byte("已收到消息！"))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func (ws *WebSocket) sendWelcomeMsg() error {
	welcome := Command{CMD: "welcome", Data: "欢迎来到叮叮爸爸王国，宇宙中最有趣的王国~"}
	data, err := json.Marshal(welcome)
	if err != nil {
		return err
	}

	ws.Conn.WriteMessage(websocket.TextMessage, data)

	return nil
}

func (ws *WebSocket) process(msg []byte) error {
	//解析命令
	var cmd, ret Command
	err := json.Unmarshal(msg, &cmd)
	if err != nil {
		return err
	}

	// ret, ok := commands[cmd.CMD]
	// if !ok {
	// 	ret = commands[""]
	// }
	// data, err := ret.Process(cmd)
	switch cmd.CMD {
	case "login":
		//todo:判断用户是否存在
		ret.CMD = "loginpwd"
		ret.Data = "请输入密码："
		ws.Player = &Player{}
		ws.Player.Name = cmd.Data
	case "loginpwd":
		//todo:检查用户密码是否正确
		ret.CMD = "online"

		ws.Player = LoadPlayer(ws.Player.Name)
		ret.Data = ws.Player.Online()
	case "passwd":
		//重设密码
	default:
		//命令交互,直接由player处理

		ret.CMD = ""
		ret.Data = ws.Player.Process(cmd.Data)
	}

	data, err := json.Marshal(ret)

	if err != nil {
		return err
	}

	ws.Conn.WriteMessage(websocket.TextMessage, data)

	return nil
}
