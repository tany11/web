package websocket

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		return origin == os.Getenv("FRONT_URL")
	},
}

type Server struct {
	clients    map[*websocket.Conn]bool
	broadcast  chan []byte
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
	mutex      sync.Mutex
}

func NewServer() *Server {
	return &Server{
		clients:    make(map[*websocket.Conn]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
	}
}

func (s *Server) Run() {
	for {
		select {
		case client := <-s.register:
			s.mutex.Lock()
			s.clients[client] = true
			s.mutex.Unlock()
		case client := <-s.unregister:
			s.mutex.Lock()
			if _, ok := s.clients[client]; ok {
				delete(s.clients, client)
				client.Close()
			}
			s.mutex.Unlock()
		case message := <-s.broadcast:
			s.mutex.Lock()
			for client := range s.clients {
				if err := client.WriteMessage(websocket.TextMessage, message); err != nil {
					log.Printf("error: %v", err)
					client.Close()
					delete(s.clients, client)
				}
			}
			s.mutex.Unlock()
		}
	}
}

func (s *Server) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	log.Printf("WebSocket接続の試行: %s", r.RemoteAddr)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket接続の失敗: %v", err)
		return
	}
	log.Printf("WebSocket接続の成功: %s", r.RemoteAddr)
	s.register <- conn

	go s.readPump(conn)
}

func (s *Server) readPump(conn *websocket.Conn) {
	defer func() {
		s.unregister <- conn
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		// メッセージを他のクライアントにブロードキャスト
		s.broadcast <- message
	}
}

func (s *Server) BroadcastMessage(message []byte) {
	s.broadcast <- message
}
