package libhttp

import (
	"net/http"
	"strconv"

	"library/service/handler"
	"library/service/models"

	"github.com/gin-gonic/gin"
)

type BorrowDetailHTTPHandler struct {
	service *handler.BorrowDetailService
}

func NewBorrowDetailHTTPHandler(s *handler.BorrowDetailService) *BorrowDetailHTTPHandler {
	return &BorrowDetailHTTPHandler{service: s}
}

// Post/borrowdetails
func (h *BorrowDetailHTTPHandler) CreateBorrowDetail(c *gin.Context) {
	var bd models.BorrowDetail
	if err := c.ShouldBindJSON(&bd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created, err := h.service.CreateBorrowDetail(&bd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

func (h *BorrowDetailHTTPHandler) GetBorrowDetailsByBorrowID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	list, err := h.service.GetBorrowDetailsByBorrowID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}
func (h *BorrowDetailHTTPHandler) ReturnBorrowDetail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.ReturnBorrowDetail(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "returned"})
}
func (h *BorrowDetailHTTPHandler) GetMemberBorrowHistory(c *gin.Context) {
	memberID, _ := strconv.Atoi(c.Param("id"))

	result, err := h.service.GetMemberBorrowHistory(memberID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot fetch history"})
		return
	}

	c.JSON(http.StatusOK, result)
}
