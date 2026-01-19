package main

import (
	"context"
	"fmt"
	"net/http"
	"github.com/quic-go/webtransport-go"
)

func main() {
	server := webtransport.Server{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	http.HandleFunc("/radio", func(w http.ResponseWriter, r *http.Request) {
		session, err := server.Upgrade(w, r)
		if err != nil {
			return
		}
		fmt.Println("Celular conectado")
		for {
			stream, err := session.AcceptUniStream(context.Background())
			if err != nil {
				break
			}
			go func(s webtransport.ReceiveStream) {
				buf := make([]byte, 4096)
				for {
					n, err := s.Read(buf)
					if err != nil {
						break
					}
					fmt.Printf("Audio: %d bytes\n", n)
				}
			}(stream)
		}
	})

	fmt.Println("Servidor en puerto 8080")
	http.ListenAndServe(":8080", nil)
}