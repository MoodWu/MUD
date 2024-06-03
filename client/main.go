package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"syscall/js"
)

var command string
var addMessage js.Value

var ws js.Value

func handleMessage(this js.Value, args []js.Value) interface{} {
	// 获取MessageEvent的data属性
	message := args[0].Get("data").String()
	//处理收到命令
	println("Received message:", message)
	Process(message)
	// Process the incoming message...
	return nil
}

func onOpen(this js.Value, args []js.Value) interface{} {
	println("WebSocket connected")
	return nil
}

func onClose(this js.Value, args []js.Value) interface{} {
	println("WebSocket disconnected")
	return nil
}

func connectWebSocket(this js.Value, args []js.Value) interface{} {
	url := args[0].String()
	addMessage = args[1]
	ws = js.Global().Get("WebSocket").New(url)
	ws.Set("onopen", js.FuncOf(onOpen))
	ws.Set("onmessage", js.FuncOf(handleMessage))
	ws.Set("onclose", js.FuncOf(onClose))
	return nil
}

func sendMessage(this js.Value, args []js.Value) interface{} {
	message := args[0].String()

	cmd := Command{CMD: command, Data: message}
	data, err := json.Marshal(cmd)
	if err != nil {
		println("send marshal error.", err)
		return nil
	}

	/*
		0: CONNECTING - 表示连接正在进行中。
		1: OPEN - 表示连接已经建立。
		2: CLOSING - 表示连接正在关闭。
		3: CLOSED - 表示连接已经关闭。
	*/
	if ws.Truthy() && ws.Get("readyState").Int() == 1 {
		println("send....")
		ws.Call("send", js.ValueOf(string(data)))
	}
	println("send:", string(data))
	return nil
}

// 处理服务器的回应
func Process(data string) {
	println("process:", data)
	command = "command"
	var cmd Command
	err := json.Unmarshal([]byte(data), &cmd)
	if err != nil {
		println("Cannot unmarshal Command.", err.Error())
		return
	}

	switch cmd.CMD {
	case "welcome":
		addMessage.Invoke(js.ValueOf("请输入用户名:"))
		command = "login"
	case "loginpwd":
		addMessage.Invoke(js.ValueOf("请输入密码:"))
		command = "loginpwd"
	case "map":
		detail := MapDetail{}
		err := json.Unmarshal([]byte(cmd.Data), &detail)
		if err != nil {
			fmt.Println("Unmarshal error.", err)
			addMessage.Invoke(js.ValueOf("出了点问题，世界被迷雾遮盖。"))
			return
		}
		mapDetail := showMap(detail)
		addMessage.Invoke(js.ValueOf(mapDetail))
	default:
		// 将回车转换为<br>
		msg := cmd.Data
		msg = strings.ReplaceAll(msg, "\n", "<br>")
		addMessage.Invoke(js.ValueOf(msg))
	}
}

func showMap(info MapDetail) string {
	ret := "<table border=\"0\">"
	//根据map的长宽设定table的行列
	for col := info.MapInfo.Width - 1; col >= 0; col-- {
		ret += "<tr>"
		for row := 0; row < info.MapInfo.Long; row++ {
			bFound := false

			for _, s := range info.MapInfo.Scenes {
				if s.X == row && s.Y == col {

					if info.MapInfo.X == row && info.MapInfo.Y == col {
						ret += "<td align=\"center\" style=\"color:red\">" + s.Name + "</td>"
					} else {
						ret += "<td align=\"center\" >" + s.Name + "</td>"
					}

					if _, ok := s.Path["east"]; ok {
						ret += "<td>----</td>"
					} else {
						ret += "<td></td>"
					}
					bFound = true
					break
				}
			}
			if !bFound {
				ret += "<td></td><td></td>"
			}
		}
		ret += "</tr>"
		ret += "<tr>"
		for row := 0; row < info.MapInfo.Long; row++ {
			bFound := false

			for _, s := range info.MapInfo.Scenes {
				if s.X == row && s.Y == col {
					if _, ok := s.Path["south"]; ok {
						ret += "<td align=\"center\">|</td><td ></td>"
					} else {
						ret += "<td></td><td></td>"
					}
					bFound = true
					break
				}
			}
			if !bFound {
				ret += "<td></td><td></td>"
			}
		}
		ret += "</tr>"
	}
	ret += "</table>"

	ret += "<div>" + strings.ReplaceAll(info.SceneDesc, "\n", "<br>") + "</div>"

	return ret
}

func registerCallbacks() {
	js.Global().Set("connectWebSocket", js.FuncOf(connectWebSocket))
	js.Global().Set("sendMessage", js.FuncOf(sendMessage))
}

func main() {
	c := make(chan struct{}, 0)

	registerCallbacks()
	<-c
}

// Start WebSocket connection
// func start(this js.Value, args []js.Value) interface{} {
// 	wsURL := args[0].String()
// 	onMessageCallback := args[1]

// 	u, err := url.Parse(wsURL)
// 	if err != nil {
// 		fmt.Println("Invalid URL:", err)
// 		return nil
// 	}

// 	// Connect to WebSocket server
// 	fmt.Println("Ready to connect:", u.String())
// 	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
// 	if err != nil {
// 		fmt.Println("WebSocket connection failed:", err, u)
// 		return nil
// 	}
// 	conn = c

// 	// Listen for messages
// 	go func() {
// 		defer conn.Close()
// 		for {
// 			_, message, err := conn.ReadMessage()
// 			if err != nil {
// 				fmt.Println("Read error:", err)
// 				break
// 			}

// 			//处理收到命令

// 			msg := Process(message)
// 			// Call the JavaScript onMessage callback
// 			onMessageCallback.Invoke(js.ValueOf(msg))
// 		}
// 	}()

// 	return nil
// }

// // Send message over WebSocket
// func sendMsg(this js.Value, args []js.Value) interface{} {
// 	if conn == nil {
// 		fmt.Println("WebSocket connection is not established.")
// 		return nil
// 	}
// 	message := args[0].String()

// 	cmd := Command{CMD: command, Data: message}
// 	data, err := json.Marshal(cmd)
// 	if err != nil {
// 		fmt.Println("Marshal err,", err)
// 		return nil
// 	}
// 	err = conn.WriteMessage(websocket.TextMessage, data)
// 	if err != nil {
// 		fmt.Println("Write error:", err)
// 	}
// 	onMessageCallback.Invoke(args[0])
// 	return nil
// }
