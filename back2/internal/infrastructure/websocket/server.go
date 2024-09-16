package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	socketio "github.com/googollee/go-socket.io"
)

type Server struct {
	server *socketio.Server
}

func NewServer() (*Server, error) {
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		log.Printf("新しい接続: %s", s.ID())
		return nil
	})

	server.OnEvent("/", "join", func(s socketio.Conn, groupID int, staffID string) {
		log.Printf("クライアントがグループに参加: %s, GroupID: %d, StaffID: %s", s.ID(), groupID, staffID)
		s.Join(fmt.Sprintf("group_%d", groupID))
		s.SetContext(map[string]interface{}{
			"groupID": groupID,
			"staffID": staffID,
		})
	})

	server.OnEvent("/", "message", func(s socketio.Conn, msg string) {
		context := s.Context().(map[string]interface{})
		groupID := context["groupID"].(int)
		log.Printf("メッセージ受信: %s, GroupID: %d", msg, groupID)
		server.BroadcastToRoom("/", fmt.Sprintf("group_%d", groupID), "message", msg)
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Printf("エラー発生: %v", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Printf("接続解除: %s, 理由: %s", s.ID(), reason)
	})

	return &Server{server: server}, nil
}

func (s *Server) Serve(w http.ResponseWriter, r *http.Request) {
	log.Println("Serve called")

	// トークンの検証
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
		return
	}

	// "Bearer "プレフィックスを削除
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secretKey := os.Getenv("JWT_SECRET_KEY")
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		log.Printf("トークン検証エラー: %v", err)
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	s.server.ServeHTTP(w, r)
}

func (s *Server) BroadcastToGroup(groupID int, data interface{}) error {
	messageBytes, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("メッセージのマーシャリングに失敗: %v", err)
	}

	roomName := fmt.Sprintf("group_%d", groupID)
	s.server.BroadcastToRoom("/", roomName, "message", string(messageBytes))

	log.Printf("グループ %d にメッセージをブロードキャスト: %s", groupID, string(messageBytes))
	return nil
}

// ... 他のメソッドは必要に応じて追加
