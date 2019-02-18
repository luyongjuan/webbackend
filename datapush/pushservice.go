package datapush

import (
	"luyongjuan/webbackend/pusher"
	"net/http"
)
//init websocket service
func DataPusherInit(){

	hub := pusher.NewHub()
	go hub.Run()

	http.HandleFunc("/op/ws", func(writer http.ResponseWriter, request *http.Request) {
		pusher.ServeWs(hub, writer, request)
	})

	return
}