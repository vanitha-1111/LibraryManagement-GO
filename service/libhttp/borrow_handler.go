package libhttp

import (
	"net/http"
	"strconv"

	"library/service/handler"
	"library/service/models"

	"github.com/gin-gonic/gin"
)

type BorrowHTTPHandler struct {
	service *handler.BorrowService
}

func NewBorrowHTTPHandler(service *handler.BorrowService) *BorrowHTTPHandler {
	return &BorrowHTTPHandler{service: service}
}

func (h *BorrowHTTPHandler) CreateBorrow(c *gin.Context) {
	var b models.Borrow
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	id, err := h.service.CreateBorrow(&b)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	b.BorrowID = id
	c.JSON(http.StatusCreated, b)
}
func (h *BorrowHTTPHandler) GetBorrowByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	borrow, err := h.service.GetBorrowByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, borrow)
}
func (h *BorrowHTTPHandler) ListBorrows(c *gin.Context) {
	list, err := h.service.ListBorrows()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}
