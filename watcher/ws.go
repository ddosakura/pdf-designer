package watcher

import (
	"encoding/json"

	"github.com/ddosakura/gklang"
	"golang.org/x/net/websocket"
)

var (
	webPages = NewSet()
)

// WsFreshHandler to fresh
func WsFreshHandler(ws *websocket.Conn) {
	defer removeFreshWebPageCallback(ws)
	addFreshWebPageCallback(ws)
	var reply string
	for {
		if err := websocket.Message.Receive(ws, &reply); err != nil {
			break
		}
	}
}

func addFreshWebPageCallback(ws *websocket.Conn) {
	webPages.Add(ws)
}

func removeFreshWebPageCallback(ws *websocket.Conn) {
	webPages.Remove(ws)
}

func callFreshWebPage(vv interface{}) {
	v, e := json.Marshal(vv)
	if e != nil {
		gklang.Log(gklang.LWarn, e)
		return
	}
	list := webPages.List()
	// fmt.Println(list, v)
	for i := range list {
		ws := list[i].(*websocket.Conn)
		if err := websocket.Message.Send(ws, v); err != nil {
			gklang.Log(gklang.LWarn, err)
			removeFreshWebPageCallback(ws)
		}
	}
}
