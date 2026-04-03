package transport

import (
	"NumbersManagmentService/internal/transport/mapper"
	"NumbersManagmentService/internal/transport/requests"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Handler struct {
	commandService CommandService
	queryService   QueryService
	logger         *zap.Logger
}

func NewHandler(cs CommandService, qs QueryService, logger *zap.Logger) *Handler {
	return &Handler{
		commandService: cs,
		queryService:   qs,
		logger:         logger,
	}
}

func (h *Handler) Import(c *fiber.Ctx) error {

	var req requests.ImportRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	cmd := mapper.ToImportCommand(req)

	res, err := h.commandService.Import(c.Context(), cmd)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(mapper.ToImportResponse(res))
}

func (h *Handler) Search(c *fiber.Ctx) error {

	var req requests.SearchRequest
	if err := c.QueryParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	query := mapper.ToSearchQuery(req)

	res, err := h.queryService.Search(c.Context(), query)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(mapper.ToSearchResponse(res))
}
