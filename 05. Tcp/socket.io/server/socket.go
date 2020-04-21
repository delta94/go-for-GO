package server

import (
	"github.com/googollee/go-socket.io"
	"log"
	"sort"
	"sync"
)

var s *socketServer

func NewSocketServer() (*socketServer, error) {
	server, err := socketio.NewServer(nil)
	if err != nil {
		return nil, err
	}
	s = &socketServer{
		Server:  server,
		clients: []*client{},
		mutex:   &sync.Mutex{},
	}

	s.Server.OnConnect("/", func(conn socketio.Conn) (err error) {
		if err = s.accept(conn); err == nil {
			log.Printf("Completed connecting with new client(id: %s)! tatal client: %d", conn.ID(), len(s.clients))
		}
		return
	})
	s.Server.OnDisconnect("/", func(conn socketio.Conn, reason string) {
		if err = s.remove(conn); err == nil {
			log.Printf("Completed disconnecting with client(id: %s)! tatal client: %d", conn.ID(), len(s.clients))
		}
	})
	return s, nil
}

func (s *socketServer) accept(conn socketio.Conn) error {
	c := &client{
		conn: conn,
		id: conn.ID(),
	}
	s.clients = append(s.clients, c)
	return nil
}

func (s *socketServer) remove(conn socketio.Conn) error {
	idx := sort.Search(len(s.clients), func(i int) bool {
		return s.clients[i].id == conn.ID()
	})
	s.clients = append(s.clients[:idx], s.clients[idx+1:]...)
	return nil
}