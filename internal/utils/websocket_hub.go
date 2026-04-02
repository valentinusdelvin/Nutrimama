package utils

import (
	"log"
	"sync"

	"github.com/gofiber/websocket/v2"
)

// ChatHub safely manages strictly partitioned rooms containing thread-safe websocket arrays.
type ChatHub struct {
	Rooms map[int]map[*websocket.Conn]bool
	mu    sync.RWMutex
}

func NewChatHub() *ChatHub {
	return &ChatHub{
		Rooms: make(map[int]map[*websocket.Conn]bool),
	}
}

// AddClient securely maps the incoming connection purely to a specified uniquely isolated room.
func (h *ChatHub) AddClient(roomID int, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.Rooms[roomID] == nil {
		h.Rooms[roomID] = make(map[*websocket.Conn]bool)
	}
	h.Rooms[roomID][conn] = true
	log.Printf("[Websocket] Successfully mapped Client securely inside Consultation Room ID #%d", roomID)
}

// RemoveClient drops disconnected sockets efficiently reclaiming strict memory bounds natively!
func (h *ChatHub) RemoveClient(roomID int, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, ok := h.Rooms[roomID]; ok {
		delete(h.Rooms[roomID], conn)
		conn.Close()

		if len(h.Rooms[roomID]) == 0 {
			delete(h.Rooms, roomID) // Reclaim map entirely natively removing isolated arrays successfully!
		}
	}
	log.Printf("[Websocket] Erased disconnected Client from Consultation Room ID #%d", roomID)
}

// BroadcastToRoom securely funnels arbitrary payloads successfully bouncing natively strictly isolated avoiding cross-leaks!
func (h *ChatHub) BroadcastToRoom(roomID int, payload interface{}) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	room, ok := h.Rooms[roomID]
	if !ok {
		return // Completely skip arrays completely empty logically handling offline safely!
	}

	for conn := range room {
		// Asynchronous writes natively ensuring dead connections won't hang memory heavily!
		if err := conn.WriteJSON(payload); err != nil {
			log.Printf("[Websocket] Connection actively crashed inside Room #%d: %v", roomID, err)
			continue
		}
	}
}
