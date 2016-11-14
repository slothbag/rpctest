package main

import (
	"net/rpc"
	"net/rpc/jsonrpc"
	"net/http"
	"golang.org/x/net/websocket"
	"time"
)

type MyRPC struct {}

func (self MyRPC) MyFunc(T1 int, T2 *bool) error {
	*T2 = true
	return nil
}


func sendUI(w http.ResponseWriter, r *http.Request) {
	//THIS SHOULD SEND UI FROM IPFS OR SOMETHING

	var html string = "<html><head></head><script language=\"javascript\">\n"
	html += "var loc = window.location;\n"
	html += "var exampleSocket = new WebSocket(\"ws://\" + loc.host + \"/conn\");\n"
	html += "exampleSocket.onopen = function (event) {\n"
	html += "   exampleSocket.send('{\"id\":1, \"method\":\"MyRPC.MyFunc\", \"params\":[1]}');\n"
	html += "}\n"
	html += "exampleSocket.onmessage = function (event) {\n"
	html += "       document.body.innerHTML += JSON.stringify(event.data) + '<br>';\n"
	html += "}\n"
	html += "</script><body></body></html>"

	w.Write([]byte(html))
}

func main() { 
        rpc.Register(new(MyRPC)) 

        http.Handle("/conn", websocket.Handler(serve)) 
        http.HandleFunc("/", sendUI)
        http.ListenAndServe(":7000", nil)
} 

func serve(ws *websocket.Conn) { 
	go sendBoo(ws)	
    jsonrpc.ServeConn(ws) 
} 

func sendBoo(ws *websocket.Conn) {
	timer := time.NewTimer(time.Second * 5)
	<- timer.C
	ws.Write([]byte("boo"))
}
