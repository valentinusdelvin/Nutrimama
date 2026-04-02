package http

import (
	"nutrimama/internal/usecase"
	"nutrimama/internal/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type ChatController struct {
	ChatUseCase *usecase.ChatUseCase
	ChatHub     *utils.ChatHub
}

func NewChatController(uc *usecase.ChatUseCase, hub *utils.ChatHub) *ChatController {
	return &ChatController{ChatUseCase: uc, ChatHub: hub}
}

func (c *ChatController) SendMessage(ctx *fiber.Ctx) error {
	role, ok := ctx.Locals("role").(string)
	if !ok {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "unauthorized"})
	}

	var userID int
	switch v := ctx.Locals("userID").(type) {
	case int:
		userID = v
	case float64:
		userID = int(v)
	}

	req := new(usecase.SendMessageRequest)
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	msg, err := c.ChatUseCase.SendMessage(userID, role, *req)
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
	}

	c.ChatHub.BroadcastToRoom(req.ConsultationID, msg)

	return utils.SuccessResponse(ctx, fiber.StatusCreated, "message stored securely and passed time-lock limits", msg)
}

func (c *ChatController) GetHistory(ctx *fiber.Ctx) error {
	role, ok := ctx.Locals("role").(string)
	if !ok {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "unauthorized"})
	}

	var userID int
	switch v := ctx.Locals("userID").(type) {
	case int:
		userID = v
	case float64:
		userID = int(v)
	}

	roomIdStr := ctx.Params("id")
	roomId, err := strconv.Atoi(roomIdStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid consultation id param"})
	}

	history, err := c.ChatUseCase.GetHistory(userID, role, roomId)
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, "fetched chat history seamlessly", history)
}

func (c *ChatController) HandleWebSocket(conn *websocket.Conn) {
	roomIdStr := conn.Params("id")
	roomId, err := strconv.Atoi(roomIdStr)
	if err != nil {
		conn.Close()
		return
	}

	tokenStr := conn.Query("token")
	if tokenStr == "" {
		conn.Close()
		return
	}

	claims, err := utils.ValidateToken(tokenStr)
	if err != nil {
		conn.Close()
		return
	}

	// Mathematically block out-of-bounds rooms before allowing memory allocations natively!
	_, err = c.ChatUseCase.ValidateSessionAccess(claims.ID, claims.Role, roomId, true)
	if err != nil {
		conn.WriteJSON(fiber.Map{"error": err.Error()})
		conn.Close()
		return
	}

	c.ChatHub.AddClient(roomId, conn)
	defer c.ChatHub.RemoveClient(roomId, conn)

	// Infinite Listen Loop
	for {
		type WSMessage struct {
			Message string `json:"message"`
		}
		var wsReq WSMessage
		if err := conn.ReadJSON(&wsReq); err != nil {
			break // Exits efficiently directly unmounting clients natively on errors!
		}

		// Map WebSocket messages natively backwards executing HTTP logic cleanly!
		sendReq := usecase.SendMessageRequest{
			ConsultationID: roomId,
			Message:        wsReq.Message,
		}
		
		savedMsg, err := c.ChatUseCase.SendMessage(claims.ID, claims.Role, sendReq)
		if err != nil {
			conn.WriteJSON(fiber.Map{"error": err.Error()})
			continue
		}

		c.ChatHub.BroadcastToRoom(roomId, savedMsg)
	}
}
