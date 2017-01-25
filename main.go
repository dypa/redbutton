package main

import (
	"net/http"
	"log"
	"strconv"
	"github.com/hoisie/redis"
	"github.com/googollee/go-socket.io"
)

const (
	serverPort = 8888
	key = "hits"
)

func main() {
	var client redis.Client
	client.Addr = "redis:6379"

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")

		so.Join("allClients")

		so.On("click", func(msg string) {
			log.Println("click")

			_, err := client.Incr(key)
			if err != nil {
				log.Fatal(err)
			}

			val, err := client.Get(key)
			if err != nil {
				log.Fatal(err)
			}

			log.Println("clicked" + string(val))

			so.Emit("pull", string(val))
			so.BroadcastTo("allClients", "pull", string(val))
		})

		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})

	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./web")))
	log.Println("Serving at localhost " + strconv.Itoa(serverPort) + "...")
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(serverPort), nil))
}