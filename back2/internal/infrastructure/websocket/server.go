package websocket

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/websocket"
)

type Server struct {
	clients    map[*websocket.Conn]Client
	groups     map[int]map[*websocket.Conn]bool
	broadcast  chan Message
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
	mutex      sync.Mutex
	upgrader   websocket.Upgrader
}

type Client struct {
	GroupID int
	StaffID string
}

type Message struct {
	GroupID int         `json:"groupId"`
	Content interface{} `json:"content"`
}

func NewServer() (*Server, error) {
	return &Server{
		clients:    make(map[*websocket.Conn]Client),
		groups:     make(map[int]map[*websocket.Conn]bool),
		broadcast:  make(chan Message),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}, nil
}

func (s *Server) Run() {
	for {
		select {
		case conn := <-s.register:
			s.mutex.Lock()
			client := s.clients[conn]
			if _, ok := s.groups[client.GroupID]; !ok {
				s.groups[client.GroupID] = make(map[*websocket.Conn]bool)
			}
			s.groups[client.GroupID][conn] = true
			s.mutex.Unlock()
		case conn := <-s.unregister:
			s.mutex.Lock()
			s.closeConnection(conn)
			s.mutex.Unlock()
		case message := <-s.broadcast:
			s.mutex.Lock()
			s.broadcastMessage(message)
			s.mutex.Unlock()
			// ... 他のケース ...
		}
	}
}

func (s *Server) broadcastMessage(message Message) {
	if connections, ok := s.groups[message.GroupID]; ok {
		for conn := range connections {
			err := conn.WriteJSON(message)
			if err != nil {
				log.Printf("警告: グループ %d のクライアントへのメッセージ送信に失敗: %v", message.GroupID, err)
				s.closeConnection(conn)
			}
		}
	}
}

func (s *Server) closeConnection(conn *websocket.Conn) {
	if client, ok := s.clients[conn]; ok {
		groupID := client.GroupID
		if group, ok := s.groups[groupID]; ok {
			delete(group, conn)
			if len(group) == 0 {
				delete(s.groups, groupID)
			}
		}
		delete(s.clients, conn)
		conn.Close()
		log.Printf("クライアントの接続をクローズしました。総接続数: %d, グループ %d の接続数: %d", len(s.clients), groupID, len(s.groups[groupID]))
	}
}

func (s *Server) Serve(w http.ResponseWriter, r *http.Request) {
	log.Println("Serve called")
	log.Printf("リクエストヘッダー: %+v", r.Header)
	log.Printf("リクエストURL: %s", r.URL.String())

	tokenString := r.URL.Query().Get("token")
	if tokenString == "" {
		log.Println("トークンが見つかりません")
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		log.Printf("トークン検証エラー: %v", err)
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		log.Println("無効なトークンクレーム")
		http.Error(w, "Invalid token claims", http.StatusUnauthorized)
		return
	}

	groupID := int(claims["group_id"].(float64))
	staffID := claims["staff_id"].(string)

	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket接続のアップグレードに失敗: %v", err)
		return
	}

	s.mutex.Lock()
	client := Client{GroupID: groupID, StaffID: staffID}
	s.clients[conn] = client
	if _, ok := s.groups[groupID]; !ok {
		s.groups[groupID] = make(map[*websocket.Conn]bool)
	}
	s.groups[groupID][conn] = true
	s.mutex.Unlock()

	s.register <- conn

	log.Printf("新しいクライアントが接続しました。総接続数: %d, グループ %d の接続数: %d", len(s.clients), groupID, len(s.groups[groupID]))

	s.readPump(conn)
}

func (s *Server) readPump(conn *websocket.Conn) {
	defer func() {
		s.unregister <- conn
	}()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("警告: 予期せぬWebSocket接続クローズ: %v", err)
			}
			break
		}
	}
}

func (s *Server) BroadcastToGroup(groupID int, data interface{}) error {
	message := Message{
		GroupID: groupID,
		Content: data, // データをそのまま使用し、追加のエンコーディングを避ける
	}

	s.broadcast <- message

	return nil
}

// ... 他のメソッドは必要に応じて追加
