package handlers

import (
	"github.com/DSC-UNSRI/backend-url-shortener/model"
	"github.com/DSC-UNSRI/backend-url-shortener/usecase"
	"github.com/gofiber/fiber/v2"
)

type LinkHandler struct {
	usecase.LinkUsecase
}

func NewLinkHandler(uc usecase.LinkUsecase) *LinkHandler {
	return &LinkHandler{
		LinkUsecase: uc,
	}
}

func (h *LinkHandler) GetAllLinks(c *fiber.Ctx) error {
	c.JSON(h.LinkUsecase.GetAllLinks())

	return nil
}

func (h *LinkHandler) Redirect(c *fiber.Ctx) error {
	shortenedLink := c.Params("shortened_link")
	originLink := h.LinkUsecase.Redirect(shortenedLink)
	if originLink == "error" {
		c.JSON(model.Response{
			Status:  500,
			Data:    nil,
			Message: "Gagal mengambil data",
		})
	} else if originLink == "" {
		c.JSON(model.Response{
			Status:  404,
			Data:    nil,
			Message: "Link not found",
		})
	} else {
		c.Redirect(originLink)
	}

	return nil
}
