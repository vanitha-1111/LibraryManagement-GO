package libhttp

import (
	"net/http"
	"strconv"

	"library/service/handler"
	"library/service/models"

	"github.com/gin-gonic/gin"
)

type MemberHTTPHandler struct {
	service *handler.MemberService
}

func NewMemberHTTPHandler(service *handler.MemberService) *MemberHTTPHandler {
	return &MemberHTTPHandler{service: service}
}

func (h *MemberHTTPHandler) CreateMember(c *gin.Context) {
	var m models.Member
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	created, err := h.service.CreateMember(&m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}

func (h *MemberHTTPHandler) ListMembers(c *gin.Context) {
	list, err := h.service.ListMembers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *MemberHTTPHandler) GetMemberByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	m, err := h.service.GetMemberByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "member not found"})
		return
	}
	c.JSON(http.StatusOK, m)
}

func (h *MemberHTTPHandler) UpdateMember(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid member id"})
		return
	}
	var m models.Member
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}
	m.MemberID = id
	updated, err := h.service.UpdateMember(&m)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)

}
func (h *MemberHTTPHandler) DeleteMember(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteMember(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to delete"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
