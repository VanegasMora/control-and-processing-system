package server

import (
	"backend-avanzada/auth"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // En producción, validar origen
	},
}

type WebSocketClient struct {
	conn   *websocket.Conn
	userID uint
	send   chan []byte
	server *Server
}

type WebSocketMessage struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

func (s *Server) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Extraer token del query string o header
	token := r.URL.Query().Get("token")
	if token == "" {
		authHeader := r.Header.Get("Authorization")
		if authHeader != "" {
			parts := strings.Split(authHeader, " ")
			if len(parts) == 2 && parts[0] == "Bearer" {
				token = parts[1]
			}
		}
	}

	if token == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Validar token
	claims, err := auth.ValidateToken(token)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	userID := claims.UserID

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading connection: %v", err)
		return
	}

	client := &WebSocketClient{
		conn:   conn,
		userID: userID,
		send:   make(chan []byte, 256),
		server: s,
	}

	s.websocketMutex.Lock()
	s.websocketClients[userID] = client
	s.websocketMutex.Unlock()

	go client.writePump()
	go client.readPump()
}

func (c *WebSocketClient) readPump() {
	defer func() {
		c.server.websocketMutex.Lock()
		delete(c.server.websocketClients, c.userID)
		c.server.websocketMutex.Unlock()
		c.conn.Close()
	}()

	for {
		_, _, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket error: %v", err)
			}
			break
		}
	}
}

func (c *WebSocketClient) writePump() {
	defer c.conn.Close()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Printf("Error writing message: %v", err)
				return
			}
		}
	}
}

func (s *Server) NotifyWebSocket(messageType string, payload interface{}) {
	msg := WebSocketMessage{
		Type:    messageType,
		Payload: payload,
	}

	data, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Error marshaling WebSocket message: %v", err)
		return
	}

	s.websocketMutex.RLock()
	defer s.websocketMutex.RUnlock()

	// Enviar a todos los clientes conectados
	for _, client := range s.websocketClients {
		select {
		case client.send <- data:
		default:
			close(client.send)
			delete(s.websocketClients, client.userID)
		}
	}
}

func (s *Server) NotifyUser(userID uint, messageType string, payload interface{}) {
	msg := WebSocketMessage{
		Type:    messageType,
		Payload: payload,
	}

	data, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Error marshaling WebSocket message: %v", err)
		return
	}

	s.websocketMutex.RLock()
	client, exists := s.websocketClients[userID]
	s.websocketMutex.RUnlock()

	if exists {
		select {
		case client.send <- data:
		default:
			close(client.send)
			s.websocketMutex.Lock()
			delete(s.websocketClients, userID)
			s.websocketMutex.Unlock()
		}
	}
}
